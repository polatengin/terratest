package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
)

const (
	AzureEnvironmentEnvName = "AZURE_ENVIRONMENT"
)

// LoadBalancerExistsE returns an LB client
func LoadBalancerExistsE(loadBalancerName, resourceGroupName, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	client, err := GetLoadBalancerClientE(subscriptionID)
	lb, err := client.Get(context.Background(), resourceGroupName, loadBalancerName, "")
	if err != nil {
		return false, err
	}

	return *lb.Name == loadBalancerName, nil
}

// GetLoadBalancerE returns a load balancer resource as specified by name.
func GetLoadBalancerE(loadBalancerName string, resourceGroupName string, subscriptionID string) (*network.LoadBalancer, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	client, err := GetLoadBalancerClientE(subscriptionID)
	lb, err := client.Get(context.Background(), resourceGroupName, loadBalancerName, "")
	if err != nil {
		return nil, err
	}

	return &lb, nil
}

// GetLoadBalancerClientE creates a storage account client.
func GetLoadBalancerClientE(subscriptionID string) (*network.LoadBalancersClient, error) {
	loadBalancerClient := network.NewLoadBalancersClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	loadBalancerClient.Authorizer = *authorizer
	return &loadBalancerClient, nil
}

// GetPublicIPAddressE will return ApplicationGateway object and an error object
func GetPublicIPAddressE(resourceGroupName, publicIPAddressName, subscriptionID string) (*network.PublicIPAddress, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	client, err := GetPublicIPAddressClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	publicIPAddress, err := client.Get(context.Background(), resourceGroupName, publicIPAddressName, "")
	if err != nil {
		return nil, err
	}
	return &publicIPAddress, nil
}

// GetPublicIPAddressClientE creates a PublicIPAddresses client
func GetPublicIPAddressClientE(subscriptionID string) (*network.PublicIPAddressesClient, error) {
	client := network.NewPublicIPAddressesClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer
	return &client, nil
}
