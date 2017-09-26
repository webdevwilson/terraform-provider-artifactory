package artifactory

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func IgnoreTestAccRemoteRepository_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_remote_repository.foobar"),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRemoteRepository_basic,
			},
			resource.TestStep{
				ResourceName:      "artifactory_remote_repository.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
