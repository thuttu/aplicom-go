package dprotocol

import (
	"encoding/binary"
	"fmt"
)

// Header represents a D protocol packet header.
type Header struct {
	Flags          HeaderFlags
	SnapshotLength uint16
	FieldSelectors FieldSelectors
	UnitID         uint64
}

// header field lengths.
const (
	lengthOfPacketIdentifier = 1
	lengthOfHeaderFlags      = 1
	lengthOfUnitID           = 7
	lengthOfSnapshotLength   = 2
	lengthOfSelectorBits     = 3
	lengthOfPacketHeader     = 14
)

// header field indices
const (
	indexOfPacketIdentifier = 0
	indexOfHeaderFlags      = indexOfPacketIdentifier + lengthOfPacketIdentifier
	indexOfUnitID           = indexOfHeaderFlags + lengthOfHeaderFlags
	indexOfSnapshotLength   = indexOfUnitID + lengthOfUnitID
	indexOfSelectorBits     = indexOfSnapshotLength + lengthOfSnapshotLength
)

// compile-time assertion on header structure.
var _ [indexOfSelectorBits + lengthOfSelectorBits]struct{} = [lengthOfPacketHeader]struct{}{}

func (p *Header) UnmarshalBinary(b []byte) error {
	if len(b) < lengthOfPacketHeader {
		return fmt.Errorf("invalid packet header length: %v", len(b))
	}
	p.Flags = HeaderFlags(b[indexOfHeaderFlags])
	switch p.Flags.Version() {
	case Version4, Version3, Version2: // supported
	default:
		return fmt.Errorf("invalid protocol version: %v", p.Flags.Version())
	}
	switch {
	case !p.Flags.HasSelectorBits():
		return fmt.Errorf("invalid protocol version: does not have selector bits")
	case !p.Flags.HasLongUnitID():
		return fmt.Errorf("invalid protocol version: does not have long unit ID")
	}
	p.UnitID = bigEndianUint56(b[indexOfUnitID:])
	p.SnapshotLength = binary.BigEndian.Uint16(b[indexOfSnapshotLength:])
	p.FieldSelectors = FieldSelectors(bigEndianUint24(b[indexOfSelectorBits:]))
	return nil
}
