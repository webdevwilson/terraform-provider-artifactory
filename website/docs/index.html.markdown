---
layout: "artifactory"
page_title: "Provider: Artifactory"
sidebar_current: "docs-artifactory-index"
description: |-
  The Artifactory provider is used to manage Artifactory repositories.
---

# Artifactory Provider

[Artifactory](https://www.jfrog.com/artifactory/) is an artifact repository manager.
The artifactory provider allows you to manage your repository and security configurations in Artifactory.

**Note: this provider requires Artifactory Pro v2.3.0 or later**.

## Example Usage

```
# Configure the Artifactory provider
provider "artifactory" {
  username = "${var.artifactory_username}"
  password = "${var.artifactory_password}"
  url      = "${var.artifactory_url}"
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `username` - (Required) Your username used to connect to Artifactory. You can
  also set this via the environment variable. `ARTIFACTORY_USER`

* `password` - (Required) Your password or an API key used to connect to Artifactory. You can
  also set this via the environment variable. `ARTIFACTORY_PASSWORD`

* `url` - (Required) The url to your Artifactory instance. This will typically be
  everything in front of the /webapp of your web console login. For instance, Artifactory
  cloud users will have a url similar to `https://youraccountname.jfrog.io/youraccountname`. You can
  also set this via the environment variable. `ARTIFACTORY_URL`