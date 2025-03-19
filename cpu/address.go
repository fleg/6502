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
func (cpu *CPU) fetchOperandAddress(am AddressMode) (uint16, bool) {
	switch am {
	case amImp, amAcc:
		return 0, false
	case amImm:
		return cpu.nextPC(), false
	case amZeP:
		return uint16(cpu.readPC()), false
	case amZeX:
		addr := cpu.readPC() + cpu.X
		return uint16(addr), false
	case amZeY:
		addr := cpu.readPC() + cpu.Y
		return uint16(addr), false
	case amRel:
		offset := int8(cpu.readPC())
		addr := uint16(int32(cpu.PC) + int32(offset))
		return addr, isPageCrossed(cpu.PC, addr)
	case amAbs:
		return cpu.readPCWord(), false
	case amAbX:
		base := cpu.readPCWord()
		addr := base + uint16(cpu.X)
		return addr, isPageCrossed(base, addr)
	case amAbY:
		base := cpu.readPCWord()
		addr := base + uint16(cpu.Y)
		return addr, isPageCrossed(base, addr)
	case amInd:
		return cpu.readWordWithoutPageCross(cpu.readPCWord()), false
	case amInX:
		addr := uint16(cpu.readPC() + cpu.X)
		return cpu.readWordWithoutPageCross(addr), false
	case amInY:
		base := cpu.readWordWithoutPageCross(uint16(cpu.readPC()))
		addr := base + uint16(cpu.Y)
		return addr, isPageCrossed(base, addr)
	default:
		panic("Unknown address mode")
	}
}

func isPageCrossed(a uint16, b uint16) bool {
	return a&0xff00 != b&0xff00
}
