package protocol

import (
	"bytes"
	"encoding/binary"

	"golang.org/x/xerrors"
)

// CellInfo Gsm info
type CellInfo struct {
	Network byte
	LAI     [3]byte
	LAC     uint16
	RSSI    int8
	CID     uint32
}

// PacketV3 Aplicom D protocol version 3
type PacketV3 struct {
	ProtocolIdentifier    uint8
	ProtocolVersion       byte
	UnitIDHighBytes       [4]byte
	SnapshotDataLength    uint16
	UnitIDLowBytes        [3]byte
	SelectorBits          SelectorBits
	EventID               uint8
	EventInformation      uint8
	Time                  int32
	GPSTime               int32
	Latitude              int32
	Longitude             int32
	DataValidity          uint8
	NrOfSatellites        uint8
	Speed                 uint8
	MaxSpeed              uint8
	Heading               uint8
	DINStatus             byte
	AD1                   uint16
	AD2                   uint16
	AD3                   uint16
	AD4                   uint16
	MainPower             uint16
	EXTBattery            uint16
	PCNT1Rate             uint16
	PCNT1Counter          uint32
	PCNT2Rate             uint16
	GPSAltitude           int16
	PCNT2Counter          uint32
	Trip1Distance         uint32
	Trip2Distance         uint32
	OutputStatus          byte
	IButtonKeyID          [6]byte
	DLKP3PADState         uint8
	FlagBits              [8]byte
	SnapshotCounter       uint16
	ExtDigitalInputStatus [2]byte
	CellInfo              CellInfo
	EventSpecificBytes    [256]byte // MAX GEOFENCE_DYNAMIC can use 256bytes,
}

const (
	selectorBitsTimestamp         = 0x000004 // Timestamp. Snapshot recording time.
	selectorBitsGPSDataValidity   = 0x000008 // GPS. Latitude, longitude, GPS time, satellite count, data validity.
	selectorBitsGPSSpeed          = 0x000010 // GPS Speed. Speed, Max speed, heading.
	selectorBitsAD                = 0x000020 // AD. AD1 – AD4.
	selectorBitsIO                = 0x000040 // IO. DIN status, output status.
	selectorBitsTrip1             = 0x000080 // Trip1.
	selectorBitsTrip2             = 0x000100 // Trip2.
	selectorBitsiButton           = 0x000200 // iButton. iButton ID.
	selectorBitsDLKP3PAD          = 0x000400 // 3PAD / DLKP state
	selectorBitsGSPExtras         = 0x000800 // GPS extras. Altitude.
	selectorBitsEvent             = 0x001000 // Event specific additional bytes. Appended to the very end of the message.
	selectorBitsFlagBits          = 0x004000 // Flag bits
	selectorBitsPower             = 0x008000 // Power
	selectorBitsPulseCounter1     = 0x010000 // Pulse counter1 values.
	selectorBitsPulseCounter2     = 0x020000 // Pulse counter2 values.
	selectorBitsCellInfo          = 0x080000 // Cell Info
	selectorBitsExtendedDINStatus = 0x100000 // Extended DIN status. Contains status of all digital inputs and ignition.
)

// SelectorBits defines which info is included in packet
type SelectorBits [3]byte

const headerNrBytes = 14

func (s *SelectorBits) hasTime() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsTimestamp > 0
}

func (s *SelectorBits) hasGPSDataValidity() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsGPSDataValidity > 0
}

func (s *SelectorBits) hasGPSSpeed() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsGPSSpeed > 0
}

func (s *SelectorBits) hasAD() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsAD > 0
}

func (s *SelectorBits) hasIO() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsIO > 0
}

func (s *SelectorBits) hasTrip1() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsTrip1 > 0
}

func (s *SelectorBits) hasTrip2() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsTrip2 > 0
}

func (s *SelectorBits) hasIButton() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsiButton > 0
}

func (s *SelectorBits) hasDLKP3PAD() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsDLKP3PAD > 0
}

func (s *SelectorBits) hasGSPExtras() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsGSPExtras > 0
}

func (s *SelectorBits) hasEvent() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsEvent > 0
}

func (s *SelectorBits) hasFlagBits() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsFlagBits > 0
}

func (s *SelectorBits) hasPower() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsPower > 0
}

func (s *SelectorBits) hasPulseCounter1() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsPulseCounter1 > 0
}

func (s *SelectorBits) hasPulseCounter2() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsPulseCounter2 > 0
}

func (s *SelectorBits) hasCellInfo() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsCellInfo > 0
}

