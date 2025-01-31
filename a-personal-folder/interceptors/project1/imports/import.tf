# First define the empty resource
import {
  to = conduktor_gateway_interceptor_v2.mask-sensitive-fields
  id = "mask-sensitive-fields" # Import "interceptor-name" Interceptor
}

# resource "conduktor_gateway_interceptor_v2" "mask-sensitive-fields" {
#   name = "mask-sensitive-fields"
#   # Leave empty for now - will be populated after import
# }

resource "conduktor_gateway_interceptor_v2" "mask-sensitive-fields" {
  name = "mask-sensitive-fields"
  spec {
    plugin_class = "io.conduktor.gateway.interceptor.FieldLevelDataMaskingPlugin"
    priority     = 100
    comment      = "Mask sensitive data"
    config = jsonencode(jsondecode(<<EOF
  {
        "topic": "^[A-Za-z]*_masked$",
        "schemaRegistryConfig": {
          "host": "http://redpanda-0:8081"
        },
        "policies": [
          {
              "name": "Mask credit card",
              "rule": {
                "type": "MASK_ALL"
              },
              "fields": [
                "profile.creditCardNumber",
                "contact.email"
              ]
          },
          {
              "name": "Partial mask phone",
              "rule": {
                "type": "MASK_FIRST_N",
                "maskingChar": "*",
                "numberOfChars": 9
              },
              "fields": [
                "contact.phone"
              ]
          }
        ]
    }
        EOF
    ))
  }
}