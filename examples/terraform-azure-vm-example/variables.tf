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
  default     = "terratest-vm"
}

variable "location" {
  description = "temp"
  type        = string
  default     = "East US"
}

variable "subnet_prefix" {
  description = "temp"
  type        = string
  default     = "10.0.17.0/24"
}

variable "private_ip" {
  description = "temp"
  type        = string
  default     = "10.0.17.4"
}

variable "vm_size" {
  description = "temp"
  type        = string
  default     = "Standard_DS1_v2"
}

variable "vm_image_sku" {
  description = "temp"
  type        = string
  default     = "2016-Datacenter"
}

variable "vm_image_version" {
  description = "temp"
  type        = string
  default     = "latest"
}

variable "disk_type" {
  description = "temp"
  type        = string
  default     = "Standard_LRS"
}


variable "user_name" {
  description = "the username to be provisioned into your vm"
  type        = string
  default     = "testadmin"
}

variable "password" {
  description = "the password to configure for ssh access"
  type        = string
  default     = "horriblepassword1234!"
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------


