curl \
  --request GET \
  --url 'http://localhost:8888/gateway/v2/service-account?vcluster=passthrough' \
  --header 'Authorization: Basic YWRtaW46Y29uZHVrdG9y'

curl \
  --request GET \
  --url 'http://localhost:8888/gateway/v2/service-account?vcluster=passthrough?name=azure-app-billing-dev' \
  --header 'Authorization: Basic YWRtaW46Y29uZHVrdG9y'


  curl \
  --request DELETE \
  --url 'http://localhost:8888/gateway/v2/service-account' \
  --header 'Authorization: Basic YWRtaW46Y29uZHVrdG9y' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name" : "azure-app-billing-dev"
}'



curl \
  --request PUT \
  --url 'http://localhost:8888/gateway/v2/service-account' \
  --header 'Authorization: Basic YWRtaW46Y29uZHVrdG9y' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "kind" : "GatewayServiceAccount",
    "apiVersion" : "gateway/v2",
    "metadata" : {
      "name" : "local-app-2",
      "vCluster" : "passthrough"
    },
    "spec" : { "type" : "LOCAL" }
  }'

curl \
  --request POST \
  --url 'http://localhost:8888/gateway/v2/token' \
  --header 'Authorization: Basic YWRtaW46Y29uZHVrdG9y' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "vCluster": "passthrough",
    "username": "local-app-finance-dev",
    "lifeTimeSeconds": 3600000
  }'