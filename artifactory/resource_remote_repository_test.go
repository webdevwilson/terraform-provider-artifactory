package artifactory

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccRemoteRepository_basic = `
resource "artifactory_remote_repository" "foobar" {
	key = "foobar-test"
    url = "https://central.maven.org"
}`

func TestAccRemoteRepository_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_remote_repository.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRemoteRepository_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "key", "foobar-test"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "url", "https://central.maven.org"),
				),
			},
		},
	})
}

const testAccRemoteRepository_full = `
resource "artifactory_remote_repository" "foobar" {
	key                                   = "foobar-test"
	package_type                          = "npm"
	url                                   = "https://registry.npmjs.org/"
	username                              = "user"
	password                              = "pass"
    proxy                                 = ""
	description                           = "desc"
	notes                                 = "notes"
	includes_pattern                      = "**/*.js"
	excludes_pattern                      = "**/*.jsx"
	repo_layout_ref                       = "npm-default"
	remote_repo_checksum_policy_type      = ""
	handle_releases                       = true
	handle_snapshots                      = true
	max_unique_snapshots                  = 15
	suppress_pom_consistency_checks       = true
	hard_fail                             = true
	offline                               = true
	blacked_out                           = false
	store_artifacts_locally               = true
	socket_timeout_millis                 = 25000
	local_address                         = ""
	retrieval_cache_period_seconds        = 15
	failed_cache_period_seconds           = 0
	missed_cache_period_seconds           = 2500
	unused_artifacts_cleanup_enabled      = false
	unused_artifacts_cleanup_period_hours = 96
	fetch_jars_eagerly                    = true
    fetch_sources_eagerly                 = true
	share_configuration                   = true
	synchronize_properties                = true
	property_sets                         = ["artifactory"]
	allow_any_host_auth                   = false
	enable_cookie_management              = true
	bower_registry_url                    = ""
	vcs_type                              = ""
	vcs_git_provider                      = ""
	vcs_git_download_url                  = ""
}`

func TestAccRemoteRepository_full(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRepositoryDestroy("artifactory_remote_repository.foobar"),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRemoteRepository_full,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "key", "foobar-test"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "package_type", "npm"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "url", "https://registry.npmjs.org/"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "username", "user"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "password", "pass"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "proxy", ""),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "description", "desc (local file cache)"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "notes", "notes"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "includes_pattern", "**/*.js"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "excludes_pattern", "**/*.jsx"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "repo_layout_ref", "npm-default"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "remote_repo_checksum_policy_type", ""),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "handle_releases", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "handle_snapshots", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "max_unique_snapshots", "15"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "suppress_pom_consistency_checks", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "hard_fail", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "offline", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "blacked_out", "false"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "store_artifacts_locally", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "socket_timeout_millis", "25000"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "local_address", ""),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "retrieval_cache_period_seconds", "15"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "failed_cache_period_seconds", "0"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "missed_cache_period_seconds", "2500"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "unused_artifacts_cleanup_enabled", "false"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "unused_artifacts_cleanup_period_hours", "96"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "fetch_jars_eagerly", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "fetch_sources_eagerly", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "share_configuration", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "synchronize_properties", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "property_sets.#", "1"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "property_sets.214975871", "artifactory"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "allow_any_host_auth", "false"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "enable_cookie_management", "true"),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "bower_registry_url", ""),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "vcs_type", ""),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "vcs_git_provider", ""),
					resource.TestCheckResourceAttr("artifactory_remote_repository.foobar", "vcs_git_download_url", ""),
				),
			},
		},
	})
}

func TestAccRemoteRepository_import(t *testing.T) {
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
