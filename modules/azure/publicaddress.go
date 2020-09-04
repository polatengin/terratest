package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetPublicAddressIP gets the private IPs of a Network INterface
func GetPublicAddressIP(t testing.TestingT, publicAddressName string, resGroupName string, subscriptionID string) string {
	IP, err := GetPublicAddressIPE(t, publicAddressName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return IP
}

// GetPublicAddressIPE gets the provate IPs of a Network INterface
func GetPublicAddressIPE(t testing.TestingT, publicAddressName string, resGroupName string, subscriptionID string) (string, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return "", err
	}

	// Create a NIC client
	pip, err := GetPublicIPAddressE(publicAddressName, resGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return *pip.IPAddress, nil
}

// GetPublicIPAddressE gets a PublicIPAddressesd
func GetPublicIPAddressE(publicIPAddressName string, resGroupName string, subscriptionID string) (*network.PublicIPAddress, error) {
	client, err := GetPublicIPAddressClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	publicIPAddress, err := client.Get(context.Background(), resGroupName, publicIPAddressName, "")
	if err != nil {
		return nil, err
	}
	return &publicIPAddress, nil
}

// GetPublicIPAddressClientE creates a PublicIPAddresses client
func GetPublicIPAddressClientE(subscriptionID string) (*network.PublicIPAddressesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	publicIPClient := network.NewPublicIPAddressesClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	publicIPClient.Authorizer = *authorizer
	return &publicIPClient, nil
}
