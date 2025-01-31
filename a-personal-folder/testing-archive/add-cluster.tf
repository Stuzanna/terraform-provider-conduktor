# terraform {
#   required_providers {
#     conduktor = {
#         source = "conduktor/conduktor"
#         version = "~> 0.1.x" # where X.Y is the current major version and minor version
#     }
#   }
# }

# terraform {
#   required_providers {
#     conduktor = {
#       source  = "terraform.local/conduktor/conduktor" # local provider
#       version = ">= 0.1.3"                            # latest version found locally in the plugin cache.
#     }
#   }
# }

# provider "conduktor" {
#   # mandatory console URL
#   console_url = "http://localhost:8080" # or env vars CDK_CONSOLE_URL or CDK_BASE_URL

#   # authentication either with api token or admin credentials
# #   api_token = "your-api-token" # or env var CDK_API_TOKEN or CDK_API_KEY
#   admin_email    = "admin@conduktor.io" # or env var CDK_ADMIN_EMAIL
#   admin_password = "testP4ss!"   # or env var CDK_ADMIN_PASSWORD

#   # optional http client TLS configuration
# #   cert     = file("path/to/cert.pem") # or env var CDK_CERT
# #   insecure = true                     # or env var CDK_INSECURE

#   # optional authentication via certificate
# #   key    = file("path/to/key.pem") # or env var CDK_KEY
# #   cacert = file("path/to/ca.pem")  # or env var CDK_CA_CERT
# }


resource "conduktor_console_kafka_cluster_v2" "local-rp" {
  provider = conduktor.console
  name     = "local-rp"
  spec {
    display_name                 = "Local RP"
    icon                         = "kafka"
    color                        = "#000000"
    bootstrap_servers            = "redpanda:9092"
    ignore_untrusted_certificate = true

    schema_registry = {
      type                         = "ConfluentLike"
      url                          = "http://redpanda:8082"
      ignore_untrusted_certificate = true
      security = {
        type = "NoSecurity"
      }
    }
  }
}

resource "conduktor_console_kafka_cluster_v2" "gateway" {
  provider = conduktor.console
  name     = "gateway"
  spec {
    display_name                 = "gateway"
    icon                         = "dog"
    color                        = "#000000"
    bootstrap_servers            = "gateway:9094"
    ignore_untrusted_certificate = true
    properties = {
      "sasl.jaas.config"  = "org.apache.kafka.common.security.plain.PlainLoginModule required username='local-app-finance-dev' password='<>';"
      "sasl.mechanism"    = "PLAIN"
      "security.protocol" = "SASL_PLAINTEXT"
    }
  }
}

# resource "conduktor_kafka_cluster_v2" "aiven" {
#   name = "aiven-cluster-tf"
#   labels = {
#     "env" = "test"
#   }
#   spec {
#     display_name      = "Aiven Cluster"
#     bootstrap_servers = "<>"
#     properties = {
#       "sasl.jaas.config"  = "org.apache.kafka.common.security.plain.PlainLoginModule required username='avnadmin' password='<>';"
#       "security.protocol" = "SASL_SSL"
#       "sasl.mechanism"    = "PLAIN"
#     }
#     icon                         = "crab"
#     ignore_untrusted_certificate = true
#     kafka_flavor = {
#       type         = "Aiven"
#       # api_token    = "<>"
#       api_token = "<>"
#       project      = "<>"
#       service_name = "<>"
#     }


#     schema_registry = {
#       type                         = "ConfluentLike"
#       url                          = "https://sr.aiven.io:8081"
#       ignore_untrusted_certificate = false
#       security = {
#         type     = "BasicAuth"
#         username = "uuuuuuu"
#         password = "ppppppp"
#       }
#     }
#   }
# }

# resource "conduktor_kafka_cluster_v2" "aiven-darren" {
#   name = "aiven-cluster-tf-darren"
#   labels = {
#     "env" = "test"
#   }
#   spec {
#     display_name      = "Aiven Cluster Darren"
#     bootstrap_servers = "<>"
#     properties = {
#       "sasl.jaas.config"  = "org.apache.kafka.common.security.plain.PlainLoginModule required username='darren-test' password='<>';"
#       "security.protocol" = "SASL_SSL"
#       "sasl.mechanism"    = "PLAIN"
#     }
#     icon                         = "crab"
#     ignore_untrusted_certificate = true
#     kafka_flavor = {
#       type         = "Aiven"
#       # api_token    = "<>"
#       api_token = "<>"
#       project      = ""
#       service_name = ""
#     }


# schema_registry = {
#   type                         = "ConfluentLike"
#   url                          = "https://sr.aiven.io:8081"
#   ignore_untrusted_certificate = false
#   security = {
#     type     = "BasicAuth"
#     username = "uuuuuuu"
#     password = "ppppppp"
#   }
# }
#   }
# }

