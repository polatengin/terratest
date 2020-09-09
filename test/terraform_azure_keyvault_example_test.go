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

func TestTerraformAzureExample(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-keyvault-example",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	keyVaultName := terraform.Output(t, terraformOptions, "key_vault_name")
	secretName := terraform.Output(t, terraformOptions, "secret_name")
	keyName := terraform.Output(t, terraformOptions, "key_name")
	certificateName := terraform.Output(t, terraformOptions, "certificate_name")

	secretExists := azure.KeyVaultSecretExists(t, keyVaultName, secretName)
	assert.True(t, secretExists, "kv-secret does not exist")

	keyExists := azure.KeyVaultKeyExists(t, keyVaultName, keyName)
	assert.True(t, keyExists, "kv-key does not exist")

	certificateExists := azure.KeyVaultCertificateExists(t, keyVaultName, certificateName)
	assert.True(t, certificateExists, "kv-cert does not exist")
}
