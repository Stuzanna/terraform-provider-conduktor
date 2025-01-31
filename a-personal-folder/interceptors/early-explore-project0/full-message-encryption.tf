resource "conduktor_gateway_interceptor_encryption" "simple" {
  provider = conduktor.gateway
  name     = "myEncryptPlugin"
  spec {
    priority = 100
    config = {
      topic = ".*"
      schema_registry_config = {
        host = "http://schema-registry:8081"
      }
      kms_config = {
        // TBD
        // solution 1. All the vault inputs flattened
        vault = {
          type    = "token"
          uri     = "http://vault:8200"
          token   = "vault-plaintext-root-token"
          version = 1
        }
        // solution 2. different objects for different each type - vault_app_role, vault_ldap, vault_aws_ec2, vault_gcp etc...
        vault_token = {
          uri     = "http://vault:8200"
          token   = "vault-plaintext-root-token"
          version = 1
        }
      }
      record_value = {
        payload = {
          key_secret_id = "vault-kms://vault:8200/transit/keys/{{record.header.test-header}}-secret-key-account-username-{{record.topic}}",
          algorithm     = "AES128_GCM"
        }
      }
    }
  }
}