package azure

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// AssertVirtualMachineExists checks if the given VM exists in the given subscription
func AssertVirtualMachineExists(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) {
	err := AssertVirtualMachineExistsE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertVirtualMachineExistsE checks if the given VM exists in the given subscription and returns an error if not found
func AssertVirtualMachineExistsE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) error {
	// Get VM Object
	_, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

// GetNicNamesForVirtualMachine gets a list of Network Interfaces for a given Azure Virtual Machine
func GetNicNamesForVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) []string {
	nicList, err := GetNicNamesForVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return nicList
}

// GetNicNamesForVirtualMachineE gets a list of Network Interfaces for a given Azure Virtual Machine with error
func GetNicNamesForVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) ([]string, error) {
	nics := []string{}

	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return nics, err
	}

	vmNICs := *vm.NetworkProfile.NetworkInterfaces
	if len(vmNICs) == 0 {
		// No VM NICs attached is still valid but returning a meaningful error
		return nics, errors.New("No network interface attached to this Virtual Machine")
	}

	// Get the attached NIC names
	for _, nic := range vmNICs {
		nics = append(nics, GetNameFromResourceID(*nic.ID))
	}
	return nics, nil
}

// GetNicCountOfVirtualMachine gets the Network Interface count of the given Azure Virtual Machine
func GetNicCountOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) int {
	nicCount, err := GetNicCountOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return nicCount
}

// GetNicCountOfVirtualMachineE gets the Network Interface count of the given Azure Virtual Machine with error
func GetNicCountOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (int, error) {
	nicCount := 0

	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return nicCount, err
	}

	return len(*vm.NetworkProfile.NetworkInterfaces), nil
}

// GetManagedDiskNamesOfVirtualMachine gets the list of Managed Disk names of the given Azure Virtual Machine
func GetManagedDiskNamesOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) []string {
	diskNames, err := GetManagedDiskNamesOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return diskNames
}

// GetManagedDiskNamesOfVirtualMachineE gets the list of Managed Disk names of the given Azure Virtual Machine with error
func GetManagedDiskNamesOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) ([]string, error) {
	diskNames := []string{}

	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return diskNames, err
	}

	// Get VM Attached Disks
	vmDisks := *vm.StorageProfile.DataDisks
	for _, v := range vmDisks {
		diskNames = append(diskNames, *v.Name)
	}

	return diskNames, nil
}

// GetManagedDiskCountOfVirtualMachine gets the Managed Disk count of the given Azure Virtual Machine
func GetManagedDiskCountOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) int {
	mngDiskCount, err := GetManagedDiskCountOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return mngDiskCount
}

// GetManagedDiskCountOfVirtualMachineE gets the Managed Disk count of the given Azure Virtual Machine with error
func GetManagedDiskCountOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (int, error) {
	mngDiskCount := -1

	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return mngDiskCount, err
	}

	return len(*vm.StorageProfile.DataDisks), nil
}

// GetOsDiskNameOfVirtualMachine gets the OS Disk Name of the given Azure Virtual Machine
func GetOsDiskNameOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) string {
	osDiskName, err := GetOsDiskNameOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return osDiskName
}

// GetOsDiskNameOfVirtualMachineE gets the OS Disk Name of the given Azure Virtual Machine with error
func GetOsDiskNameOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (string, error) {
	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return *vm.StorageProfile.OsDisk.Name, nil
}

// GetAvailabilitySetIDOfVirtualMachine gets the Availability Set ID of the given Azure Virtual Machine
func GetAvailabilitySetIDOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) string {
	adminUser, err := GetAvailabilitySetIDOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return adminUser
}

// GetAvailabilitySetIDOfVirtualMachineE gets the Availability Set ID of the given Azure Virtual Machine with error
func GetAvailabilitySetIDOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (string, error) {
	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return GetNameFromResourceID(*vm.AvailabilitySet.ID), nil
}

// VMImage represents the storage image for the given Azure Virtual Machine
type VMImage struct {
	Publisher string
	Offer     string
	SKU       string
	Version   string
}

// GetImageOfVirtualMachine gets the VM Image of the given Azure Virtual Machine
func GetImageOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) VMImage {
	adminUser, err := GetImageOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return adminUser
}

// GetImageOfVirtualMachineE gets the VM Image  of the given Azure Virtual Machine with error
func GetImageOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (VMImage, error) {
	vmImage := VMImage{}

	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return vmImage, err
	}

	vmImage.Publisher = *vm.StorageProfile.ImageReference.Publisher
	vmImage.Offer = *vm.StorageProfile.ImageReference.Offer
	vmImage.SKU = *vm.StorageProfile.ImageReference.Sku
	vmImage.Version = *vm.StorageProfile.ImageReference.Version

	return vmImage, nil
}

// GetAdminUserOfVirtualMachine gets the admin username of the given Azure Virtual Machine
func GetAdminUserOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) string {
	adminUser, err := GetAdminUserOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return adminUser
}