func (s *SelectorBits) hasExtendedDINStatus() bool {
	return binary.BigEndian.Uint32(append([]byte{0}, s[:]...))&selectorBitsExtendedDINStatus > 0
}

func (p *PacketV3) setUnitID(b []byte) int {
	i := 0
	p.UnitIDHighBytes[0] = b[i]
	i++
	p.UnitIDHighBytes[1] = b[i]
	i++
	p.UnitIDHighBytes[2] = b[i]
	i++
	p.UnitIDHighBytes[3] = b[i]
	i++
	p.UnitIDLowBytes[0] = b[i]
	i++
	p.UnitIDLowBytes[1] = b[i]
	i++
	p.UnitIDLowBytes[2] = b[i]
	i++
	return i
}

func (p *PacketV3) setSelectorBits(b []byte) int {
	i := 0
	p.SelectorBits[0] = b[i]
	i++
	p.SelectorBits[1] = b[i]
	i++
	p.SelectorBits[2] = b[i]
	i++
	return i
}

func (p *PacketV3) setGPSDataValidity(b []byte) int {
	i := 0
	p.GPSTime = int32(binary.BigEndian.Uint32(b[i : i+4]))
	i += 4
	p.Latitude = int32(binary.BigEndian.Uint32(b[i : i+4]))
	i += 4
	p.Longitude = int32(binary.BigEndian.Uint32(b[i : i+4]))
	i += 4
	p.NrOfSatellites = b[i]
	i++
	return i
}

func (p *PacketV3) setSpeed(b []byte) int {
	i := 0
	p.Speed = b[i]
	i++
	p.MaxSpeed = b[i]
	i++
	p.Heading = b[i]
	i++
	return i
}

func (p *PacketV3) setAD(b []byte) int {
	i := 0
	p.AD1 = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	p.AD2 = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	p.AD3 = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	p.AD4 = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	return i
}

func (p *PacketV3) setPower(b []byte) int {
	i := 0
	p.MainPower = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	p.EXTBattery = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	return i
}

func (p *PacketV3) setPulseCounter1(b []byte) int {
	i := 0
	p.PCNT1Rate = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	p.PCNT1Counter = binary.BigEndian.Uint32(b[i : i+4])
	i += 4
	return i
}

func (p *PacketV3) setPulseCounter2(b []byte) int {
	i := 0
	p.PCNT2Rate = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	p.PCNT2Counter = binary.BigEndian.Uint32(b[i : i+4])
	i += 4
	return i
}

func (p *PacketV3) setIButton(b []byte) int {
	i := 0
	p.IButtonKeyID[0] = b[i]
	i++
	p.IButtonKeyID[1] = b[i]
	i++
	p.IButtonKeyID[2] = b[i]
	i++
	p.IButtonKeyID[3] = b[i]
	i++
	p.IButtonKeyID[4] = b[i]
	i++
	p.IButtonKeyID[5] = b[i]
	i++
	return i
}

func (p *PacketV3) setFlagBits(b []byte) int {
	i := 0
	p.FlagBits[0] = b[i]
	i++
	p.FlagBits[1] = b[i]
	i++
	p.FlagBits[2] = b[i]
	i++
	p.FlagBits[3] = b[i]
	i++
	p.FlagBits[4] = b[i]
	i++
	p.FlagBits[5] = b[i]
	i++
	p.FlagBits[6] = b[i]
	i++
	p.FlagBits[7] = b[i]
	i++
	return i
}

func (p *PacketV3) setCellInfo(b []byte) int {
	i := 0
	p.CellInfo.Network = b[i]
	i++
	p.CellInfo.LAI[0] = b[i]
	i++
	p.CellInfo.LAI[1] = b[i]
	i++
	p.CellInfo.LAI[2] = b[i]
	i++
	p.CellInfo.LAC = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	p.CellInfo.CID = binary.BigEndian.Uint32(b[i : i+4])
	i += 4
	p.CellInfo.RSSI = int8(b[i])
	i++
	return i
}

func (p *PacketV3) setExtDigitalInputStatus(b []byte) int {
	i := 0
	p.ExtDigitalInputStatus[0] = b[i]
	i++
	p.ExtDigitalInputStatus[1] = b[i]
	i++
	return i
}

func (p *PacketV3) setGPSData(b []byte) int {
	i := 0
	if p.SelectorBits.hasGPSDataValidity() {
		p.DataValidity = b[i]
		i++
	}
	if p.SelectorBits.hasTime() {
		p.Time = int32(binary.BigEndian.Uint32(b[i : i+4]))
		i += 4
	}
	if p.SelectorBits.hasGPSDataValidity() {
		i += p.setGPSDataValidity(b[i:])
	}
	if p.SelectorBits.hasGPSSpeed() {
		i += p.setSpeed(b[i:])
	}
	return i
}

