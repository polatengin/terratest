
output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "vm_name" {
  value = azurerm_virtual_machine.main.name
}

output "virtual_network_name" {
  value = azurerm_virtual_network.main.name
}

output "subnet_name" {
  value = azurerm_subnet.main.name
}

output "public_ip_name" {
  value = azurerm_public_ip.main.name
}

output "network_interface_name" {
  value = azurerm_network_interface.main.name
}

output "availability_set_name" {
  value = azurerm_availability_set.main.name
}

output "managed_disk_name" {
  value = azurerm_managed_disk.main.name
}