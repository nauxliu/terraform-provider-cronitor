package cronitor

import (
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProviders = map[string]terraform.ResourceProvider{
		"cronitor": Provider(),
	}
}

func testAccPreCheck(t *testing.T) {
	variables := []string{
		ApiKeyEnvName,
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("`%s` must be set for acceptance tests!", variable)
		}
	}
}
