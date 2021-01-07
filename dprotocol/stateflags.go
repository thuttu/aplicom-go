package dprotocol

type StateFlags uint32

func (s StateFlags) Has(flag StateFlag) bool {
	return StateFlag(s)&flag == flag
}
