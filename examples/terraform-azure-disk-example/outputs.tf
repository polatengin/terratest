output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "disk_name" {
  value = azurerm_virtual_machine.main.storage_os_disk[0].name
}

output "disk_type" {
  value = azurerm_virtual_machine.main.storage_os_disk[0].managed_disk_type
}
