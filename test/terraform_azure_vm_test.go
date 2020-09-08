// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestTerraformAzureVmExample(t *testing.T) {
	t.Parallel()

	// Subscription ID, leave blank if available as an Environment Var
	subID := ""

	// Initiate VM instance information
	uniquePrefix := "terratest"
	uniqueSuffix := strings.ToLower(random.UniqueId())
	expectedResourceGroup := fmt.Sprintf("%s-rg-%s", uniquePrefix, uniqueSuffix)
	expectedName := fmt.Sprintf("%s-vm-%s", uniquePrefix, uniqueSuffix)
	expectedVnetName := fmt.Sprintf("%s-vnet-%s", uniquePrefix, uniqueSuffix)
	expectedSubnetName := fmt.Sprintf("%s-snt-%s-int", uniquePrefix, uniqueSuffix)
	expectedNicName := fmt.Sprintf("%s-nic-%s", uniquePrefix, uniqueSuffix)
	//expectedIPConfigName := fmt.Sprintf("%s-ip-%s", uniquePrefix, uniqueSuffix)
	expectedAvsName := fmt.Sprintf("%s-avs-%s", uniquePrefix, uniqueSuffix)
	expectedOSDiskName := fmt.Sprintf("%s-vm-%s-osdisk", uniquePrefix, uniqueSuffix)
	expectedDisk01Name := fmt.Sprintf("%s-vm-%s-disk01", uniquePrefix, uniqueSuffix)

	expectedResourceGroup = "terratest-rg-01"
	expectedName = "terratest-vm-01"
	expectedVnetName = "terratest-vnet-01"
	expectedSubnetName = "terratest-snt-01-internal"
	expectedNicName = "terratest-nic-01"
	//expectedIPConfigName = "terratest-pip-01-internal"
	expectedAvsName = "terratest-avs-01"
	expectedOSDiskName = "terratest-vm-01-osdisk"
	expectedDisk01Name = "terratest-vm-01-disk-01"

	expectedVmAdminUser := "testadmin"
	expectedVmSize := "Standard_DS1_v2"
	var expectedAvsFaultDomainCount int32
	expectedAvsFaultDomainCount = 2
	expectedImageSKU := "2016-Datacenter"
	expectedImageVersion := "latest"

	expectedOSDiskTypeString := "Standard_LRS"
	expectedManagedDiskCount := 1
	expectedNicCount := 1
	expectedPrivateIPAddress := "10.0.17.4"
	expectedSubnetAddressRange := "10.0.17.0/24"
	expectedPublicAddressName := "terratest-pip-01-external"

	// Configure Terraform setting up a path to Terraform code.
	/* terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-vm-example",

		// Variables to pass to our Terraform code using -var options
		// "username" and "password" should not be passed from here in a production scenario.
		Vars: map[string]interface{}{
			"resource_group_name": expectedResourceGroup,
			"user_name":           expectedVmAdminUser,
			"vm_name":             expectedName,
			"vm_size":             expectedVmSize,
			"vm_license":          "Windows_Server",
			"vm_image_publisher":  "MicrosoftWindowsServer",
			"vm_image_offer":      "WindowsServer",
			"vm_image_sku":        expectedImageSKU,
			"vm_image_version":    expectedImageVersion,
			"disk_type":           "Standard_LRS",
			"managed_disk_count":  1,
			"disk_01_size":        10,
			"nic_count":           1,
			"vnet_name":           expectedVnetName,
			"subnet_name":         expectedSubnetName,
			"nic_name":            expectedNicName,
			"ip_config_name":      expectedPipName,
			"avs_name":            expectedAvsName,
			"osdisk_name":         expectedOSDiskName,
			"disk_01_name":        expectedDisk01Name,
		},

		// Environment variables to set when running Terraform
		// EnvVars: map[string]string{
		//  	"ARM_SUBSCRIPTION_ID": subID,
		// },
	} */

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	// defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	// terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	vmName := expectedName                     // terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := expectedResourceGroup // terraform.Output(t, terraformOptions, "resource_group_name")
	expectedVMSize := compute.VirtualMachineSizeTypes(expectedVmSize)

	t.Run("Strategies", func(t *testing.T) {
		t.Parallel()

		// Check the VM Size directly
		actualVMSize := azure.GetSizeOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedVMSize, actualVMSize)

		// Check the VM Size by object ref
		vmRef := azure.GetVirtualMachine(t, vmName, resourceGroupName, subID)
		actualVMSize = vmRef.HardwareProfile.VMSize
		assert.Equal(t, expectedVMSize, actualVMSize)

		// Check the VM Size by instance getter
		vmInstance := azure.GetVirtualMachineInstance(t, vmName, resourceGroupName, subID)
		actualVMSize = vmInstance.GetVirtualMachineSize()
		assert.Equal(t, expectedVMSize, actualVMSize)
	})

	t.Run("VirtualMachines", func(t *testing.T) {
		t.Parallel()

		// Get a list of all VMs and confirm one (or more) VMs exist
		vmList := azure.GetVirtualMachinesList(t, resourceGroupName, subID)
		assert.True(t, len(vmList) > 0)
		assert.Contains(t, vmList, expectedName)

		// Get all VMs by ref (warning: pointer deref painc if vm is not in list!)
		vmsByRef := azure.GetVirtualMachines(t, resourceGroupName, subID)
		assert.True(t, len(*vmsByRef) > 0)
		thisVm := (*vmsByRef)[expectedName]
		assert.Equal(t, expectedVMSize, thisVm.HardwareProfile.VMSize)
	})

	t.Run("Information", func(t *testing.T) {
		t.Parallel()

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
		t.Parallel()

		// Check the Availability Set
		actualAvsName := azure.GetAvailabilitySetIDOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.True(t, strings.EqualFold(expectedAvsName, actualAvsName))

		// Check the Availability set fault domain counts
		actualAvsFaultDomainCount := azure.GetFaultDomainCountOfAvailabilitySet(t, expectedAvsName, resourceGroupName, subID)
		assert.Equal(t, expectedAvsFaultDomainCount, actualAvsFaultDomainCount)

		actualVMsInAvs := azure.GetVMsOfAvailabilitySet(t, expectedAvsName, resourceGroupName, subID)
		assert.Contains(t, actualVMsInAvs, strings.ToUpper(vmName))
	})

	t.Run("Disk", func(t *testing.T) {
		t.Parallel()

		// Check the OS Disk name
		actualOSDiskName := azure.GetOsDiskNameOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedOSDiskName, actualOSDiskName)

		// Check the OS Disk Storage Type (SKU)
		expectedOSDiskType := compute.DiskStorageAccountTypes(expectedOSDiskTypeString)
		actualOSDiskType := azure.GetTypeOfDisk(t, expectedOSDiskName, resourceGroupName, subID)
		assert.Equal(t, expectedOSDiskType, actualOSDiskType)

		// Check the Managed Disk count
		actualManagedDiskCount := azure.GetManagedDiskCountOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedManagedDiskCount, actualManagedDiskCount)

		// Check the VM Managed Disk exists in the list of all VM Managed Disks
		actualManagedDiskNames := azure.GetManagedDiskNamesOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Contains(t, actualManagedDiskNames, expectedDisk01Name)
	})

	t.Run("NetworkInterface", func(t *testing.T) {
		t.Parallel()

		// Check the Network Interface count
		actualNicCount := azure.GetNicCountOfVirtualMachine(t, vmName, resourceGroupName, subID)
		assert.Equal(t, expectedNicCount, actualNicCount)

		// Check the VM Network Interface exists in the list of all VM Network Interfaces
		actualNics := azure.GetVMNetworkInterfacesList(t, vmName, resourceGroupName, subID)
		assert.Contains(t, actualNics, expectedNicName)

		// Check the Private IP
		actualNicIPs := azure.GetNetworkInterfacePrivateIPs(t, expectedNicName, resourceGroupName, subID)
		assert.Contains(t, actualNicIPs, expectedPrivateIPAddress)

		// Check the Public IP exists
		actualPublicIP := azure.GetPublicAddressIP(t, expectedPublicAddressName, resourceGroupName, subID)
		assert.NotNil(t, actualPublicIP)
	})

	t.Run("Vnet&Subnet", func(t *testing.T) {
		t.Parallel()

		// Check the VM Subnet IP Range
		actualSubnetAddressRange := azure.GetSubnetIPRange(t, expectedVnetName, expectedSubnetName, resourceGroupName, subID)
		assert.Equal(t, expectedSubnetAddressRange, actualSubnetAddressRange)

		// Check the Subnet exists in the Virtual Network Subnets
		actualVnetSubnets := azure.GetVNetSubnets(t, resourceGroupName, expectedVnetName, subID)
		assert.NotNil(t, actualVnetSubnets[expectedVnetName])

		// Check the Private IP is in the Subnet Range
		actualVMNicIPInSubnet := azure.CheckIPInSubnet(t, resourceGroupName, expectedVnetName, expectedSubnetName, expectedPrivateIPAddress, subID)
		assert.True(t, actualVMNicIPInSubnet)
	})
}
