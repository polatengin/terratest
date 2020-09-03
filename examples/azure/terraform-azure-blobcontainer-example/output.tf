output "subscription_id" {
  value = data.azurerm_client_config.current.subscription_id
}

output "id" {
  value = azurerm_storage_account.storage_account.id
}
