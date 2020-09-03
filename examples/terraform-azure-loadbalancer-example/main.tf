provider "azurerm" {
  version = "=1.31.0"
}

# ---------------------------------------------------------------------------------------------------------------------
# PIN TERRAFORM VERSION TO >= 0.12
# The examples have been upgraded to 0.12 syntax
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = ">= 0.12"
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A RESOURCE GROUP
# See test/terraform_azure_example_test.go for how to write automated tests for this code.
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "main" {
  name     = "${var.prefix}-resources"
  location = "East US"
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY LOAD BALANCER WITH PUBLIC IP 
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_public_ip" "main" {
  name                    = "${var.prefix}-pip"
  resource_group_name = azurerm_resource_group.main.name
  location            = azurerm_resource_group.main.location
  allocation_method       = "Static"
  ip_version              = "IPv4"
  sku                     = "Standard"
  idle_timeout_in_minutes = "4"
}

resource "azurerm_lb" "main01" {
  name                = "${var.prefix}-lb01"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  sku                 = "Standard"

    frontend_ip_configuration {
      name                 = "${var.prefix}-frontendip01"
      public_ip_address_id = azurerm_public_ip.main.id
    }
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY LOAD BALANCER WITH PRIVATE IP 
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_virtual_network" "main" {
  name                = "${var.prefix}-vnet01"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  address_space       = ["10.200.0.0/21"]
  dns_servers         = ["10.200.0.5", "10.200.0.6"]

}

resource "azurerm_subnet" "main" {
  name                = "${var.prefix}-subnet01"
  resource_group_name = azurerm_resource_group.main.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefix     = "10.200.2.0/25"
}

resource "azurerm_lb" "main02" {
  name                = "${var.prefix}-lb02"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  sku                 = "Standard"

    frontend_ip_configuration {
      name                 = "${var.prefix}-frontendip02"
      subnet_id                     = azurerm_subnet.main.id
      private_ip_address            = "10.200.2.10"
      private_ip_address_allocation = "Static"
    }
}
