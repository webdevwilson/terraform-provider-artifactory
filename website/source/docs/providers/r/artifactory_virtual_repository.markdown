---
layout: "artifactory"
page_title: "Artifactory: artifactory_virtual_repository"
sidebar_current: "docs-artifactory-virtual-repository"
description: |-
  Provides support for setting up virtual repositories in Artifactory
---

# artifactory\_virtual_repository

Provides support for setting up virtual repositories in Artifactory.

Virtual repositories are used to mirror public repositories, while adding
push capabilities. Configure a virtual repository to pull from multiple repositories
and push to a single local repository.

## Example Usage

```
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
* `debian_trivial_layout` - (Optional) When set, the repository will use the deprecated trivial layout. 
Defaults to `false`.
* `artifactory_requests_can_retrieve_remote_artifacts` - (Optional) Determines whether artifact 
requests coming from other instance of Artifactory can be fulfilled by accessing this virtual 
repository's remote repositories, or by only accessing its caches (default). Defaults to `false`.
* `key_pair` - (Optional)
* `pom_repository_references_cleanup_policy` - (Optional) Should be one of `discard_active_reference`, 
`discard_any_reference`, `nothing`.

