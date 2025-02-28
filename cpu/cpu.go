package cpu

import "fmt"

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

	Memory         Memory
	TotalTicks     uint64
	CurrentOpTicks uint8
	CurrentOp      *Operation
	IsPageCrossed  bool
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

func (cpu *CPU) Tick() bool {
	cpu.TotalTicks += 1

	if cpu.CurrentOpTicks > 0 {
		cpu.CurrentOpTicks -= 1
		return false
	}

	cpu.nextOp()

	cpu.CurrentOpTicks = cpu.CurrentOp.Ticks - 1
	if cpu.hasExtraTickOnPageCross() {
		cpu.CurrentOpTicks += 1
	}

	return true
}

func (cpu *CPU) nextOp() {
	opcode := cpu.Memory.Read(cpu.PC)
	op := opcode2op[opcode]

	if op == nil {
		panic(fmt.Sprintf("Unknown opcode 0x%02x", opcode))
	}

	cpu.CurrentOp = op

	op.Do(cpu)

	cpu.PC += op.Size
}

func (cpu *CPU) getOperand(am AddressMode) uint8 {
	if am == amAccumulator {
		return cpu.A
	}

	addr := cpu.getOperandAddress(am)

	cpu.IsPageCrossed = addr&0x00ff == 0x00ff

	return cpu.Memory.Read(addr)
}

func (cpu *CPU) getOperandAddress(am AddressMode) uint16 {
	switch am {
	case amImmediate:
		return cpu.PC + 1
	case amZeroPage:
		return uint16(cpu.Memory.Read(cpu.PC + 1))
	case amZeroPageX:
		return uint16(cpu.Memory.Read(cpu.PC+1)+cpu.X) % 0x100
	case amZeroPageY:
		return uint16(cpu.Memory.Read(cpu.PC+1)+cpu.Y) % 0x100
	case amRelative:
		return uint16(int32(cpu.PC) + 2 + int32(int8(cpu.Memory.Read(cpu.PC+1))))
	case amAbsolute:
		return cpu.Memory.ReadWord(cpu.PC + 1)
	case amAbsoluteX:
		return cpu.Memory.ReadWord(cpu.PC + 1 + uint16(cpu.X))
	case amAbsoluteY:
		return cpu.Memory.ReadWord(cpu.PC + 1 + uint16(cpu.Y))
	case amIndirect:
		return cpu.Memory.ReadWord(cpu.Memory.ReadWord(cpu.PC + 1))
	case amIndirectX:
		return cpu.Memory.ReadWord((uint16(cpu.Memory.Read(cpu.PC+1)) + uint16(cpu.X)) % 0x100)
	case amIndirectY:
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
	case amAbsoluteX, amAbsoluteY, amIndirectY:
		return true
	}

	return false
}
