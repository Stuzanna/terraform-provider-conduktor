terraform {
  required_providers {
    conduktor = {
      source  = "terraform.local/conduktor/conduktor" # local provider
      version = "= 0.0.11"                            # latest version found locally in the plugin cache.
    }
  }
}

# GATEWAY
# provider "conduktor" {
#   mode = "gateway"
#   # mandatory gateway URL
#   base_url = "http://localhost:8888" # or env vars CDK_GATEWAY_BASE_URL or CDK_BASE_URL

#   # authentication with admin credentials
#   admin_user     = "admin"          # or env var CDK_GATEWAY_USER or CDK_ADMIN_USER
#   admin_password = "admin-password" # or env var CDK_GATEWAY_PASSWORD or CDK_ADMIN_PASSWORD

#   # optional http client TLS configuration
# #   cert     = file("path/to/cert.pem") # or env var CDK_GATEWAY_CERT or CDK_CERT
#   insecure = true                     # or env var CDK_GATEWAY_INSECURE or CDK_INSECURE

#   # optional authentication via certificate
# #   key    = file("path/to/key.pem") # or env var CDK_GATEWAY_KEY or CDK_KEY
# #   cacert = file("path/to/ca.pem")  # or env var CDK_GATEWAY_CACERT or CDK_CACERT
# }

# CONSOLE ONLY
# provider "conduktor" {
#   mode     = "console"
#   base_url = "http://localhost:8080"
#   admin_user     = "admin@conduktor.io"
#   admin_password = "testP4ss!"
#   insecure = true
# }

# BOTH
provider "conduktor" {
  alias    = "console"
  mode     = "console"
  base_url = "http://localhost:8080"

  #   api_token = "your-api-token"
  admin_user     = "admin@conduktor.io"
  admin_password = "testP4ss!"

  insecure = true
}

provider "conduktor" {
  alias    = "gateway"
  mode     = "gateway"
  base_url = "http://localhost:8888"

  admin_user     = "admin"
  admin_password = "conduktor"

  insecure = true
}

# ORIGINAL BLOCK
# provider "conduktor" {
#   # mandatory console URL
#   console_url = "http://localhost:8080" # or env vars CDK_CONSOLE_URL or CDK_BASE_URL

#   # authentication either with api token or admin credentials
# #   api_token = "your-api-token" # or env var CDK_API_TOKEN or CDK_API_KEY
#   admin_user    = "admin@conduktor.io" # or env var CDK_ADMIN_EMAIL
#   admin_password = "testP4ss!"   # or env var CDK_ADMIN_PASSWORD

#   # optional http client TLS configuration
# #   cert     = file("path/to/cert.pem") # or env var CDK_CERT
# #   insecure = true                     # or env var CDK_INSECURE

#   # optional authentication via certificate
# #   key    = file("path/to/key.pem") # or env var CDK_KEY
# #   cacert = file("path/to/ca.pem")  # or env var CDK_CA_CERT
# }