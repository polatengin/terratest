# Terraform Azure Example

This folder contains a simple Terraform module that deploys resources in [Azure](https://azure.microsoft.com/) to demonstrate
how you can use Terratest to write automated tests for your Azure Terraform code. This module deploys the following resources:
- Resource Group
- Public IP Resource
- Virtual Network (vnet) with a Subnet
- Load Balancer, associated with Public IP
- Load Balancer with Private IP, associated with vnet

For more information on the above resources, you may visit the official Azure website for more information:
- Resource Groups: https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/manage-resource-groups-portal#what-is-a-resource-group
- Public IPs: https://docs.microsoft.com/en-us/azure/virtual-network/public-ip-addresses
- Virtual Networks: https://docs.microsoft.com/en-us/azure/virtual-network/virtual-networks-overview
- Subnets: https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-manage-subnet
- Load Balancers: https://docs.microsoft.com/en-us/azure/load-balancer/load-balancer-overview

Check out [/test/terraform_azure_loadbalancer_example_test.go](/test/terraform_azure_loadbalancer_example_test.go) to see how you can write
automated tests for this module.

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
1. Make sure [the azure-sdk-for-go versions match](#check-go-dependencies) in [/test/go.mod](/test/go.mod) and in [test/terraform_azure_loadbalancer_example_test.go](/test/terraform_azure_loadbalancer_example_test.go).
1. `go build terraform_azure_loadbalancer_example_test.go`
1. `go test -v -run TestTerraformAzureLoadBalancerExample`

## Check Go Dependencies

Check that the `github.com/Azure/azure-sdk-for-go` version in your generated `go.mod` for this test matches the version in the terratest [go.mod](https://github.com/gruntwork-io/terratest/blob/master/go.mod) file.  

> This was tested with **go1.14.1**.

### Check Azure-sdk-for-go version

Let's make sure [go.mod](https://github.com/gruntwork-io/terratest/blob/master/go.mod) includes the appropriate [azure-sdk-for-go version](https://github.com/Azure/azure-sdk-for-go/releases/tag/v38.1.0):

```go
require (
    ...
    github.com/Azure/azure-sdk-for-go v38.1.0+incompatible
    ...
)
```

We should check that [test/terraform_azure_example_test.go](/test/terraform_azure_example_test.go) includes the corresponding [azure-sdk-for-go package](https://github.com/Azure/azure-sdk-for-go/tree/master/services/compute/mgmt/2019-07-01/compute):

```go
import (
    "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
    ...
)
```

If we make changes to either the **go.mod** or the **go test file**, we should make sure that the go build command works still.

```powershell
go build terraform_azure_example_test.go
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

## Load Balancer Module APIs
* `LoadBalancerExistsE` checks if the given Load Balancer exists in the given subscription and returns true/false
* `GetLoadBalancerE` checks if the given Load Balancer exists in the given subscription and returns a Load Balancer resource (or nil if not found)
* `GetLoadBalancerClientE` checks if the given Load Balancer exists in the given subscription and returns a Load Balancer client object (or nil if not found)
* `GetPublicIPAddressE` checks if the given Public IP Address resource exists in the given subscription and returns a Public IP Address object (or nil if not found)
* `GetPublicIPAddressClientE` checks if the given Public IP Address resource exists in the given subscription and returns a Public IP Address client object (or nil if not found)