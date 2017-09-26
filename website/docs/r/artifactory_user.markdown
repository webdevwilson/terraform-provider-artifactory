---
layout: "artifactory"
page_title: "Artifactory: artifactory_user"
sidebar_current: "docs-artifactory-user"
description: |-
  Provides support for creating users in Artifactory
---

# artifactory\_user

Provides support for creating users in Artifactory. 

**This resource does not allow setting of the user's password**. 

Instead, a random password is generated for each user. The user should do a 
_forgot my password_ to reset their password. On updates, the password is set, 
then immediately expired. This should trigger an email to the user if Artifactory is configured to.

## Example Usage

```
resource "artifact_user" "walter" {
    name     = "walter.sobchak"
    email    = "walter.sobchak@sobchaksecurity.com"
    is_admin = true
    groups   = [ "readers", "publishers" ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user.
* `email` - (Required) The user's email address.
* `is_admin` - (Optional) Does this user have admin privileges. Default `false`.
* `is_updatable` - (Optional) Can this user update their profile?. Cannot be `false` 
when is_admin is set to `true`. Default `true`.
* `groups` - (Optional) An array of groups this user belongs to.
* `realm` - (Computed) The realm the user belongs to.