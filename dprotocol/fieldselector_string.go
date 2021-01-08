// Code generated by "stringer -type FieldSelector -trimprefix FieldSelector"; DO NOT EDIT.

package dprotocol

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FieldSelectorGPSFlags-8]
	_ = x[FieldSelectorTime-4]
	_ = x[FieldSelectorGPS-8]
	_ = x[FieldSelectorGPSSpeed-16]
	_ = x[FieldSelectorAnalogInput-32]
	_ = x[FieldSelectorIO-64]
	_ = x[FieldSelectorTrip1-128]
	_ = x[FieldSelectorTrip2-256]
	_ = x[FieldSelectorIButton-512]
	_ = x[FieldSelectorDriverLogKeypad-1024]
	_ = x[FieldSelectorGPSExtras-2048]
	_ = x[FieldSelectorEventSpecificBytes-4096]
	_ = x[FieldSelectorSnapshotCounter-8192]
	_ = x[FieldSelectorFlags-16384]
	_ = x[FieldSelectorPower-32768]
	_ = x[FieldSelectorPulseCounter1-65536]
	_ = x[FieldSelectorPulseCounter2-131072]
	_ = x[FieldSelectorCellInfo-524288]
	_ = x[FieldSelectorExtendedDigitalInputs-1048576]
}

const _FieldSelector_name = "TimeGPSFlagsGPSSpeedAnalogInputIOTrip1Trip2IButtonDriverLogKeypadGPSExtrasEventSpecificBytesSnapshotCounterFlagsPowerPulseCounter1PulseCounter2CellInfoExtendedDigitalInputs"

var _FieldSelector_map = map[FieldSelector]string{
	4:       _FieldSelector_name[0:4],
	8:       _FieldSelector_name[4:12],
	16:      _FieldSelector_name[12:20],
	32:      _FieldSelector_name[20:31],
	64:      _FieldSelector_name[31:33],
	128:     _FieldSelector_name[33:38],
	256:     _FieldSelector_name[38:43],
	512:     _FieldSelector_name[43:50],
	1024:    _FieldSelector_name[50:65],
	2048:    _FieldSelector_name[65:74],
	4096:    _FieldSelector_name[74:92],
	8192:    _FieldSelector_name[92:107],
	16384:   _FieldSelector_name[107:112],
	32768:   _FieldSelector_name[112:117],
	65536:   _FieldSelector_name[117:130],
	131072:  _FieldSelector_name[130:143],
	524288:  _FieldSelector_name[143:151],
	1048576: _FieldSelector_name[151:172],
}

func (i FieldSelector) String() string {
	if str, ok := _FieldSelector_map[i]; ok {
		return str
	}
	return "FieldSelector(" + strconv.FormatInt(int64(i), 10) + ")"
}