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

	Memory        Memory
	TotalTicks    uint64
	CurrentOp     *Operation
	IsPageCrossed bool
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
	cpu.nextOp()

	cpu.TotalTicks += uint64(cpu.CurrentOp.Ticks)

	if cpu.hasExtraTickOnPageCross() {
		cpu.TotalTicks += 1
	}
}

func (cpu *CPU) nextOp() {
	opcode := cpu.Memory.Read(cpu.PC)
	op := opcode2op(opcode)

	cpu.CurrentOp = op

	op.Do(cpu)

	cpu.PC += op.Size
}

func (cpu *CPU) getOperand(am AddressMode) uint8 {
	if am == amAcc {
		return cpu.A
	}

	addr := cpu.getOperandAddress(am)

	cpu.IsPageCrossed = addr&0x00ff == 0x00ff

	return cpu.Memory.Read(addr)
}

func (cpu *CPU) getOperandAddress(am AddressMode) uint16 {
	switch am {
	case amImm:
		return cpu.PC + 1
	case amZeP:
		return uint16(cpu.Memory.Read(cpu.PC + 1))
	case amZeX:
		return uint16(cpu.Memory.Read(cpu.PC+1)+cpu.X) % 0x100
	case amZeY:
		return uint16(cpu.Memory.Read(cpu.PC+1)+cpu.Y) % 0x100
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
		return cpu.Memory.ReadWord((uint16(cpu.Memory.Read(cpu.PC+1)) + uint16(cpu.X)) % 0x100)
	case amInY:
		return cpu.Memory.ReadWord((uint16(cpu.Memory.Read(cpu.PC+1)))%0x100) + uint16(cpu.Y)
	default:
		panic("Unknown address mode")
	}
}

func (cpu *CPU) hasExtraTickOnPageCross() bool {
	if !cpu.IsPageCrossed {
		return false
	}

	switch cpu.CurrentOp.AddressMode {
	case amAbX, amAbY, amInY:
		return true
	}

	return false
}
