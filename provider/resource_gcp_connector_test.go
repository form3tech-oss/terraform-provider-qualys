package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGcpConnector_basic(t *testing.T) {
	projectId := uuid.New()

	ts := testServer()
	defer ts.Close()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			service := gcp.NewService(os.Getenv("QUALYS_URL"), os.Getenv("QUALYS_USERNAME"), os.Getenv("QUALYS_PASSWORD"))
			for _, rs := range state.RootModule().Resources {
				_, err := service.Get(rs.Primary.ID)
				if err.Error() != "Not Found" {
					return fmt.Errorf("expected 'Not Found', unknown error: %s", err)
				}
			}
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
