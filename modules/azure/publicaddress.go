package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// AssertPublicAddressExists checks for an Azure Public Address
func AssertPublicAddressExists(t testing.TestingT, publicAddressName string, resGroupName string, subscriptionID string) {
	err := AssertPublicAddressExistsE(publicAddressName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertPublicAddressExistsE checks for an Azure Public Address with error
func AssertPublicAddressExistsE(publicAddressName string, resGroupName string, subscriptionID string) error {
	// Get the Public Address
	_, err := GetPublicIPAddressE(publicAddressName, resGroupName, subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

// GetPublicAddressIP gets the IP of a Public IP Address
func GetPublicAddressIP(t testing.TestingT, publicAddressName string, resGroupName string, subscriptionID string) string {
	IP, err := GetPublicAddressIPE(t, publicAddressName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return IP
}

// GetPublicAddressIPE gets the IP of a Public IP Address with error
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

// GetPublicIPAddressE gets a Public IP Addresses in the specified Azure Resource Group
func GetPublicIPAddressE(publicIPAddressName string, resGroupName string, subscriptionID string) (*network.PublicIPAddress, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetPublicIPAddressClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Public IP Address
	pip, err := client.Get(context.Background(), resGroupName, publicIPAddressName, "")
	if err != nil {
		return nil, err
	}
	return &pip, nil
}

// GetPublicIPAddressClientE creates a Public IP Addresses client in the specified Azure Subscription
func GetPublicIPAddressClientE(subscriptionID string) (*network.PublicIPAddressesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Public IP Address client
	client := network.NewPublicIPAddressesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}
