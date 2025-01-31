
# Commands
tf init
tf plan
tf apply
tf state list
tf plan -generate-config-out='generated.tf'

# Structure
Normal to order folders by project. Note not recursive so dev/subfolder would be ignored and need a new init to make a new root.
Reusability comes from using modules where required.
Imports to bring existing infra to TF state.

project/
  ├── modules/
  │   └── common-interceptors/
  │       ├── main.tf
  │       ├── variables.tf
  │       └── outputs.tf
  ├── dev/
  │   ├── provider.tf
  │   ├── variables.tf
  │   └── main.tf
  └── prod/
      ├── provider.tf
      ├── variables.tf
      └── main.tf

`tf init` within the project dev or prod.

Always need provider, have one in each project.

# Multi-provider
Alternative is multi. In this mode you must add provider property per resource.

```h
# per resource
resource "foo" "bar" {
  provider = conduktor.console # or conduktor.gateway
  # ...
}
```

```h
provider "conduktor" {
  alias    = "console"
  mode     = "console"
  base_url = "http://localhost:8080"

  api_token = "your-api-token"
  #admin_user     = "admin@my-org.com"
  #admin_password = "admin-password"

  insecure = true
}

provider "conduktor" {
  alias    = "gateway"
  mode     = "gateway"
  base_url = "http://localhost:8888"

  admin_user     = "admin"
  admin_password = "admin-password"

  insecure = true
}
```

# Imports

Need an import statement and the resource to send it to.

```
import {
  to = conduktor_gateway_interceptor_v2.mask-sensitive-fields
  id = "mask-sensitive-fields"
}
```
This will need to be already defined, matching what you're importing. With this approach (pre 1.5) you have to best guess then check the diff with another `tf plan` after the import to see the difference. 
Alternative you can get it to generate it to a file, then copy paste that into your main.tf or other file.
`tf plan -generate-config-out='generated.tf'`
[Official docs](https://developer.hashicorp.com/terraform/language/import/generating-configuration), experimental still (as of 1.10) so they can make changes to it.

See [import.tf](stu-wd/interceptors/37A-json-block/project1/imports/import.tf) for pre-filled ready for import.
See [main.tf](stu-wd/interceptors/37A-json-block/project1/imports/imports-generated/main.tf) for generated content that was pasted into file. Both the copy paste and the auto-gen.