output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "availability_set_name" {
  value = azurerm_availability_set.main.name
}

output "availability_set_fdc" {
  value = azurerm_availability_set.main.platform_fault_domain_count
}

output "vm_name_01" {
  value = azurerm_virtual_machine.main01.name
}

output "vm_name_02" {
  value = azurerm_virtual_machine.main02.name
}

