package artifactory

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func IgnoreTestAccUser_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckUserDestroy("artifactory_user.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccUser_full,
			},
			resource.TestStep{
				ResourceName:      "artifactory_user.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
