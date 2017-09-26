package artifactory

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func IgnoreTestAccLocalRepository_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_local_repository.foobar"),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLocalRepository_basic,
			},
			resource.TestStep{
				ResourceName:      "artifactory_local_repository.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
