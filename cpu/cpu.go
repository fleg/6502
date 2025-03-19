package cpu

const (
	resetVector = 0xfffc
	irqVector   = 0xfffe
	stackInit   = 0xfd
	stackBase   = 0x0100
	magic       = 0xee
)

type CPU struct {
	PC uint16
	SP uint8
	A  uint8
	X  uint8
	Y  uint8
	PS Flags

	memory     Memory
	totalTicks uint64
}

func New(mem Memory) *CPU {
	return &CPU{
		memory: mem,
	}
}

func NewWithRAM() *CPU {
	return New(NewRAM())
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
	return cpu.read(cpu.nextPC())
}

func (cpu *CPU) readPCWord() uint16 {
	lo := cpu.readPC()
	hi := cpu.readPC()

	return word(lo, hi)
}

func (cpu *CPU) readWord(addr uint16) uint16 {
	lo := cpu.read(addr)
	hi := cpu.read(addr + 1)

	return word(lo, hi)
}

func (cpu *CPU) readWordWithoutPageCross(addr uint16) uint16 {
	lo := cpu.read(addr)

	if addr&0x00ff == 0x00ff {
		addr = addr & 0xff00
	} else {
		addr += 1
	}

	hi := cpu.read(addr)

	return word(lo, hi)
}

func (cpu *CPU) Step() {
	opcode := cpu.readPC()
	op := opcode2op(opcode)

	operand := cpu.fetchOperand(op.AddressMode)
	op.Do(cpu, operand)

	cpu.totalTicks += uint64(op.Ticks)
	// TODO handle extra page cross ticks
}

func (cpu *CPU) fetchOperand(am AddressMode) *Operand {
	return &Operand{
		Address:     cpu.fetchOperandAddress(am),
		AddressMode: am,
	}
}

func (cpu *CPU) readOperand(operand *Operand) uint8 {
	if operand.AddressMode == amAcc {
		return cpu.A
	}

	return cpu.read(operand.Address)
}

func (cpu *CPU) writeOperand(operand *Operand, val uint8) {
	if operand.AddressMode == amAcc {
		cpu.A = val
	} else {
		cpu.write(operand.Address, val)
	}
}

func word(lo uint8, hi uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}
