output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "loadbalancer01_name" {
  value = azurerm_lb.main01.name
}

output "lb01_feconfig" {
  value = azurerm_lb.main01.frontend_ip_configuration[0].name
}

output "pip_forlb01" {
  value = azurerm_public_ip.main.name
}

output "loadbalancer02_name" {
  value = azurerm_lb.main02.name
}
output "feIPConfig_forlb02" {
  value = azurerm_lb.main02.frontend_ip_configuration[0].private_ip_address
}

output "feSubnet_forlb02" {
  value = azurerm_lb.main02.frontend_ip_configuration[0].subnet_id
}
