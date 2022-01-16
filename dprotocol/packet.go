package dprotocol

import (
	"encoding/binary"
	"fmt"
	"time"
)

// Packet represents a D protocol packet.
type Packet struct {
	// Header is the packet header.
	Header Header

	// EventID is the event ID.
	EventID EventID

	// EventInformation contains additional, event-dependent information.
	EventInformation uint8

	// GPSFlags contains the GPS flags.
	GPSFlags GPSFlags

	// Time is the time the snapshot was recorded.
	// Does not indicate snapshot sending time.
	// Snapshot recorded time starts running from zero if the device doesn't have
	// GPS fix at startup and it is updated to correct time when GPS fix is received.
	Time time.Time

	// GPSTime is the time when GPS position data was recorded..
	GPSTime time.Time

	// LatitudeMicroDegrees is the GPS latitude in millionths of a degree.
	// Southbound is negative, northbound is positive.
	LatitudeMicroDegrees uint32

	// LongitudeMicroDegrees is the GPS longitude in millionths of a degree.
	// Westbound is negative, eastbound is positive.
	LongitudeMicroDegrees uint32

	// NumSatellites is the number of visible satellites.
	NumSatellites uint8

	// SpeedKph is the current speed in km/h.
	// Note: Wrap-around in speeds over 255km/h.
	SpeedKph uint8

	// MaximumSpeedKph is the maximum detected speed since last event in km/h.
	MaximumSpeedKph uint8

	// HeadingHalfDegrees is the vehicle heading in degrees / 2. Multiply value by 2 to get degrees.
	// For example, 260 degrees is sent as a value of 130.
	// 0 or 360 degrees equals heading to north.
	HeadingHalfDegrees uint8

	// DigitalInputs contains the digital input status.
	DigitalInputs DigitalInputs

	// AnalogInput1VoltageMilliVolts is the voltage of analog input AD1 in millivolts.
	AnalogInput1VoltageMilliVolts uint16

	// AnalogInput2VoltageMilliVolts is the voltage of analog input AD2 in millivolts.
	AnalogInput2VoltageMilliVolts uint16

	// AnalogInput3VoltageMilliVolts is the voltage of analog input AD3 in millivolts.
	AnalogInput3VoltageMilliVolts uint16

	// AnalogInput4VoltageMilliVolts is the voltage of analog input AD4 in millivolts.
	AnalogInput4VoltageMilliVolts uint16

	// MainPowerVoltageMilliVolts is the voltage of main power in millivolts.
	MainPowerVoltageMilliVolts uint16

	// ExternalBatteryVoltageMilliVolts is the voltage of external battery in millivolts.
	ExternalBatteryVoltageMilliVolts uint16

	// PulseCounter1Rate is the latest pulse rate of counter channel 1.
	// Unit depends on configuration. Default is pulses per second (PPS).
	PulseCounter1Rate uint16

	// PulseCounter1 is the counter for pulse counter channel 1.
	// Unit depends on configuration. Default is number of pulses.
	PulseCounter1 uint32

	// PulseCounter2Rate is the latest pulse rate of counter channel 2.
	// Unit depends on configuration. Default is pulses per second (PPS).
	PulseCounter2Rate uint16

	// PulseCounter2 is the counter for pulse counter channel 2.
	// Unit depends on configuration. Default is number of pulses.
	PulseCounter2 uint32

	// Trip1DistanceMeters is the distance traveled in meters since trip meter was reset.
	Trip1DistanceMetres uint32

	// Trip2DistanceMeters is the distance traveled in meters since trip meter was reset.
	Trip2DistanceMetres uint32

	// OutputStatus contains the status of outputs.
	OutputStatus OutputStatus

	// IButtonKeyID is the iButton key ID without family code or checksum.
	// iButton key ID is all zeroes if driver was not logged in at the time the snapshot was recorded.
	IButtonKeyID uint64

	// DriverLogKeypad is the Driver Log Keypad (DLKP) / 3PAD button pressed state.
	DriverLogKeypad DriverLogKeypad

	// GPSAltituteMeters is the GPS altitude in meters. Antenna height above/below mean sea level.
	GPSAltitudeMetres int16

	// SnapshotCounter is the transport-based snapshot counter.
	// Counter is incremented for every snapshot at message formatting time, and it is destination specific.
	// Note that formatting time is not send time.
	SnapshotCounter uint16

	// StateFlags contains the state flags.
	StateFlags StateFlags

	// UserDefinedFlags contains the user-definable flag bits.
	UserDefinedFlags uint32

	// CellInfoFlags contains the cell info.
	CellInfoFlags CellInfoFlags

	// LocationAreaIdentity contains the Location Area Identity (LAI).
	LocationAreaIdentity uint32

	// LocationAreaCode contains the Location Area Code (LAC).
	LocationAreaCode uint16

	// GSMCellID is the GSM cell ID.
	GSMCellID uint32

	// RSSI is the Received Signal Strength Indicator (RSSI).
	RSSI int8

	// EventSpecificBytes contains the event-specific bytes.
	EventSpecificBytes []byte
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

// UnmarshalBinary unmarshals the packet from the provided bytes.
func (p *Packet) UnmarshalBinary(b []byte) error {
	if err := p.Header.UnmarshalBinary(b); err != nil {
		return fmt.Errorf("unmarshal packet: %w", err)
	}
	if len(b) < lengthOfPacketHeader+p.Header.FieldSelectors.SnapshotLength() {
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
		p.SpeedKph = b[i]
		i += lengthOfSpeed
		p.MaximumSpeedKph = b[i]
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
		p.RSSI = int8(b[i])
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
