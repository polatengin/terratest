# Terraform Azure Virtual Machine Example

This folder contains a complete Terraform VM module that deploys resources in [Azure](https://azure.microsoft.com/) to demonstrate
how you can use Terratest to write automated tests for your Azure Virtual Machine Terraform code. This module deploys these resources:

* A [Virtual Machine](https://azure.microsoft.com/en-us/services/virtual-machines/) and gives that VM the following:
    * `Virtual Machine Name` with the value specified in the `vm_name` variable.
    * `Managed Disk` with the value specified in the `managed_disk_name` variable.
    * `Availability Set` with the value specified in the `availability_set_name` variable.
* A [Virtual Network](https://azure.microsoft.com/en-us/services/virtual-network/) that gives the VM the following:
    * `Virtual Network Name` with the value specified in the `virtual_network_name` variable.
    * `Subnet` with the value specified in the `subnet_name` variable.
    * `Public Address` with the value specified in the `public_ip_name` variable.
    * `Network Interface` with the value specified in the `network_interface_name` variable.

Check out [test/terraform_azure_vm_test.go](/test/terraform_azure_vm_test.go) to see how you can write
automated tests for this module.

Note that the Virtual Machine madule creates a Microsoft Windows Server Image with and availability set and networking sample configurations for
demonstration purposes.

**WARNING**: This module and the automated tests for it deploy real resources into your Azure account which can cost you
money. The resources are all part of the [Azure Free Account](https://azure.microsoft.com/en-us/free/), so if you haven't used that up,
it should be free, but you are completely responsible for all Azure charges.

## Running this module manually

1. Sign up for [Azure](https://azure.microsoft.com/).
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest).
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Run `terraform init`.
1. Run `terraform apply`.
1. When you're done, run `terraform destroy`.

## Running automated tests against this module

1. Sign up for [Azure](https://azure.microsoft.com/).
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest).
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. [Review environment variables](#review-environment-variables).
1. Install [Golang](https://golang.org/) and make sure this code is checked out into your `GOPATH`.
1. `cd test`
1. Make sure [the azure-sdk-for-go versions match](#check-go-dependencies) in [/test/go.mod](/test/go.mod) and in [test/terraform_azure_example_test.go](/test/terraform_azure_example_test.go).
1. `go build terraform_azure_vm_test.go`
1. `go test -v -run TestTerraformAzureVmExample -timeout 20m` 
    * Note the extra -timeout flag of 20 minutes ensures proper Azure resource removal time.

## Test Module APIs

* `AssertVirtualMachineExists` checks if the given VM exists in the given subscription
* `GetAdminUserOfVirtualMachine` gets the admin username of the given Azure Virtual Machine
* `GetAvailabilitySetIDOfVirtualMachine` gets the Availability Set ID of the given Azure Virtual Machine
* `GetImageOfVirtualMachine` gets the VM Image of the given Azure Virtual Machine
* `GetManagedDiskCountOfVirtualMachine` gets the Managed Disk count of the given Azure Virtual Machine
* `GetManagedDiskNamesOfVirtualMachine` gets the list of Managed Disk names of the given Azure Virtual Machine
* `GetNicCountOfVirtualMachine` gets the Network Interface count of the given Azure Virtual Machine
* `GetNicNamesForVirtualMachine` gets a list of Network Interfaces for a given Azure Virtual Machine
* `GetOsDiskNameOfVirtualMachine` gets the OS Disk Name of the given Azure Virtual Machine
* `GetSizeOfVirtualMachine` gets the size type of the given Azure Virtual Machine
* `GetTagsForVirtualMachine` gets the tags of the given Virtual Machine as a map
* `GetVirtualMachine` gets a virtual machine in the specified resource group
* `GetVirtualMachineClientE` creates a Azure Virtual Machine client in the specified Azure Subscription
* `GetVirtualMachineE` gets a Virtual Machine in the specified Azure Resource Group
* `GetVirtualMachineInstance` gets a local virtual machine instance in the specified resource group
* `GetVirtualMachineInstanceE` gets a local virtual machine instance in the specified resource group with error
* `GetVirtualMachines` gets all virtual machine objects in the specified resource group
* `GetVirtualMachinesE` gets all virtual machine objects in the specified resource group with error
* `GetVirtualMachineSize` gets vm size
* `GetVirtualMachinesList` gets a list of all virtual machines in the specified resource group
* `GetVirtualMachinesListE` gets a list of all virtual machines in the specified resource group with error



## Check Go Dependencies

Check that the `github.com/Azure/azure-sdk-for-go` version in your generated `go.mod` for this test matches the version in the terratest [go.mod](https://github.com/gruntwork-io/terratest/blob/master/go.mod) file.  

> This was tested with **go1.14.4**.

### Check Azure-sdk-for-go version

Let's make sure [go.mod](https://github.com/gruntwork-io/terratest/blob/master/go.mod) includes the appropriate [azure-sdk-for-go version](https://github.com/Azure/azure-sdk-for-go/releases/tag/v38.1.0):

```go
require (
    ...
    github.com/Azure/azure-sdk-for-go v38.1.0+incompatible
    ...
)
```

If we make changes to either the **go.mod** or the **go test file**, we should make sure that the go build command works still.

```powershell
go build terraform_azure_vm_test.go
```

## Review Environment Variables

As part of configuring terraform for Azure, we'll want to check that we have set the appropriate [credentials](https://docs.microsoft.com/en-us/azure/terraform/terraform-install-configure?toc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fterraform%2Ftoc.json&bc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fbread%2Ftoc.json#set-up-terraform-access-to-azure) and also that we set the [environment variables](https://docs.microsoft.com/en-us/azure/terraform/terraform-install-configure?toc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fterraform%2Ftoc.json&bc=https%3A%2F%2Fdocs.microsoft.com%2Fen-us%2Fazure%2Fbread%2Ftoc.json#configure-terraform-environment-variables) on the testing host.

```bash
export ARM_CLIENT_ID=your_app_id
export ARM_CLIENT_SECRET=your_password
export ARM_SUBSCRIPTION_ID=your_subscription_id
export ARM_TENANT_ID=your_tenant_id
```

Note, in a Windows environment, these should be set as **system environment variables**.  We can use a PowerShell console with administrative rights to update these environment variables:

```powershell
[System.Environment]::SetEnvironmentVariable("ARM_CLIENT_ID",$your_app_id,[System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ARM_CLIENT_SECRET",$your_password,[System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ARM_SUBSCRIPTION_ID",$your_subscription_id,[System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ARM_TENANT_ID",$your_tenant_id,[System.EnvironmentVariableTarget]::Machine)
```

