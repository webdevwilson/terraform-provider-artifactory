resource "artifactory_remote_repository" "npm_public" {
  key             = "registry.npmjs.org"
  package_type    = "npm"
  description     = "Proxy public npm registry"
  repo_layout_ref = "npm-default"
  url             = "https://registry.npmjs.org/"

  property_sets = [
    "artifactory",
  ]
}

resource "artifactory_local_repository" "npm_private" {
  key             = "npm-local"
  package_type    = "npm"
  description     = "Private npm registry"
  repo_layout_ref = "npm-default"
}

resource "artifactory_virtual_repository" "npm" {
  key                                                = "npm"
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

  repositories = [
    "${artifactory_remote_repository.npm_public.key}",
    "${artifactory_local_repository.npm_private.key}",
  ]
}
