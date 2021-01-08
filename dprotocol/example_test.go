package dprotocol

import "time"

func getExamplePacket() *Packet {
	return &Packet{
		Header:                     getExampleHeader(),
		EventID:                    0x32,
		GPSFlags:                   0xd0,
		Time:                       time.Unix(0x450ef906, 0).UTC(),
		GPSTime:                    time.Unix(0x450ef906, 0).UTC(),
		LatitudeMicroDegrees:       0x03b20b80,
		LongitudeMicroDegrees:      0x03b20b80,
		NumSatellites:              0x10,
		SpeedKph:                   0x2c,
		MaximumSpeedKph:            0x2c,
		HeadingHalfDegrees:         0x19,
		DigitalInputs:              0x87,
		MainPowerVoltageMilliVolts: 0x3756,
		PulseCounter1Rate:          0x32,
		PulseCounter1:              0x1520,
		PulseCounter2:              0x30aa2,
		Trip1DistanceMetres:        0x3e9b5,
		Trip2DistanceMetres:        0x6240,
		OutputStatus:               0x3,
		IButtonKeyID:               0xbcb17ab,
		DriverLogKeypad:            0xff,
		GPSAltitudeMetres:          0x66,
		SnapshotCounter:            0xc9,
		CellInfoFlags:              0x1,
		LocationAreaIdentity:       0x5fab,
		LocationAreaCode:           0x12c6,
		GSMCellID:                  0x24e1,
		RSSI:                       0x17,
		EventSpecificBytes:         []byte{0x00, 0xb7, 0x01, 0xce, 0x0a, 0x28},
	}
}

func getExampleHeader() Header {
	return Header{
		Flags:          HeaderFlags(0xc3),
		UnitID:         0x0143720729d684,
		SnapshotLength: 0x5d,
		FieldSelectors: 0xbffff,
	}
}

func getExampleData() []byte {
	return []byte{
		// 'D' protocol identifier
		0x44,
		// protocol version identifier (both long unit ID and field selector bits are included).
		0xc3,
		// extended unit id
		0x01, 0x43, 0x72, 0x07,
		// unit ID (of value 355632002225796)
		0x29, 0xD6, 0x84,
		// payload data length: 93 bytes.
		0x00, 0x5d,
		// selector bits (14 byte header).
		0x0b, 0xff, 0xff,
		// event ID. ACC_RAPID_ACCELERATION.
		0x32,
		// event information.
		0x00,
		// fix validity (current fix and data fix are valid, GPS fix when dual-mode GPS-GLONASS receiver is available)
		0xd0,
		// time
		0x45, 0x0e, 0xf9, 0x06,
		// GPS time
		0x45, 0x0e, 0xf9, 0x06,
		// latitude 62.000000
		0x03, 0xb2, 0x0b, 0x80,
		// longitude 62.000000
		0x03, 0xb2, 0x0b, 0x80,
		// number of satellites
		0x10,
		// speed 44km/h
		0x2c,
		// max speed 44km/h
		0x2c,
		// heading 50 degrees
		0x19,
		// input status. IGN and DINs 1-3 on, others off
		0x87,
		// AD1 0mV
		0x00, 0x00,
		// AD2 0mV
		0x00, 0x00,
		// AD3 0mV
		0x00, 0x00,
		// AD4 0mV
		0x00, 0x00,
		// main power 14166 mV
		0x37, 0x56,
		// external power 0mV
		0x00, 0x00,
		// PCNT1 rate, 50 pulses per second.
		0x00, 0x32,
		// PCNT1 counter, 5408 pulses received.
		0x00, 0x00, 0x15, 0x20,
		// PCNT2 rate, 0 pulses per second.
		0x00, 0x00,
		// PCNT2 counter, 199330 pulses received.
		0x00, 0x03, 0x0a, 0xa2,
		// trip 1 distance 256437 meters
		0x00, 0x03, 0xe9, 0xb5,
		// trip 2 distance 25152 meters
		0x00, 0x00, 0x62, 0x40,
		// output status: Out 1 active, Out 2 off.
		0x03,
		// iButton key ID 00000BCB17AB
		0x00, 0x00, 0x0b, 0xcb, 0x17, 0xab,
		// DLKP not available
		0xff,
		// GPS altitude 102 meters above sea level
		0x00, 0x66,
		// snapshot counter, 201
		0x00, 0xc9,
		// flag bits
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// cell info, validity 1, LAI 24491, LAC 4806, CID 9441, RSSI 23
		0x01, 0x00, 0x5f, 0xab, 0x12, 0xc6, 0x00, 0x00, 0x24, 0xe1, 0x17,
		// event specific data: ACC_RAPID_ACCELERATION: max acceleration 183mG, max speed 46,2km/h, duration 2600 ms
		0x00, 0xb7, 0x01, 0xce, 0x0a, 0x28,
	}
}
