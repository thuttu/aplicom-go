package protocol

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvert_ConvertToIMEI(t *testing.T) {
	UnitIDHighBytes := [4]byte{0x01, 0x43, 0x72, 0x07}
	UnitIDLowBytes := [3]byte{0x29, 0xD6, 0x84}
	imie := ConvertToIMEI(UnitIDHighBytes, UnitIDLowBytes)
	require.Equal(t, "355632002225796", imie)
}

func Test_ConvertToIMEI(t *testing.T) {
	laiData := [3]byte{0x01, 0x5f, 0xab}
	lai := ConvertLaiUint32(laiData)
	require.Equal(t, uint32(0x15fab), lai)
}
