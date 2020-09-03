// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureLogAnalyticsExample(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-loganalytics-example",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	assert.True(t, resourceGroupName != "")
	workspaceName := terraform.Output(t, terraformOptions, "loganalytics_workspace_name")
	assert.True(t, workspaceName != "")
	expectedSku := terraform.Output(t, terraformOptions, "loganalytics_workspace_sku")
	assert.True(t, expectedSku != "")
	days := terraform.Output(t, terraformOptions, "loganalytics_workspace_retention")
	assert.True(t, days != "")
	// expectedDays, _ := strconv.ParseInt(days, 10, 32)

	exists, err := azure.LogAnalyticsWorkspaceExistsE(workspaceName, resourceGroupName, "")
	assert.NoError(t, err, "Log analytics workspace error")
	assert.True(t, exists, "Log analytics workspace name mismatch")

	// actualSku, err := azure.GetLogAnalyticsSkuE(workspaceName, resourceGroupName, "")
	// assert.NoError(t, err, "Log analytics workspace sku error")
	// assert.Equal(t, expectedSku, actualSku, "Log analytics workspace sku mismatch")

	// actualDays, err := azure.GetLogAnalyticsRetentionInDaysE(workspaceName, resourceGroupName, "")
	// assert.NoError(t, err, "Log analytics workspace retention period days error")
	// assert.Equal(t, expectedDays, actualDays, "Log analytics workspace retention period days mismatch")
}
