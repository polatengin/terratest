package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetDiskType gets the Type of the given Azure Disk
func GetDiskType(t testing.TestingT, diskName string, resGroupName string, subscriptionID string) compute.DiskStorageAccountTypes {
	diskType, err := GetDiskTypeE(diskName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return diskType
}

// GetDiskTypeE gets the Type of the given Azure Disk
func GetDiskTypeE(diskName string, resGroupName string, subscriptionID string) (compute.DiskStorageAccountTypes, error) {
	client, err := GetDiskClientE(subscriptionID)
	if err != nil {
		return "", err
	}

	disk, err := client.Get(context.Background(), resGroupName, diskName)
	if err != nil {
		return "", err
	}

	return disk.Sku.Name, nil
}

// GetDiskE gets a Disk in the specified resource group
func GetDiskE(resGroupName string, diskName string, subscriptionID string) (*compute.Disk, error) {
	client, err := GetDiskClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	disk, err := client.Get(context.Background(), resGroupName, diskName)
	if err != nil {
		return nil, err
	}

	return &disk, nil
}

// GetDiskClientE creates a new Disk client
func GetDiskClientE(subscriptionID string) (*compute.DisksClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	diskClient := compute.NewDisksClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	diskClient.Authorizer = *authorizer
	return &diskClient, nil
}
