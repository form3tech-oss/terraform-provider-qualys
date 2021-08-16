package provider

import (
	"fmt"
	"github.com/form3tech-oss/terraform-provider-qualys/cloudview/gcp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"testing"
)

func TestDataSourceGcpConnector_basic(t *testing.T) {
	ts := testServer()
	defer ts.Close()

	service := gcp.NewService(os.Getenv("QUALYS_URL"), os.Getenv("QUALYS_USERNAME"), os.Getenv("QUALYS_PASSWORD"))
	opt := gcp.NewConnectorConfig().
		WithName("test name").
		WithDescription("test description").
		WithProjectId("projectid").
		WithCredentialsJSON("{\"foo\" = \"bar\"}")
	resp, err := service.Create(opt)
	if err != nil {
		t.Fatalf("Create() err = %v", err)
	}

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		PreCheck:   func() { testAccPreCheck(t) },
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getHclDataSourceConfig(resp.ConnectorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "id", resp.ConnectorID),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "connector_id", resp.ConnectorID),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "name", "test name"),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "description", "test description"),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "project_id", "projectid"),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "cloud_provider", "gcp"),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "state", ""),
					resource.TestCheckResourceAttrSet("data.qualys_gcp_connector.test", "last_synced_on"),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "total_assets", "1"),
					resource.TestCheckResourceAttr("data.qualys_gcp_connector.test", "groups.#", "0"),
				),
			},
		},
	})
}

func getHclDataSourceConfig(connectorId string) string {
	return fmt.Sprintf(`
data "qualys_gcp_connector" "test" {
	connector_id          = "%s"
}
`, connectorId)
}
