package dprotocol

import (
	"encoding/binary"
	"fmt"
	"time"
)

type Packet struct {
	Header                           Header
	EventID                          EventID
	EventInformation                 uint8
	GPSFlags                         GPSFlags
	Time                             time.Time
	GPSTime                          time.Time
	LatitudeMicroDegrees             uint32
	LongitudeMicroDegrees            uint32
	NumSatellites                    uint8
	SpeedKilometresPerHour           uint8
	MaximumSpeedKilometresPerHour    uint8
	HeadingHalfDegrees               uint8
	DigitalInputs                    DigitalInputs
	AnalogInput1VoltageMilliVolts    uint16
	AnalogInput2VoltageMilliVolts    uint16
	AnalogInput3VoltageMilliVolts    uint16
	AnalogInput4VoltageMilliVolts    uint16
	MainPowerVoltageMilliVolts       uint16
	ExternalBatteryVoltageMilliVolts uint16
	PulseCounter1Rate                uint16
	PulseCounter1                    uint32
	PulseCounter2Rate                uint16
	PulseCounter2                    uint32
	Trip1DistanceMetres              uint32
	Trip2DistanceMetres              uint32
	OutputStatus                     OutputStatus
	IButtonKeyID                     uint64
	DriverLogKeypad                  DriverLogKeypad
	GPSAltitudeMetres                int16
	SnapshotCounter                  uint16
	StateFlags                       StateFlags
	UserDefinedFlags                 uint32
	CellInfoFlags                    CellInfoFlags
	LocationAreaIdentity             uint32
	LocationAreaCode                 uint16
	GSMCellID                        uint32
	ReceivedSignalStrengthIndicator  int8
	EventSpecificBytes               []byte
}

const indexOfSnapshot = lengthOfPacketHeader

// snapshot field lengths.
const (
	lengthOfEventID                         = 1
	lengthOfEventInformation                = 1
	lengthOfGPSFlags                        = 1
	lengthOfTime                            = 4
	lengthOfGPSTime                         = 4
	lengthOfLatitude                        = 4
	lengthOfLongitude                       = 4
	lengthOfNumSatellites                   = 1
	lengthOfSpeed                           = 1
	lengthOfMaximumSpeed                    = 1
	lengthOfHeading                         = 1
	lengthOfDigitalInputs                   = 1
	lengthOfAnalogInput1                    = 2
	lengthOfAnalogInput2                    = 2
	lengthOfAnalogInput3                    = 2
	lengthOfAnalogInput4                    = 2
	lengthOfMainPower                       = 2
	lengthOfExternalBattery                 = 2
	lengthOfPulseCounter1Rate               = 2
	lengthOfPulseCounter1                   = 4
	lengthOfPulseCounter2Rate               = 2
	lengthOfPulseCounter2                   = 4
	lengthOfTrip1Distance                   = 4
	lengthOfTrip2Distance                   = 4
	lengthOfOutputStatus                    = 1
	lengthOfIButtonKeyID                    = 6
	lengthOfDriverLogKeypadState            = 1
	lengthOfGPSAltitude                     = 2
	lengthOfSnapshotCounter                 = 2
	lengthOfStateFlags                      = 4
	lengthOfUserDefinedFlags                = 4
	lengthOfCellInfoFlags                   = 1
	lengthOfLocationAreaIdentity            = 3
	lengthOfLocationAreaCode                = 2
	lengthOfGSMCellID                       = 4
	lengthOfReceivedSignalStrengthIndicator = 1
	lengthOfExtendedDigitalInputs           = 2
)

