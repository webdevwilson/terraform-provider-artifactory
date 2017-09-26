package artifactory

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func IgnoreTestAccGroup_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGroupDestroy("artifactory_group.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGroup_full,
			},
			resource.TestStep{
				ResourceName:      "artifactory_group.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
