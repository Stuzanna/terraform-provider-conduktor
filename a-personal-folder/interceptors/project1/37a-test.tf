resource "conduktor_gateway_interceptor_v2" "topic-policy" {
  # provider = conduktor.gateway
  name = "enforce-partition-limit2"
  spec {
    plugin_class = "io.conduktor.gateway.interceptor.safeguard.CreateTopicPolicyPlugin"
    priority     = 1
    config = jsonencode({
      topic = "myprefix-.*"
      numPartition = {
        min    = 1
        max    = 5
        action = "INFO"
      }
    })
  }
  scope {
    vcluster = "passthrough" # This targets only the "prod-cluster" vcluster
  }
}

# resource "conduktor_gateway_interceptor_v2" "field-encryption" {
#   provider = conduktor.gateway
#   name = "field-encryption"
#   spec {
#     plugin_class = "io.conduktor.gateway.interceptor.EncryptPlugin"
#     priority     = 1
#     config = jsonencode({
#         "topic": ".*",
#         "recordValue": {
#           "fields": [
#             {
#               "fieldName": "password",
#               "keySecretId": "password-secret",
#               "algorithm": "AES128_GCM"
#             }
#           ]
#         }
#       })
#   }
# }