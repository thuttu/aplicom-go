package dprotocol

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanPackets(t *testing.T) {
	sc := bufio.NewScanner(bytes.NewReader(getExampleData()))
	sc.Split(ScanPackets)
	require.True(t, sc.Scan())
	require.Equal(t, getExampleData(), sc.Bytes())
	require.False(t, sc.Scan())
	require.Nil(t, sc.Err())
}
