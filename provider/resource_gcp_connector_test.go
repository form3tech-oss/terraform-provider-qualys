package provider

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGcpConnector_basic(t *testing.T) {
	projectId := uuid.New()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			// TODO
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: getHclConfig(projectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "name", "test name"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "description", "test description"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "gcp_credentials_json", "{\"foo\":\"bar\"}"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "project_id", projectId.String()),
				),
			},
		},
	})
}

func getHclConfig(projectId uuid.UUID) string {
	return fmt.Sprintf(`
resource "qualys_gcp_connector" "test" {
	name                  = "test name"
    description           = "test description"
    gcp_credentials_json  = jsonencode({foo = "bar"})
    project_id            = "%s"
}
`, projectId.String())
}

func testAccCheckConnectorExists() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return nil
	}
}
