resource "conduktor_gateway_interceptor_v2" "mask-sensitive-fields-gen" {
  name = "mask-sensitive-fields"
  spec {
    comment = "Mask sensitive data"
    config = jsonencode({
      policies = [{
        fields = ["profile.creditCardNumber", "contact.email"]
        name   = "Mask credit card"
        rule = {
          type = "MASK_ALL"
        }
        }, {
        fields = ["contact.phone"]
        name   = "Partial mask phone"
        rule = {
          maskingChar   = "*"
          numberOfChars = 9
          type          = "MASK_FIRST_N"
        }
      }]
      schemaRegistryConfig = {
        host = "http://redpanda-0:8081"
      }
      topic = "^[A-Za-z]*_masked$"
    })
    plugin_class = "io.conduktor.gateway.interceptor.FieldLevelDataMaskingPlugin"
    priority     = 100
  }
}

# __generated__ by Terraform from "mask-sensitive-fields   "
resource "conduktor_gateway_interceptor_v2" "mask-sensitive-fields-gen2" {
  name = "mask-sensitive-fields2"
  scope {
    vcluster = "passthrough"
  }
  spec {
    comment = "Mask sensitive data"
    config = jsonencode({
      policies = [{
        fields = ["profile.creditCardNumber", "contact.email"]
        name   = "Mask credit card"
        rule = {
          type = "MASK_ALL"
        }
        }, {
        fields = ["contact.phone"]
        name   = "Partial mask phone"
        rule = {
          maskingChar   = "*"
          numberOfChars = 9
          type          = "MASK_FIRST_N"
        }
      }]
      schemaRegistryConfig = {
        host = "http://redpanda-0:8081"
      }
      topic = "^[A-Za-z]*_masked$"
    })
    plugin_class = "io.conduktor.gateway.interceptor.FieldLevelDataMaskingPlugin"
    priority     = 100
  }
}
