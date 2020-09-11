# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# ARM_CLIENT_ID
# ARM_CLIENT_SECRET
# ARM_SUBSCRIPTION_ID
# ARM_TENANT_ID

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "prefix" {
  description = "temp"
  type        = string
  default     = "terratest-avs"
}

variable "location" {
  description = "The region"
  type        = string
  default     = "East US"
}

variable "disk_type" {
  description = "The managed disk type"
  type        = string
  default     = "Standard_LRS"
}


