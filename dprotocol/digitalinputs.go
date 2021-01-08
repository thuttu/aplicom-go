package dprotocol

// DigitalInputs represents the current digital input status.
type DigitalInputs uint16

// IsSet returns true if the provided digital input is currently set (logical high).
func (d DigitalInputs) IsSet(input DigitalInput) bool {
	return DigitalInput(d)&input == input
}
