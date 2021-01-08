package dprotocol

// Output represents an output status.
type Output uint8

//go:generate stringer -type Output -trimprefix Output

// Output values.
const (
	OutputOff      Output = 0b00
	OutputInactive Output = 0b10
	OutputActive   Output = 0b11
)
