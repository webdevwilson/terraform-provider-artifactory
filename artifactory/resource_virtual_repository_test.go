package artifactory

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccVirtualRepository_basic = `
resource "artifactory_virtual_repository" "foobar" {
    key          = "acctest-virtual-basic"
}`

func TestAccVirtualRepository_basic(t *testing.T) {
	resourceName := "artifactory_virtual_repository.foobar"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy(resourceName),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVirtualRepository_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "key", "acctest-virtual-basic"),
				),
			},
		},
	})
}

const testAccVirtualRepository_full = `
resource "artifactory_remote_repository" "npm_public" {
  key               = "acctest-virtual-npm-public"
  package_type      = "npm"
  description       = "Proxy public npm registry"
  repo_layout_ref   = "npm-default"
  url               = "https://registry.npmjs.org/"
  property_sets = [
    "artifactory"
  ]
}

resource "artifactory_local_repository" "npm_private" {
    key              = "acctest-virtual-npm-private"
    package_type     = "npm"
    repo_layout_ref = "npm-default"
}

resource "artifactory_virtual_repository" "npm" {
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
    default_deployment_repo                            = "${artifactory_local_repository.npm_private.key}"
    debian_trivial_layout                              = false
    repositories                                       = [
        "${artifactory_remote_repository.npm_public.key}",
        "${artifactory_local_repository.npm_private.key}"
    ]
}`

func TestAccVirtualRepository_full(t *testing.T) {
	resourceName := "artifactory_virtual_repository.npm"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy(resourceName),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVirtualRepository_full,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "key", "acctest-virtual-full"),
					resource.TestCheckResourceAttr(resourceName, "package_type", "npm"),
					resource.TestCheckResourceAttr(resourceName, "description", "desc"),
					resource.TestCheckResourceAttr(resourceName, "notes", "the notes"),
					resource.TestCheckResourceAttr(resourceName, "includes_pattern", "**/*"),
					resource.TestCheckResourceAttr(resourceName, "excludes_pattern", "**/*.tgz"),
					resource.TestCheckResourceAttr(resourceName, "repo_layout_ref", "npm-default"),
					resource.TestCheckResourceAttr(resourceName, "artifactory_requests_can_retrieve_remote_artifacts", "false"),
					resource.TestCheckResourceAttr(resourceName, "key_pair", "keypair"),
					resource.TestCheckResourceAttr(resourceName, "pom_repository_references_cleanup_policy", "discard_any_reference"),
					resource.TestCheckResourceAttr(resourceName, "default_deployment_repo", "acctest-virtual-npm-private"),
					resource.TestCheckResourceAttr(resourceName, "repositories.#", "2"),
				),
			},
		},
	})
}
