package dprotocol

// EventID represents an event ID.
type EventID uint8

//go:generate stringer -type EventID -trimprefix EventID

// EventID values.
const (
	EventIDInputChanged                   EventID = 0x02
	EventIDOutputChanged                  EventID = 0x03
	EventIDGpsStatusChanged               EventID = 0x05
	EventIDIgnOn                          EventID = 0x07
	EventIDIgnOff                         EventID = 0x08
	EventIDNetChanged                     EventID = 0x09
	EventIDIButton                        EventID = 0x0B
	EventIDPowerSupplyChanged             EventID = 0x10
	EventIDTempOK                         EventID = 0x11
	EventIDTempHigh                       EventID = 0x12
	EventIDTempLow                        EventID = 0x13
	EventIDBatteryLow                     EventID = 0x14
	EventIDBatteryOK                      EventID = 0x15
	EventIDTachoEvent                     EventID = 0x17
	EventIDAdSampling                     EventID = 0x1e
	EventIDFuelSampling                   EventID = 0x1f
	EventIDAccOrientation                 EventID = 0x26
	EventIDSystemInfo                     EventID = 0x28
	EventIDAccelerometerRapidAcceleration EventID = 0x32
	EventIDAccelerometerHashBraking       EventID = 0x33
	EventIDAccelerometerLeftAcceleration  EventID = 0x34
	EventIDAccelerometerRightAcceleration EventID = 0x35
	EventIDSoftwareStart                  EventID = 0x65
	EventIDSoftwareStop                   EventID = 0x66
	EventIDStartMoving                    EventID = 0x68
	EventIDStopMoving                     EventID = 0x69
	EventIDDirectionChanged               EventID = 0x6a
	EventIDGeofence                       EventID = 0x6b
	EventIDGeofenceDynamic                EventID = 0x6c
	EventIDAlarmActive                    EventID = 0x6d
	EventIDScheduledEvent                 EventID = 0x6e
	EventIDSpeedLimit                     EventID = 0x6f
	EventIDDistanceTraveled               EventID = 0x70
	EventIDADThreshold                    EventID = 0x71
	EventIDFMSOverspeed                   EventID = 0x72
	EventIDFMSOvertemp                    EventID = 0x73
	EventIDFMSOverrevolutions             EventID = 0x74
	EventIDFMSHarshBraking                EventID = 0x75
	EventIDFMSCruiseControl               EventID = 0x76
	EventIDDataEvent                      EventID = 0x77
	EventIDCommFail                       EventID = 0x78
	EventIDFlagChanged                    EventID = 0x79
	EventIDDLKPEmergencyPress             EventID = 0x7a
	EventIDDLKPStateChanged               EventID = 0x7b
	EventIDGPSHarshBraking                EventID = 0x7c
	EventIDGPSRapidAcceleration           EventID = 0x7d
	EventIDCommSessionClosed              EventID = 0x81
	EventIDPCNTData                       EventID = 0x82
	EventIDPCNTLimit                      EventID = 0x83
	EventIDPCNTState                      EventID = 0x84
	EventIDVoiceCall                      EventID = 0x8c
	EventIDHistogramEvent                 EventID = 0x8e
	EventIDFuelLevelEvent                 EventID = 0x8f
	EventIDGarminEvent                    EventID = 0xa0
	EventIDCardDL                         EventID = 0xaa
	EventIDTachoDL                        EventID = 0xab
	EventIDTemperatureDataEvent           EventID = 0xba
	EventIDEBSDataEvent                   EventID = 0xbc
	EventIDCOPUpdate                      EventID = 0xff
)
