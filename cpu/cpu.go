package cpu

const (
	resetVector = 0xfffc
	irqVector   = 0xfffe
	stackInit   = 0xfd
	stackBase   = 0x0100
	magic       = 0xee
)

// The processor is little endian and expects
// addresses to be stored in memory least significant byte first.
type CPU struct {
	PC uint16
	SP uint8
	A  uint8
	X  uint8
	Y  uint8
	PS Flags

	Memory     Memory
	TotalTicks uint64
}

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) Reset() {
	cpu.PC = cpu.readWord(resetVector)
	cpu.SP = stackInit
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.PS = 0

	cpu.setFlag(flagInterrupt, true)
}

func (cpu *CPU) nextPC() uint16 {
	addr := cpu.PC
	cpu.PC += 1

	return addr
}

func (cpu *CPU) readPC() uint8 {
	return cpu.Memory.Read(cpu.nextPC())
}

func (cpu *CPU) readPCWord() uint16 {
	lo := cpu.readPC()
	hi := cpu.readPC()

	return word(lo, hi)
}

func (cpu *CPU) readWord(addr uint16) uint16 {
	lo := cpu.Memory.Read(addr)
	hi := cpu.Memory.Read(addr + 1)

	return word(lo, hi)
}

func (cpu *CPU) readWordWithoutPageCross(addr uint16) uint16 {
	lo := cpu.Memory.Read(addr)

	if addr&0x00ff == 0x00ff {
		addr = addr & 0xff00
	} else {
		addr += 1
	}

	hi := cpu.Memory.Read(addr)

	return word(lo, hi)
}

func (cpu *CPU) Step() {
	opcode := cpu.readPC()
	op := opcode2op(opcode)

	operand := cpu.fetchOperand(op.AddressMode)
	op.Do(cpu, operand)

	cpu.TotalTicks += uint64(op.Ticks)
	// TODO handle extra page cross ticks
}

func (cpu *CPU) fetchOperand(am AddressMode) *Operand {
	switch am {
	case amImp, amAcc:
		return &Operand{AddressMode: am}
	default:
		return &Operand{
			Address:     cpu.fetchOperandAddress(am),
			AddressMode: am,
		}
	}
}

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
		return cpu.readWord(addr)
	case amInY:
		zp := uint16(cpu.readPC())
		sum := uint16(cpu.Memory.Read(zp)) + uint16(cpu.Y)
		carry := uint8(sum >> 8)
		lo := uint8(sum & 0x00ff)
		hi := cpu.Memory.Read(zp+1) + carry
		return word(lo, hi)
	default:
		panic("Unknown address mode")
	}
}

func (cpu *CPU) readOperand(operand *Operand) uint8 {
	if operand.AddressMode == amAcc {
		return cpu.A
	}

	return cpu.Memory.Read(operand.Address)
}

func (cpu *CPU) writeOperand(operand *Operand, val uint8) {
	if operand.AddressMode == amAcc {
		cpu.A = val
	} else {
		cpu.Memory.Write(operand.Address, val)
	}
}

func word(lo uint8, hi uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}
