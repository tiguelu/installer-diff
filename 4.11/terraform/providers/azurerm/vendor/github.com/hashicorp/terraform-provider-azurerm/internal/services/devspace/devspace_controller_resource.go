package devspace

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/devspaces/mgmt/2019-04-01/devspaces"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/devspace/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/devspace/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceDevSpaceController() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceDevSpaceControllerCreate,
		Read:   resourceDevSpaceControllerRead,
		Update: resourceDevSpaceControllerUpdate,
		Delete: resourceDevSpaceControllerDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.ControllerID(id)
			return err
		}),

		DeprecationMessage: `DevSpace Controllers are deprecated and will be retired on 31 October 2023 - at this time the Azure API does not allow new Controllers to be provisioned, but existing DevSpace Controllers can continue to be used.

Since these are deprecated and can no longer be provisioned, version 3.0 of the Azure Provider will remove support for DevSpace Controllers.`,

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DevSpaceName(),
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupName(),

			"sku_name": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"S1",
				}, false),
			},

			"target_container_host_resource_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceID,
			},

			"target_container_host_credentials_base64": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringIsBase64,
			},

			"tags": tags.Schema(),

			"data_plane_fqdn": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"host_suffix": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDevSpaceControllerCreate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DevSpace.ControllersClient
	subscriptionid := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for DevSpace Controller creation")

	id := parse.NewControllerID(subscriptionid, d.Get("resource_group_name").(string), d.Get("name").(string))

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.Name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing %s: %s", id, err)
			}
		}

		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_devspace_controller", id.ID())
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	t := d.Get("tags").(map[string]interface{})

	sku, err := expandControllerSkuName(d.Get("sku_name").(string))
	if err != nil {
		return fmt.Errorf("expanding `sku_name` for %s: %v", id, err)
	}

	controller := devspaces.Controller{
		Location: &location,
		Sku:      sku,
		ControllerProperties: &devspaces.ControllerProperties{
			TargetContainerHostResourceID:        utils.String(d.Get("target_container_host_resource_id").(string)),
			TargetContainerHostCredentialsBase64: utils.String(d.Get("target_container_host_credentials_base64").(string)),
		},
		Tags: tags.Expand(t),
	}

	future, err := client.Create(ctx, id.ResourceGroup, id.Name, controller)
	if err != nil {
		return fmt.Errorf("creating %s: %+v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for creation of %s: %+v", id, err)
	}

	d.SetId(id.ID())

	return resourceDevSpaceControllerRead(d, meta)
}

func resourceDevSpaceControllerUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DevSpace.ControllersClient
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for DevSpace Controller updating")

	id, err := parse.ControllerID(d.Id())
	if err != nil {
		return err
	}
	params := devspaces.ControllerUpdateParameters{
		Tags: tags.Expand(d.Get("tags").(map[string]interface{})),
	}

	result, err := client.Update(ctx, id.ResourceGroup, id.Name, params)
	if err != nil {
		return fmt.Errorf("updating DevSpace Controller %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	if result.ID == nil {
		return fmt.Errorf("Cannot read DevSpace Controller %q (Resource Group %q) ID", id.Name, id.ResourceGroup)
	}
	d.SetId(id.ID())

	return resourceDevSpaceControllerRead(d, meta)
}

func resourceDevSpaceControllerRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DevSpace.ControllersClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.ControllerID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] DevSpace Controller %q was not found in Resource Group %q - removing from state!", id.Name, id.ResourceGroup)
			d.SetId("")
			return nil
		}

		return fmt.Errorf("making Read request on DevSpace Controller %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	d.Set("name", id.Name)
	d.Set("resource_group_name", id.ResourceGroup)
	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	if sku := resp.Sku; sku != nil {
		d.Set("sku_name", sku.Name)
	}

	if props := resp.ControllerProperties; props != nil {
		d.Set("host_suffix", props.HostSuffix)
		d.Set("data_plane_fqdn", props.DataPlaneFqdn)
		d.Set("target_container_host_resource_id", props.TargetContainerHostResourceID)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceDevSpaceControllerDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).DevSpace.ControllersClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.ControllerID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.Name)
	if err != nil {
		return fmt.Errorf("deleting DevSpace Controller %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the deletion of DevSpace Controller %q (Resource Group %q): %+v", id.Name, id.ResourceGroup, err)
	}

	return nil
}

func expandControllerSkuName(skuName string) (*devspaces.Sku, error) {
	var tier devspaces.SkuTier
	switch skuName[0:1] {
	case "S":
		tier = devspaces.Standard
	default:
		return nil, fmt.Errorf("sku_name %s has unknown sku tier %s", skuName, skuName[0:1])
	}

	return &devspaces.Sku{
		Name: utils.String(skuName),
		Tier: tier,
	}, nil
}
