resource "azurerm_resource_group" "main" {
  name     =  "${var.prefix}-resources"
  location = var.location
}

resource "random_id" "server" {
  byte_length = 8
}

data "azurerm_client_config" "current" {}

data "azurerm_key_vault_access_policy" "contributor" {
  name = "Key, Secret, & Certificate Management"
}

resource "azurerm_key_vault" "kvexample" {
  name                        = format("%s%s", "kv", random_id.server.hex)
  location                    = azurerm_resource_group.main.location
  resource_group_name         = azurerm_resource_group.main.name
  enabled_for_disk_encryption = true
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  soft_delete_enabled         = true
  soft_delete_retention_days  = 7
  purge_protection_enabled    = false

  sku_name = "standard"

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    key_permissions = [
      "create",
      "get",
      "list",
      "delete",
    ]

    secret_permissions = [
      "set",
      "get",
      "list",
      "delete",
    ]

    certificate_permissions = [
      "create",
      "delete",
      "deleteissuers",
      "get",
      "getissuers",
      "import",
      "list",
      "listissuers",
      "managecontacts",
      "manageissuers",
      "setissuers",
      "update",
    ]
  }
}

resource "azurerm_key_vault_secret" "example" {
  name         = "secret1"
  value        = "mysecret"
  key_vault_id = azurerm_key_vault.kvexample.id
}

resource "azurerm_key_vault_key" "example" {
  name         = "key1"
  key_vault_id = azurerm_key_vault.kvexample.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
  ]
}

resource "azurerm_key_vault_certificate" "example" {
  name         = "certificate1"
  key_vault_id = azurerm_key_vault.kvexample.id

  certificate {
    contents = filebase64("example.pfx")
    password = "password"
  }

  certificate_policy {
    issuer_parameters {
      name = "Self"
    }

    key_properties {
      exportable = true
      key_size   = 2048
      key_type   = "RSA"
      reuse_key  = false
    }

    secret_properties {
      content_type = "application/x-pkcs12"
    }
  }
}
