package dprotocol

// Version represents a D protocol version.
type Version uint8

const (
	// Version1 is deprecated.
	Version1 Version = 1
	// Version2 is legacy, used by A1 up to 10.10.x and A9 up to 2.10.x.
	Version2 Version = 2
	// Version3 is used by A1 10.20.x and A9 NEX.
	Version3 Version = 3
	// Version 4 can be used with A11.
	Version4 Version = 4
)
