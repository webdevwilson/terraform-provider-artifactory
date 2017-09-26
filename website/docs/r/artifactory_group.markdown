---
layout: "artifactory"
page_title: "Artifactory: artifactory_group"
sidebar_current: "docs-artifactory-group"
description: |-
  Provides support for creating groups in Artifactory
---

# artifactory\_group

Provides support for creating groups in Artifactory. 

**This resource requires Artifactory Pro v2.4.0 or later**.

## Example Usage

```
resource "artifact_group" "developers" {
    name      = "developers"
    auto_join = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group.
* `auto_join` - (Optional) Should new user's be automatically added to this group? Default `false`.
* `realm` - (Optional) The name of the realm associated with this group (e.g. ARTIFACTORY, CROWD).
* `realm_attributes` - (Optional) Realm attributes for use by LDAP.