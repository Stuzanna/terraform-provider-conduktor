# ---
# '{
#     "kind" : "GatewayServiceAccount",
#     "apiVersion" : "gateway/v2",
#     "metadata" : {
#       "name" : "local-app-finance-dev",
#       "vCluster" : "passthrough"
#     },
#     "spec" : { "type" : "LOCAL" }
#   }'
# ---

#  Local SA no vcluster, should this need passthrough? No auto adds.
resource "conduktor_gateway_service_account_v2" "local_sa2" {
  provider = conduktor.gateway
  name     = "local-app-finance-dev"
  spec {
    type = "LOCAL"
  }
}

# resource "conduktor_gateway_service_account_v2" "external_sa" {
#   provider = conduktor.gateway
#   name     = "azure-app-billing-dev"
#   vcluster = "passthrough"
#   spec {
#     type           = "EXTERNAL"
#     external_names = ["externalName"]
#   }
# }