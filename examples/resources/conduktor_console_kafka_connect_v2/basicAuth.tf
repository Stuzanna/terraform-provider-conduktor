resource "conduktor_console_kafka_cluster_v2" "minimal" {
  name = "mini-cluster"
  spec = {
    display_name      = "Minimal Cluster"
    bootstrap_servers = "localhost:9092"
  }
}

resource "conduktor_console_kafka_connect_v2" "basic" {
  name    = "basic-connect"
  cluster = conduktor_console_kafka_cluster_v2.minimal.name
  labels = {
    description   = "This is a complex connect using basic authentication"
    documentation = "https://docs.mycompany.com/complex-connect"
    env           = "dev"
  }
  spec = {
    display_name = "Basic Connect server"
    urls         = "http://localhost:8083"
    headers = {
      X-PROJECT-HEADER = "value"
      Cache-Control    = "no-cache"
    }
    ignore_untrusted_certificate = false
    security = {
      basic_auth = {
        username = "user"
        password = "password"
      }
    }
  }
}
