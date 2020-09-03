// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"fmt"
	"strings"

	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureBlobContainerExample(t *testing.T) {
	t.Parallel()

	_random := strings.ToLower(random.UniqueId())

	expectedContainerName := fmt.Sprintf("bc_%s", _random)
	expectedResourceName := fmt.Sprintf("tmpaci%s", _random)
	expectedResourceGroupName := fmt.Sprintf("tmp-rg-%s", _random)

	terraformOptions := &terraform.Options{
		TerraformDir: "../../examples/terraform-azure-aci-example",
		Vars: map[string]interface{}{
			"blob_container_name": expectedContainerName,
			"account_name":        expectedResourceName,
			"resource_group_name": expectedResourceGroupName,
		},
	}
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	client := azure.GetBlobContainerClient(t, expectedContainerName, expectedResourceName, expectedResourceGroupName, "")

	assert := assert.New(t)

	assert.Equal(&expectedContainerName, *client.Name)
}
