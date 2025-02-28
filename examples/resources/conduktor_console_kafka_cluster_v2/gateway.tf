resource "conduktor_console_kafka_cluster_v2" "gateway" {
  name = "gateway-cluster"
  labels = {
    "env" = "prod"
  }
  spec = {
    display_name      = "Gateway Cluster"
    bootstrap_servers = "gateway:6969"
    properties = {
      "sasl.jaas.config"  = "org.apache.kafka.common.security.plain.PlainLoginModule required username='admin' password='admin-secret';"
      "security.protocol" = "SASL_SSL"
      "sasl.mechanism"    = "PLAIN"
    }
    icon                         = "shield-blank"
    ignore_untrusted_certificate = true
    kafka_flavor = {
      gateway = {
        url                          = "http://gateway:8888"
        user                         = "admin"
        password                     = "admin"
        virtual_cluster              = "passthrough"
        ignore_untrusted_certificate = true
      }
    }
    schema_registry = {
      confluent_like = {
        url                          = "http://localhost:8081"
        ignore_untrusted_certificate = true
        security = {
          bearer_token = {
            token = "auth-token"
          }
        }
      }
    }
  }
}
