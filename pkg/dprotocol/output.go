package dprotocol

type Output uint8

//go:generate gobin -m -run golang.org/x/tools/cmd/stringer -type Output -trimprefix Output

const (
	OutputOff      Output = 0b00
	OutputInactive Output = 0b10
	OutputActive   Output = 0b11
)
