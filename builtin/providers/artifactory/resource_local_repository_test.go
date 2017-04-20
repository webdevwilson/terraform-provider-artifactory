package artifactory

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccLocalRepository_basic = `
resource "artifactory_local_repository" "foobar" {
	key 	     = "foobar-test"
	package_type = "docker"
}`

func TestAccLocalRepository_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_local_repository.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLocalRepository_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "key", "foobar-test"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "package_type", "docker"),
				),
			},
		},
	})
}

const testAccLocalRepository_full = `
resource "artifactory_local_repository" "foobar" {
    key                             = "foobar-test"
    package_type                    = "npm"
	description                     = "desc"
	notes                           = "the notes"
	includes_pattern                = "**/*"
	excludes_pattern                = "**/*.tgz"
	repo_layout_ref                 = "npm-default"
	handle_releases                 = true
	handle_snapshots                = true
	max_unique_snapshots            = 25
	debian_trivial_layout           = false
	checksum_policy_type            = "client-checksums"
	max_unique_tags                 = 100
	snapshot_version_behavior       = "unique"
	suppress_pom_consistency_checks = true
	blacked_out                     = false
	property_sets                   = [ "artifactory" ]
	archive_browsing_enabled        = false
	calculate_yum_metadata          = false
	yum_root_depth                  = 0
	docker_api_version              = "V2"    
}`

func TestAccLocalRepository_full(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_local_repository.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLocalRepository_full,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "key", "foobar-test"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "package_type", "npm"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "description", "desc"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "notes", "the notes"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "includes_pattern", "**/*"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "excludes_pattern", "**/*.tgz"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "repo_layout_ref", "npm-default"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "handle_releases", "true"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "handle_snapshots", "true"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "max_unique_snapshots", "25"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "debian_trivial_layout", "false"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "checksum_policy_type", "client-checksums"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "max_unique_tags", "100"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "snapshot_version_behavior", "unique"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "suppress_pom_consistency_checks", "true"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "blacked_out", "false"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "property_sets.#", "1"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "property_sets.214975871", "artifactory"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "archive_browsing_enabled", "false"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "calculate_yum_metadata", "false"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "yum_root_depth", "0"),
					resource.TestCheckResourceAttr("artifactory_local_repository.foobar", "docker_api_version", "V2"),
				),
			},
		},
	})
}

func TestAccLocalRepository_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_local_repository.foobar"),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLocalRepository_basic,
			},
			resource.TestStep{
				ResourceName:      "artifactory_local_repository.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
