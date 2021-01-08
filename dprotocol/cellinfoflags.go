package dprotocol

// CellInfoFlags represents the cell info flags field.
type CellInfoFlags uint8

const (
	maskIsRegisteredToNetwork = 0b00000001
	maskRadioAccessTechnology = 0b00000110
)

// IsRegisteredToNetwork returns true if the device is registered ot the network.
func (c CellInfoFlags) IsRegisteredToNetwork() bool {
	return c&maskIsRegisteredToNetwork > 0
}

// RadioAccessTechnology returns the current Radio Access Technology (RAT).
func (c CellInfoFlags) RadioAccessTechnology() RadioAccessTechnology {
	return RadioAccessTechnology(c & maskRadioAccessTechnology)
}
