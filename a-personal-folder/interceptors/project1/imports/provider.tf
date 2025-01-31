terraform {
  required_providers {
    conduktor = {
      source  = "terraform.local/conduktor/conduktor" # local provider
      version = "0.0.11"                              # latest version found locally in the plugin cache.
    }
  }
}

# gateway only, not multi
provider "conduktor" {
  mode = "gateway"
  # mandatory gateway URL
  base_url = "http://localhost:8888" # or env vars CDK_GATEWAY_BASE_URL or CDK_BASE_URL

  # authentication with admin credentials
  admin_user     = "admin"     # or env var CDK_GATEWAY_USER or CDK_ADMIN_USER
  admin_password = "conduktor" # or env var CDK_GATEWAY_PASSWORD or CDK_ADMIN_PASSWORD

  # optional http client TLS configuration
  # cert     = file("path/to/cert.pem") # or env var CDK_GATEWAY_CERT or CDK_CERT
  # insecure = true                     # or env var CDK_GATEWAY_INSECURE or CDK_INSECURE

  # optional authentication via certificate
  # key    = file("path/to/key.pem") # or env var CDK_GATEWAY_KEY or CDK_KEY
  # cacert = file("path/to/ca.pem")  # or env var CDK_GATEWAY_CACERT or CDK_CACERT
}
