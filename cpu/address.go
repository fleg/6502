package cpu

type AddressMode uint8

const (
	amImpilict AddressMode = iota
	amAccumulator
	amImmediate
	amZeroPage
	amZeroPageX
	amZeroPageY
	amRelative
	amAbsolute
	amAbsoluteX
	amAbsoluteY
	amIndirect
	amIndirectX
	amIndirectY
)
