package gokore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisteredEndpointsGetByIndexAndSortOrder(t *testing.T) {
	t1 := NewTestEndpoint("dogs", "dog")
	t2 := NewTestEndpoint("cats", "cat")
	t3 := NewTestEndpoint("cows", "cow")

	RegisterEndpoint(t1)
	RegisterEndpoint(t2)
	RegisterEndpoint(t3)

	// expected order: cats, cows, dogs
	require.Equal(t, 3, GetEndpointCount())

	tt, err := GetEndpointByIndex(0)
	require.NoError(t, err)
	assert.Equal(t, "cats", tt.Title())

	tt, err = GetEndpointByIndex(1)
	require.NoError(t, err)
	assert.Equal(t, "cows", tt.Title())

	tt, err = GetEndpointByIndex(2)
	require.NoError(t, err)
	assert.Equal(t, "dogs", tt.Title())

	_, err = GetEndpointByIndex(-1)
	require.Error(t, err)

	_, err = GetEndpointByIndex(11)
	require.Error(t, err)
}

func TestRegisteredEndpointsGetByTitleSuccess(t *testing.T) {
	t1 := NewTestEndpoint("dogs", "dog")
	t2 := NewTestEndpoint("cats", "cat")
	t3 := NewTestEndpoint("cows", "cow")

	RegisterEndpoint(t1)
	RegisterEndpoint(t2)
	RegisterEndpoint(t3)

	tt, err := GetEndpointByTitle("dogs")
	require.NoError(t, err)
	assert.Equal(t, "dogs", tt.Title())

	_, err = GetEndpointByTitle("dog")
	require.Error(t, err)
}
