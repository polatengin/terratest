# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

variable "client_id" {
  description = "The Service Principal Client Id for AKS to modify Azure resources."
}
variable "client_secret" {
  description = "The Service Principal Client Password for AKS to modify Azure resources."
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "account_name" {
  description = "The name to set for the Storage Account"
}

variable "blob_container_name" {
  description = "The name to set for the Blob Container"
}

variable "resource_group_name" {
  description = "The name to set for the resource group."
  default     = "storage-rg"
}

variable "location" {
  description = "The location to set for the Resource Group"
  default     = "Central US"
}
