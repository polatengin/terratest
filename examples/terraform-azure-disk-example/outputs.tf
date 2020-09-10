output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "disk_name" {
  value = azurerm_managed_disk.main.name
}

output "disk_type" {
  value = azurerm_managed_disk.main.storage_account_type
}
