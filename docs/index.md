---
page_title: "Provider: Conduktor"
subcategory: ""
description: |-
    The Conduktor provider is used to interact with the resources supported by Conduktor. The provider needs to be configured with the proper credentials before it can be used.
---

# Conduktor Provider

The Conduktor provider is used to interact with the resources supported by Conduktor. The provider needs to be configured with the proper credentials before it can be used.

> [!WARNING]
> - The Conduktor Terraform provider is currently in **Alpha**.
> - It does not support all Console and Gateway resources yet. See our [resources roadmap](https://github.com/conduktor/terraform-provider-conduktor/blob/main/README.md#resources-roadmap).
> - Let us know if you have [feedback](https://product.conduktor.help/c/74-terraform-provider) or wish to be a design partner.

## Example Usage

### Console client only

```terraform
provider "conduktor" {
  mode = "console"
  # mandatory console URL
  base_url = "http://localhost:8080" # or env vars CDK_CONSOLE_BASE_URL or CDK_BASE_URL

  # authentication either with api token or admin credentials
  api_token = "your-api-token" # or env var CDK_API_TOKEN or CDK_API_KEY
  #admin_user     = "admin@my-org.com" # or env var CDK_CONSOLE_USER or CDK_ADMIN_EMAIL or CDK_ADMIN_USER
  #admin_password = "admin-password"   # or env var CDK_CONSOLE_PASSWORD or CDK_ADMIN_PASSWORD

  # optional http client TLS configuration
  cert     = file("path/to/cert.pem") # or env var CDK_CONSOLE_CERT or CDK_CERT
  insecure = true                     # or env var CDK_CONSOLE_INSECURE or CDK_INSECURE

  # optional authentication via certificate
  key    = file("path/to/key.pem") # or env var CDK_CONSOLE_KEY or CDK_KEY
  cacert = file("path/to/ca.pem")  # or env var CDK_CONSOLE_CA_CERT CDK_CA_CERT
}
```

### Gateway client only

```terraform
provider "conduktor" {
  mode = "gateway"
  # mandatory gateway URL
  base_url = "http://localhost:8888" # or env vars CDK_GATEWAY_BASE_URL or CDK_BASE_URL

  # authentication with admin credentials
  admin_user     = "admin"          # or env var CDK_GATEWAY_USER or CDK_ADMIN_USER
  admin_password = "admin-password" # or env var CDK_GATEWAY_PASSWORD or CDK_ADMIN_PASSWORD

  # optional http client TLS configuration
  cert     = file("path/to/cert.pem") # or env var CDK_GATEWAY_CERT or CDK_CERT
  insecure = true                     # or env var CDK_GATEWAY_INSECURE or CDK_INSECURE

  # optional authentication via certificate
  key    = file("path/to/key.pem") # or env var CDK_GATEWAY_KEY or CDK_KEY
  cacert = file("path/to/ca.pem")  # or env var CDK_GATEWAY_CACERT or CDK_CACERT
}
```

### Multi client configuration using [terraform alias](https://developer.hashicorp.com/terraform/language/providers/configuration#alias-multiple-provider-configurations)

```terraform
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

# And how to use them with example resources
resource "conduktor_console_user_v2" "user" {
  provider = conduktor.console
  # ...
}

resource "conduktor_gateway_service_account_v2" "gateway_sa" {
  provider = conduktor.gateway
  # ...
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `mode` (String) The mode for the Terraform provider. When using one provider, can be set to either `console` or `gateway`. Alternatively for multi-provider configuration, set both providers with an an alias, and set the relevant alias at the resource level. May also be set using environment variable `CDK_PROVIDER_MODE`.  
See [documentation](https://github.com/conduktor/terraform-provider-conduktor/blob/main/docs/index.md#multi-client-configuration-using-terraform-alias) for more information.

### Optional

- `admin_password` (String, Sensitive) The password of the admin user. May be set using environment variable `CDK_CONSOLE_PASSWORD` or `CDK_ADMIN_PASSWORD` for Console, `CDK_GATEWAY_PASSWORD` or `CDK_ADMIN_PASSWORD` for Gateway. Required if admin_user is set. If not provided, the API token will be used to authenticater.
- `admin_user` (String) The login credentials of the admin user. May be set using environment variable `CDK_CONSOLE_USER`, `CDK_ADMIN_EMAIL` or `CDK_ADMIN_USER` for Console, `CDK_GATEWAY_USER` or `CDK_ADMIN_USER` for Gateway. Required if admin_password is set. If not provided and `mode` is Console, the API token will be used to authenticate.
- `api_token` (String, Sensitive) The API token to authenticate with the Conduktor Console API. May be set using environment variable `CDK_API_TOKEN` or `CDK_API_KEY`. If not provided, admin_user and admin_password will be used to authenticate. See [documentation](https://docs.conduktor.io/platform/reference/api-reference/#generate-an-api-key) for more information. Not used if `mode` is Gateway.
- `base_url` (String) The URL of either Conduktor Console or Gateway, depending on the `mode`. May be set using environment variable `CDK_CONSOLE_BASE_URL` or `CDK_BASE_URL` for Console, `CDK_GATEWAY_BASE_URL` or `CDK_BASE_URL` for Gateway. Required either here or in the environment.
- `cacert` (String) Root CA certificate in PEM format to verify the Conduktor certificate. May be set using environment variable `CDK_CONSOLE_CACERT` or `CDK_CACERT` for Console, `CDK_GATEWAY_CACERT` or `CDK_CACERT` for Gateway. If not provided, the system's root CA certificates will be used.
- `cert` (String) Cert in PEM format to authenticate using client certificates. May be set using environment variable `CDK_CONSOLE_CERT` or `CDK_CERT` for Console, `CDK_GATEWAY_CERT` or `CDK_CERT` for Gateway. Must be used with key. If key is provided, cert is required. Useful when Console is behind a reverse proxy with client certificate authentication.
- `insecure` (Boolean) Skip TLS verification flag. May be set using environment variable `CDK_CONSOLE_INSECURE` or `CDK_INSECURE` for Console, `CDK_GATEWAY_INSECURE` or `CDK_INSECURE` for Gateway.
- `key` (String) Key in PEM format to authenticate using client certificates. May be set using environment variable `CDK_CONSOLE_KEY` or `CDK_KEY` for Console, `CDK_GATEWAY_KEY` or `CDK_KEY` for Gateway. Must be used with cert. If cert is provided, key is required. Useful when Console is behind a reverse proxy with client certificate authentication.


