package protocol

import (
	"encoding/binary"
	"strconv"
)

func ConvertToIMEI(high [4]byte, low [3]byte) string {
	var value int64
	var uuid [8]byte
	uuid[0] = 0x00
	uuid[1] = high[0]
	uuid[2] = high[1]
	uuid[3] = high[2]
	uuid[4] = high[3]
	uuid[5] = low[0]
	uuid[6] = low[1]
	uuid[7] = low[2]
	value = int64(binary.BigEndian.Uint64(uuid[:]))
	return strconv.FormatInt(value, 10)
}

func ConvertLaiUint32(b [3]byte) uint32 {
	return (uint32(b[2]) | uint32(b[1])<<8 | uint32(b[0])<<16)
}
