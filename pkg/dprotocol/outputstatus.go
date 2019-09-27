package dprotocol

type OutputStatus uint8

func (o OutputStatus) Output1() Output {
	return Output(o & 0b11)
}

func (o OutputStatus) Output2() Output {
	return Output((o >> 2) & 0b11)
}
