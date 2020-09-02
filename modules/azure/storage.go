package azure

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	"github.com/Azure/go-autorest/autorest/azure"
)

const (
	AzureEnvironmentEnvName = "AZURE_ENVIRONMENT"
)

// StorageAccountExistsE indicates whether the storage account name exactly matches; otherwise false.
func StorageAccountExistsE(storageAccountName, resourceGroupName, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return false, err
	}
	resourceGroupName, err2 := getTargetAzureResourceGroupName((resourceGroupName))
	if err2 != nil {
		return false, err2
	}
	_, err3 := GetStorageAccountClientE(subscriptionID)
	if err3 != nil {
		return false, err3
	}
	storageAccount, err4 := GetStorageAccountPropertyE(storageAccountName, resourceGroupName, subscriptionID)
	if err4 != nil {
		return false, nil
	}
	return *storageAccount.Name == storageAccountName, nil
}

// BlobContainerExistsE returns true if the container name exactly matches; otherwise false
func BlobContainerExistsE(containerName, storageAccountName, resourceGroupName, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return false, err
	}
	resourceGroupName, err2 := getTargetAzureResourceGroupName((resourceGroupName))
	if err2 != nil {
		return false, err2
	}
	client, err := GetBlobContainersClientE(subscriptionID)
	if err != nil {
		return false, err
	}
	container, err := client.Get(context.Background(), resourceGroupName, storageAccountName, containerName)
	if err != nil {
		return false, err
	}

	return containerName == *container.Name, nil
}

// StorageContainerHasPublicAccessE indicates whether a storage container has public access; otherwise false.
func StorageContainerHasPublicAccessE(containerName, storageAccountName, resourceGroupName, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return false, err
	}
	resourceGroupName, err2 := getTargetAzureResourceGroupName((resourceGroupName))
	if err2 != nil {
		return false, err2
	}
	client, err := GetBlobContainersClientE(subscriptionID)
	if err != nil {
		return false, err
	}
	container, err := client.Get(context.Background(), resourceGroupName, storageAccountName, containerName)
	if err != nil {
		return false, err
	}

	return (string(container.PublicAccess) != "None"), nil
}

// GetStorageAccountKindE returns one of Storage, StorageV2, BlobStorage, FileStorage, or BlockBlobStorage; otherwise error.
func GetStorageAccountKindE(storageAccountName, resourceGroupName, subscriptionID string) (string, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return "", err
	}
	resourceGroupName, err2 := getTargetAzureResourceGroupName((resourceGroupName))
	if err2 != nil {
		return "", err2
	}
	storageAccount, err := GetStorageAccountPropertyE(storageAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		return "", err
	}
	return string(storageAccount.Kind), nil
}

// GetStorageAccountSkuTierE returns the storage account sku tier as Standard or Premium; otherwise error.
func GetStorageAccountSkuTierE(storageAccountName, resourceGroupName, subscriptionID string) (string, error) {
	storageAccount, err := GetStorageAccountPropertyE(storageAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		return "", err
	}
	return string(storageAccount.Sku.Tier), nil
}

// GetBlobContainerE returns Blob container
func GetBlobContainerE(containerName, storageAccountName, resourceGroupName, subscriptionID string) (*storage.BlobContainer, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	resourceGroupName, err2 := getTargetAzureResourceGroupName((resourceGroupName))
	if err2 != nil {
		return nil, err2
	}
	client, err := GetBlobContainersClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	container, err := client.Get(context.Background(), resourceGroupName, storageAccountName, containerName)
	if err != nil {
		return nil, err
	}
	return &container, nil
}

// GetStorageAccountPropertyE return StorageAccount that matches the parameter.
func GetStorageAccountPropertyE(storageAccountName, resourceGroupName, subscriptionID string) (*storage.Account, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	resourceGroupName, err2 := getTargetAzureResourceGroupName((resourceGroupName))
	if err2 != nil {
		return nil, err2
	}
	client, err := GetStorageAccountClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	account, err := client.GetProperties(context.Background(), resourceGroupName, storageAccountName, "")
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetStorageAccountClientE creates a storage account client.
func GetStorageAccountClientE(subscriptionID string) (*storage.AccountsClient, error) {
	storageAccountClient := storage.NewAccountsClient(os.Getenv(AzureSubscriptionID))
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	storageAccountClient.Authorizer = *authorizer
	return &storageAccountClient, nil
}

// GetBlobContainersClientE creates a storage container client.
func GetBlobContainersClientE(subscriptionID string) (*storage.BlobContainersClient, error) {
	blobContainerClient := storage.NewBlobContainersClient(os.Getenv(AzureSubscriptionID))
	authorizer, err := NewAuthorizer()

	if err != nil {
		return nil, err
	}
	blobContainerClient.Authorizer = *authorizer
	return &blobContainerClient, nil
}

// GetStorageUriSuffix returns the proper storage URI suffix for the configured Azure environment
func GetStorageUriSuffix() (*string, error) {
	envName := os.Getenv(AzureEnvironmentEnvName)
	if envName == "" {
		envName = "AzurePublicCloud"
	}
	env, err := azure.EnvironmentFromName(envName)
	if err != nil {
		return nil, err
	}
	return &env.StorageEndpointSuffix, nil
}

// GetStorageAccountPrimaryBlobEndpointE gets the storage account blob endpoint as URI string; otherwise error.
func GetStorageAccountPrimaryBlobEndpointE(storageAccountName, resourceGroupName, subscriptionID string) (string, error) {
	storageAccount, err := GetStorageAccountPropertyE(storageAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return *storageAccount.AccountProperties.PrimaryEndpoints.Blob, nil
}

// BuildStorageDNSStringE builds and returns the storage account dns string if the storage account exists; otherwise error.
func BuildStorageDNSStringE(storageAccountName, resourceGroupName, subscriptionID string) (string, error) {
	retval, err := StorageAccountExistsE(storageAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		return "", err
	}
	if retval {
		storageSuffix, _ := GetStorageUriSuffix()
		return fmt.Sprintf("https://%s.blob.%s/", storageAccountName, *storageSuffix), nil
	} else {
		return "", StorageAccountNameNotFound{}
	}
}
