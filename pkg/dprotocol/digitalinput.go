package dprotocol

type DigitalInput uint16

//go:generate gobin -m -run golang.org/x/tools/cmd/stringer -type DigitalInput -trimprefix DigitalInput

const (
	DigitalInput1        DigitalInput = 0b0000_0000_0000_0001
	DigitalInput2        DigitalInput = 0b0000_0000_0000_0010
	DigitalInput3        DigitalInput = 0b0000_0000_0000_0100
	DigitalInput4        DigitalInput = 0b0000_0000_0000_1000
	DigitalInput5        DigitalInput = 0b0000_0000_0001_0000
	DigitalInput6        DigitalInput = 0b0000_0000_0010_0000
	DigitalInputIgnition DigitalInput = 0b0000_0000_1000_0000
	DigitalInput7        DigitalInput = 0b0000_0001_0000_0000
	DigitalInput8        DigitalInput = 0b0000_0010_0000_0000
	DigitalInput9        DigitalInput = 0b0000_0100_0000_0000
	DigitalInput10       DigitalInput = 0b0000_1000_0000_0000
)
