package citrixsdx

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
// var providerFactories = map[string]func() (*schema.Provider, error){
// 	"citrixsdx": func() (*schema.Provider, error) {
// 		return New("dev")(), nil
// 	},
// }

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	requiredEnvVariables := []string{
		"CITRIXSDX_HOST",
		"CITRIXSDX_USERNAME",
		"CITRIXSDX_PASSWORD",
		"CITRIXSDX_SSL_VERIFY",
	}
	for _, envVar := range requiredEnvVariables {
		if v := os.Getenv(envVar); v == "" {
			t.Fatalf("%s must be set for acceptance tests", envVar)
		}
	}
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"citrixsdx": testAccProvider,
	}
}
