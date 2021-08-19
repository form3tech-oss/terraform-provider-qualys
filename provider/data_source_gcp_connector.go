package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"

	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGCPConnector() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGCPConnectorRead,

		Schema: map[string]*schema.Schema{
			"cloud_provider": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connector_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"last_synced_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_assets": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGCPConnectorRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Reading gcp connector %q ", d.Get("connector_id"))

	service := meta.(*gcp.ConnectorService)
	connector, err := service.Get(d.Get("connector_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(connector.ConnectorID)

	return combineErrors(
		d.Set("cloud_provider", connector.Provider),
		d.Set("connector_id", connector.ConnectorID),
		d.Set("description", connector.Description),
		d.Set("groups", connector.Groups),
		d.Set("last_synced_on", connector.LastSyncedOn),
		d.Set("project_id", connector.Project),
		d.Set("name", connector.Name),
		d.Set("total_assets", connector.TotalAssets),
		d.Set("state", connector.State),
	)
}
