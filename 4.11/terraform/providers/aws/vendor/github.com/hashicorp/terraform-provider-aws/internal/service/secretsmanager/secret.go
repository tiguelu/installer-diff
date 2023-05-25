package secretsmanager

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tfiam "github.com/hashicorp/terraform-provider-aws/internal/service/iam"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func ResourceSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretCreate,
		Read:   resourceSecretRead,
		Update: resourceSecretUpdate,
		Delete: resourceSecretDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_overwrite_replica_secret": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name_prefix"},
				ValidateFunc:  validSecretName,
			},
			"name_prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name"},
				ValidateFunc:  validSecretNamePrefix,
			},
			"policy": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: verify.SuppressEquivalentPolicyDiffs,
				StateFunc: func(v interface{}) string {
					json, _ := structure.NormalizeJsonString(v)
					return json
				},
			},
			"recovery_window_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
				ValidateFunc: validation.Any(
					validation.IntBetween(7, 30),
					validation.IntInSlice([]int{0}),
				),
			},
			"replica": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      secretReplicaHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_accessed_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"rotation_enabled": {
				Deprecated: "Use the aws_secretsmanager_secret_rotation resource instead",
				Type:       schema.TypeBool,
				Computed:   true,
			},
			"rotation_lambda_arn": {
				Deprecated: "Use the aws_secretsmanager_secret_rotation resource instead",
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
			},
			"rotation_rules": {
				Deprecated: "Use the aws_secretsmanager_secret_rotation resource instead",
				Type:       schema.TypeList,
				Computed:   true,
				Optional:   true,
				MaxItems:   1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"automatically_after_days": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},

		CustomizeDiff: verify.SetTagsDiff,
	}
}

func resourceSecretCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).SecretsManagerConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(d.Get("tags").(map[string]interface{})))

	var secretName string
	if v, ok := d.GetOk("name"); ok {
		secretName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		secretName = resource.PrefixedUniqueId(v.(string))
	} else {
		secretName = resource.UniqueId()
	}

	input := &secretsmanager.CreateSecretInput{
		Description:                 aws.String(d.Get("description").(string)),
		Name:                        aws.String(secretName),
		ForceOverwriteReplicaSecret: aws.Bool(d.Get("force_overwrite_replica_secret").(bool)),
	}

	if len(tags) > 0 {
		input.Tags = Tags(tags.IgnoreAWS())
	}

	if v, ok := d.GetOk("kms_key_id"); ok {
		input.KmsKeyId = aws.String(v.(string))
	}

	if v, ok := d.GetOk("replica"); ok && v.(*schema.Set).Len() > 0 {
		input.AddReplicaRegions = expandSecretsManagerSecretReplicas(v.(*schema.Set).List())
	}

	log.Printf("[DEBUG] Creating Secrets Manager Secret: %s", input)

	// Retry for secret recreation after deletion
	var output *secretsmanager.CreateSecretOutput
	err := resource.Retry(PropagationTimeout, func() *resource.RetryError {
		var err error
		output, err = conn.CreateSecret(input)
		// Temporarily retry on these errors to support immediate secret recreation:
		// InvalidRequestException: You can’t perform this operation on the secret because it was deleted.
		// InvalidRequestException: You can't create this secret because a secret with this name is already scheduled for deletion.
		if tfawserr.ErrMessageContains(err, secretsmanager.ErrCodeInvalidRequestException, "scheduled for deletion") || tfawserr.ErrMessageContains(err, secretsmanager.ErrCodeInvalidRequestException, "was deleted") {
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if tfresource.TimedOut(err) {
		output, err = conn.CreateSecret(input)
	}
	if err != nil {
		return fmt.Errorf("error creating Secrets Manager Secret: %w", err)
	}

	d.SetId(aws.StringValue(output.ARN))

	if v, ok := d.GetOk("policy"); ok && v.(string) != "" && v.(string) != "{}" {
		policy, err := structure.NormalizeJsonString(v.(string))

		if err != nil {
			return fmt.Errorf("policy (%s) is invalid JSON: %w", v.(string), err)
		}

		input := &secretsmanager.PutResourcePolicyInput{
			ResourcePolicy: aws.String(policy),
			SecretId:       aws.String(d.Id()),
		}

		err = resource.Retry(tfiam.PropagationTimeout, func() *resource.RetryError {
			_, err := conn.PutResourcePolicy(input)
			if tfawserr.ErrMessageContains(err, secretsmanager.ErrCodeMalformedPolicyDocumentException,
				"This resource policy contains an unsupported principal") {
				return resource.RetryableError(err)
			}
			if err != nil {
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if tfresource.TimedOut(err) {
			_, err = conn.PutResourcePolicy(input)
		}
		if err != nil {
			return fmt.Errorf("error setting Secrets Manager Secret %q policy: %w", d.Id(), err)
		}
	}

	if v, ok := d.GetOk("rotation_lambda_arn"); ok && v.(string) != "" {
		input := &secretsmanager.RotateSecretInput{
			RotationLambdaARN: aws.String(v.(string)),
			RotationRules:     expandSecretsManagerRotationRules(d.Get("rotation_rules").([]interface{})),
			SecretId:          aws.String(d.Id()),
		}

		log.Printf("[DEBUG] Enabling Secrets Manager Secret rotation: %s", input)
		err := resource.Retry(1*time.Minute, func() *resource.RetryError {
			_, err := conn.RotateSecret(input)
			if err != nil {
				// AccessDeniedException: Secrets Manager cannot invoke the specified Lambda function.
				if tfawserr.ErrCodeEquals(err, "AccessDeniedException") {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if tfresource.TimedOut(err) {
			_, err = conn.RotateSecret(input)
		}
		if err != nil {
			return fmt.Errorf("error enabling Secrets Manager Secret %q rotation: %w", d.Id(), err)
		}
	}

	return resourceSecretRead(d, meta)
}

func resourceSecretRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).SecretsManagerConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &secretsmanager.DescribeSecretInput{
		SecretId: aws.String(d.Id()),
	}

	outputRaw, err := tfresource.RetryWhenNotFound(PropagationTimeout, func() (interface{}, error) {
		return conn.DescribeSecret(input)
	})

	if !d.IsNewResource() && tfawserr.ErrCodeEquals(err, secretsmanager.ErrCodeResourceNotFoundException) {
		log.Printf("[WARN] Secrets Manager Secret (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("error reading Secrets Manager Secret (%s): %w", d.Id(), err)
	}

	output := outputRaw.(*secretsmanager.DescribeSecretOutput)

	if output == nil {
		return fmt.Errorf("error reading Secrets Manager Secret (%s): empty response", d.Id())
	}

	d.Set("arn", output.ARN)
	d.Set("description", output.Description)
	d.Set("kms_key_id", output.KmsKeyId)
	d.Set("name", output.Name)

	if err := d.Set("replica", flattenSecretsManagerSecretReplicas(output.ReplicationStatus)); err != nil {
		return fmt.Errorf("error setting replica: %w", err)
	}

	pIn := &secretsmanager.GetResourcePolicyInput{
		SecretId: aws.String(d.Id()),
	}
	log.Printf("[DEBUG] Reading Secrets Manager Secret policy: %s", pIn)
	pOut, err := conn.GetResourcePolicy(pIn)
	if err != nil {
		return fmt.Errorf("error reading Secrets Manager Secret policy: %w", err)
	}

	if pOut.ResourcePolicy != nil {
		policyToSet, err := verify.PolicyToSet(d.Get("policy").(string), aws.StringValue(pOut.ResourcePolicy))

		if err != nil {
			return err
		}

		d.Set("policy", policyToSet)
	} else {
		d.Set("policy", "")
	}

	d.Set("rotation_enabled", output.RotationEnabled)

	if aws.BoolValue(output.RotationEnabled) {
		d.Set("rotation_lambda_arn", output.RotationLambdaARN)
		if err := d.Set("rotation_rules", flattenSecretsManagerRotationRules(output.RotationRules)); err != nil {
			return fmt.Errorf("error setting rotation_rules: %w", err)
		}
	} else {
		d.Set("rotation_lambda_arn", "")
		d.Set("rotation_rules", []interface{}{})
	}

	tags := KeyValueTags(output.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

	//lintignore:AWSR002
	if err := d.Set("tags", tags.RemoveDefaultConfig(defaultTagsConfig).Map()); err != nil {
		return fmt.Errorf("error setting tags: %w", err)
	}

	if err := d.Set("tags_all", tags.Map()); err != nil {
		return fmt.Errorf("error setting tags_all: %w", err)
	}

	return nil
}

func resourceSecretUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).SecretsManagerConn

	if d.HasChange("replica") {
		o, n := d.GetChange("replica")

		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		err := removeSecretsManagerSecretReplicas(conn, d.Id(), os.Difference(ns).List())

		if err != nil {
			return fmt.Errorf("error deleting Secrets Manager Secret replica: %w", err)
		}

		err = addSecretsManagerSecretReplicas(conn, d.Id(), d.Get("force_overwrite_replica_secret").(bool), ns.Difference(os).List())

		if err != nil {
			return fmt.Errorf("error adding Secrets Manager Secret replica: %w", err)
		}
	}

	if d.HasChanges("description", "kms_key_id") {
		input := &secretsmanager.UpdateSecretInput{
			Description: aws.String(d.Get("description").(string)),
			SecretId:    aws.String(d.Id()),
		}

		if v, ok := d.GetOk("kms_key_id"); ok {
			input.KmsKeyId = aws.String(v.(string))
		}

		log.Printf("[DEBUG] Updating Secrets Manager Secret: %s", input)
		_, err := conn.UpdateSecret(input)
		if err != nil {
			return fmt.Errorf("error updating Secrets Manager Secret: %w", err)
		}
	}

	if d.HasChange("policy") {
		if v, ok := d.GetOk("policy"); ok && v.(string) != "" && v.(string) != "{}" {
			policy, err := structure.NormalizeJsonString(v.(string))
			if err != nil {
				return fmt.Errorf("policy contains an invalid JSON: %w", err)
			}
			input := &secretsmanager.PutResourcePolicyInput{
				ResourcePolicy: aws.String(policy),
				SecretId:       aws.String(d.Id()),
			}

			log.Printf("[DEBUG] Setting Secrets Manager Secret resource policy; %#v", input)
			err = resource.Retry(tfiam.PropagationTimeout, func() *resource.RetryError {
				_, err := conn.PutResourcePolicy(input)
				if tfawserr.ErrMessageContains(err, secretsmanager.ErrCodeMalformedPolicyDocumentException,
					"This resource policy contains an unsupported principal") {
					return resource.RetryableError(err)
				}
				if err != nil {
					return resource.NonRetryableError(err)
				}
				return nil
			})
			if tfresource.TimedOut(err) {
				_, err = conn.PutResourcePolicy(input)
			}
			if err != nil {
				return fmt.Errorf("error setting Secrets Manager Secret %q policy: %w", d.Id(), err)
			}
		} else {
			input := &secretsmanager.DeleteResourcePolicyInput{
				SecretId: aws.String(d.Id()),
			}

			log.Printf("[DEBUG] Removing Secrets Manager Secret policy: %#v", input)
			_, err := conn.DeleteResourcePolicy(input)
			if err != nil {
				return fmt.Errorf("error removing Secrets Manager Secret %q policy: %w", d.Id(), err)
			}
		}
	}

	if d.HasChanges("rotation_lambda_arn", "rotation_rules") {
		if v, ok := d.GetOk("rotation_lambda_arn"); ok && v.(string) != "" {
			input := &secretsmanager.RotateSecretInput{
				RotationLambdaARN: aws.String(v.(string)),
				RotationRules:     expandSecretsManagerRotationRules(d.Get("rotation_rules").([]interface{})),
				SecretId:          aws.String(d.Id()),
			}

			log.Printf("[DEBUG] Enabling Secrets Manager Secret rotation: %s", input)
			err := resource.Retry(1*time.Minute, func() *resource.RetryError {
				_, err := conn.RotateSecret(input)
				if err != nil {
					// AccessDeniedException: Secrets Manager cannot invoke the specified Lambda function.
					if tfawserr.ErrCodeEquals(err, "AccessDeniedException") {
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			if tfresource.TimedOut(err) {
				_, err = conn.RotateSecret(input)
			}
			if err != nil {
				return fmt.Errorf("error updating Secrets Manager Secret %q rotation: %w", d.Id(), err)
			}
		} else {
			input := &secretsmanager.CancelRotateSecretInput{
				SecretId: aws.String(d.Id()),
			}

			log.Printf("[DEBUG] Cancelling Secrets Manager Secret rotation: %s", input)
			_, err := conn.CancelRotateSecret(input)
			if err != nil {
				return fmt.Errorf("error cancelling Secret Manager Secret %q rotation: %w", d.Id(), err)
			}
		}
	}

	if d.HasChange("tags_all") {
		o, n := d.GetChange("tags_all")
		if err := UpdateTags(conn, d.Id(), o, n); err != nil {
			return fmt.Errorf("error updating tags: %w", err)
		}
	}

	return resourceSecretRead(d, meta)
}

func resourceSecretDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).SecretsManagerConn

	if v, ok := d.GetOk("replica"); ok && v.(*schema.Set).Len() > 0 {
		err := removeSecretsManagerSecretReplicas(conn, d.Id(), v.(*schema.Set).List())

		if err != nil {
			return fmt.Errorf("error deleting Secrets Manager Secret replica: %w", err)
		}
	}

	input := &secretsmanager.DeleteSecretInput{
		SecretId: aws.String(d.Id()),
	}

	recoveryWindowInDays := d.Get("recovery_window_in_days").(int)
	if recoveryWindowInDays == 0 {
		input.ForceDeleteWithoutRecovery = aws.Bool(true)
	} else {
		input.RecoveryWindowInDays = aws.Int64(int64(recoveryWindowInDays))
	}

	log.Printf("[DEBUG] Deleting Secrets Manager Secret: %s", input)
	_, err := conn.DeleteSecret(input)
	if err != nil {
		if tfawserr.ErrCodeEquals(err, secretsmanager.ErrCodeResourceNotFoundException) {
			return nil
		}
		return fmt.Errorf("error deleting Secrets Manager Secret: %w", err)
	}

	return nil
}

func removeSecretsManagerSecretReplicas(conn *secretsmanager.SecretsManager, id string, tfList []interface{}) error {
	if len(tfList) == 0 {
		return nil
	}

	input := &secretsmanager.RemoveRegionsFromReplicationInput{
		SecretId: aws.String(id),
	}

	var regions []string

	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]interface{})

		if !ok {
			continue
		}

		regions = append(regions, tfMap["region"].(string))
	}

	input.RemoveReplicaRegions = aws.StringSlice(regions)

	log.Printf("[DEBUG] Removing Secrets Manager Secret Replicas: %s", input)

	_, err := conn.RemoveRegionsFromReplication(input)

	if err != nil {
		if tfawserr.ErrCodeEquals(err, secretsmanager.ErrCodeResourceNotFoundException) {
			return nil
		}

		return err
	}

	return nil
}

func addSecretsManagerSecretReplicas(conn *secretsmanager.SecretsManager, id string, forceOverwrite bool, tfList []interface{}) error {
	if len(tfList) == 0 {
		return nil
	}

	input := &secretsmanager.ReplicateSecretToRegionsInput{
		SecretId:                    aws.String(id),
		ForceOverwriteReplicaSecret: aws.Bool(forceOverwrite),
		AddReplicaRegions:           expandSecretsManagerSecretReplicas(tfList),
	}

	log.Printf("[DEBUG] Removing Secrets Manager Secret Replica: %s", input)

	_, err := conn.ReplicateSecretToRegions(input)

	if err != nil {
		return err
	}

	return nil
}

func expandSecretsManagerSecretReplica(tfMap map[string]interface{}) *secretsmanager.ReplicaRegionType {
	if tfMap == nil {
		return nil
	}

	apiObject := &secretsmanager.ReplicaRegionType{}

	if v, ok := tfMap["kms_key_id"].(string); ok && v != "" {
		apiObject.KmsKeyId = aws.String(v)
	}

	if v, ok := tfMap["region"].(string); ok && v != "" {
		apiObject.Region = aws.String(v)
	}

	return apiObject
}

func expandSecretsManagerSecretReplicas(tfList []interface{}) []*secretsmanager.ReplicaRegionType {
	if len(tfList) == 0 {
		return nil
	}

	var apiObjects []*secretsmanager.ReplicaRegionType

	for _, tfMapRaw := range tfList {
		tfMap, ok := tfMapRaw.(map[string]interface{})

		if !ok {
			continue
		}

		apiObject := expandSecretsManagerSecretReplica(tfMap)

		if apiObject == nil {
			continue
		}

		apiObjects = append(apiObjects, apiObject)
	}

	return apiObjects
}

func flattenSecretsManagerSecretReplica(apiObject *secretsmanager.ReplicationStatusType) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.KmsKeyId; v != nil {
		tfMap["kms_key_id"] = aws.StringValue(v)
	}

	if v := apiObject.LastAccessedDate; v != nil {
		tfMap["last_accessed_date"] = aws.TimeValue(v).Format(time.RFC3339)
	}

	if v := apiObject.Region; v != nil {
		tfMap["region"] = aws.StringValue(v)
	}

	if v := apiObject.Status; v != nil {
		tfMap["status"] = aws.StringValue(v)
	}

	if v := apiObject.StatusMessage; v != nil {
		tfMap["status_message"] = aws.StringValue(v)
	}

	return tfMap
}

func flattenSecretsManagerSecretReplicas(apiObjects []*secretsmanager.ReplicationStatusType) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var tfList []interface{}

	for _, apiObject := range apiObjects {
		if apiObject == nil {
			continue
		}

		tfList = append(tfList, flattenSecretsManagerSecretReplica(apiObject))
	}

	return tfList
}

func secretReplicaHash(v interface{}) int {
	var buf bytes.Buffer

	m := v.(map[string]interface{})

	if v, ok := m["kms_key_id"].(string); ok {
		buf.WriteString(fmt.Sprintf("%s-", v))
	}

	if v, ok := m["region"].(string); ok {
		buf.WriteString(fmt.Sprintf("%s-", v))
	}

	return create.StringHashcode(buf.String())
}
