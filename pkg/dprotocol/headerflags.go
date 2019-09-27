package dprotocol

type HeaderFlags uint8

const (
	versionMask    HeaderFlags = 0b00001111
	selectorMask   HeaderFlags = 0b01000000
	longUnitIDMask HeaderFlags = 0b10000000
)

func (p HeaderFlags) Version() Version {
	return Version(p & versionMask)
}

func (p HeaderFlags) HasSelectorBits() bool {
	return p&selectorMask == selectorMask
}

func (p HeaderFlags) HasLongUnitID() bool {
	return p&longUnitIDMask == longUnitIDMask
}
