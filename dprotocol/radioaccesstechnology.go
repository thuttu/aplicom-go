package dprotocol

// RadioAccessTechnology represents a Radio Access Technology (RAT).
type RadioAccessTechnology uint8

//go:generate stringer -type RadioAccessTechnology -trimprefix RadioAccessTechnology

// RadioAccessTechnology values.
const (
	RadioAccessTechnologyUnknown RadioAccessTechnology = 0b000
	RadioAccessTechnology2G      RadioAccessTechnology = 0b010
	RadioAccessTechnology3G      RadioAccessTechnology = 0b100
)
