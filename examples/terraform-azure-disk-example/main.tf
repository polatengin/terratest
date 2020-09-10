# ---
# DEPLOY A DATA DISK TO A VIRTUAL MACHINE
# ---

provider "azurerm" {
  version = "=1.31.0"
}

terraform {
  required_version = ">= 0.12"
}

resource "azurerm_resource_group" "main" {
  name     = "${var.prefix}-rg-01"
  location = var.location
}

resource "azurerm_managed_disk" "main" {
  name                 = "${var.prefix}-disk"
  location             = azurerm_resource_group.main.location
  resource_group_name  = azurerm_resource_group.main.name
  storage_account_type = var.disk_type
  create_option        = "Empty"
  disk_size_gb         = 10
}
