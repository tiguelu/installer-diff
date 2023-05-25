package elasticache

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func ResourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserGroupCreate,
		Read:   resourceUserGroupRead,
		Update: resourceUserGroupUpdate,
		Delete: resourceUserGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		CustomizeDiff: verify.SetTagsDiff,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"engine": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"REDIS"}, false),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
			"user_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

var resourceUserGroupPendingStates = []string{
	"creating",
	"modifying",
}

func resourceUserGroupCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ElastiCacheConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(d.Get("tags").(map[string]interface{})))

	input := &elasticache.CreateUserGroupInput{
		Engine:      aws.String(d.Get("engine").(string)),
		UserGroupId: aws.String(d.Get("user_group_id").(string)),
	}

	if v, ok := d.GetOk("user_ids"); ok {
		input.UserIds = flex.ExpandStringSet(v.(*schema.Set))
	}

	// Tags are currently only supported in AWS Commercial.
	if len(tags) > 0 && meta.(*conns.AWSClient).Partition == endpoints.AwsPartitionID {
		input.Tags = Tags(tags.IgnoreAWS())
	}

	out, err := conn.CreateUserGroup(input)
	if err != nil {
		return fmt.Errorf("error creating ElastiCache User Group: %w", err)
	}

	d.SetId(aws.StringValue(out.UserGroupId))

	stateConf := &resource.StateChangeConf{
		Pending:    resourceUserGroupPendingStates,
		Target:     []string{"active"},
		Refresh:    resourceUserGroupStateRefreshFunc(d.Get("user_group_id").(string), conn),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		MinTimeout: 10 * time.Second,
		Delay:      30 * time.Second, // Wait 30 secs before starting
	}

	log.Printf("[INFO] Waiting for ElastiCache User Group (%s) to be available", d.Id())
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error creating ElastiCache User Group: %w", err)
	}

	return resourceUserGroupRead(d, meta)

}

func resourceUserGroupRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ElastiCacheConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	resp, err := FindElastiCacheUserGroupByID(conn, d.Id())
	if !d.IsNewResource() && (tfresource.NotFound(err) || tfawserr.ErrCodeEquals(err, elasticache.ErrCodeUserGroupNotFoundFault)) {
		d.SetId("")
		log.Printf("[DEBUG] ElastiCache User Group (%s) not found", d.Id())
		return nil
	}

	if err != nil && !tfawserr.ErrCodeEquals(err, elasticache.ErrCodeUserGroupNotFoundFault) {
		return fmt.Errorf("error describing ElastiCache User Group (%s): %w", d.Id(), err)
	}

	d.Set("arn", resp.ARN)
	d.Set("engine", resp.Engine)
	d.Set("user_ids", resp.UserIds)
	d.Set("user_group_id", resp.UserGroupId)

	// Tags are currently only supported in AWS Commercial.
	if meta.(*conns.AWSClient).Partition == endpoints.AwsPartitionID {
		tags, err := ListTags(conn, aws.StringValue(resp.ARN))

		if err != nil {
			return fmt.Errorf("error listing tags for ElastiCache User (%s): %w", aws.StringValue(resp.ARN), err)
		}

		tags = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

		//lintignore:AWSR002
		if err := d.Set("tags", tags.RemoveDefaultConfig(defaultTagsConfig).Map()); err != nil {
			return fmt.Errorf("error setting tags: %w", err)
		}

		if err := d.Set("tags_all", tags.Map()); err != nil {
			return fmt.Errorf("error setting tags_all: %w", err)
		}
	} else {
		d.Set("tags", nil)
		d.Set("tags_all", nil)
	}

	return nil
}

func resourceUserGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ElastiCacheConn
	hasChange := false

	if d.HasChangesExcept("tags", "tags_all") {
		req := &elasticache.ModifyUserGroupInput{
			UserGroupId: aws.String(d.Get("user_group_id").(string)),
		}

		if d.HasChange("user_ids") {
			o, n := d.GetChange("user_ids")
			usersRemove := o.(*schema.Set).Difference(n.(*schema.Set))
			usersAdd := n.(*schema.Set).Difference(o.(*schema.Set))

			if usersAdd.Len() > 0 {
				req.UserIdsToAdd = flex.ExpandStringSet(usersAdd)
				hasChange = true
			}
			if usersRemove.Len() > 0 {
				req.UserIdsToRemove = flex.ExpandStringSet(usersRemove)
				hasChange = true
			}
		}

		if hasChange {
			_, err := conn.ModifyUserGroup(req)
			if err != nil {
				return fmt.Errorf("error updating ElastiCache User Group (%q): %w", d.Id(), err)
			}
			stateConf := &resource.StateChangeConf{
				Pending:    resourceUserGroupPendingStates,
				Target:     []string{"active"},
				Refresh:    resourceUserGroupStateRefreshFunc(d.Get("user_group_id").(string), conn),
				Timeout:    d.Timeout(schema.TimeoutCreate),
				MinTimeout: 10 * time.Second,
				Delay:      30 * time.Second, // Wait 30 secs before starting
			}

			log.Printf("[INFO] Waiting for ElastiCache User Group (%s) to be available", d.Id())
			_, err = stateConf.WaitForState()
			if err != nil {
				return fmt.Errorf("error updating ElastiCache User Group (%q): %w", d.Id(), err)
			}
		}
	}

	// Tags are currently only supported in AWS Commercial.
	if d.HasChange("tags_all") && meta.(*conns.AWSClient).Partition == endpoints.AwsPartitionID {
		o, n := d.GetChange("tags_all")

		if err := UpdateTags(conn, d.Get("arn").(string), o, n); err != nil {
			return fmt.Errorf("error updating ElastiCache User Group (%s) tags: %w", d.Get("arn").(string), err)
		}
	}

	return resourceUserGroupRead(d, meta)
}

func resourceUserGroupDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ElastiCacheConn

	input := &elasticache.DeleteUserGroupInput{
		UserGroupId: aws.String(d.Id()),
	}

	_, err := conn.DeleteUserGroup(input)
	if err != nil && !tfawserr.ErrCodeEquals(err, elasticache.ErrCodeUserGroupNotFoundFault) {
		return fmt.Errorf("error deleting ElastiCache User Group: %w", err)
	}
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{},
		Refresh:    resourceUserGroupStateRefreshFunc(d.Get("user_group_id").(string), conn),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		MinTimeout: 10 * time.Second,
		Delay:      30 * time.Second, // Wait 30 secs before starting
	}

	log.Printf("[INFO] Waiting for ElastiCache User Group (%s) to be available", d.Id())
	_, err = stateConf.WaitForState()
	if err != nil {
		if tfawserr.ErrCodeEquals(err, elasticache.ErrCodeUserGroupNotFoundFault) || tfawserr.ErrCodeEquals(err, elasticache.ErrCodeInvalidUserGroupStateFault) {
			return nil
		}
		return fmt.Errorf("ElastiCache User Group cannot be deleted: %w", err)
	}

	return nil
}

func resourceUserGroupStateRefreshFunc(id string, conn *elasticache.ElastiCache) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		v, err := FindElastiCacheUserGroupByID(conn, id)

		if err != nil {
			log.Printf("Error on retrieving ElastiCache User Group when waiting: %s", err)
			return nil, "", err
		}

		if v == nil {
			return nil, "", nil
		}

		if v.Status != nil {
			log.Printf("[DEBUG] ElastiCache User Group status for instance %s: %s", id, *v.Status)
		}

		return v, *v.Status, nil
	}
}
