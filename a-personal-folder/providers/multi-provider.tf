provider "conduktor" {
  alias    = "console"
  mode     = "console"
  base_url = "http://localhost:8080"

  # api_token = "your-api-token"
  admin_user     = "admin@conduktor.io"
  admin_password = "adminP4ss!"

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