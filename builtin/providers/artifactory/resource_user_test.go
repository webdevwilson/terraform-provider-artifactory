package artifactory

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccUser_basic = `
resource "artifactory_user" "foobar" {
	name  = "the.dude"
    email = "the.dude@domain.com"
}`

func TestAccUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckUserDestroy("artifactory_user.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccUser_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_user.foobar", "name", "the.dude"),
					resource.TestCheckResourceAttr("artifactory_user.foobar", "email", "the.dude@domain.com"),
					resource.TestCheckResourceAttr("artifactory_user.foobar", "is_admin", "false"),
					resource.TestCheckResourceAttr("artifactory_user.foobar", "is_editable", "true"),
				),
			},
		},
	})
}

const testAccUser_full = `
resource "artifactory_user" "foobar" {
	name        = "walter"
    email       = "walter.sobchak@domain.com"
    is_admin    = true
    is_editable = true
    groups      = [ "readers", "developers" ]
}`

func TestAccUser_full(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckUserDestroy("artifactory_user.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccUser_full,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_user.foobar", "name", "walter"),
					resource.TestCheckResourceAttr("artifactory_user.foobar", "email", "walter.sobchak@domain.com"),
					resource.TestCheckResourceAttr("artifactory_user.foobar", "is_admin", "true"),
					resource.TestCheckResourceAttr("artifactory_user.foobar", "is_editable", "true"),
					resource.TestCheckResourceAttr("artifactory_user.foobar", "groups.#", "2"),
				),
			},
		},
	})
}

func TestAccUser_import(t *testing.T) {
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

func testAccCheckUserDestroy(id string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(Client)
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found %s", id)
		}

		_, err := client.GetUser(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("User %s still exists", rs.Primary.ID)
		}

		return nil
	}
}
