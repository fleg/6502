package cpu

const (
	resetVector = 0xfffc
	stackInit   = 0xfd
)

// The processor is little endian and expects addresses to be stored in memory least significant byte first.
type CPU struct {
	PC uint16
	SP uint8
	A  uint8
	X  uint8
	Y  uint8
	PS Flags

	Memory     Memory
	TotalTicks uint64
	CurrentOp  *Operation
}

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) Reset() {
	cpu.PC = cpu.Memory.ReadWord(resetVector)
	cpu.SP = stackInit
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.PS = 0

	cpu.setFlag(flagInterrupt, true)
}

func (cpu *CPU) Step() {
	opcode := cpu.Memory.Read(cpu.PC)
	op := opcode2op(opcode)

	cpu.CurrentOp = op
	cpu.TotalTicks += uint64(cpu.CurrentOp.Ticks)
	// TODO handle extra page cross ticks

	op.Do(cpu)

	cpu.PC += op.Size
}

func (cpu *CPU) getOperand(am AddressMode) uint8 {
	if am == amAcc {
		return cpu.A
	}

	addr := cpu.getOperandAddress(am)

	return cpu.Memory.Read(addr)
}

func (cpu *CPU) getOperandAddress(am AddressMode) uint16 {
	switch am {
	case amImm:
		return cpu.PC + 1
	case amZeP:
		return uint16(cpu.Memory.Read(cpu.PC + 1))
	case amZeX:
		addr := cpu.Memory.Read(cpu.PC+1) + cpu.X
		return uint16(addr)
	case amZeY:
		addr := cpu.Memory.Read(cpu.PC+1) + cpu.Y
		return uint16(addr)
	case amRel:
		return uint16(int32(cpu.PC) + 2 + int32(int8(cpu.Memory.Read(cpu.PC+1))))
	case amAbs:
		return cpu.Memory.ReadWord(cpu.PC + 1)
	case amAbX:
		return cpu.Memory.ReadWord(cpu.PC+1) + uint16(cpu.X)
	case amAbY:
		return cpu.Memory.ReadWord(cpu.PC+1) + uint16(cpu.Y)
	case amInd:
		return cpu.Memory.ReadWord(cpu.Memory.ReadWord(cpu.PC + 1))
	case amInX:
		loAddr := cpu.Memory.Read(cpu.PC+1) + cpu.X
		lo := cpu.Memory.Read(uint16(loAddr))
		hi := cpu.Memory.Read(uint16(loAddr + 1))
		return uint16(hi)<<8 | uint16(lo)
	case amInY:
		zp := uint16(cpu.Memory.Read(cpu.PC + 1))
		lo := uint16(cpu.Memory.Read(zp)) + uint16(cpu.Y)
		carry := lo >> 8
		hi := uint16(cpu.Memory.Read(zp+1)) + carry
		return uint16(hi)<<8 | uint16(lo&0x00ff)
	default:
		panic("Unknown address mode")
	}
}
