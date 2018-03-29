# terraform-provider-artifactory

Manage [Artifactory](http://jfrog.io) with Terraform

[![CircleCI](https://circleci.com/gh/webdevwilson/terraform-provider-artifactory.svg?style=svg)](https://circleci.com/gh/webdevwilson/terraform-provider-artifactory)

## Installation via Homebrew

```bash
brew tap drewsonne/tap
brew install terraform-provider-artifactory
tf-install-provider artifactory
```
See details at https://github.com/drewsonne/homebrew-tap/blob/master/terraform-provider-artifactory.rb

## Provider

```hcl
provider "artifactory" {
  username = "${var.artifactory_username}"
  password = "${var.artifactory_password}"
  url      = "${var.artifactory_url}"
}
```

* `username` - (Required) Your username used to connect to Artifactory. You can
  also set this via the environment variable. `ARTIFACTORY_USER`

* `password` - (Required) Your password or an API key used to connect to Artifactory. You can
  also set this via the environment variable. `ARTIFACTORY_PASSWORD`

* `url` - (Required) The url to your Artifactory instance. This will typically be
  everything in front of the /webapp of your web console login. For instance, Artifactory
  cloud users will have a url similar to `https://youraccountname.jfrog.io/youraccountname`. You can
  also set this via the environment variable. `ARTIFACTORY_URL`
  
## Resources

### artifactory\_group

Provides support for creating groups in Artifactory. 

**This resource requires Artifactory Pro v2.4.0 or later**.

#### Example Usage

```hcl
resource "artifactory_group" "developers" {
    name      = "developers"
    auto_join = true
}
```

#### Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group.
* `auto_join` - (Optional) Should new user's be automatically added to this group? Default `false`.
* `realm` - (Optional) The name of the realm associated with this group (e.g. ARTIFACTORY, CROWD).
* `realm_attributes` - (Optional) Realm attributes for use by LDAP.

---

### artifactory\_local_repository

Provides support for setting up local repositories in Artifactory.

#### Example Usage

```
# A Sample NPM registry
resource "artifactory_local_repository" "localnpm" {
    key                   = "npm-local"
    package_type          = "npm"
    repository_layout_ref = "npm-default"
    property_sets         = [ "artifactory" ]
}

# A Docker repository
resource "artifactory_local_repository" "docker" {
    key                   = "docker-local"
    package_type          = "docker"
    repository_layout_ref = "simple-default"
    docker_api_version    = "V2"
    property_sets         = [ "artifactory" ]
}
```

#### Argument Reference

The following arguments are supported:

* `key` - (Required) The key of the repository.
* `package_type` - (Optional) The type of the repository. One of (`maven`, `gradle`, `ivy`, `sbt`,
`nuget`, `gems`, `npm`, `bower`, `debian`, `composer`, `pypi`, `docker`, `vagrant`, `gitlfs`, `yum`, 
`conan`, or `generic`). Default is `generic`.
* `description` - (Optional) Description of the repository.
* `notes` - (Optional) Notes about the repository.
* `repo_layout_ref` - (Optional) The layout of the repository. The full
list of available layouts is available in the Artifactory UI. Defaults to `maven-2-default`.
* `includes_pattern` - (Optional) List of artifact patterns to include when evaluating artifact 
requests in the form of x/y/\**/z/*. When used, only artifacts matching one of the include 
patterns are served. Defaults to `**/*`.
* `excludes_pattern` - (Optional) List of artifact patterns to exclude when evaluating artifact 
requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
* `handle_releases` - (Optional) Defaults to `true`.
* `handle_snapshots` - (Optional) Defaults to `true`.
* `max_unique_snapshots` - (Optional)
* `max_unique_tags` - (Optional)	
* `debian_trivial_layout` - (Optional) When set, the repository will use the deprecated trivial layout. 
Defaults to `false`.
* `checksum_policy_type` - (Optional) One of `client-checksums`, or `server-generated-checksums`. Defaults to `client-checksums`.
* `snapshot_version_behavior` - (Optional) One of `unique`, `non-unique`, or `deployer`. Defaults to `non-unique`.
* `suppress_pom_consistency_checks` - (Optional) Defaults to `false`.
* `blacked_out` - (Optional) When set, the repository does not participate in artifact resolution and 
new artifacts cannot be deployed. Defaults to `false`.
* `property_sets` - (Optional) List of property sets to apply to the repository.
* `archive_browsing_enabled` - (Optional) When set, you may view content such as HTML or Javadoc 
files directly from Artifactory. This may not be safe and therefore requires strict content 
moderation to prevent malicious users from uploading content that may compromise 
security (e.g., cross-site scripting attacks). Defaults to `false`.
* `calculate_yum_metadata` - (Optional) Defaults to `false`.
* `yum_root_depth` - (Optional) Defaults to `0`.
* `docker_api_version` - (Optional) Docker API compatibility. Must be `V1` or `V2`. Defaults to `V2`.

---

### artifactory\_remote_repository

Provides support for setting up remote repositories in Artifactory.

A remote repository serves as a caching proxy for a repository managed at a remote URL 
(which may itself be another Artifactory remote repository).  

Artifacts are stored and updated in remote repositories according to various configuration 
parameters that control the caching and proxying behavior. You can remove artifacts 
from a remote repository cache but you cannot actually deploy a new artifact into a 
remote repository.


#### Example Usage

```hcl
# Manage your repository
resource "artifactory_remote_repository" "publicnpm" {
  key               = "registry.npmjs.org"
  package_type      = "npm"
  description       = "Proxy public npm registry"
  repo_layout_ref   = "npm-default"
  url               = "https://registry.npmjs.org/"
  property_sets = [
    "artifactory"
  ]
}
```

#### Argument Reference

The following arguments are supported:

* `key` - (Required) The key of the repository.
* `package_type` - (Optional) The type of the repository. One of (`maven`, `gradle`, `ivy`, `sbt`,
`nuget`, `gems`, `npm`, `bower`, `debian`, `composer`, `pypi`, `docker`, `vagrant`, `gitlfs`, `yum`, 
`conan`, or `generic`). Default is `generic`.
* `repositories` - (Optional) The upstream repositories to pull from.
* `default_deployment_repo` - (Optional) The local repo this repository will push to.
* `description` - (Optional) Description of the repository.
* `notes` - (Optional) Notes about the repository.
* `repo_layout_ref` - (Optional) The layout of the repository. The full
list of available layouts is available in the Artifactory UI. Defaults to `maven-2-default`.
* `includes_pattern` - (Optional) List of artifact patterns to include when evaluating artifact 
requests in the form of x/y/\**/z/*. When used, only artifacts matching one of the include 
patterns are served. Defaults to `**/*`.
* `excludes_pattern` - (Optional) List of artifact patterns to exclude when evaluating artifact 
requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
* `handle_releases` - (Optional) Defaults to `true`.
* `handle_snapshots` - (Optional) Defaults to `true`.
* `max_unique_snapshots` - (Optional)
* `suppress_pom_consistency_checks` - (Optional)
* `url` - (Optional) URL of the upstream remote repository.
* `username` - (Optional) The username to use to authenticate to the upstream remote repository.
* `password` - (Optional) The password to use to authenticate to the upstream remote repository.
* `proxy` - (Optional)
* `remote_repo_checksum_policy_type` - (Optional)
* `hard_fail` - (Optional)
* `offline` - (Optional) If set, Artifactory does not try to fetch remote artifacts. 
Only locally-cached artifacts are retrieved. Defaults to `false`.
* `blacked_out` - (Optional) When set, the repository does not participate in artifact resolution and 
new artifacts cannot be deployed. Defaults to `false`.
* `property_sets` - (Optional) List of property sets to apply to the repository.
* `store_artifacts_locally` - (Optional) When set, the repository should store cached artifacts 
locally. When not set, artifacts are not stored locally, and direct repository-to-client streaming 
is used. This can be useful for multi-server setups over a high-speed LAN, with one Artifactory 
caching certain data on central storage, and streaming it directly to satellite pass-though 
Artifactory servers. Defaults to `true`.
* `socket_timeout_millis` - (Optional) Network timeout (in ms) to use when establishing a connection 
and for unanswered requests. Timing out on a network operation is considered a retrieval failure. 
Defaults to `15000`.
* `local_address` - (Optional) The local address to be used when creating connections. Useful for 
specifying the interface to use on systems with multiple network interfaces.
* `retrieval_cache_period_seconds` - (Optional) This value refers to the number of seconds to 
cache metadata files before checking for newer versions on remote server. A value of 0 indicates 
no caching. Defaults to `600`.
* `failed_cache_period_seconds` - (Optional) Defaults to `0`.
* `missed_cache_period_seconds` - (Optional) The number of seconds to cache artifact retrieval misses (artifact not found).     
A value of 0 indicates no caching. Defaults to `1800`.
* `unused_artifacts_cleanup_enabled` - (Optional) Defaults to `false`.
* `unused_artifacts_cleanup_period_hours` - (Optional) The number of hours to wait before an 
artifact is deemed "unused" and eligible for cleanup from the repository.
A value of `0` means automatic cleanup of cached artifacts is disabled. Defaults to `0`.
* `fetch_jars_eagerly` - (Optional) Defaults to `false`.
* `fetch_sources_eagerly` - (Optional) Defaults to `false`.
* `share_configuration` - (Optional) Defaults to `false`.
* `synchronize_properties` - (Optional) When set, remote artifacts are fetched along with their 
properties. Defaults to `false`.
* `allow_any_host_auth` - (Optional) Defaults to `false`.
* `enable_cookie_management` - (Optional) Enables cookie management if the remote repository uses 
cookies to manage client state. Defaults to `false`.
* `bower_registry_url` - (Optional)
* `vcs_type` - (Optional) When present, must be `GIT`.
* `vcs_git_provider` - (Optional) Should be one of `GITHUB`, `BITBUCKET`,
`STASH`, `ARTIFACTORY`, `CUSTOM`. Defaults to `GITHUB`.
* `vcs_git_download_url` - (Optional)
			
---

### artifactory\_user

Provides support for creating users in Artifactory. 

**This resource does not allow setting of the user's password**. 

Instead, a random password is generated for each user. The user should do a 
_forgot my password_ to reset their password. On updates, the password is set, 
then immediately expired. This should trigger an email to the user if Artifactory is configured to.

#### Example Usage

```hcl
resource "artifactory_user" "walter" {
    name     = "walter.sobchak"
    email    = "walter.sobchak@sobchaksecurity.com"
    is_admin = true
    groups   = [ "readers", "publishers" ]
}
```

#### Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user.
* `email` - (Required) The user's email address.
* `is_admin` - (Optional) Does this user have admin privileges. Default `false`.
* `is_updatable` - (Optional) Can this user update their profile?. Cannot be `false` 
when is_admin is set to `true`. Default `true`.
* `groups` - (Optional) An array of groups this user belongs to.
* `realm` - (Computed) The realm the user belongs to.

---

### artifactory\_virtual_repository

Provides support for setting up virtual repositories in Artifactory.

Virtual repositories are used to mirror public repositories, while adding
push capabilities. Configure a virtual repository to pull from multiple repositories
and push to a single local repository.

#### Example Usage

```hcl
# Create a virtual repository that mirrors public, and pushes to a local
resource "artifactory_virtual_repository" "gems" {
    key                     = "gems"
    package_type            = "gem"
    description             = "Gem Repository"
    repositories            = [
      "org-rubygems-cache",
      "gems-local"
    ]
    repo_layout_ref         = "simple-default"
    default_deployment_repo = "gems-local"
}
```

#### Argument Reference

The following arguments are supported:

* `key` - (Required) The key of the repository.
* `package_type` - (Optional) The type of the repository. One of (`maven`, `gradle`, `ivy`, `sbt`,
`nuget`, `gems`, `npm`, `bower`, `debian`, `composer`, `pypi`, `docker`, `vagrant`, `gitlfs`, `yum`, 
`conan`, or `generic`). Default is `generic`.
* `repositories` - (Optional) The upstream repositories to pull from.
* `default_deployment_repo` - (Optional) The local repo this repository will push to.
* `description` - (Optional) Description of the repository.
* `notes` - (Optional) Notes about the repository.
* `repo_layout_ref` - (Optional) The layout of the repository. The full
list of available layouts is available in the Artifactory UI. Defaults to `maven-2-default`.
* `includes_pattern` - (Optional) List of artifact patterns to include when evaluating artifact 
requests in the form of x/y/\**/z/*. When used, only artifacts matching one of the include 
patterns are served. Defaults to `**/*`.
* `excludes_pattern` - (Optional) List of artifact patterns to exclude when evaluating artifact 
requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
* `debian_trivial_layout` - (Optional) When set, the repository will use the deprecated trivial layout. 
Defaults to `false`.
* `artifactory_requests_can_retrieve_remote_artifacts` - (Optional) Determines whether artifact 
requests coming from other instance of Artifactory can be fulfilled by accessing this virtual 
repository's remote repositories, or by only accessing its caches (default). Defaults to `false`.
* `key_pair` - (Optional)
* `pom_repository_references_cleanup_policy` - (Optional) Should be one of `discard_active_reference`, 
`discard_any_reference`, `nothing`.


