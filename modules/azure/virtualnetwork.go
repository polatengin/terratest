package azure

import (
	"context"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// CheckIPInSubnet gets the subnet and checks if the IP is in its range
func CheckIPInSubnet(t testing.TestingT, resGroupName string, vnetName string, subnetName string, IP string, subscriptionID string) bool {
	inRange, err := CheckIPInSubnetE(t, resGroupName, vnetName, subnetName, IP, subscriptionID)
	require.NoError(t, err)

	return inRange
}

// CheckIPInSubnetE gets the subnet and checks if the IP is in its range
func CheckIPInSubnetE(t testing.TestingT, resGroupName string, vnetName string, subnetName string, IP string, subscriptionID string) (bool, error) {
	envSubnetRange, err := GetSubnetIPRangeE(vnetName, subnetName, resGroupName, subscriptionID)
	if err != nil {
		return false, err
	}

	inRange, err := CheckIPInSubnetPrefixE(&IP, &envSubnetRange)
	if err != nil {
		return false, err
	}

	return inRange, nil
}

// CheckIPInSubnetPrefixE checks to see if an IP is contained within a given subnet prefix
func CheckIPInSubnetPrefixE(ipAddress *string, subnetPrefix *string) (bool, error) {
	ip := net.ParseIP(*ipAddress)

	if ip == nil {
		return false, fmt.Errorf("Failed to parse IP address %s", *ipAddress)
	}

	_, ipNet, err := net.ParseCIDR(*subnetPrefix)
	if err != nil {
		return false, fmt.Errorf("Failed to parse subnet range %s", *subnetPrefix)
	}

	return ipNet.Contains(ip), nil
}

// GetSubnetIPRange gets the range of IPs for a subnet
func GetSubnetIPRange(t testing.TestingT, vnetName string, subnetName string, resGroupName string, subscriptionID string) string {
	IPRange, err := GetSubnetIPRangeE(vnetName, subnetName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return IPRange
}

// GetSubnetIPRangeE gets the range of IPs for a subnet
func GetSubnetIPRangeE(vnetName string, subnetName string, resGroupName string, subscriptionID string) (string, error) {
	envSubnets, err := GetVNetSubnetsE(resGroupName, vnetName, subscriptionID)
	if err != nil {
		return "", err
	}

	return envSubnets[subnetName], nil
}

// GetVNetSubnets gets all virtual network subclients name, and address prefix
func GetVNetSubnets(t testing.TestingT, resGroupName string, vnetName string, subscriptionID string) map[string]string {
	subnets, err := GetVNetSubnetsE(resGroupName, vnetName, subscriptionID)
	require.NoError(t, err)

	return subnets
}

// GetVNetSubnetsE gets all virtual network subclients name, and address prefix
func GetVNetSubnetsE(resGroupName string, vnetName string, subscriptionID string) (map[string]string, error) {
	client, err := GetSubnetClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	subnets, err := client.List(context.Background(), resGroupName, vnetName)
	if err != nil {
		return nil, err
	}

	subNetDetails := make(map[string]string)
	for _, v := range subnets.Values() {
		subnetName := v.Name
		subNetAddressPrefix := v.AddressPrefix

		subNetDetails[string(*subnetName)] = string(*subNetAddressPrefix)
	}
	return subNetDetails, nil
}

// GetSubnetClientE creates a virtual network subnet client
func GetSubnetClientE(subscriptionID string) (*network.SubnetsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	subNetClient := network.NewSubnetsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	subNetClient.Authorizer = *authorizer
	return &subNetClient, nil
}

// GetVirtualNetworkE gets virtual network object
func GetVirtualNetworkE(t testing.TestingT, vnetName string, resGroupName string, subscriptionID string) (*network.VirtualNetwork, error) {

	client, err := GetVirtualNetworksClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	virtualNetwork, err := client.Get(context.Background(), resGroupName, vnetName, "")
	if err != nil {
		return nil, err
	}
	return &virtualNetwork, nil
}

// GetVirtualNetworksClientE creates a virtual network client
func GetVirtualNetworksClientE(subscriptionID string) (*network.VirtualNetworksClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	vnClient := network.NewVirtualNetworksClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	vnClient.Authorizer = *authorizer
	return &vnClient, nil
}
