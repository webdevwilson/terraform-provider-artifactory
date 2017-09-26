package artifactory

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccGroup_basic = `
resource "artifactory_group" "foobar" {
	name  = "acctest-basic"
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
					resource.TestCheckResourceAttr("artifactory_group.foobar", "name", "acctest-basic"),
				),
			},
		},
	})
}

const testAccGroup_full = `
resource "artifactory_group" "foobar" {
	name             = "acctest-full"
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
					resource.TestCheckResourceAttr("artifactory_group.foobar", "name", "acctest-full"),
					resource.TestCheckResourceAttr("artifactory_group.foobar", "auto_join", "true"),
				),
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
