package artifactory

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"artifactory": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ARTIFACTORY_USERNAME"); v == "" {
		t.Fatal("ARTIFACTORY_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("ARTIFACTORY_PASSWORD"); v == "" {
		t.Fatal("ARTIFACTORY_PASSWORD must be set for acceptance tests")
	}
	if v := os.Getenv("ARTIFACTORY_URL"); v == "" {
		t.Fatal("ARTIFACTORY_URL must be set for acceptance tests")
	}
}

func testAccCheckRepositoryDestroy(id string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(Client)
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found %s", id)
		}

		err := client.GetRepository(rs.Primary.ID, nil)

		if err == nil {
			return fmt.Errorf("Repository %s still exists", rs.Primary.ID)
		}

		return nil
	}
}
