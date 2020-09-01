# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# ARM_CLIENT_ID
# ARM_CLIENT_SECRET
# ARM_SUBSCRIPTION_ID
# ARM_TENANT_ID

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "prefix" {
  description = "The prefix that will be attached to all resources deployed."
  type        = string
  default     = "testprefix"
}

variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
}

#variable "storage_account_name" {
#  description = "The name of the storage account."
#  type        = string
#  default     = "terrateststore7"
#}