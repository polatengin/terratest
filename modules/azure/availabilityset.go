package azure

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

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
		tmp := strings.Split(*vm.ID, "/")
		if len(tmp) > 0 {
			vms = append(vms, tmp[len(tmp)-1])
		}
	}

	return vms, nil
}

// GetFaultDomainCountOfAvailabilitySet gets the Availability Set Platform Fault Domain Count
func GetFaultDomainCountOfAvailabilitySet(t testing.TestingT, avsName string, resGroupName string, subscriptionID string) int32 {
	avsFaultDomainCount, err := GetFaultDomainCountOfAvailabilitySetE(avsName, resGroupName, subscriptionID)
	require.NoError(t, err)

	return avsFaultDomainCount
}

// GetFaultDomainCountOfAvailabilitySetE gets the Availability Set Platform Fault Domain Count
func GetFaultDomainCountOfAvailabilitySetE(avsName string, resGroupName string, subscriptionID string) (int32, error) {
	client, err := GetAvailabilitySetClientE(subscriptionID)
	if err != nil {
		return -1, err
	}

	avs, err := client.Get(context.Background(), resGroupName, avsName)
	if err != nil {
		return -1, err
	}

	return *avs.PlatformFaultDomainCount, nil
}

// GetAvailabilitySetE gets an Availability Set in the specified resource group
func GetAvailabilitySetE(avsName string, resGroupName string, subscriptionID string) (*compute.AvailabilitySet, error) {
	client, err := GetAvailabilitySetClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	avs, err := client.Get(context.Background(), resGroupName, avsName)
	if err != nil {
		return nil, err
	}

	return &avs, nil
}

// GetAvailabilitySetClientE creates a new Availability Set client
func GetAvailabilitySetClientE(subscriptionID string) (*compute.AvailabilitySetsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	avsClient := compute.NewAvailabilitySetsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	avsClient.Authorizer = *authorizer
	return &avsClient, nil
}
