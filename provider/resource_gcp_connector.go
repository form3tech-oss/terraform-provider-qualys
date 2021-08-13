package provider

import (
	"log"

	gcp "github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGCPConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceGCPConnectorCreate,
		Read:   resourceGCPConnectorRead,
		Update: resourceGCPConnectorUpdate,
		Delete: resourceGCPConnectorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"connector_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_credentials_json": {
				Type:      schema.TypeString,
				Sensitive: true,
				Required:  true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_provider": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGCPConnectorCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] create connector %q", d.Get("name").(string))

	opt := gcp.NewConnectorConfig().
		WithName(d.Get("name").(string)).
		WithDescription(d.Get("description").(string)).
		WithConfigFile(d.Get("gcp_credentials_json").(string)).
		WithProjectId(d.Get("project_id").(string))

	service := meta.(*gcp.ConnectorService)
	connector, err := service.Create(opt)
	if err != nil {
		return err
	}

	d.SetId(connector.ConnectorID)

	return resourceGCPConnectorRead(d, meta)
}

func resourceGCPConnectorRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading gcp connector %q ", d.Id())

	service := meta.(*gcp.ConnectorService)
	connector, err := service.Get(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	return resourceDataFromConnector(d, connector)
}

func resourceGCPConnectorUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] update gcp connector %s", d.Id())

	opt := gcp.NewConnectorConfig().
		WithName(d.Get("name").(string)).
		WithDescription(d.Get("description").(string)).
		WithConfigFile(d.Get("gcp_credentials_json").(string)).
		WithProjectId(d.Get("project_id").(string))

	service := meta.(*gcp.ConnectorService)
	err := service.Update(d.Id(), opt)
	if err != nil {
		return err
	}

	return resourceGCPConnectorRead(d, meta)
}

func resourceGCPConnectorDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Delete qualys connector %s", d.Id())

	service := meta.(*gcp.ConnectorService)
	return service.Delete([]string{d.Id()})
}

func resourceDataFromConnector(d *schema.ResourceData, connector *gcp.Connector) error {
	// TODO wtf is this doing?
	_, ok := d.GetOk("gcp_credentials_json")
	if !ok {
		if err := d.Set("gcp_credentials_json", ""); err != nil {
			return err
		}
	}

	return combineErrors(
		d.Set("connector_id", connector.ConnectorID),
		d.Set("name", connector.Name),
		d.Set("description", connector.Description),
		d.Set("project_id", connector.Project),
	)
}
