package cpu

type AddressMode uint8

const (
	AmImp AddressMode = iota
	AmAcc
	AmImm
	AmZeP
	AmZeX
	AmZeY
	AmRel
	AmAbs
	AmAbX
	AmAbY
	AmInd
	AmInX
	AmInY
)

// The processor is little endian and expects
// addresses to be stored in memory least significant byte first.
func (cpu *CPU) fetchOperandAddress(am AddressMode) (uint16, bool) {
	switch am {
	case AmImp, AmAcc:
		return 0, false
	case AmImm:
		return cpu.nextPC(), false
	case AmZeP:
		return uint16(cpu.readPC()), false
	case AmZeX:
		addr := cpu.readPC() + cpu.X
		return uint16(addr), false
	case AmZeY:
		addr := cpu.readPC() + cpu.Y
		return uint16(addr), false
	case AmRel:
		offset := int8(cpu.readPC())
		addr := uint16(int32(cpu.PC) + int32(offset))
		return addr, isPageCrossed(cpu.PC, addr)
	case AmAbs:
		return cpu.readPCWord(), false
	case AmAbX:
		base := cpu.readPCWord()
		addr := base + uint16(cpu.X)
		return addr, isPageCrossed(base, addr)
	case AmAbY:
		base := cpu.readPCWord()
		addr := base + uint16(cpu.Y)
		return addr, isPageCrossed(base, addr)
	case AmInd:
		return cpu.readWordWithoutPageCross(cpu.readPCWord()), false
	case AmInX:
		addr := uint16(cpu.readPC() + cpu.X)
		return cpu.readWordWithoutPageCross(addr), false
	case AmInY:
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
