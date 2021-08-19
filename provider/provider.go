package provider

import (
	"context"
	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
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

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	baseURL := d.Get("base_url").(string)
	return gcp.NewService(baseURL, username, password), []diag.Diagnostic{}
}

func combineErrors(errs ...error) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, e := range errs {
		if e == nil {
			diags = append(diags, diag.FromErr(e)...)
		}
	}

	return diags
}
