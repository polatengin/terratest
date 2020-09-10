// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureVmExample(t *testing.T) {
	t.Parallel()

	subID := "" // Subscription ID, leave blank if available as an Environment Var
	prefix := "terratest-vm"
	expectedVmAdminUser := "testadmin"
	expectedVMSize := "Standard_DS1_v2"
	expectedImageSKU := "2016-Datacenter"
	expectedImageVersion := "latest"
	expectedDiskType := "Standard_LRS"
	expectedSubnetAddressRange := "10.0.17.0/24"
	expectedPrivateIPAddress := "10.0.17.4"
	var expectedAvsFaultDomainCount int32 = 2
	expectedManagedDiskCount := 1
	expectedNicCount := 1

	// Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-vm-example",

		// Variables to pass to our Terraform code using -var options
		// "username" and "password" should not be passed from here in a production scenario.
		Vars: map[string]interface{}{
			"prefix":           prefix,
			"user_name":        expectedVmAdminUser,
			"vm_size":          expectedVMSize,
			"vm_image_sku":     expectedImageSKU,
			"vm_image_version": expectedImageVersion,
			"disk_type":        expectedDiskType,
			"private_ip":       expectedPrivateIPAddress,
			"subnet_prefix":    expectedSubnetAddressRange,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	vnetName := terraform.Output(t, terraformOptions, "virtual_network_name")
	subnetName := terraform.Output(t, terraformOptions, "subnet_name")
	publicIpName := terraform.Output(t, terraformOptions, "public_ip_name")
	nicName := terraform.Output(t, terraformOptions, "network_interface_name")
	avsName := terraform.Output(t, terraformOptions, "availability_set_name")
	osDiskName := prefix + "-osdisk"
	diskName := terraform.Output(t, terraformOptions, "managed_disk_name")

	t.Run("Strategies", func(t *testing.T) {
		// Check the VM Size directly
		actualVMSize := azure.GetSizeOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedVMSize, string(actualVMSize))

		// Check the VM Size by object ref
		vmRef := azure.GetVirtualMachine(t, vmName, resourceGroupName, subID)
		actualVMSize = vmRef.HardwareProfile.VMSize
		assert.Equal(t, expectedVMSize, string(actualVMSize))

		// Check the VM Size by instance getter
		vmInstance := azure.GetVirtualMachineInstance(t, vmName, resourceGroupName, subID)
		actualVMSize = vmInstance.GetVirtualMachineSize()
		assert.Equal(t, expectedVMSize, string(actualVMSize))
	})

	t.Run("VirtualMachines", func(t *testing.T) {
		// Get a list of all VMs and confirm one (or more) VMs exist
		vmList := azure.GetVirtualMachinesList(t, resourceGroupName, subID)
		assert.True(t, len(vmList) > 0)
		assert.Contains(t, vmList, vmName)

		// Get all VMs by ref (warning: pointer deref painc if vm is not in list!)
		vmsByRef := azure.GetVirtualMachines(t, resourceGroupName, subID)
		assert.True(t, len(*vmsByRef) > 0)
		thisVm := (*vmsByRef)[vmName]
		assert.Equal(t, expectedVMSize, string(thisVm.HardwareProfile.VMSize))
	})

	t.Run("Information", func(t *testing.T) {
		// Check the Virtual Machine exists
		azure.AssertVirtualMachineExists(t, vmName, resourceGroupName, subID)

		// Check the Admin User
		actualVmAdminUser := azure.GetAdminUserOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedVmAdminUser, actualVmAdminUser)

		// Check the Storage Image reference
		actualImage := azure.GetImageOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedImageSKU, actualImage.SKU)
		assert.Equal(t, expectedImageVersion, actualImage.Version)
	})

	t.Run("AvailablitySet", func(t *testing.T) {
		// Check the Availability Set
		actualAvsName := azure.GetAvailabilitySetIDOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.True(t, strings.EqualFold(avsName, actualAvsName))

		// Check the Availability set fault domain counts
		actualAvsFaultDomainCount := azure.GetFaultDomainCountOfAvailabilitySet(t, avsName, resourceGroupName, subID)
		assert.Equal(t, expectedAvsFaultDomainCount, actualAvsFaultDomainCount)

		actualVMsInAvs := azure.GetVMsOfAvailabilitySet(t, avsName, resourceGroupName, subID)
		assert.Contains(t, actualVMsInAvs, strings.ToUpper(vmName))
	})

	t.Run("Disk", func(t *testing.T) {
		// Check the OS Disk name
		actualOSDiskName := azure.GetOsDiskNameOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, osDiskName, actualOSDiskName)

		// Check the OS Disk Storage Type (SKU)
		actualOSDiskType := azure.GetTypeOfDisk(t, osDiskName, resourceGroupName, subID)
		assert.Equal(t, expectedDiskType, string(actualOSDiskType))

		// Check the Managed Disk count
		actualManagedDiskCount := azure.GetManagedDiskCountOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedManagedDiskCount, actualManagedDiskCount)

		// Check the VM Managed Disk exists in the list of all VM Managed Disks
		actualManagedDiskNames := azure.GetManagedDiskNamesOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Contains(t, actualManagedDiskNames, diskName)
	})

	t.Run("NetworkInterface", func(t *testing.T) {
		// Check the Network Interface count
		actualNicCount := azure.GetNicCountOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedNicCount, actualNicCount)

		// Check the VM Network Interface exists in the list of all VM Network Interfaces
		actualNics := azure.GetNicNamesForVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Contains(t, actualNics, nicName)

		// Check the Private IP
		actualNicIPs := azure.GetNetworkInterfacePrivateIPs(t, nicName, resourceGroupName, subID)
		assert.Contains(t, actualNicIPs, expectedPrivateIPAddress)

		// Check the Public IP exists
		actualPublicIP := azure.GetPublicAddressIP(t, publicIpName, resourceGroupName, subID)
		assert.NotNil(t, actualPublicIP)
	})

	t.Run("Vnet&Subnet", func(t *testing.T) {
		// Check the VM Subnet IP Range
		actualSubnetAddressRange := azure.GetSubnetIPRange(t, vnetName, subnetName, resourceGroupName, subID)
		assert.Equal(t, expectedSubnetAddressRange, actualSubnetAddressRange)

		// Check the Subnet exists in the Virtual Network Subnets
		actualVnetSubnets := azure.GetVNetSubnets(t, resourceGroupName, vnetName, subID)
		assert.NotNil(t, actualVnetSubnets[vnetName])

		// Check the Private IP is in the Subnet Range
		actualVMNicIPInSubnet := azure.CheckPrivateIPInSubnet(t, resourceGroupName, vnetName, subnetName, expectedPrivateIPAddress, subID)
		assert.True(t, actualVMNicIPInSubnet)
	})
}