func (p *PacketV3) setPowerData(b []byte) int {
	i := 0
	if p.SelectorBits.hasIO() {
		p.DINStatus = b[i]
		i++
	}
	if p.SelectorBits.hasAD() {
		i += p.setAD(b[i:])
	}
	if p.SelectorBits.hasPower() {
		i += p.setPower(b[i:])
	}
	if p.SelectorBits.hasPulseCounter1() {
		i += p.setPulseCounter1(b[i:])
	}
	if p.SelectorBits.hasPulseCounter2() {
		i += p.setPulseCounter2(b[i:])
	}

	return i
}

func (p *PacketV3) setDistanceData(b []byte) int {
	i := 0
	if p.SelectorBits.hasTrip1() {
		p.Trip1Distance = binary.BigEndian.Uint32(b[i : i+4])
		i += 4
	}
	if p.SelectorBits.hasTrip2() {
		p.Trip2Distance = binary.BigEndian.Uint32(b[i : i+4])
		i += 4
	}
	if p.SelectorBits.hasIO() {
		p.OutputStatus = b[i]
		i++
	}
	return i
}

func (p *PacketV3) setExternalToolData(b []byte) int {
	i := 0
	if p.SelectorBits.hasIButton() {
		i += p.setIButton(b[i:])
	}
	if p.SelectorBits.hasDLKP3PAD() {
		p.DLKP3PADState = b[i]
		i++
	}
	return i
}

func (p *PacketV3) setEventData(b []byte) int {
	i := 0
	p.EventID = b[i]
	i++
	p.EventInformation = b[i]
	i++
	return i
}

func (p *PacketV3) setPacketHeaderData(b []byte) int {
	i := 0
	p.ProtocolIdentifier = b[i]
	i++
	p.ProtocolVersion = b[i]
	i++
	i += p.setUnitID(b[i:])
	p.SnapshotDataLength = binary.BigEndian.Uint16(b[i : i+2]) // not including header 14 bytes
	i += 2
	i += p.setSelectorBits(b[i:])
	return i
}

// UnmarshalBinary  Unmarshal binary data
func (p *PacketV3) UnmarshalBinary(b []byte) error {
	// Cannot use binary.Read since protocol optimizes struct, no fixed values
	i := p.setPacketHeaderData(b)
	i += p.setEventData(b[i:])
	i += p.setGPSData(b[i:])
	i += p.setPowerData(b[i:])
	i += p.setDistanceData(b[i:])
	i += p.setExternalToolData(b[i:])
	if p.SelectorBits.hasGSPExtras() {
		p.GPSAltitude = int16(binary.BigEndian.Uint16(b[i : i+2]))
		i += 2
	}
	p.SnapshotCounter = binary.BigEndian.Uint16(b[i : i+2])
	i += 2
	if p.SelectorBits.hasFlagBits() {
		i += p.setFlagBits(b[i:])
	}
	if p.SelectorBits.hasCellInfo() {
		i += p.setCellInfo(b[i:])
	}
	if p.SelectorBits.hasExtendedDINStatus() {
		i += p.setExtDigitalInputStatus(b[i:])
	}
	if p.SelectorBits.hasEvent() { // Event specific bytes
		bytesLeft := int(p.SnapshotDataLength) + headerNrBytes - i
		eventLength := len(p.EventSpecificBytes[:])
		if bytesLeft < eventLength {
			copy(p.EventSpecificBytes[:], b[i:i+bytesLeft])
		} else {
			copy(p.EventSpecificBytes[:], b[i:i+eventLength])
		}
	}
	return nil
}

// ScanAplicomPackets Scan aplicom binary data
func (p *PacketV3) ScanAplicomPackets(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, nil, nil
	}
	startIndex := bytes.IndexByte(data, 0x44)
	if startIndex == -1 {
		return len(data), nil, nil
	}
	if startIndex != 0 {
		return startIndex, nil, nil
	}
	if len(data) < 11 { // protocol v2 and upwards
		return 0, nil, nil
	}
	protocolVersion := data[1] & 0x0F
	length := 0
	if protocolVersion < 3 {
		return 0, nil, xerrors.Errorf("protocol version not supported")
	}
	length = int(binary.BigEndian.Uint16(data[9:11])) + headerNrBytes
	if len(data) < length {
		return 0, nil, nil
	}
	return length, data[:length], nil
}
