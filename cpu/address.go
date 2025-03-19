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

// The processor is little endian and expects
// addresses to be stored in memory least significant byte first.
func (cpu *CPU) fetchOperandAddress(am AddressMode) uint16 {
	switch am {
	case amImp, amAcc:
		return 0
	case amImm:
		return cpu.nextPC()
	case amZeP:
		return uint16(cpu.readPC())
	case amZeX:
		addr := cpu.readPC() + cpu.X
		return uint16(addr)
	case amZeY:
		addr := cpu.readPC() + cpu.Y
		return uint16(addr)
	case amRel:
		offset := int8(cpu.readPC())
		return uint16(int32(cpu.PC) + int32(offset))
	case amAbs:
		return cpu.readPCWord()
	case amAbX:
		return cpu.readPCWord() + uint16(cpu.X)
	case amAbY:
		return cpu.readPCWord() + uint16(cpu.Y)
	case amInd:
		return cpu.readWordWithoutPageCross(cpu.readPCWord())
	case amInX:
		addr := uint16(cpu.readPC() + cpu.X)
		return cpu.readWordWithoutPageCross(addr)
	case amInY:
		addr := cpu.readWordWithoutPageCross(uint16(cpu.readPC()))
		return addr + uint16(cpu.Y)
	default:
		panic("Unknown address mode")
	}
}
