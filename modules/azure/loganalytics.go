package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2020-03-01-preview/operationalinsights"
)

// LogAnalyticsWorkspaceExistsE indicates whether the log analytics workspace exists; otherwise false
func LogAnalyticsWorkspaceExistsE(workspaceName, resourceGroupName, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return false, err
	}
	resourceGroupName, err2 := getTargetAzureResourceGroupName(resourceGroupName)
	if err2 != nil {
		return false, err2
	}
	client, err3 := GetLogAnalyticsWorkspacesClientE(subscriptionID)
	if err3 != nil {
		return false, err3
	}
	ws, err4 := client.Get(context.Background(), resourceGroupName, workspaceName)
	if err4 != nil {
		return false, err4
	}
	return (workspaceName == *ws.Name), nil
}

// // GetLogAnalyticsSkuE gets the log analytics sku
// func GetLogAnalyticsSkuE(workspaceName, resourceGroupName, subscriptionID string) (string, error) {
// 	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
// 	if err != nil {
// 		return "", err
// 	}
// 	resourceGroupName, err2 := getTargetAzureResourceGroupName(resourceGroupName)
// 	if err2 != nil {
// 		return "", err2
// 	}
// 	client, err3 := GetLogAnalyticsWorkspacesClientE(subscriptionID)
// 	if err3 != nil {
// 		return "", err3
// 	}
// 	ws, err4 := client.Get(context.Background(), resourceGroupName, workspaceName)
// 	if err4 != nil {
// 		return "", err4
// 	}
// 	return string(ws.Sku.Name), nil
// }

// // GetLogAnalyticsRetentionInDaysE get the log analytics retention period in days.
// func GetLogAnalyticsRetentionInDaysE(workspaceName, resourceGroupName, subscriptionID string) (int32, error) {
// 	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
// 	if err != nil {
// 		return -1, err
// 	}
// 	resourceGroupName, err2 := getTargetAzureResourceGroupName((resourceGroupName))
// 	if err2 != nil {
// 		return -1, err2
// 	}
// 	client, err3 := GetLogAnalyticsWorkspacesClientE(subscriptionID)
// 	if err3 != nil {
// 		return -1, err3
// 	}
// 	ws, err4 := client.Get(context.Background(), resourceGroupName, workspaceName)
// 	if err4 != nil {
// 		return -1, err4
// 	}
// 	return *ws.RetentionInDays, nil
// }

// GetLogAnalyticsWorkspacesClientE get the workspaces client for log analytics
func GetLogAnalyticsWorkspacesClientE(subscriptionID string) (*operationalinsights.WorkspacesClient, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	workspacesClient := operationalinsights.NewWorkspacesClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	workspacesClient.Authorizer = *authorizer
	return &workspacesClient, nil
}
