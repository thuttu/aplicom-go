package dprotocol

import (
	"testing"

	"gotest.tools/v3/assert"
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
	t.Parallel()
	var actual Packet
	assert.NilError(t, actual.UnmarshalBinary(getExampleData()))
	assert.DeepEqual(t, getExamplePacket(), &actual)
}
