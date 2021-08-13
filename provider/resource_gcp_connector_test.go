package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGcpConnector_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			// TODO
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: "TODO",
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("foo.bar", "foo", "bar")),
			},
		},
	})
}
