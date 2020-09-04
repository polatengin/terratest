package azure

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetVMNetworkInterfacesList gets a list of Network Interfaces for a given Azure Virtual Machine
func GetVMNetworkInterfacesList(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) []string {
	nicList, err := GetVMNetworkInterfacesListE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return nicList
}

// GetVMNetworkInterfacesListE gets a list of Network Interfaces for a given Azure Virtual Machine
func GetVMNetworkInterfacesListE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) ([]string, error) {
	nics := []string{}

	theVM, err := GetVirtualMachineE(vmName, resGroupName, "")
	if err != nil {
		return nil, err
	}

	theNICs := *theVM.NetworkProfile.NetworkInterfaces
	if len(theNICs) == 0 {
		return nil, errors.New("No network interface attached to this Virtual Machine")
	}

	for _, nic := range theNICs {
		tmp := strings.Split(*nic.ID, "/")
		if len(tmp) > 0 {
			nics = append(nics, tmp[len(tmp)-1])
		}
	}
	return nics, nil
}

// GetNicCountOfVirtualMachine gets the Managed Disk count of the given Azure Virtual Machine
func GetNicCountOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) int {
	nicCount, err := GetNicCountOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return nicCount
}

// GetNicCountOfVirtualMachineE gets the Managed Disk count of the given Azure Virtual Machine
func GetNicCountOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (int, error) {
	nicCount := 0

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nicCount, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return nicCount, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return nicCount, err
	}

	return len(*vm.NetworkProfile.NetworkInterfaces), nil
}

// GetManagedDiskNamesOfVirtualMachine gets the list of Managed Disk name of the given Azure Virtual Machine
func GetManagedDiskNamesOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) []string {
	mngDiskNames, err := GetManagedDiskNamesOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return mngDiskNames
}

// GetManagedDiskNamesOfVirtualMachineE gets the Managed Disk count of the given Azure Virtual Machine
func GetManagedDiskNamesOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) ([]string, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return nil, err
	}

	vmDisks := *vm.StorageProfile.DataDisks
	diskNames := []string{}

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

// GetManagedDiskCountOfVirtualMachineE gets the Managed Disk count of the given Azure Virtual Machine
func GetManagedDiskCountOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (int, error) {
	mngDiskCount := 0

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return mngDiskCount, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return mngDiskCount, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return mngDiskCount, err
	}

	return len(*vm.StorageProfile.DataDisks), nil
}

// GetOsDiskNameOfVirtualMachine gets the Availability Set ID of the given Azure Virtual Machine
func GetOsDiskNameOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) string {
	osDiskName, err := GetOsDiskNameOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return osDiskName
}

// GetOsDiskNameOfVirtualMachineE gets the Availability Set ID of the given Azure Virtual Machine
func GetOsDiskNameOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (string, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return "", err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return "", err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return "", err
	}

	return *vm.StorageProfile.OsDisk.Name, nil
}

// VMImage represents the storage image for the given Azure Virtual Machine
type VMImage struct {
	Publisher string
	Offer     string
	SKU       string
	Version   string
}

// GetImageOfVirtualMachine gets the Availability Set ID of the given Azure Virtual Machine
func GetImageOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) VMImage {
	adminUser, err := GetImageOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return adminUser
}

// GetImageOfVirtualMachineE gets the Availability Set ID of the given Azure Virtual Machine
func GetImageOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (VMImage, error) {
	vmImage := VMImage{}

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return vmImage, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return vmImage, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return vmImage, err
	}

	vmImage.Publisher = *vm.StorageProfile.ImageReference.Publisher
	vmImage.Offer = *vm.StorageProfile.ImageReference.Offer
	vmImage.SKU = *vm.StorageProfile.ImageReference.Sku
	vmImage.Version = *vm.StorageProfile.ImageReference.Version

	return vmImage, nil
}

// GetAvailabilitySetIDOfVirtualMachine gets the Availability Set ID of the given Azure Virtual Machine
func GetAvailabilitySetIDOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) string {
	adminUser, err := GetAvailabilitySetIDOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return adminUser
}

// GetAvailabilitySetIDOfVirtualMachineE gets the Availability Set ID of the given Azure Virtual Machine
func GetAvailabilitySetIDOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (string, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return "", err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return "", err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return "", err
	}

	avsID := string(*vm.AvailabilitySet.ID)
	tmp := strings.Split(avsID, "/")
	if !(len(tmp) > 0) {
		return "", errors.New("Availability set ID not found")
	}

	return tmp[len(tmp)-1], nil
}

