package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceGcpConnector_basic(t *testing.T) {
	projectId := uuid.New()

	ts := testServer()
	defer ts.Close()

	resource.Test(t, resource.TestCase{
		IsUnitTest:        true,
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckResourcesDestroyed,
		Steps: []resource.TestStep{
			{
				Config: getHclConfig(projectId, "test name"),
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

func TestResourceGcpConnector_update(t *testing.T) {
	projectId := uuid.New()

	ts := testServer()
	defer ts.Close()

	resource.Test(t, resource.TestCase{
		IsUnitTest:        true,
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckResourcesDestroyed,
		Steps: []resource.TestStep{
			{
				Config: getHclConfig(projectId, "test name"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "name", "test name"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "description", "test description"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "gcp_credentials_json", "{\"foo\":\"bar\"}"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "project_id", projectId.String()),
				),
			},
			{
				Config: getHclConfig(projectId, "test name update"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "name", "test name update"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "description", "test description"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "gcp_credentials_json", "{\"foo\":\"bar\"}"),
					resource.TestCheckResourceAttr("qualys_gcp_connector.test", "project_id", projectId.String()),
				),
			},
		},
	})
}

func getHclConfig(projectId uuid.UUID, name string) string {
	return fmt.Sprintf(`
resource "qualys_gcp_connector" "test" {
    name                  = "%s"
    description           = "test description"
    gcp_credentials_json  = jsonencode({foo = "bar"})
    project_id            = "%s"
}
`, name, projectId.String())
}

func testAccCheckResourcesDestroyed(state *terraform.State) error {
	service := gcp.NewService(os.Getenv("QUALYS_URL"), os.Getenv("QUALYS_USERNAME"), os.Getenv("QUALYS_PASSWORD"))
	for _, rs := range state.RootModule().Resources {
		_, err := service.Get(rs.Primary.ID)
		if err.Error() != "Not Found" {
			return fmt.Errorf("expected 'Not Found', unknown error: %s", err)
		}
	}
	return nil
}
