package artifactory

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccVirtualRepository_basic = `
resource "artifactory_virtual_repository" "foobar" {
	key 	     = "acctest-virtual-basic"
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
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "key", "acctest-virtual-basic"),
				),
			},
		},
	})
}

const testAccVirtualRepository_full = `
resource "artifactory_remote_repository" "npm" {
	key = "registry.npmjs.org"
    url = "https://registry.npmjs.org"
}

resource "artifactory_local_repository" "npm-local" {
	key 	     = "acctest-npm-local"
	package_type = "npm"
}

resource "artifactory_virtual_repository" "foobar" {
    key                                                = "acctest-virtual-full"
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
    default_deployment_repo                            = "acctest-npm-local"
	debian_trivial_layout                              = false
    repositories                                       = ["registry.npmjs.org"]
	depends_on										   = [
		"artifactory_remote_repository.npm",
		"artifactory_local_repository.npm-local"
	]
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
					resource.TestCheckResourceAttr("artifactory_virtual_repository.foobar", "key", "acctest-virtual-full"),
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
