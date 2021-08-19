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
		Description: "Returns the details of the connector instances defined by the `connector_id`",
		ReadContext: dataSourceGCPConnectorRead,

		Schema: map[string]*schema.Schema{
			"cloud_provider": {
				Description: "The cloud provider associated with this connector",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connector_id": {
				Description: "The unique ID for this connector instance",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A string describing this connector instance",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"groups": {
				Description: "Groups that this connector belong to",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "Name of group",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"uuid": {
							Description: "ID of group",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"last_synced_on": {
				Description: "Last sync timestamp",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the connector",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"project_id": {
				Description: "GCP project id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"total_assets": {
				Description: "Total assets associated with this connector",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"state": {
				Description: "State of the connector",
				Type:        schema.TypeString,
				Computed:    true,
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
