package main

import (
	provider "github.com/form3tech-oss/terraform-provider-qualys/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider})
}
