// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformAzureStorageExample(t *testing.T) {
	t.Parallel()

	// storage::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-storage-example",
	}

	// storage::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// storage::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// storage::tag::3:: Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")

	// storage::tag::5 Set expected variables for test
	expectedSkuTier := "Standard"
	expectedKind := "StorageV2"
	containerName := "container1"

	// happy path tests

	//storage account exists
	exists, err := azure.StorageAccountExistsE(storageAccountName, resourceGroupName, "")
	require.NoError(t, err)
	assert.True(t, exists)

	//blob endpoint matches
	blobEndpoint, err := azure.GetStorageAccountPrimaryBlobEndpointE(storageAccountName, resourceGroupName, "")
	require.NoError(t, err)
	builtEndpointString, err := azure.BuildStorageDNSStringE(storageAccountName, resourceGroupName, "")
	require.NoError(t, err)
	assert.Equal(t, builtEndpointString, blobEndpoint, "Blob endpoint URI mismatch.")

	//sku tier
	storageSkuTier, err := azure.GetStorageAccountSkuTierE(storageAccountName, resourceGroupName, "")
	require.NoError(t, err)
	assert.Equal(t, expectedSkuTier, storageSkuTier, "Storage SKU Tier mismatch.")

	//kind
	kind, err := azure.GetStorageAccountKindE(storageAccountName, resourceGroupName, "")
	require.NoError(t, err)
	assert.Equal(t, expectedKind, kind, "Storage kind mismatch.")

	//container exists
	containerExists, err := azure.BlobContainerExistsE(containerName, storageAccountName, resourceGroupName, "")
	require.NoError(t, err)
	assert.True(t, containerExists, "Blob storage container does not exist.")

	//container public access denied
	publicAccess, err := azure.StorageContainerHasPublicAccessE(containerName, storageAccountName, resourceGroupName, "")
	require.NoError(t, err)
	assert.False(t, publicAccess, "Blob container has public access.")
}
