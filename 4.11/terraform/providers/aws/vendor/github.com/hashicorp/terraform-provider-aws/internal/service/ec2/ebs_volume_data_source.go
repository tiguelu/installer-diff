package ec2

import (
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func DataSourceEBSVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEBSVolumeRead,

		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"most_recent": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encrypted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"iops": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multi_attach_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"volume_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"outpost_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(),
			"throughput": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceEBSVolumeRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).EC2Conn

	filters, filtersOk := d.GetOk("filter")

	params := &ec2.DescribeVolumesInput{}
	if filtersOk {
		params.Filters = BuildFiltersDataSource(filters.(*schema.Set))
	}

	log.Printf("[DEBUG] Reading EBS Volume: %s", params)
	resp, err := conn.DescribeVolumes(params)
	if err != nil {
		return err
	}

	filteredVolumes := resp.Volumes[:]

	var volume *ec2.Volume
	if len(filteredVolumes) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVolumes) > 1 {
		recent := d.Get("most_recent").(bool)
		log.Printf("[DEBUG] aws_ebs_volume - multiple results found and `most_recent` is set to: %t", recent)
		if recent {
			volume = mostRecentVolume(filteredVolumes)
		} else {
			return fmt.Errorf("Your query returned more than one result. Please try a more " +
				"specific search criteria, or set `most_recent` attribute to true.")
		}
	} else {
		// Query returned single result.
		volume = filteredVolumes[0]
	}

	log.Printf("[DEBUG] aws_ebs_volume - Single Volume found: %s", *volume.VolumeId)
	return volumeDescriptionAttributes(d, meta.(*conns.AWSClient), volume)
}

type volumeSort []*ec2.Volume

func (a volumeSort) Len() int      { return len(a) }
func (a volumeSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a volumeSort) Less(i, j int) bool {
	itime := aws.TimeValue(a[i].CreateTime)
	jtime := aws.TimeValue(a[j].CreateTime)
	return itime.Unix() < jtime.Unix()
}

func mostRecentVolume(volumes []*ec2.Volume) *ec2.Volume {
	sortedVolumes := volumes
	sort.Sort(volumeSort(sortedVolumes))
	return sortedVolumes[len(sortedVolumes)-1]
}

func volumeDescriptionAttributes(d *schema.ResourceData, client *conns.AWSClient, volume *ec2.Volume) error {
	d.SetId(aws.StringValue(volume.VolumeId))
	d.Set("volume_id", volume.VolumeId)

	arn := arn.ARN{
		Partition: client.Partition,
		Region:    client.Region,
		Service:   ec2.ServiceName,
		AccountID: client.AccountID,
		Resource:  fmt.Sprintf("volume/%s", d.Id()),
	}
	d.Set("arn", arn.String())

	d.Set("availability_zone", volume.AvailabilityZone)
	d.Set("encrypted", volume.Encrypted)
	d.Set("iops", volume.Iops)
	d.Set("kms_key_id", volume.KmsKeyId)
	d.Set("size", volume.Size)
	d.Set("snapshot_id", volume.SnapshotId)
	d.Set("volume_type", volume.VolumeType)
	d.Set("outpost_arn", volume.OutpostArn)
	d.Set("multi_attach_enabled", volume.MultiAttachEnabled)
	d.Set("throughput", volume.Throughput)

	if err := d.Set("tags", KeyValueTags(volume.Tags).IgnoreAWS().IgnoreConfig(client.IgnoreTagsConfig).Map()); err != nil {
		return fmt.Errorf("error setting tags: %w", err)
	}

	return nil
}
