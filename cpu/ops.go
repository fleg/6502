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

	0xa0: {"ldy", amImm, 2, 2, ldy},
	0xa4: {"ldy", amZeP, 2, 3, ldy},
	0xb4: {"ldy", amZeX, 2, 4, ldy},
	0xac: {"ldy", amAbs, 3, 4, ldy},
	0xbc: {"ldy", amAbX, 3, 4, ldy},

	0x85: {"sta", amZeP, 2, 3, sta},
	0x95: {"sta", amZeX, 2, 4, sta},
	0x8d: {"sta", amAbs, 3, 4, sta},
	0x9d: {"sta", amAbX, 3, 5, sta},
	0x99: {"sta", amAbY, 3, 5, sta},
	0x81: {"sta", amInX, 2, 6, sta},
	0x91: {"sta", amInY, 2, 6, sta},

	0x86: {"stx", amZeP, 2, 3, stx},
	0x96: {"stx", amZeY, 2, 4, stx},
	0x8e: {"stx", amAbs, 3, 4, stx},

	0x84: {"sty", amZeP, 2, 3, sty},
	0x94: {"sty", amZeY, 2, 4, sty},
	0x8c: {"sty", amAbs, 3, 4, sty},

	0xaa: {"tax", amImp, 1, 2, tax},
	0xa8: {"tay", amImp, 1, 2, tay},
	0xba: {"tsx", amImp, 1, 2, tsx},
	0x8a: {"txa", amImp, 1, 2, txa},
	0x9a: {"txs", amImp, 1, 2, txs},
	0x98: {"tya", amImp, 1, 2, tya},

	0x48: {"pha", amImp, 1, 3, pha},
	0x68: {"pla", amImp, 1, 3, pla},
	0x08: {"php", amImp, 1, 3, php},
	0x28: {"plp", amImp, 1, 3, plp},

	0x4c: {"jmp", amAbs, 3, 3, jmp},
	0x6c: {"jmp", amInd, 3, 5, jmp},

	0x20: {"jsr", amAbs, 3, 6, jsr},
	0x60: {"rts", amImp, 1, 6, rts},
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
	val := cpu.fetchOp()

	cpu.A = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func ldx(cpu *CPU) {
	val := cpu.fetchOp()

	cpu.X = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func ldy(cpu *CPU) {
	val := cpu.fetchOp()

	cpu.Y = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func sta(cpu *CPU) {
	addr := cpu.fetchOpAddress()

	cpu.Memory.Write(addr, cpu.A)
}

func stx(cpu *CPU) {
	addr := cpu.fetchOpAddress()

	cpu.Memory.Write(addr, cpu.X)
}

func sty(cpu *CPU) {
	addr := cpu.fetchOpAddress()

	cpu.Memory.Write(addr, cpu.Y)
}

func tax(cpu *CPU) {
	cpu.X = cpu.A

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func tay(cpu *CPU) {
	cpu.Y = cpu.A

	cpu.updateZeroFlag(cpu.Y)
	cpu.updateNegativeFlag(cpu.Y)
}

func tsx(cpu *CPU) {
	cpu.X = cpu.SP

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func txa(cpu *CPU) {
	cpu.A = cpu.X

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func txs(cpu *CPU) {
	cpu.SP = cpu.X
}

func tya(cpu *CPU) {
	cpu.A = cpu.X

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func pha(cpu *CPU) {
	cpu.Memory.Write(stackBase|uint16(cpu.SP), cpu.A)
	cpu.SP -= 1
}

func pla(cpu *CPU) {
	cpu.SP += 1
	cpu.A = cpu.Memory.Read(stackBase | uint16(cpu.SP))

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func php(cpu *CPU) {
	cpu.Memory.Write(stackBase|uint16(cpu.SP), uint8(cpu.PS))
	cpu.SP -= 1
}

func plp(cpu *CPU) {
	cpu.SP += 1
	cpu.PS = Flags(cpu.Memory.Read(stackBase | uint16(cpu.SP)))
}

func jmp(cpu *CPU) {
	addr := cpu.fetchOpAddress()

	cpu.PC = addr
}

func jsr(cpu *CPU) {
	ret := cpu.PC + 1

	cpu.PC = cpu.fetchOpAddress()
	cpu.Memory.Write(stackBase|uint16(cpu.SP), uint8((ret&0xff00)>>8))
	cpu.Memory.Write(stackBase|uint16(cpu.SP-1), uint8(ret&0x00ff))
	cpu.SP -= 2
}

func rts(cpu *CPU) {
	cpu.SP += 2
	cpu.PC = 1 + cpu.readWord(stackBase|uint16(cpu.SP-1))
}
