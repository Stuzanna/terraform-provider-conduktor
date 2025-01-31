module "interceptors" {
  source      = "../modules/common-interceptors"
  environment = "dev"
}

# Additional environment-specific interceptors
resource "conduktor_gateway_interceptor_v2" "topic-policy" {
  name = "${var.environment}-enforce-partition-limit2"
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
    vcluster = "passthrough"
  }
}