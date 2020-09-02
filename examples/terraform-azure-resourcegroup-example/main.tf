resource "azurerm_resource_group" "main" {
  name     =  "${var.prefix}-${var.resourceGroupName}"
  location = var.location
}