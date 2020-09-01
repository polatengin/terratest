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

resource "azurerm_storage_account" "storageaccount" {
  name                     = lower(random_string.default.result)
  resource_group_name      = azurerm_resource_group.main.name
  location                 = azurerm_resource_group.main.location
  account_kind             = "StorageV2"
  account_tier             = "Standard"
  account_replication_type = "GRS"
}

resource "azurerm_storage_container" "container" {
  name = "container1"
  storage_account_name  = azurerm_storage_account.storageaccount.name
  container_access_type = "private"
}
