package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"

	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGCPConnector() *schema.Resource {
	return &schema.Resource{
		Description:   "A Qualys connector used for scanning GCP project assets",
		CreateContext: resourceGCPConnectorCreate,
		ReadContext:   resourceGCPConnectorRead,
		UpdateContext: resourceGCPConnectorUpdate,
		DeleteContext: resourceGCPConnectorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"connector_id": {
				Description: "The unique ID for this connector instance",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the connector",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A string describing this connector instance",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"gcp_credentials_json": {
				Description: "The JSON credentials file for a GCP service account",
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
			},
			"project_id": {
				Description: "GCP project id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_provider": {
				Description: "The cloud provider associated with this connector",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceGCPConnectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] create connector %q", d.Get("name").(string))

	opt := gcp.NewConnectorConfig().
		WithName(d.Get("name").(string)).
		WithDescription(d.Get("description").(string)).
		WithCredentialsJSON(d.Get("gcp_credentials_json").(string)).
		WithProjectId(d.Get("project_id").(string))

	service := meta.(*gcp.ConnectorService)
	connector, err := service.Create(opt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(connector.ConnectorID)

	return resourceGCPConnectorRead(ctx, d, meta)
}

func resourceGCPConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Reading gcp connector %q ", d.Id())

	service := meta.(*gcp.ConnectorService)
	connector, err := service.Get(d.Id())
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	return resourceDataFromConnector(ctx, d, connector)
}

func resourceGCPConnectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] update gcp connector %s", d.Id())

	opt := gcp.NewConnectorConfig().
		WithName(d.Get("name").(string)).
		WithDescription(d.Get("description").(string)).
		WithCredentialsJSON(d.Get("gcp_credentials_json").(string)).
		WithProjectId(d.Get("project_id").(string))

	service := meta.(*gcp.ConnectorService)
	err := service.Update(d.Id(), opt)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGCPConnectorRead(ctx, d, meta)
}

func resourceGCPConnectorDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Delete qualys connector %s", d.Id())

	service := meta.(*gcp.ConnectorService)
	return diag.FromErr(service.Delete([]string{d.Id()}))
}

func resourceDataFromConnector(_ context.Context, d *schema.ResourceData, connector *gcp.Connector) diag.Diagnostics {
	_, ok := d.GetOk("gcp_credentials_json")
	if !ok {
		if err := d.Set("gcp_credentials_json", ""); err != nil {
			return diag.FromErr(err)
		}
	}

	return combineErrors(
		d.Set("connector_id", connector.ConnectorID),
		d.Set("name", connector.Name),
		d.Set("description", connector.Description),
		d.Set("project_id", connector.Project),
	)
}
