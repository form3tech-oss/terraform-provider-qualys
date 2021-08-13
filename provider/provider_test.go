package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"qualys": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

}

func TestProviderImpl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("FORM3_CLIENT_ID"); v == "" {
		t.Fatal("FORM3_CLIENT_ID must be set for acceptance tests")
	}

	if v := os.Getenv("FORM3_CLIENT_SECRET"); v == "" {
		t.Fatal("FORM3_CLIENT_SECRET must be set for acceptance tests")
	}

	if v := os.Getenv("FORM3_ORGANISATION_ID"); v == "" {
		t.Fatal("FORM3_ORGANISATION_ID must be set for acceptance tests")
	}
}
