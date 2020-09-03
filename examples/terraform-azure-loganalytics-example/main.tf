// Resources

resource "random_string" "default" {
  length = 8  
  lower = true
  number = false
  special = false
}

resource "azurerm_resource_group" "main" {
  name     =  "${var.prefix}-resources"
  location = var.location
}

resource "azurerm_log_analytics_workspace" "logws" {
  name                = lower(random_string.default.result)
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

