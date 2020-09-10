package azure

import (
	"errors"
	"strings"
)

// GetNameFromResourceID gets the Name from an Azure Resource ID
func GetNameFromResourceID(resourceID string) string {
	lastValue, err := GetSliceLastValueE(resourceID, "/")
	if err != nil {
		return ""
	}
	return lastValue
}

// GetSliceLastValueE will take a source string and returns the last value when split by the seperaror char
func GetSliceLastValueE(source string, seperator string) (string, error) {
	if !(len(source) == 0 || len(seperator) == 0 || !strings.Contains(source, seperator)) {
		tmp := strings.Split(source, seperator)
		return tmp[len(tmp)-1], nil
	}
	return "", errors.New("invalid input or no slice available")
}

// GetSliceIndexValueE will take a source string and returns the value at the given index when split by the seperaror char
func GetSliceIndexValueE(source string, seperator string, index int) (string, error) {
	if !(len(source) == 0 || len(seperator) == 0 || !strings.Contains(source, seperator) || index < 0) {
		tmp := strings.Split(source, seperator)
		if !(len(tmp) >= index) {
			return "", errors.New("index out of slice range")
		}
		return tmp[index], nil
	}
	return "", errors.New("invalid input or no slice available")
}
