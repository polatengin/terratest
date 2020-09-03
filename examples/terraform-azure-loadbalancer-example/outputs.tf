output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "loadbalancer_name" {
  value = azurerm_lb.main01.name
}