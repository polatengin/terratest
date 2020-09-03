package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"

	"github.com/stretchr/testify/require"
)

// GetBlobContainerClient is a helper function that will setup a Bloc Container from Azure Storage Account client on your behalf
// containerName - required to find the Blob Container
// resourceName - required to find the Storage Account
// resGroupName - use an empty string if you have the AZURE_RES_GROUP_NAME environment variable set
// subscriptionId - use an empty string if you have the ARM_SUBSCRIPTION_ID environment variable set
func GetBlobContainerClient(t *testing.T, containerName string, resourceName string, resGroupName string, subscriptionID string) *storage.BlobContainer {
	resource, err := getBlobContainerClientE(containerName, resourceName, resGroupName, subscriptionID)

	require.NoError(t, err)

	return resource
}

func getBlobContainerClientE(containerName string, resourceName string, resGroupName string, subscriptionID string) (*storage.BlobContainer, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	managedServicesClient := storage.NewBlobContainersClient(subscriptionID)
	authorizer, err := NewAuthorizer()

	if err != nil {
		return nil, err
	}

	managedServicesClient.Authorizer = *authorizer

	resource, err := managedServicesClient.Get(context.Background(), resGroupName, resourceName, containerName)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
