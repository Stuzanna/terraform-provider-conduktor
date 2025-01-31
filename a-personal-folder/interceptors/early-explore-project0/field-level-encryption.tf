resource "conduktor_gateway_interceprtor_encryption" "field_level_encryption" {
  provider = conduktor.gateway
  name     = "fieldLevelEncryption"
  spec {
    priority = 100
    config = {
      topic = ".*"
      schema_registry_config = {
        host = "http://schema-registry:8081"
      }
      kms_config = {
        vault = {
          uri     = "http://vault:8200"
          token   = "vault-plaintext-root-token"
          version = 1
        }
      }
      record_value = {
        value_fields = [
          {
            field_name    = "password"
            key_secret_id = "vault-kms://vault:8200/transit/keys/password-secret"
            algorithm     = "AES128_GCM"
          },
          {
            field_name    = "visa",
            key_secret_id = "vault-kms://vault:8200/transit/keys/{{record.header.test-header}}-visa-secret-{{record.key}}-{{record.value.username}}-{{record.value.education.account.accountId}}"
            algorithm     = "AES128_GCM"
          },
          {
            field_name    = "education.account.username"
            key_secret_id = "azure-kms://my-key-vault.vault.azure.net/keys/conduktor-gateway/4ceb7a4d1f3e4738b23bea870ae8745d"
            algorithm     = "AES128_GCM"
          }
        ]
      }
    }
  }
}