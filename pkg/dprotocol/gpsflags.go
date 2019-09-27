package dprotocol

type GPSFlags uint8

func (g GPSFlags) Has(flag GPSFlag) bool {
	return GPSFlag(g)&flag == flag
}
