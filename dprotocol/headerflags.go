package dprotocol

// HeaderFlags contains the header flags for a D protocol packet.
type HeaderFlags uint8

const (
	versionMask    HeaderFlags = 0b00001111
	selectorMask   HeaderFlags = 0b01000000
	longUnitIDMask HeaderFlags = 0b10000000
)

// Version returns the protocol version.
func (p HeaderFlags) Version() Version {
	return Version(p & versionMask)
}

// HasSelectorBits returns true if the packet has selector bits.
func (p HeaderFlags) HasSelectorBits() bool {
	return p&selectorMask == selectorMask
}

// HasLongUnitID returns true if the packet has a long unit ID..
func (p HeaderFlags) HasLongUnitID() bool {
	return p&longUnitIDMask == longUnitIDMask
}
