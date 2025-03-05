package cpu

const (
	resetVector = 0xfffc
	irqVector   = 0xfffe
	stackInit   = 0xfd
	stackBase   = 0x0100
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
	CurrentOp  *Op
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

func (cpu *CPU) Step() {
	opcode := cpu.readPC()
	op := opcode2op(opcode)

	cpu.CurrentOp = op
	cpu.TotalTicks += uint64(cpu.CurrentOp.Ticks)
	// TODO handle extra page cross ticks

	op.Do(cpu)
}

func (cpu *CPU) fetchOp() uint8 {
	if cpu.CurrentOp.AddressMode == amAcc {
		return cpu.A
	}

	addr := cpu.fetchOpAddress()

	return cpu.Memory.Read(addr)
}

func (cpu *CPU) fetchOpAddress() uint16 {
	switch cpu.CurrentOp.AddressMode {
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
		return cpu.readWord(cpu.readPCWord())
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

func word(lo uint8, hi uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}
