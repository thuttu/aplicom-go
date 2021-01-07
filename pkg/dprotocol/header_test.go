package dprotocol

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestHeader_UnmarshalBinary(t *testing.T) {
	t.Parallel()
	var actual Header
	assert.NilError(t, actual.UnmarshalBinary(getExampleData()))
	assert.DeepEqual(t, getExampleHeader(), actual)
}
