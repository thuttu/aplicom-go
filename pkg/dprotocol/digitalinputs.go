package dprotocol

type DigitalInputs uint16

func (d DigitalInputs) IsHigh(input DigitalInput) bool {
	return DigitalInput(d)&input == input
}