// GetAdminUserOfVirtualMachine gets the admin username of the given Azure Virtual Machine
func GetAdminUserOfVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) string {
	adminUser, err := GetAdminUserOfVirtualMachineE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return adminUser
}

// GetAdminUserOfVirtualMachineE gets the admin username of the given Azure Virtual Machine
func GetAdminUserOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (string, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return "", err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return "", err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
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

// GetSizeOfVirtualMachineE gets the size type of the given Azure Virtual Machine
func GetSizeOfVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (compute.VirtualMachineSizeTypes, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return "", err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return "", err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
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

// GetTagsForVirtualMachineE gets the tags of the given Virtual Machine as a map
func GetTagsForVirtualMachineE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) (map[string]string, error) {
	// Setup a blank map to populate and return
	tags := make(map[string]string)

	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return tags, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return tags, err
	}

	// Get the details of the target virtual machine
	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return tags, err
	}

	// Range through existing tags and populate above map accordingly
	for k, v := range vm.Tags {
		tags[k] = *v
	}

	return tags, nil
}

// GetVirtualMachinesList gets a list of all virtual machines in the specified resource group
func GetVirtualMachinesList(t testing.TestingT, resGroupName string, subscriptionID string) []string {
	vms, err := GetVirtualMachinesListE(resGroupName, subscriptionID)
	require.NoError(t, err)
	return vms
}

// GetVirtualMachinesListE gets all virtual machines in the specified resource group
func GetVirtualMachinesListE(resourceGroupName string, subscriptionID string) ([]string, error) {
	vmDetails := []string{}

	vmClient, err := GetVirtualMachineClient(subscriptionID)
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

// GetVirtualMachines gets all virtual machines in the specified resource group
func GetVirtualMachines(t testing.TestingT, resGroupName string, subscriptionID string) *map[string]compute.VirtualMachineProperties {
	vms, err := GetVirtualMachinesE(resGroupName, subscriptionID)
	require.NoError(t, err)
	return vms
}

// GetVirtualMachinesE gets all virtual machines in the specified resource group
func GetVirtualMachinesE(resourceGroupName string, subscriptionID string) (*map[string]compute.VirtualMachineProperties, error) {
	vmClient, err := GetVirtualMachineClient(subscriptionID)
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
// Option to retrieve the VM Instance and related getter //
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

// GetVirtualMachineInstanceE gets a local virtual machine instance in the specified resource group
func GetVirtualMachineInstanceE(vmName string, resGroupName string, subscriptionID string) (*Instance, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return nil, err
	}

	return &Instance{&vm}, nil
}

// GetVirtualMachineSize gets vm size
func (vm *Instance) GetVirtualMachineSize() compute.VirtualMachineSizeTypes {
	return vm.VirtualMachineProperties.HardwareProfile.VMSize
}

// ******************************** //
// Option to retrieve the VM Object //
// ******************************** //

// GetVirtualMachine gets a virtual machine in the specified resource group
func GetVirtualMachine(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) *compute.VirtualMachine {
	vm, err := GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	require.NoError(t, err)
	return vm
}

// GetVirtualMachineE gets a virtual machine in the specified resource group
func GetVirtualMachineE(vmName string, resGroupName string, subscriptionID string) (*compute.VirtualMachine, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	vm, err := vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return nil, err
	}

	return &vm, nil
}

// AssertVirtualMachineExists checks if the given VM exists in the given subscription and fail the test if it does not
func AssertVirtualMachineExists(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) {
	err := AssertVirtualMachineExistsE(t, vmName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertVirtualMachineExistsE checks if the given VM exists in the given subscription and fail the test if it does not
func AssertVirtualMachineExistsE(t testing.TestingT, vmName string, resGroupName string, subscriptionID string) error {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return err
	}

	// Create a VM client
	vmClient, err := GetVirtualMachineClient(subscriptionID)
	if err != nil {
		return err
	}

	_, err = vmClient.Get(context.Background(), resGroupName, vmName, compute.InstanceView)
	if err != nil {
		return err
	}

	return nil
}

// GetVirtualMachineClient is a helper function that will setup an Azure Virtual Machine client on your behalf
func GetVirtualMachineClient(subscriptionID string) (*compute.VirtualMachinesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a VM client
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	vmClient.Authorizer = *authorizer

	return &vmClient, nil
}

// ********************************** //
// VM attached resources and clients  //
// ********************************** //

// PrettyPrint will print the contents of the obj
func PrettyPrint(data interface{}) string {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Sprintln(err)
	}
	return fmt.Sprintf("%s \n", p)
}
