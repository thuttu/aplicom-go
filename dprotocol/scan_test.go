package dprotocol

import (
	"bufio"
	"bytes"
	"testing"

	"gotest.tools/v3/assert"
)

func TestScanPackets(t *testing.T) {
	t.Parallel()
	sc := bufio.NewScanner(bytes.NewReader(getExampleData()))
	sc.Split(ScanPackets)
	assert.Assert(t, sc.Scan())
	assert.DeepEqual(t, getExampleData(), sc.Bytes())
	assert.Assert(t, !sc.Scan())
	assert.NilError(t, sc.Err())
}
