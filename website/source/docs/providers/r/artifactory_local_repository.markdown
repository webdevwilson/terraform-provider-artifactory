---
layout: "artifactory"
page_title: "Artifactory: artifactory_local_repository"
sidebar_current: "docs-artifactory-local-repository"
description: |-
  Provides support for setting up local repositories in Artifactory
---

# artifactory\_local_repository

Provides support for setting up local repositories in Artifactory.

## Example Usage

```
# A Sample NPM registry
resource "artifactory_local_repository" "localnpm" {
    key                   = "npm-local"
    package_type          = "npm"
    repository_layout_ref = "npm-default"
    property_sets         = [ "artifactory" ]
}

# A Docker repository
resource "artifact_local_repository" "docker" {
    key                   = "docker-local"
    package_type          = "docker"
    repository_layout_ref = "simple-default"
    docker_api_version    = "V2"
    property_sets         = [ "artifactory" ]
}
```

## Argument Reference

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