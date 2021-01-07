package dprotocol

import (
	"bytes"
	"fmt"
)

const packetIdentifier = 'D'

func ScanPackets(data []byte, _ bool) (advance int, token []byte, err error) {
	startIndex := bytes.IndexByte(data, packetIdentifier)
	switch {
	case startIndex == -1:
		return len(data), nil, nil // no packet identifier, discard all data
	case startIndex != 0:
		return startIndex, nil, nil // discard until packet identifier
	case len(data) < lengthOfPacketHeader:
		return 0, nil, nil // wait for more data
	}
	var header Header
	if err := header.UnmarshalBinary(data); err != nil {
		return 0, nil, fmt.Errorf("scan packets: %w", err)
	}
	packetLength := lengthOfPacketHeader + int(header.SnapshotLength)
	if len(data) < packetLength {
		return 0, nil, nil
	}
	return packetLength, data[:packetLength], nil
}
