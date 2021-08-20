package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]func() (*schema.Provider, error){
		"qualys": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

}

func TestProviderImpl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("LOCAL"); v == "" {
		if v := os.Getenv("QUALYS_USERNAME"); v == "" {
			t.Fatal("QUALYS_USERNAME must be set for acceptance tests")
		}

		if v := os.Getenv("QUALYS_PASSWORD"); v == "" {
			t.Fatal("QUALYS_PASSWORD must be set for acceptance tests")
		}

		if v := os.Getenv("QUALYS_URL"); v == "" {
			t.Fatal("QUALYS_URL must be set for acceptance tests")
		}
		return
	}
}
