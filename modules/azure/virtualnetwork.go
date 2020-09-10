package azure

import (
	"context"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// AssertSubnetExists checks for a Subnet
func AssertSubnetExists(t testing.TestingT, subnetName string, vnetName string, resGroupName string, subscriptionID string) {
	err := AssertSubnetExistsE(t, subnetName, vnetName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertSubnetExistsE checks for a Virtual Network with Error
func AssertSubnetExistsE(t testing.TestingT, subnetName string, vnetName string, resGroupName string, subscriptionID string) error {
	// Get the Network Interface client
	_, err := GetSubnetE(t, subnetName, vnetName, resGroupName, subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

// AssertVirtualNetworkExists checks for an Azure Virtual Network
func AssertVirtualNetworkExists(t testing.TestingT, vnetName string, resGroupName string, subscriptionID string) {
	err := AssertVirtualNetworkExistsE(t, vnetName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertVirtualNetworkExistsE checks for an Azure Virtual Network with Error
func AssertVirtualNetworkExistsE(t testing.TestingT, vnetName string, resGroupName string, subscriptionID string) error {
	// Get the Network Interface
	_, err := GetVirtualNetworkE(t, vnetName, resGroupName, subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

// CheckPrivateIPInSubnet gets the subnet and checks if the IP is in its range
func CheckPrivateIPInSubnet(t testing.TestingT, resGroupName string, vnetName string, subnetName string, IP string, subscriptionID string) bool {
	inRange, err := CheckPrivateIPInSubnetE(t, resGroupName, vnetName, subnetName, IP, subscriptionID)
	require.NoError(t, err)

	return inRange
}

// CheckPrivateIPInSubnetE gets the subnet and checks if the IP is in its range
func CheckPrivateIPInSubnetE(t testing.TestingT, resGroupName string, vnetName string, subnetName string, IP string, subscriptionID string) (bool, error) {
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

// GetVNetSubnetList gets a list of all virtual network subnets
func GetVNetSubnetList(t testing.TestingT, vnetName string, resGroupName string, subscriptionID string) []string {
	subnets, err := GetVNetSubnetListsE(vnetName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return subnets
}

// GetVNetSubnetListsE gets a list of all virtual network subnets with Error
func GetVNetSubnetListsE(vnetName string, resGroupName string, subscriptionID string) ([]string, error) {
	client, err := GetSubnetClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	subnets, err := client.List(context.Background(), resGroupName, vnetName)
	if err != nil {
		return nil, err
	}

	subnetList := []string{}
	for _, v := range subnets.Values() {
		subnetList = append(subnetList, *v.Name)
	}
	return subnetList, nil
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

// GetVNetDNSServerIPList gets a list of all Virtual Network DNS server IPs
func GetVNetDNSServerIPList(t testing.TestingT, vnetName string, resGroupName string, subscriptionID string) []string {
	vnetDNSIPs, err := GetVNetDNSServerIPListE(t, vnetName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return vnetDNSIPs
}

// GetVNetDNSServerIPListE gets a list of all Virtual Network DNS server IPs with Error
func GetVNetDNSServerIPListE(t testing.TestingT, vnetName string, resGroupName string, subscriptionID string) ([]string, error) {
	// Get Virtual Network
	vnet, err := GetVirtualNetworkE(t, vnetName, resGroupName, subscriptionID)
	if err != nil {
		return nil, err
	}

	return *vnet.DhcpOptions.DNSServers, nil
}

// GetSubnetRange gets the Subnet IPv4 Range
func GetSubnetRange(t testing.TestingT, subnetName string, vnetName string, resGroupName string, subscriptionID string) string {
	vnetDNSIPs, err := GetSubnetRangeE(t, subnetName, vnetName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return vnetDNSIPs
}

// GetSubnetRangeE gets the Subnet IPv4 Range with Error
func GetSubnetRangeE(t testing.TestingT, subnetName string, vnetName string, resGroupName string, subscriptionID string) (string, error) {
	// Get Subnet
	subnet, err := GetSubnetE(t, subnetName, vnetName, resGroupName, subscriptionID)
	if err != nil {
		return "", err
	}

	return *subnet.AddressPrefix, nil
}

// GetSubnetE gets a subnet
func GetSubnetE(t testing.TestingT, subnetName string, vnetName string, resGroupName string, subscriptionID string) (*network.Subnet, error) {
	// Validate Azure Resource Group
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetSubnetClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Subnet
	subnet, err := client.Get(context.Background(), resGroupName, vnetName, subnetName, "")
	if err != nil {
		return nil, err
	}

	return &subnet, nil
}

// GetSubnetClientE creates a subnet client
func GetSubnetClientE(subscriptionID string) (*network.SubnetsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Subnet client
	client := network.NewSubnetsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}

// GetVirtualNetworkE gets Virtual Network in the specified Azure Resource Group
func GetVirtualNetworkE(t testing.TestingT, vnetName string, resGroupName string, subscriptionID string) (*network.VirtualNetwork, error) {
	// Validate Azure Resource Group
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetVirtualNetworksClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Virtual Network
	vnet, err := client.Get(context.Background(), resGroupName, vnetName, "")
	if err != nil {
		return nil, err
	}
	return &vnet, nil
}

// GetVirtualNetworksClientE creates a virtual network client in the specified Azure Subscription
func GetVirtualNetworksClientE(subscriptionID string) (*network.VirtualNetworksClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Virtual Network client
	client := network.NewVirtualNetworksClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}
