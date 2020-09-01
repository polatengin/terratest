// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"fmt"
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
	skuTier := "Standard"
	kind := "StorageV2"
	accessLevel := "None"
	containerName := "container1"

	// storage::tag::4:: Look up the size of the given Virtual Machine and ensure it matches the output.

	storageAccount, err := azure.GetStorageAccountPropertyE(storageAccountName, resourceGroupName, "")
	require.NoError(t, err)

	storageSuffix, _ := azure.GetStorageUriSuffix()
	expectedDns := fmt.Sprintf("https://%s.blob.%s/", storageAccountName, *storageSuffix)
	assert.Equal(t, expectedDns, string(*storageAccount.AccountProperties.PrimaryEndpoints.Blob))

	assert.Equal(t, storageAccountName, *storageAccount.Name, "Storage account name not found.")
	assert.Equal(t, skuTier, string(storageAccount.Sku.Tier), "Storage account SKU tier mismatch.")
	assert.Equal(t, kind, string(storageAccount.Kind), "Storage account kind mismatch.")

	container, err := azure.GetBlobContainerE(storageAccountName, containerName, resourceGroupName, "")
	require.NoError(t, err)

	assert.Equal(t, containerName, *container.Name, "Storage container name mismatch.")
	assert.Equal(t, accessLevel, string(container.PublicAccess), "Storage container access level not private.")
}
