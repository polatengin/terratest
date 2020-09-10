// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureNetworkingExample(t *testing.T) {
	t.Parallel()

	subID := "" // Subscription ID, leave blank if available as an Environment Var
	prefix := "terratest-net"
	expectedSubnetRange := "10.0.20.0/24"
	expectedPrivateIP := "10.0.20.5"
	expectedDnsIp01 := "10.0.0.5"
	expectedDnsIp02 := "10.0.0.6"

	// Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// Relative path to the Terraform dir
		TerraformDir: "../examples/terraform-azure-network-example",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"prefix":        prefix,
			"subnet_prefix": expectedSubnetRange,
			"private_ip":    expectedPrivateIP,
			"dns_ip_01":     expectedDnsIp01,
			"dns_ip_02":     expectedDnsIp02,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	virtualNetworkName := terraform.Output(t, terraformOptions, "virtual_network_name")
	subnetName := terraform.Output(t, terraformOptions, "subnet_name")
	publicAddressName := terraform.Output(t, terraformOptions, "public_address_name")
	nicInternalName := terraform.Output(t, terraformOptions, "network_interface_internal")
	nicExternalName := terraform.Output(t, terraformOptions, "network_interface_external")

	t.Run("VirtualNetwork", func(t *testing.T) {
		// Check the Virtual Network exists
		azure.AssertVirtualNetworkExists(t, virtualNetworkName, resourceGroupName, subID)

		// Check the Virtual Network Subnet
		actualSubnets := azure.GetVNetSubnetList(t, virtualNetworkName, resourceGroupName, subID)
		assert.Contains(t, actualSubnets, subnetName)

		// Check the Virtual Network DNS Server IPs
		actualDNSIPs := azure.GetVNetDNSServerIPList(t, virtualNetworkName, resourceGroupName, subID)
		assert.Contains(t, actualDNSIPs, expectedDnsIp01)
		assert.Contains(t, actualDNSIPs, expectedDnsIp02)
	})

	t.Run("Subnet", func(t *testing.T) {
		// Check the Subnet exists
		azure.AssertSubnetExists(t, subnetName, virtualNetworkName, resourceGroupName, subID)

		// Check Subnet Address Range
		actualSubnetRange := azure.GetSubnetRange(t, subnetName, virtualNetworkName, resourceGroupName, subID)
		assert.Equal(t, expectedSubnetRange, actualSubnetRange)
	})

	t.Run("PublicAddress", func(t *testing.T) {
		// Check the Public Address IP
		actualPublicIP := azure.GetPublicAddressIP(t, publicAddressName, resourceGroupName, subID)
		assert.True(t, len(actualPublicIP) > 0)
	})

	t.Run("NIC", func(t *testing.T) {
		// Check the Network Interface exists
		azure.AssertNetworkInterfaceExists(t, nicInternalName, resourceGroupName, subID)
		azure.AssertNetworkInterfaceExists(t, nicExternalName, resourceGroupName, subID)

		// Check the Private IP
		actualPrivateIPs := azure.GetNetworkInterfacePrivateIPs(t, nicInternalName, resourceGroupName, subID)
		assert.Contains(t, actualPrivateIPs, expectedPrivateIP)
	})

	t.Run("NIC_PublicAddress", func(t *testing.T) {
		// Check the internal network interface does NOT have a public IP
		actualPublicIPsFail := azure.GetNetworkInterfacePublicIPs(t, nicInternalName, resourceGroupName, subID)
		assert.True(t, len(actualPublicIPsFail) == 0)

		// Check the external network interface has a public IP
		actualPublicIPs := azure.GetNetworkInterfacePublicIPs(t, nicExternalName, resourceGroupName, subID)
		assert.True(t, len(actualPublicIPs) == 1)
	})

	t.Run("VirtualNetwork_Subnet_NIC_IP", func(t *testing.T) {
		// Check the private IP is in the subnet range
		checkPrivateIpInSubnet := azure.CheckPrivateIPInSubnet(t, resourceGroupName, virtualNetworkName, subnetName, expectedPrivateIP, subID)
		assert.True(t, checkPrivateIpInSubnet)
	})
}