// GetAdminUserOfVirtualMachineE gets the admin username of the given Azure Virtual Machine with error
func GetAdminUserOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (string, error) {
	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return string(*vm.OsProfile.AdminUsername), nil
}

// GetSizeOfVirtualMachine gets the size type of the given Azure Virtual Machine
func GetSizeOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) compute.VirtualMachineSizeTypes {
	size, err := GetSizeOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return size
}

// GetSizeOfVirtualMachineE gets the size type of the given Azure Virtual Machine with error
func GetSizeOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (compute.VirtualMachineSizeTypes, error) {
	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return vm.VirtualMachineProperties.HardwareProfile.VMSize, nil
}

// GetTagsForVirtualMachine gets the tags of the given Virtual Machine as a map
func GetTagsForVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) map[string]string {
	tags, err := GetTagsForVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return tags
}

// GetTagsForVirtualMachineE gets the tags of the given Virtual Machine as a map with error
func GetTagsForVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (map[string]string, error) {
	// Setup a blank map to populate and return
	tags := make(map[string]string)

	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return nil, err
	}

	// Range through existing tags and populate above map accordingly
	for k, v := range vm.Tags {
		tags[k] = *v
	}

	return tags, nil
}

// ***************************************************** //
// Get multiple Virtual Machines from a Resource Group
// ***************************************************** //

// GetVirtualMachinesList gets a list of all virtual machines in the specified resource group
func GetVirtualMachinesList(t testing.TestingT, resGroupName string, subscriptionID string) []string {
	vms, err := GetVirtualMachinesListE(resGroupName, subscriptionID)
	require.NoError(t, err)
	return vms
}

// GetVirtualMachinesListE gets a list of all virtual machines in the specified resource group with error
func GetVirtualMachinesListE(resourceGroupName string, subscriptionID string) ([]string, error) {
	vmDetails := []string{}

	vmClient, err := GetVirtualMachineClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	vms, err := vmClient.List(context.Background(), resourceGroupName)
	if err != nil {
		return nil, err
	}

	for _, v := range vms.Values() {
		vmDetails = append(vmDetails, *v.Name)
	}
	return vmDetails, nil
}

// GetVirtualMachines gets all virtual machine objects in the specified resource group
func GetVirtualMachines(t testing.TestingT, resGroupName string, subscriptionID string) *map[string]compute.VirtualMachineProperties {
	vms, err := GetVirtualMachinesE(resGroupName, subscriptionID)
	require.NoError(t, err)
	return vms
}

// GetVirtualMachinesE gets all virtual machine objects in the specified resource group with error
func GetVirtualMachinesE(resourceGroupName string, subscriptionID string) (*map[string]compute.VirtualMachineProperties, error) {
	vmClient, err := GetVirtualMachineClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	vms, err := vmClient.List(context.Background(), resourceGroupName)
	if err != nil {
		return nil, err
	}

	vmDetails := make(map[string]compute.VirtualMachineProperties)
	for _, v := range vms.Values() {
		machineName := v.Name
		vmProperties := v.VirtualMachineProperties
		vmDetails[*machineName] = *vmProperties
	}
	return &vmDetails, nil
}

// ***************************************************** //
// Get VM Instance and sample property getter
// ***************************************************** //

// Instance of the VM
type Instance struct {
	*compute.VirtualMachine
}

// GetVirtualMachineInstance gets a local virtual machine instance in the specified resource group
func GetVirtualMachineInstance(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) *Instance {
	vm, err := GetVirtualMachineInstanceE(vmName, resGroupName, subscriptionID)
	require.NoError(t, err)
	return vm
}

// GetVirtualMachineInstanceE gets a local virtual machine instance in the specified resource group with error
func GetVirtualMachineInstanceE(vmName string, resGroupName string, subscriptionID string) (*Instance, error) {
	// Get VM Object
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return nil, err
	}

	return &Instance{vm}, nil
}

// GetVirtualMachineSize gets vm size
func (vm *Instance) GetVirtualMachineSize() compute.VirtualMachineSizeTypes {
	return vm.VirtualMachineProperties.HardwareProfile.VMSize
}

// ******************************** //
// Get the base VM Object
// ******************************** //

// GetVirtualMachine gets a virtual machine in the specified resource group
func GetVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) *compute.VirtualMachine {
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	require.NoError(t, err)
	return vm
}

// GetVirtualMachineE gets a Virtual Machine in the specified Azure Resource Group
func GetVirtualMachineE(vmName string, resGroupName string, subscriptionID string) (*compute.VirtualMachine, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetVirtualMachineClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	vm, err := client.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return nil, err
	}

	return &vm, nil
}

// GetVirtualMachineClientE creates a Azure Virtual Machine client in the specified Azure Subscription
func GetVirtualMachineClientE(subscriptionID string) (*compute.VirtualMachinesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the VM client
	client := compute.NewVirtualMachinesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}
