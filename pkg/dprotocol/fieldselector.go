package dprotocol

// FieldSelector is a mask which enables optimization of sent data by dropping out unnecessary fields.
type FieldSelector uint32

//go:generate stringer -type FieldSelector -trimprefix FieldSelector

const (
	FieldSelectorGPSFlags              FieldSelector = 0x000008
	FieldSelectorTime                  FieldSelector = 0x000004
	FieldSelectorGPS                   FieldSelector = 0x000008
	FieldSelectorGPSSpeed              FieldSelector = 0x000010
	FieldSelectorAnalogInput           FieldSelector = 0x000020
	FieldSelectorIO                    FieldSelector = 0x000040
	FieldSelectorTrip1                 FieldSelector = 0x000080
	FieldSelectorTrip2                 FieldSelector = 0x000100
	FieldSelectorIButton               FieldSelector = 0x000200
	FieldSelectorDriverLogKeypad       FieldSelector = 0x000400
	FieldSelectorGPSExtras             FieldSelector = 0x000800
	FieldSelectorEventSpecificBytes    FieldSelector = 0x001000
	FieldSelectorSnapshotCounter       FieldSelector = 0x002000
	FieldSelectorFlags                 FieldSelector = 0x004000
	FieldSelectorPower                 FieldSelector = 0x008000
	FieldSelectorPulseCounter1         FieldSelector = 0x010000
	FieldSelectorPulseCounter2         FieldSelector = 0x020000
	FieldSelectorCellInfo              FieldSelector = 0x080000
	FieldSelectorExtendedDigitalInputs FieldSelector = 0x100000
)
