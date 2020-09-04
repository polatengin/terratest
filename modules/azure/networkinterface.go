package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetNetworkInterfacePrivateIPs gets the private IPs of a Network INterface
func GetNetworkInterfacePrivateIPs(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) []string {
	IPs, err := GetNetworkInterfacePrivateIPsE(t, nicName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return IPs
}

// GetNetworkInterfacePrivateIPsE gets the provate IPs of a Network INterface
func GetNetworkInterfacePrivateIPsE(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) ([]string, error) {
	IPs := []string{}

	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a NIC client
	nic, err := GetNetworkInterfaceE(nicName, resGroupName, subscriptionID)
	if err != nil {
		return nil, err
	}

	for _, pip := range *nic.IPConfigurations {
		IPs = append(IPs, *pip.PrivateIPAddress)
	}

	return IPs, nil
}

// GetNetworkInterfaceE gets a Network Interface in the specified Azure Resource Group
func GetNetworkInterfaceE(nicName string, resGroupName string, subscriptionID string) (*network.Interface, error) {
	client, err := GetNetworkInterfaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	nic, err := client.Get(context.Background(), resGroupName, nicName, "")
	if err != nil {
		return nil, err
	}

	return &nic, nil
}

// GetNetworkInterfaceClientE creates a new Network Interface client
func GetNetworkInterfaceClientE(subscriptionID string) (*network.InterfacesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	nicClient := network.NewInterfacesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	nicClient.Authorizer = *authorizer
	return &nicClient, nil
}