# resource "conduktor_kafka_cluster_v2" "confluent-tf" {
#   name = "confluent-cluster-tf"
#   labels = {
#     "env" = "staging"
#   }
#   spec {
#     display_name      = "Confluent Cluster TF"
#     bootstrap_servers = "pkc-l6wr6.europe-west2.gcp.confluent.cloud:9092"
#     properties = {
#       "sasl.jaas.config"  = "org.apache.kafka.common.security.plain.PlainLoginModule required username='<>' password='<>';"
#       "security.protocol" = "SASL_SSL"
#       "sasl.mechanism"    = "PLAIN"
#     }
#     icon                         = "kafka"
#     ignore_untrusted_certificate = false
#     kafka_flavor = {
#       type                     = "Confluent"
#       key                      = "<>"
#       secret                   = "<>"
#       confluent_environment_id = "<>"
#       confluent_cluster_id     = "<>"
#     }
#     schema_registry = {
#       type                         = "ConfluentLike"
#       url                          = "<>"
#       ignore_untrusted_certificate = false
#       security = {
#         username = "<>"
#         password = "<>"
#         type = "BasicAuth"
#         }
#     }
#   }
# }

# resource "conduktor_kafka_cluster_v2" "aws_msk_tf" {
#   name = "aws-cluster-tf"
#   labels = {
#     "env" = "prod"
#   }
#   spec {
#     display_name      = "AWS MSK Cluster TF"
#     bootstrap_servers = "<>"
#     properties = {
#       "sasl.jaas.config"                   = "software.amazon.msk.auth.iam.IAMLoginModule required awsRoleArn='arn:aws:iam::<number>:role/<role>';"
#       "sasl.client.callback.handler.class" = "software.amazon.msk.auth.iam.IAMClientCallbackHandler"
#       "security.protocol"                  = "SASL_SSL"
#       "sasl.mechanism"                     = "AWS_MSK_IAM"
#     }
#     icon                         = "kafka"
#     color                        = "#FF0000"
#     ignore_untrusted_certificate = true
#     schema_registry = {
#       type          = "Glue"
#       region        = "eu-west-1"
#       registry_name = "default-registry"
#       security = {
#         type          = "Credentials"
#         access_key_id = ""
#         secret_key    = ""
#       }
#     }
#   }
# }

# resource "conduktor_kafka_cluster_v2" "aws_msk_tf_by_key" {
#   name = "aws-cluster-tf2-by_key"
#   labels = {
#     "env" = "prod"
#   }
#   spec {
#     display_name      = "AWS MSK Cluster TF by key"
#     bootstrap_servers = "<>"
#     properties = {
#       "sasl.jaas.config"                   = "software.amazon.msk.auth.iam.IAMLoginModule required;"
#       "sasl.client.callback.handler.class" = "software.amazon.msk.auth.iam.IAMClientCallbackHandler"
#       "security.protocol"                  = "SASL_SSL"
#       "sasl.mechanism"                     = "AWS_MSK_IAM"
#     }
#     icon                         = "kafka"
#     color                        = "#FF0000"
#     ignore_untrusted_certificate = true
#     schema_registry = {
#       type          = "Glue"
#       region        = "eu-west-1"
#       registry_name = "default-registry"
#       security = {
#         type          = "Credentials"
#         access_key_id = ""
#         secret_key    = ""
#       }
#     }
#   }
# }

# resource "conduktor_kafka_cluster_v2" "gateway-cluster-TF1" {
#   name = "gateway-cluster-tf1"
#   labels = {
#     "env" = "prod"
#   }
#   spec {
#     display_name      = "Gateway Cluster TF 1"
#     bootstrap_servers = "conduktor-gateway:6969"
#     properties = {
#       "sasl.jaas.config"  = "org.apache.kafka.common.security.plain.PlainLoginModule required username='stu' password='<>';"
#       "security.protocol" = "SASL_PLAINTEXT"
#       "sasl.mechanism"    = "PLAIN"
#     }
#     icon                         = "shield-blank"
#     ignore_untrusted_certificate = true
#     kafka_flavor = {
#       type                         = "Gateway"
#       url                          = "http://conduktor-gateway:8088"
#       user                         = "admin"
#       password                     = "conduktor"
#       virtual_cluster              = ""
#       ignore_untrusted_certificate = true
#     }
#     schema_registry = {
#       type                         = "ConfluentLike"
#       url                          = "http://localhost:8084"
#       ignore_untrusted_certificate = true
#       security = {
#         type  = "BearerToken"
#         token = "auth-token"
#       }
#     }
#   }
# }
