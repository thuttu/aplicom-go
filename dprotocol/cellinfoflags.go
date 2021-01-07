package dprotocol

type CellInfoFlags uint8

const (
	maskIsRegisteredToNetwork = 0b00000001
	maskRadioAccessTechnology = 0b00000110
)

func (c CellInfoFlags) IsRegisteredToNetwork() bool {
	return c&maskIsRegisteredToNetwork > 0
}

func (c CellInfoFlags) RadioAccessTechnology() RadioAccessTechnology {
	return RadioAccessTechnology(c & maskRadioAccessTechnology)
}
