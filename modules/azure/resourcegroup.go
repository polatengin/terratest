package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
)

// ResourceGroupExistsE indicates whether a resource group exists within a subscription
func ResourceGroupExistsE(resourceGroupName, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return false, err
	}
	client, err := GetResourceGroupClientE(subscriptionID)
	if err != nil {
		return false, err
	}
	rg, err := client.Get(context.Background(), resourceGroupName)
	if err != nil {
		return false, err
	}
	return (resourceGroupName == *rg.Name), nil
}

//GetResourceGroupClientE gets a resource group client in a subscription
func GetResourceGroupClientE(subscriptionID string) (*resources.GroupsClient, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	resourceGroupClient := resources.NewGroupsClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	resourceGroupClient.Authorizer = *authorizer
	return &resourceGroupClient, nil
}
