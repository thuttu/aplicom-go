package dprotocol

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkPacket_UnmarshalBinary(b *testing.B) {
	var packet Packet
	data := getExampleData()
	for i := 0; i < b.N; i++ {
		if err := packet.UnmarshalBinary(data); err != nil {
			b.Fatal(err)
		}
	}
}

func TestPacket_UnmarshalBinary(t *testing.T) {
	var actual Packet
	require.NoError(t, actual.UnmarshalBinary(getExampleData()))
	require.Equal(t, getExamplePacket(), actual)
}
