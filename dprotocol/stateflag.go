package dprotocol

// StateFlag represents a state flag.
type StateFlag uint32

//go:generate stringer -type StateFlag -trimprefix StateFlag

const (
	StateFlagCardDL         StateFlag = 0x01
	StateFlagTachographDL   StateFlag = 0x02
	StateFlagIO7            StateFlag = 0x08
	StateFlagIO8            StateFlag = 0x10
	StateFlagIO9            StateFlag = 0x20
	StateFlagIO10           StateFlag = 0x40
	StateFlagIO1            StateFlag = 0x0100
	StateFlagIO2            StateFlag = 0x0200
	StateFlagIO3            StateFlag = 0x0400
	StateFlagIO4            StateFlag = 0x0800
	StateFlagIO5            StateFlag = 0x1000
	StateFlagIO6            StateFlag = 0x2000
	StateFlagProxyUsage     StateFlag = 0x4000
	StateFlagGSMGPRSJamming StateFlag = 0x010000
	StateFlagGPSJamming     StateFlag = 0x020000
	StateFlag3PADExtraLED1  StateFlag = 0x080000
	StateFlag3PADExtraLED2  StateFlag = 0x100000
	StateFlag3PADButtonLED1 StateFlag = 0x200000
	StateFlag3PADButtonLED2 StateFlag = 0x400000
	StateFlag3PADButtonLED3 StateFlag = 0x800000
)
