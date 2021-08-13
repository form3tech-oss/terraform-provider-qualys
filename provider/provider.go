package provider

import (
	"errors"
	"strings"

	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUALYS_USERNAME", ""),
				Description: "The BasicAuth username to connect to Qualys API.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUALYS_PASSWORD", ""),
				Description: "The BasicAuth password to connect to Qualys API.",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUALYS_URL", ""),
				Description: "The Qualys Base API URL",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"qualys_gcp_connector": dataSourceGCPConnector(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"qualys_gcp_connector": resourceGCPConnector(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	baseURL := d.Get("base_url").(string)
	return gcp.NewService(baseURL, username, password), nil
}

func combineErrors(errs ...error) error {
	builder := &strings.Builder{}

	for _, e := range errs {
		if e == nil {
			continue
		}
		builder.WriteString(e.Error())
	}

	if builder.Len() > 0 {
		return errors.New(builder.String())
	}
	return nil
}
