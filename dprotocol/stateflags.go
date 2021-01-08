package dprotocol

// StateFlags contains the state flags.
type StateFlags uint32

// Has returns true if the provided state flag is set (logical high).
func (s StateFlags) Has(flag StateFlag) bool {
	return StateFlag(s)&flag == flag
}
