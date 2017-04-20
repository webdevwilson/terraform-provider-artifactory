package artifactory

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccVirtualRepository_basic = `
resource "artifactory_virtual_repository" "foobar" {
	key 	     = "foobar-test"
}`

func TestAccVirtualRepository_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_virtual_repository.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVirtualRepository_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "key", "foobar-test"),
				),
			},
		},
	})
}

const testAccVirtualRepository_full = `
resource "artifactory_virtual_repository" "foobar" {
    key                                                = "foobar-test"
    package_type                                       = "npm"
	description                                        = "desc"
	notes                                              = "the notes"
	includes_pattern                                   = "**/*"
	excludes_pattern                                   = "**/*.tgz"
    debian_trivial_layout                              = false
    repo_layout_ref                                    = "npm-default"
    artifactory_requests_can_retrieve_remote_artifacts = false
	key_pair                                           = "keypair"
    pom_repository_references_cleanup_policy           = "discard_any_reference"
    default_deployment_repo                            = "npm-local"
	debian_trivial_layout                              = false
    repositories                                       = ["registry.npmjs.org"]
}`

func TestAccVirtualRepository_full(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_virtual_repository.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVirtualRepository_full,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "key", "foobar-test"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "package_type", "npm"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "description", "desc"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "notes", "the notes"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "includes_pattern", "**/*"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "excludes_pattern", "**/*.tgz"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "repo_layout_ref", "npm-default"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "artifactory_requests_can_retrieve_remote_artifacts", "false"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "key_pair", "keypair"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "pom_repository_references_cleanup_policy", "discard_any_reference"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "default_deployment_repo", "npm-local"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "repositories.#", "1"),
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "repositories.619022263", "registry.npmjs.org"),
				),
			},
		},
	})
}

func TestAccVirtualRepository_import(t *testing.T) {
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
