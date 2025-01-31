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
    vcluster = "passthrough" # This targets only the "prod-cluster" vcluster
  }
}

resource "conduktor_gateway_interceptor_v2" "audit_log" {
  name = "${var.environment}-audit-log"
  spec {
    plugin_class = "io.conduktor.gateway.interceptor.AuditPlugin"
    priority     = 200
    config = jsonencode({
      topic = ".*"
    })
  }
}