# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A BLOC CONTAINER ON A STORAGE ACCOUNT
# This is an example of how to deploy an Azure Storage Account with a Blob Container
# ---------------------------------------------------------------------------------------------------------------------

# ------------------------------------------------------------------------------
# CONFIGURE OUR AZURE CONNECTION
# ------------------------------------------------------------------------------

provider "azurerm" {
  version = "=2.22.0"
  features {}
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A RESOURCE GROUP
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A BLOB CONTAINER
# ---------------------------------------------------------------------------------------------------------------------

data "azurerm_client_config" "current" {
}

resource "azurerm_storage_account" "storage_account" {
  name                     = var.account_name
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "Testing"
  }
}

resource "azurerm_storage_container" "blob_container" {
  name                  = var.blob_container_name
  storage_account_name  = azurerm_storage_account.storage_account.name
  container_access_type = "private"
}
