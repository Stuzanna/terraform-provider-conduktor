terraform {
  required_providers {
    conduktor = {
      source  = "terraform.local/conduktor/conduktor" # local provider
      version = "0.3.0"                               # latest version found locally in the plugin cache.
    }
  }
}

provider "conduktor" {
  # mandatory console URL
  console_url = "http://localhost:8080" # or env vars CDK_CONSOLE_URL or CDK_BASE_URL

  # authentication either with api token or admin credentials
  #   api_token = "your-api-token" # or env var CDK_API_TOKEN or CDK_API_KEY
  admin_email    = "admin@conduktor.io" # or env var CDK_ADMIN_EMAIL
  admin_password = "adminP4ss!"         # or env var CDK_ADMIN_PASSWORD

  # optional http client TLS configuration
  #   cert     = file("path/to/cert.pem") # or env var CDK_CERT
  #   insecure = true                     # or env var CDK_INSECURE

  # optional authentication via certificate
  #   key    = file("path/to/key.pem") # or env var CDK_KEY
  #   cacert = file("path/to/ca.pem")  # or env var CDK_CA_CERT
}

resource "conduktor_kafka_cluster_v2" "minimal" {
  name = "minimal-kafka-cluster"
  spec {
    display_name      = "Minimal Kafka Cluster"
    bootstrap_servers = "kafka1:9092"
  }
}

# resource "conduktor_kafka_connect_v2" "simple" {
#   name    = "simple-connect-cluster"
#   cluster = conduktor_kafka_cluster_v2.minimal.name
#   spec {
#     display_name = "Simple Connect Cluster"
#     urls         = "http://kafka-connect:8083"
#   }
# }


resource "conduktor_kafka_connect_v2" "basic" {
  name    = "simple-connect-cluster"
  cluster = conduktor_kafka_cluster_v2.minimal.name
  labels = {
    description   = "This is a complex connect using basic authentication"
    documentation = "https://docs.mycompany.com/complex-connect"
    env           = "dev"
  }
  spec {
    display_name = "Simple Connect Cluster"
    urls         = "http://kafka-connect:8083"
    headers = {
      X-PROJECT-HEADER = "value"
      Cache-Control    = "no-cache"
    }
    ignore_untrusted_certificate = false
    security = {
      type     = "BasicAuth"
      username = "user"
      password = "password"
    }
  }
}
# kafka-connect:8083