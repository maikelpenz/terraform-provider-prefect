---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "prefect_project Resource - terraform-provider-prefect"
subcategory: ""
description: |-
  Project resource.
---

# prefect_project (Resource)

Project resource.

## Example Usage

```terraform
resource "prefect_project" "test" {
  name        = "test"
  description = "my test project"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `description` (String) Description

### Read-Only

- `id` (String) Server generated UUID

