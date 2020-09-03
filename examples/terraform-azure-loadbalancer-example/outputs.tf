output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "loadbalancer01_name" {
  value = azurerm_lb.main01.name
}

output "loadbalancer02_name" {
  value = azurerm_lb.main02.name
}