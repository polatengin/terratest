output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "virtual_network_name" {
  value = azurerm_virtual_network.main.name
}

output "subnet_name" {
  value = azurerm_subnet.main.name
}

output "public_address_name" {
  value = azurerm_public_ip.main.name
}

output "network_interface_internal" {
  value = azurerm_network_interface.internal.name
}

output "network_interface_external" {
  value = azurerm_network_interface.external.name
}


