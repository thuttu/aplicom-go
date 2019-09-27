package dprotocol

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeader_UnmarshalBinary(t *testing.T) {
	var actual Header
	require.NoError(t, actual.UnmarshalBinary(getExampleData()))
	require.Equal(t, getExampleHeader(), actual)
}
