package cpu

type Operation struct {
	Code        uint8
	Name        string
	AddressMode AddressMode
	Size        uint16
	Ticks       uint8
	Do          func(*CPU)
}

var ops = []Operation{
	//opcode name addrmode size ticks do
	{0xea, "nop", amImmediate, 1, 2, nop},

	{0xa9, "lda", amImmediate, 2, 2, lda},
	{0xa5, "lda", amZeroPage, 2, 3, lda},
	{0xb5, "lda", amZeroPageX, 2, 4, lda},
	{0xad, "lda", amAbsolute, 3, 4, lda},
	{0xbd, "lda", amAbsoluteX, 3, 4, lda},
	{0xd9, "lda", amAbsoluteY, 3, 4, lda},
	{0xa1, "lda", amIndirectX, 2, 6, lda},
	{0xb1, "lda", amIndirectY, 2, 5, lda},
}

var opcode2op [256]*Operation

func InitOpcodes() {
	for _, op := range ops {
		opcode2op[op.Code] = &op
	}
}

func nop(_ *CPU) {}

func lda(cpu *CPU) {
	val := cpu.getOperand(cpu.CurrentOp.AddressMode)

	cpu.A = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}
