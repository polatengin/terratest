package azure

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// AssertAvailabilitySetExists checks for an Azure Availability Set
func AssertAvailabilitySetExists(t testing.TestingT, avsName string, resGroupName string, subscriptionID string) {
	err := AssertAvailabilitySetExistsE(t, avsName, resGroupName, subscriptionID)
	require.NoError(t, err)
}

// AssertAvailabilitySetExistsE checks for an Azure Availability Set with error
func AssertAvailabilitySetExistsE(t testing.TestingT, avsName string, resGroupName string, subscriptionID string) error {
	// Get the Network Interface
	_, err := GetAvailabilitySetE(avsName, resGroupName, subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

// GetVMsOfAvailabilitySet gets the VMs in the Availability Set
func GetVMsOfAvailabilitySet(t testing.TestingT, avsName string, resGroupName string, subscriptionID string) []string {
	vms, err := GetVMsOfAvailabilitySetE(avsName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return vms
}

// GetVMsOfAvailabilitySetE gets the VMs in the Availability Set
func GetVMsOfAvailabilitySetE(avsName string, resGroupName string, subscriptionID string) ([]string, error) {
	client, err := GetAvailabilitySetClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	avs, err := client.Get(context.Background(), resGroupName, avsName)
	if err != nil {
		return nil, err
	}

	vms := []string{}
	for _, vm := range *avs.VirtualMachines {
		if err == nil {
			vms = append(vms, strings.ToLower(GetNameFromResourceID(*vm.ID)))
		}
	}

	return vms, nil
}

// GetFaultDomainCountOfAvailabilitySet gets the Availability Set Platform Fault Domain Count as a string
func GetFaultDomainCountOfAvailabilitySet(t testing.TestingT, avsName string, resGroupName string, subscriptionID string) string {
	avsFaultDomainCount, err := GetFaultDomainCountOfAvailabilitySetE(avsName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return avsFaultDomainCount
}

// GetFaultDomainCountOfAvailabilitySetE gets the Availability Set Platform Fault Domain Count as a string with error
func GetFaultDomainCountOfAvailabilitySetE(avsName string, resGroupName string, subscriptionID string) (string, error) {
	client, err := GetAvailabilitySetClientE(subscriptionID)
	if err != nil {
		return "", err
	}

	avs, err := client.Get(context.Background(), resGroupName, avsName)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", *avs.PlatformFaultDomainCount), nil
}

// GetAvailabilitySetE gets an Availability Set in the specified Azure Resource Group
func GetAvailabilitySetE(avsName string, resGroupName string, subscriptionID string) (*compute.AvailabilitySet, error) {
	// Validate resource group name and subscription ID
	resGroupName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetAvailabilitySetClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Availability Set
	avs, err := client.Get(context.Background(), resGroupName, avsName)
	if err != nil {
		return nil, err
	}

	return &avs, nil
}

// GetAvailabilitySetClientE creates a new Availability Set client in the specified Azure Subscription
func GetAvailabilitySetClientE(subscriptionID string) (*compute.AvailabilitySetsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Availability Set client
	client := compute.NewAvailabilitySetsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}
