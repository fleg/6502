package cpu

import "fmt"

type Operation struct {
	Name        string
	AddressMode AddressMode
	Size        uint16
	Ticks       uint8
	Do          func(*CPU)
}

var ops = [256]*Operation{
	//opcode: name addrmode size ticks do
	0xea: {"nop", amImm, 1, 2, nop},

	0xa9: {"lda", amImm, 2, 2, lda},
	0xa5: {"lda", amZeP, 2, 3, lda},
	0xb5: {"lda", amZeX, 2, 4, lda},
	0xad: {"lda", amAbs, 3, 4, lda},
	0xbd: {"lda", amAbX, 3, 4, lda},
	0xb9: {"lda", amAbY, 3, 4, lda},
	0xa1: {"lda", amInX, 2, 6, lda},
	0xb1: {"lda", amInY, 2, 5, lda},

	0xa2: {"ldx", amImm, 2, 2, ldx},
	0xa6: {"ldx", amZeP, 2, 3, ldx},
	0xb6: {"ldx", amZeY, 2, 4, ldx},
	0xae: {"ldx", amAbs, 3, 4, ldx},
	0xbe: {"ldx", amAbY, 3, 4, ldx},
}

func opcode2op(opcode uint8) *Operation {
	op := ops[opcode]

	if op == nil {
		panic(fmt.Sprintf("Unknown opcode 0x%02x", opcode))
	}

	return op
}

func nop(_ *CPU) {}

func lda(cpu *CPU) {
	val := cpu.getOperand()

	cpu.A = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func ldx(cpu *CPU) {
	val := cpu.getOperand()

	cpu.X = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}
