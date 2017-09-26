package artifactory

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func IgnoreTestAccVirtualRepository_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_virtual_repository.foobar"),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVirtualRepository_full,
			},
			resource.TestStep{
				ResourceName:      "artifactory_virtual_repository.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
