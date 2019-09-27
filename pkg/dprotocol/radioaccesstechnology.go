package dprotocol

type RadioAccessTechnology uint8

//go:generate gobin -m -run golang.org/x/tools/cmd/stringer -type RadioAccessTechnology -trimprefix RadioAccessTechnology

const (
	RadioAccessTechnologyUnknown RadioAccessTechnology = 0b000
	RadioAccessTechnology2G      RadioAccessTechnology = 0b010
	RadioAccessTechnology3G      RadioAccessTechnology = 0b100
)
