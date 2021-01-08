package dprotocol

import (
	"bytes"
	"testing"

	"gotest.tools/v3/assert"
)

func TestScanner_ScanPacket(t *testing.T) {
	t.Parallel()
	sc := NewScanner(bytes.NewReader(getExampleData()))
	assert.Assert(t, sc.ScanPacket())
	assert.DeepEqual(t, getExamplePacket(), sc.Packet())
	assert.DeepEqual(t, getExampleData(), sc.Bytes())
	assert.Assert(t, !sc.ScanPacket())
	assert.NilError(t, sc.Err())
}
