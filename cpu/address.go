package cpu

type AddressMode uint8

const (
	amImp AddressMode = iota
	amAcc
	amImm
	amZeP
	amZeX
	amZeY
	amRel
	amAbs
	amAbX
	amAbY
	amInd
	amInX
	amInY
)
