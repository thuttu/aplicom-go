package dprotocol

// DriverLogKeypad represents the status of the Driver Log Keypad (DLKP).
type DriverLogKeypad uint8

// DriverLogKeypad values.
const (
	DriverLogKeypadNotConnected DriverLogKeypad = 0xff
)

// IsConnected returns true if the driver log keypad is connected.
func (d DriverLogKeypad) IsConnected() bool {
	return d != DriverLogKeypadNotConnected
}
