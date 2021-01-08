package dprotocol

// OutputStatus represents the output status flags.
type OutputStatus uint8

// Output1 returns the output status of output 1.
func (o OutputStatus) Output1() Output {
	return Output(o & 0b11)
}

// Output2 returns the output status of output 2.
func (o OutputStatus) Output2() Output {
	return Output((o >> 2) & 0b11)
}
