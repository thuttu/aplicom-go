package dprotocol

type DriverLogKeypad uint8

const (
	DriverLogKeypadNotConnected DriverLogKeypad = 0xff
)

func (d DriverLogKeypad) IsConnected() bool {
	return d != DriverLogKeypadNotConnected
}
