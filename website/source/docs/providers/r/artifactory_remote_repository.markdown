---
layout: "artifactory"
page_title: "Artifactory: artifactory_remote_repository"
sidebar_current: "docs-artifactory-remote-repository"
description: |-
  Provides support for setting up remote repositories in Artifactory
---

# artifactory\_remote_repository

Provides support for setting up remote repositories in Artifactory.

A remote repository serves as a caching proxy for a repository managed at a remote URL 
(which may itself be another Artifactory remote repository).  

Artifacts are stored and updated in remote repositories according to various configuration 
parameters that control the caching and proxying behavior. You can remove artifacts 
from a remote repository cache but you cannot actually deploy a new artifact into a 
remote repository.


## Example Usage

```
# Manage your repository
resource "artifactory_remote_repository" "publicnpm" {
  key               = "registry.npmjs.org"
  packageType       = "npm"
  description       = "Proxy public npm registry"
  repoLayoutRef     = "npm-default"
  url               = "https://registry.npmjs.org/"
  property_sets = [
    "artifactory"
  ]
}
```

## Argument Reference

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
			