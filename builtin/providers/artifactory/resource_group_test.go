package artifactory

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccGroup_basic = `
resource "artifactory_group" "foobar" {
	name  = "developers"
}`

func TestAccGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGroupDestroy("artifactory_group.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGroup_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_group.foobar", "name", "developers"),
				),
			},
		},
	})
}

const testAccGroup_full = `
resource "artifactory_group" "foobar" {
	name             = "developers"
    auto_join        = true
}`

func TestAccGroup_full(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGroupDestroy("artifactory_group.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGroup_full,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_group.foobar", "name", "developers"),
					resource.TestCheckResourceAttr("artifactory_group.foobar", "auto_join", "true"),
				),
			},
		},
	})
}

func TestAccGroup_import(t *testing.T) {
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

func testAccCheckGroupDestroy(id string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(Client)
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found %s", id)
		}

		_, err := client.GetGroup(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Group %s still exists", rs.Primary.ID)
		}

		return nil
	}
}
