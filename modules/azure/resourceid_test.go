package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNameFromResourceID(t *testing.T) {
	// set slice variables
	sliceSource := "this/is/a/long/slash/separated/string/ResourceID"
	sliceResult := "ResourceID"
	sliceNotFound := "noresourcepresent"

	// verify success
	resultSuccess := GetNameFromResourceID(sliceSource)
	assert.Equal(t, sliceResult, resultSuccess)

	// verify error when seperator not found
	resultBadSeperator := GetNameFromResourceID(sliceNotFound)
	assert.Equal(t, "", resultBadSeperator)
}

func TestGetSliceLastValue(t *testing.T) {
	// set slice variables
	sliceSeperator := "/"
	sliceBadSeperator := "*"
	sliceNotFound := "noslicepresent"
	sliceSource := "this/is/a/long/slash/separated/string/success"
	sliceResult := "success"

	// verify success
	resultSuccess, err := GetSliceLastValueE(sliceSource, sliceSeperator)
	require.NoError(t, err)
	assert.Equal(t, sliceResult, resultSuccess)

	// verify error when seperator not found
	resultBadSeperator, err := GetSliceLastValueE(sliceSource, sliceBadSeperator)
	require.Error(t, err)
	assert.Equal(t, err.Error(), "invalid input or no slice available")
	assert.Equal(t, "", resultBadSeperator)

	// verify error when slice does not have seperator
	resultNotFound, err := GetSliceLastValueE(sliceNotFound, sliceSeperator)
	require.Error(t, err)
	assert.Equal(t, err.Error(), "invalid input or no slice available")
	assert.Equal(t, "", resultNotFound)
}

func TestGetSliceIndexValue(t *testing.T) {
	// set slice variables
	sliceSource := "this/is/a/long/slash/separated/string/success"
	sliceSeperator := "/"
	sliceSelectorNeg := -1
	sliceSelector0 := 0
	sliceSelector4 := 4
	sliceSelector7 := 7
	sliceSelector10 := 10
	sliceResult0 := "this"
	sliceResult4 := "slash"
	sliceResult7 := "success"

	// verify success index 0
	resultSuccess0, err := GetSliceIndexValueE(sliceSource, sliceSeperator, sliceSelector0)
	require.NoError(t, err)
	assert.Equal(t, sliceResult0, resultSuccess0)

	// verify success index 4
	resultSuccess4, err := GetSliceIndexValueE(sliceSource, sliceSeperator, sliceSelector4)
	require.NoError(t, err)
	assert.Equal(t, sliceResult4, resultSuccess4)

	// verify success index 7
	resultSuccess7, err := GetSliceIndexValueE(sliceSource, sliceSeperator, sliceSelector7)
	require.NoError(t, err)
	assert.Equal(t, sliceResult7, resultSuccess7)

	// verify error with negative index
	resultNegIndex, err := GetSliceIndexValueE(sliceSource, sliceSeperator, sliceSelectorNeg)
	require.Error(t, err)
	assert.Equal(t, err.Error(), "invalid input or no slice available")
	assert.Equal(t, "", resultNegIndex)

	// verify error when seperator not found
	resultBadIndex10, err := GetSliceIndexValueE(sliceSource, sliceSeperator, sliceSelector10)
	require.Error(t, err)
	assert.Equal(t, err.Error(), "index out of slice range")
	assert.Equal(t, "", resultBadIndex10)
}
