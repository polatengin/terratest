// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureAvailabilitySetExample(t *testing.T) {
	t.Parallel()

	// Subscription ID, leave blank if available as an Environment Var
	subID := ""
	prefix := "terratest-avs"

	// Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// Relative path to the Terraform dir
		TerraformDir: "../examples/terraform-azure-availabilityset-example",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"prefix": prefix,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	// defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	// terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	availabilitySetName := terraform.Output(t, terraformOptions, "availability_set_name")
	availabilitySetFDC := terraform.Output(t, terraformOptions, "availability_set_fdc")
	vmName01 := terraform.Output(t, terraformOptions, "vm_name_01")
	vmName02 := terraform.Output(t, terraformOptions, "vm_name_02")

	// Check the Availability Set
	actualAvsName := azure.GetAvailabilitySetIDOfVirtualMachine(t, vmName01, resourceGroupName, subID)
	assert.True(t, strings.EqualFold(availabilitySetName, actualAvsName))

	// Check the Availability set fault domain counts
	actualAvsFaultDomainCount := azure.GetFaultDomainCountOfAvailabilitySet(t, availabilitySetName, resourceGroupName, subID)
	assert.Equal(t, availabilitySetFDC, actualAvsFaultDomainCount)

	// Check VMs exist in the Availability Set
	actualVMsInAvs := azure.GetVMsOfAvailabilitySet(t, availabilitySetName, resourceGroupName, subID)
	assert.Contains(t, actualVMsInAvs, strings.ToLower(vmName01))
	assert.Contains(t, actualVMsInAvs, strings.ToLower(vmName02))

}
