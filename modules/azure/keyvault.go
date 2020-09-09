package azure

import (
	"context"
	"fmt"
	"testing"

	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/require"
)

func KeyVaultSecretExists(t *testing.T, keyVaultName string, secretName string) bool {
	result, err := KeyVaultSecretExistsE(keyVaultName, secretName)
	require.NoError(t, err)
	return result
}

func KeyVaultKeyExists(t *testing.T, keyVaultName string, keyName string) bool {
	result, err := KeyVaultKeyExistsE(keyVaultName, keyName)
	require.NoError(t, err)
	return result
}

func KeyVaultCertificateExists(t *testing.T, keyVaultName string, certificateName string) bool {
	result, err := KeyVaultCertificateExistsE(keyVaultName, certificateName)
	require.NoError(t, err)
	return result
}

// KeyVaultCertificateExistsE indicates whether a certificate exists in key vault; otherwise false
func KeyVaultCertificateExistsE(keyVaultName, certificateName string) (bool, error) {
	keyVaultSuffix, err := GetKeyVaultUriSuffixE()
	if err != nil {
		return false, err
	}
	client, err := GetKeyVaultClientE()
	if err != nil {
		return false, err
	}
	var maxVersionsCount int32 = 1
	versions, err := client.GetCertificateVersions(context.Background(),
		fmt.Sprintf("https://%s.%s", keyVaultName, keyVaultSuffix),
		certificateName,
		&maxVersionsCount)
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	items := versions.Values()

	if len(items) > 0 {
		return true, nil
	}
	return false, nil
}

// KeyVaultKeyExistsE indicates whether a key exists in the key vault; otherwise false
func KeyVaultKeyExistsE(keyVaultName, keyName string) (bool, error) {
	keyVaultSuffix, err := GetKeyVaultUriSuffixE()
	if err != nil {
		return false, err
	}
	client, err := GetKeyVaultClientE()
	if err != nil {
		return false, err
	}
	var maxVersionsCount int32 = 1
	versions, err := client.GetKeyVersions(context.Background(),
		fmt.Sprintf("https://%s.%s", keyVaultName, keyVaultSuffix),
		keyName,
		&maxVersionsCount)
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	items := versions.Values()

	if len(items) > 0 {
		return true, nil
	}
	return false, nil
}

// KeyVaultSecretExistsE indicates whether a secret exists in the key vault; otherwise false
func KeyVaultSecretExistsE(keyVaultName, secretName string) (bool, error) {
	client, err := GetKeyVaultClientE()
	if err != nil {
		return false, err
	}
	keyVaultSuffix, err := GetKeyVaultUriSuffixE()
	if err != nil {
		return false, err
	}
	var maxVersionsCount int32 = 1
	versions, err := client.GetSecretVersions(context.Background(),
		fmt.Sprintf("https://%s.%s", keyVaultName, keyVaultSuffix),
		secretName,
		&maxVersionsCount)
	if err != nil {
		return false, err
	}
	items := versions.Values()

	if len(items) > 0 {
		return true, nil
	}
	return false, nil
}

// GetKeyVaultClientE creates a KeyVault client
func GetKeyVaultClientE() (*keyvault.BaseClient, error) {
	kvClient := keyvault.New()
	authorizer, err := NewKeyVaultAuthorizerE()
	if err != nil {
		return nil, err
	}
	kvClient.Authorizer = *authorizer
	return &kvClient, nil
}

// GetKeyVaultUriSuffixE returns the proper KeyVault URI suffix for the configured Azure environment.
func GetKeyVaultUriSuffixE() (string, error) {
	env, err := azure.EnvironmentFromName("AzurePublicCloud")
	if err != nil {
		return "", err
	}
	return env.KeyVaultDNSSuffix, nil
}

// NewKeyVaultAuthorizerE will return Authorizer for KeyVault
func NewKeyVaultAuthorizerE() (*autorest.Authorizer, error) {
	authorizer, err := kvauth.NewAuthorizerFromCLI()
	return &authorizer, err
}
