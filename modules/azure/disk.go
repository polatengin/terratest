package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// AssertDiskExists checks for an Azure Managed Disk
func AssertDiskExists(t testing.TestingT, diskName string, resGroupName string, subscriptionID string) {
	err := AssertDiskExistsE(diskName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertDiskExistsE checks for an Azure Managed Disk with error
func AssertDiskExistsE(diskName string, resGroupName string, subscriptionID string) error {
	// Get the Network Interface client
	_, err := GetDiskE(diskName, resGroupName, subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

// GetTypeOfDisk gets the Type of the given Managed Disk
func GetTypeOfDisk(t testing.TestingT, diskName string, resGroupName string, subscriptionID string) compute.DiskStorageAccountTypes {
	avsFaultDomainCount, err := GetTypeOfDiskE(diskName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return avsFaultDomainCount
}

// GetTypeOfDiskE gets the Type of the given Managed Disk with error
func GetTypeOfDiskE(diskName string, resGroupName string, subscriptionID string) (compute.DiskStorageAccountTypes, error) {
	// Get disk object
	disk, err := GetDiskE(diskName, resGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return disk.Sku.Name, nil
}

// GetDiskE gets a Disk in the specified Azure Resource Group
func GetDiskE(diskName string, resGroupName string, subscriptionID string) (*compute.Disk, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetDiskClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Disk
	disk, err := client.Get(context.Background(), resGroupName, diskName)
	if err != nil {
		return nil, err
	}

	return &disk, nil
}

// GetDiskClientE creates a new Disk client in the specified Azure Subscription
func GetDiskClientE(subscriptionID string) (*compute.DisksClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Disk client
	client := compute.NewDisksClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}
