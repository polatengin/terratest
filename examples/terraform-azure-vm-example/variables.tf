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

variable "resource_group_name" {
  description = "temp"
  type        = string
  default     = "terratest-rg-01"
}

variable "location" {
  description = "temp"
  type        = string
  default     = "East US"
}

variable "vnet_name" {
  description = "temp"
  type        = string
  default     = "terratest-vnet-01"
}

variable "vnet_address" {
  description = "temp"
  type        = string
  default     = "10.0.0.0/16"
}

variable "subnet_name" {
  description = "temp"
  type        = string
  default     = "terratest-snt-01-internal"
}

variable "subnet_address" {
  description = "temp"
  type        = string
  default     = "10.0.17.0/24"
}

variable "public_address_name" {
  description = "temp"
  type        = string
  default     = "terratest-pip-01-external"
}

variable "nic_name" {
  description = "temp"
  type        = string
  default     = "terratest-nic-01"
}

variable "ip_config_name" {
  description = "temp"
  type        = string
  default     = "terratest-pip-01-internal"
}

variable "private_ip" {
  description = "temp"
  type        = string
  default     = "10.0.17.4"
}

variable "avs_name" {
  description = "temp"
  type        = string
  default     = "terratest-avs-01"
}

variable "vm_name" {
  description = "temp"
  type        = string
  default     = "terratest-vm-01"
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

variable "vm_size" {
  description = "temp"
  type        = string
  default     = "Standard_DS1_v2"
}

variable "vm_license" {
  description = "temp"
  type        = string
  default     = "Windows_Server"
}

variable "vm_image_publisher" {
  description = "temp"
  type        = string
  default     = "MicrosoftWindowsServer"
}

variable "vm_image_offer" {
  description = "temp"
  type        = string
  default     = "WindowsServer"
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

variable "disk_caching" {
  description = "temp"
  type        = string
  default     = "ReadWrite"
}

variable "osdisk_name" {
  description = "temp"
  type        = string
  default     = "terratest-vm-01-osdisk"
}

variable "osdisk_create_option" {
  description = "temp"
  type        = string
  default     = "FromImage"
}

variable "disk_01_name" {
  description = "temp"
  type        = string
  default     = "terratest-vm-01-disk-01"
}

variable "disk_01_create_option" {
  description = "temp"
  type        = string
  default     = "Empty"
}

variable "disk_01_size" {
  description = "temp"
  type        = number
  default     = 10
}

variable "disk_01_lun" {
  description = "temp"
  type        = string
  default     = "10"
}


# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------


