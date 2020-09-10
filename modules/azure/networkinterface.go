package azure

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// AssertNetworkInterfaceExists checks for an Azure Network Interface
func AssertNetworkInterfaceExists(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) {
	err := AssertNetworkInterfaceExistsE(t, nicName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertNetworkInterfaceExistsE checks for an Azure Network Interface with Error
func AssertNetworkInterfaceExistsE(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) error {
	// Get the Network Interface
	_, err := GetNetworkInterfaceE(nicName, resGroupName, subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

// GetNetworkInterfacePublicIPs gets the Public IPs of a Network Interface configs
func GetNetworkInterfacePublicIPs(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) []string {
	IPs, err := GetNetworkInterfacePublicIPsE(t, nicName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return IPs
}

// GetNetworkInterfacePublicIPsE gets the Public IPs of a Network Interface configs with error
func GetNetworkInterfacePublicIPsE(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) ([]string, error) {
	publicIPs := []string{}

	// Get the Network Interface client
	nic, err := GetNetworkInterfaceE(nicName, resGroupName, subscriptionID)
	if err != nil {
		return publicIPs, err
	}

	// Get the Public IPs from each configuration
	for _, IPConfiguration := range *nic.IPConfigurations {
		// String conversion to avoid null pointer deref panic (may be a better way, reflect also panic'ed)
		var byteIPConfig []byte
		byteIPConfig, err := json.Marshal(IPConfiguration)
		if err == nil {
			stringIPConfig := string(byteIPConfig)

			// Check for public Ip Address in string representation of NIC
			if strings.Contains(stringIPConfig, "publicIPAddress") {
				publicAddressID := GetNameFromResourceID(*IPConfiguration.PublicIPAddress.ID)

				// Get the Public Ip from the Public Address resource
				publicIP := GetPublicAddressIP(t, publicAddressID, resGroupName, subscriptionID)
				publicIPs = append(publicIPs, publicIP)
			}
		}
	}

	return publicIPs, nil
}

// GetNetworkInterfacePrivateIPs gets the Private IPs of a Network Interface configs
func GetNetworkInterfacePrivateIPs(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) []string {
	IPs, err := GetNetworkInterfacePrivateIPsE(t, nicName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return IPs
}

// GetNetworkInterfacePrivateIPsE gets the Private IPs of a Network Interface configs with error
func GetNetworkInterfacePrivateIPsE(t testing.TestingT, nicName string, resGroupName string, subscriptionID string) ([]string, error) {
	privateIPs := []string{}

	// Get the Network Interface client
	nic, err := GetNetworkInterfaceE(nicName, resGroupName, subscriptionID)
	if err != nil {
		return privateIPs, err
	}

	// Get the Private IPs from each configuration
	for _, IPConfiguration := range *nic.IPConfigurations {
		privateIPs = append(privateIPs, *IPConfiguration.PrivateIPAddress)
	}

	return privateIPs, nil
}

// GetNetworkInterfaceE gets a Network Interface in the specified Azure Resource Group
func GetNetworkInterfaceE(nicName string, resGroupName string, subscriptionID string) (*network.Interface, error) {
	// Validate Azure Resource Group
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetNetworkInterfaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Network Interface
	nic, err := client.Get(context.Background(), resGroupName, nicName, "")
	if err != nil {
		return nil, err
	}

	return &nic, nil
}

// GetNetworkInterfaceClientE creates a new Network Interface client in the specified Azure Subscription
func GetNetworkInterfaceClientE(subscriptionID string) (*network.InterfacesClient, error) {
	// Validate Azure Subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the NIC client
	client := network.NewInterfacesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}