func (p *Packet) UnmarshalBinary(b []byte) error {
	if err := p.Header.UnmarshalBinary(b); err != nil {
		return fmt.Errorf("unmarshal packet: %w", err)
	}
	if len(b) < lengthOfPacketHeader+p.Header.FieldSelectors.Length() {
		return fmt.Errorf("unmarshal packet: insufficient data for selectors: %v", p.Header.FieldSelectors)
	}
	i := indexOfSnapshot
	p.EventID = EventID(b[i])
	i += lengthOfEventID
	p.EventInformation = b[i]
	i += lengthOfEventInformation
	if p.Header.FieldSelectors.Has(FieldSelectorGPSFlags) {
		p.GPSFlags = GPSFlags(b[i])
		i += lengthOfGPSFlags
	}
	if p.Header.FieldSelectors.Has(FieldSelectorTime) {
		p.Time = time.Unix(int64(binary.BigEndian.Uint32(b[i:])), 0).UTC()
		i += lengthOfTime
	}
	if p.Header.FieldSelectors.Has(FieldSelectorGPS) {
		p.GPSTime = time.Unix(int64(binary.BigEndian.Uint32(b[i:])), 0).UTC()
		i += lengthOfGPSTime
		p.LatitudeMicroDegrees = binary.BigEndian.Uint32(b[i:])
		i += lengthOfLatitude
		p.LongitudeMicroDegrees = binary.BigEndian.Uint32(b[i:])
		i += lengthOfLongitude
		p.NumSatellites = b[i]
		i += lengthOfNumSatellites
	}
	if p.Header.FieldSelectors.Has(FieldSelectorGPSSpeed) {
		p.SpeedKilometresPerHour = b[i]
		i += lengthOfSpeed
		p.MaximumSpeedKilometresPerHour = b[i]
		i += lengthOfMaximumSpeed
		p.HeadingHalfDegrees = b[i]
		i += lengthOfHeading
	}
	if p.Header.FieldSelectors.Has(FieldSelectorIO) {
		p.DigitalInputs = DigitalInputs(b[i])
		i += lengthOfDigitalInputs
	}
	if p.Header.FieldSelectors.Has(FieldSelectorAnalogInput) {
		p.AnalogInput1VoltageMilliVolts = binary.BigEndian.Uint16(b[i:])
		i += lengthOfAnalogInput1
		p.AnalogInput2VoltageMilliVolts = binary.BigEndian.Uint16(b[i:])
		i += lengthOfAnalogInput2
		p.AnalogInput3VoltageMilliVolts = binary.BigEndian.Uint16(b[i:])
		i += lengthOfAnalogInput3
		p.AnalogInput4VoltageMilliVolts = binary.BigEndian.Uint16(b[i:])
		i += lengthOfAnalogInput4
	}
	if p.Header.FieldSelectors.Has(FieldSelectorPower) {
		p.MainPowerVoltageMilliVolts = binary.BigEndian.Uint16(b[i:])
		i += lengthOfMainPower
		p.ExternalBatteryVoltageMilliVolts = binary.BigEndian.Uint16(b[i:])
		i += lengthOfExternalBattery
	}
	if p.Header.FieldSelectors.Has(FieldSelectorPulseCounter1) {
		p.PulseCounter1Rate = binary.BigEndian.Uint16(b[i:])
		i += lengthOfPulseCounter1Rate
		p.PulseCounter1 = binary.BigEndian.Uint32(b[i:])
		i += lengthOfPulseCounter1
	}
	if p.Header.FieldSelectors.Has(FieldSelectorPulseCounter2) {
		p.PulseCounter2Rate = binary.BigEndian.Uint16(b[i:])
		i += lengthOfPulseCounter2Rate
		p.PulseCounter2 = binary.BigEndian.Uint32(b[i:])
		i += lengthOfPulseCounter2
	}
	if p.Header.FieldSelectors.Has(FieldSelectorTrip1) {
		p.Trip1DistanceMetres = binary.BigEndian.Uint32(b[i:])
		i += lengthOfTrip1Distance
	}
	if p.Header.FieldSelectors.Has(FieldSelectorTrip2) {
		p.Trip2DistanceMetres = binary.BigEndian.Uint32(b[i:])
		i += lengthOfTrip2Distance
	}
	if p.Header.FieldSelectors.Has(FieldSelectorIO) {
		p.OutputStatus = OutputStatus(b[i])
		i += lengthOfOutputStatus
	}
	if p.Header.FieldSelectors.Has(FieldSelectorIButton) {
		p.IButtonKeyID = bigEndianUint48(b[i:])
		i += lengthOfIButtonKeyID
	}
	if p.Header.FieldSelectors.Has(FieldSelectorDriverLogKeypad) {
		p.DriverLogKeypad = DriverLogKeypad(b[i])
		i += lengthOfDriverLogKeypadState
	}
	if p.Header.FieldSelectors.Has(FieldSelectorGPSExtras) {
		p.GPSAltitudeMetres = int16(binary.BigEndian.Uint16(b[i:]))
		i += lengthOfGPSAltitude
	}
	if p.Header.FieldSelectors.Has(FieldSelectorSnapshotCounter) {
		p.SnapshotCounter = binary.BigEndian.Uint16(b[i:])
		i += lengthOfSnapshotCounter
	}
	if p.Header.FieldSelectors.Has(FieldSelectorFlags) {
		p.StateFlags = StateFlags(binary.BigEndian.Uint32(b[i:]))
		i += lengthOfStateFlags
		p.UserDefinedFlags = binary.BigEndian.Uint32(b[i:])
		i += lengthOfUserDefinedFlags
	}
	if p.Header.FieldSelectors.Has(FieldSelectorCellInfo) {
		p.CellInfoFlags = CellInfoFlags(b[i])
		i += lengthOfCellInfoFlags
		p.LocationAreaIdentity = bigEndianUint24(b[i:])
		i += lengthOfLocationAreaIdentity
		p.LocationAreaCode = binary.BigEndian.Uint16(b[i:])
		i += lengthOfLocationAreaCode
		p.GSMCellID = binary.BigEndian.Uint32(b[i:])
		i += lengthOfGSMCellID
		p.ReceivedSignalStrengthIndicator = int8(b[i])
		i += lengthOfReceivedSignalStrengthIndicator
	}
	if p.Header.FieldSelectors.Has(FieldSelectorExtendedDigitalInputs) {
		p.DigitalInputs = DigitalInputs(binary.BigEndian.Uint16(b[i:]))
		i += lengthOfExtendedDigitalInputs
	}
	if p.Header.FieldSelectors.Has(FieldSelectorEventSpecificBytes) {
		p.EventSpecificBytes = b[i:] // remainder of packet
	}
	return nil
}
