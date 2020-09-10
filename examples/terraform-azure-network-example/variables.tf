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

variable "prefix" {
  description = "temp"
  type        = string
  default     = "terratest-net"
}

variable "location" {
  description = "temp"
  type        = string
  default     = "East US"
}

variable "subnet_prefix" {
  description = "temp"
  type        = string
  default     = "10.0.20.0/24"
}

variable "private_ip" {
  description = "temp"
  type        = string
  default     = "10.0.20.5"
}

variable "dns_ip_01" {
  description = "temp"
  type        = string
  default     = "10.0.0.5"
}

variable "dns_ip_02" {
  description = "temp"
  type        = string
  default     = "10.0.0.6"
}