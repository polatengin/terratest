package azure

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
)

const (
	AzureEnvironmentEnvName = "AZURE_ENVIRONMENT"
)

// ref: StorageAccountExistsE
// GetLoadBalancerE returns an LB client
func GetLoadBalancerE(loadBalancerName, resourceGroupName, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	client, err := GetLoadBalancerClientE(subscriptionID)
	lb, err := client.Get(context.Background(), resourceGroupName, loadBalancerName, "")
	if err != nil {
		return false, err
	}

	return *lb.Name == loadBalancerName, nil
}

// ref: GetStorageAccountClientE
// GetStorageAccountClientE creates a storage account client.
func GetLoadBalancerClientE(subscriptionID string) (*network.LoadBalancersClient, error) {
	loadBalancerClient := network.NewLoadBalancersClient(os.Getenv(AzureSubscriptionID))
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	loadBalancerClient.Authorizer = *authorizer
	return &loadBalancerClient, nil
}
