output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "storage_account_name" {
  value = azurerm_storage_account.storageaccount.name
}

output "storage_account_account_tier" {
  value = azurerm_storage_account.storageaccount.account_tier
}

output "storage_account_account_kind" {
  value = azurerm_storage_account.storageaccount.account_kind
}

output "storage_account_account_replication_type" {
  value = azurerm_storage_account.storageaccount.account_replication_type
}

output "storage_container_name" {
  value = azurerm_storage_container.container.name
}
