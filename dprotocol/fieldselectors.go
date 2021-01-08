package dprotocol

// FieldSelectors is a mask which enables optimization of sent data by dropping out unnecessary fields.
type FieldSelectors uint32

// Has returns true if the provided FieldSelector is included among the field selectors.
func (f FieldSelectors) Has(selector FieldSelector) bool {
	return FieldSelector(f)&selector == selector
}

// SnapshotLength returns the length in bytes of a snapshot with the current field selectors.
func (f FieldSelectors) SnapshotLength() int {
	result := lengthOfEventID + lengthOfEventInformation
	if f.Has(FieldSelectorGPSFlags) {
		result += lengthOfGPSFlags
	}
	if f.Has(FieldSelectorTime) {
		result += lengthOfTime
	}
	if f.Has(FieldSelectorGPS) {
		result += lengthOfGPSTime + lengthOfLatitude + lengthOfLongitude + lengthOfNumSatellites
	}
	if f.Has(FieldSelectorGPSSpeed) {
		result += lengthOfSpeed + lengthOfMaximumSpeed + lengthOfHeading
	}
	if f.Has(FieldSelectorIO) {
		result += lengthOfDigitalInputs + lengthOfOutputStatus
	}
	if f.Has(FieldSelectorAnalogInput) {
		result += lengthOfAnalogInput1 + lengthOfAnalogInput2 + lengthOfAnalogInput3 + lengthOfAnalogInput4
	}
	if f.Has(FieldSelectorPower) {
		result += lengthOfMainPower + lengthOfExternalBattery
	}
	if f.Has(FieldSelectorPulseCounter1) {
		result += lengthOfPulseCounter1Rate + lengthOfPulseCounter1
	}
	if f.Has(FieldSelectorPulseCounter2) {
		result += lengthOfPulseCounter2Rate + lengthOfPulseCounter2
	}
	if f.Has(FieldSelectorTrip1) {
		result += lengthOfTrip1Distance
	}
	if f.Has(FieldSelectorTrip2) {
		result += lengthOfTrip2Distance
	}
	if f.Has(FieldSelectorIButton) {
		result += lengthOfIButtonKeyID
	}
	if f.Has(FieldSelectorDriverLogKeypad) {
		result += lengthOfDriverLogKeypadState
	}
	if f.Has(FieldSelectorGPSExtras) {
		result += lengthOfGPSAltitude
	}
	if f.Has(FieldSelectorSnapshotCounter) {
		result += lengthOfSnapshotCounter
	}
	if f.Has(FieldSelectorFlags) {
		result += lengthOfStateFlags + lengthOfUserDefinedFlags
	}
	if f.Has(FieldSelectorCellInfo) {
		result += lengthOfCellInfoFlags + lengthOfLocationAreaIdentity + lengthOfLocationAreaCode +
			lengthOfGSMCellID + lengthOfReceivedSignalStrengthIndicator
	}
	if f.Has(FieldSelectorExtendedDigitalInputs) {
		result += lengthOfExtendedDigitalInputs
	}
	return result
}
