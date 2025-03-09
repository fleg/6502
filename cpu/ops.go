package cpu

type Op struct {
	Name        string
	AddressMode AddressMode
	Size        uint16
	Ticks       uint8
	Do          func(*CPU, *Operand)
}

type Operand struct {
	Address     uint16
	AddressMode AddressMode
}

func nop(_ *CPU, _ *Operand) {}

func lda(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.A = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func ldx(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.X = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func ldy(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.Y = val
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func sta(cpu *CPU, operand *Operand) {
	cpu.writeOperand(operand, cpu.A)
}

func stx(cpu *CPU, operand *Operand) {
	cpu.writeOperand(operand, cpu.X)
}

func sty(cpu *CPU, operand *Operand) {
	cpu.writeOperand(operand, cpu.Y)
}

func tax(cpu *CPU, _ *Operand) {
	cpu.X = cpu.A

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func tay(cpu *CPU, _ *Operand) {
	cpu.Y = cpu.A

	cpu.updateZeroFlag(cpu.Y)
	cpu.updateNegativeFlag(cpu.Y)
}

func tsx(cpu *CPU, _ *Operand) {
	cpu.X = cpu.SP

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func txa(cpu *CPU, _ *Operand) {
	cpu.A = cpu.X

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func txs(cpu *CPU, _ *Operand) {
	cpu.SP = cpu.X
}

func tya(cpu *CPU, _ *Operand) {
	cpu.A = cpu.X

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func pha(cpu *CPU, _ *Operand) {
	cpu.push(cpu.A)
}

func pla(cpu *CPU, _ *Operand) {
	cpu.A = cpu.pop()

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func php(cpu *CPU, _ *Operand) {
	cpu.push(uint8(cpu.PS))
}

func plp(cpu *CPU, _ *Operand) {
	cpu.PS = Flags(cpu.pop())
}

func jmp(cpu *CPU, operand *Operand) {
	cpu.PC = operand.Address
}

func jsr(cpu *CPU, operand *Operand) {
	ret := cpu.PC - 1

	cpu.PC = operand.Address
	cpu.pushWord(ret)
}

func rts(cpu *CPU, _ *Operand) {
	cpu.PC = 1 + cpu.popWord()
}

func brk(cpu *CPU, _ *Operand) {
	// note: if an IRQ happens at the same time as a BRK instruction,
	// the BRK instruction is ignored

	cpu.nextPC()
	cpu.pushWord(cpu.PC)
	cpu.push(uint8(cpu.PS | flagBreak | flagUnused))
	cpu.setFlag(flagInterrupt, true)
	cpu.PC = cpu.readWord(irqVector)
}

func rti(cpu *CPU, _ *Operand) {
	cpu.PS = Flags(cpu.pop())
	cpu.PC = cpu.popWord()

	cpu.setFlag(flagUnused, true)
	cpu.setFlag(flagBreak, false)
}

func branch(cpu *CPU, operand *Operand, flag Flags, isSet bool) {
	if isSet == cpu.getFlag(flag) {
		cpu.PC = operand.Address
	}
}

func bcs(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagCarry, true)
}

func bcc(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagCarry, false)
}

func beq(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagZero, true)
}

func bne(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagZero, false)
}

func bmi(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagNegative, true)
}

func bpl(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagNegative, false)
}

func bvs(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagOverflow, true)
}

func bvc(cpu *CPU, operand *Operand) {
	branch(cpu, operand, flagOverflow, false)
}

func clc(cpu *CPU, _ *Operand) {
	cpu.setFlag(flagCarry, false)
}

func cld(cpu *CPU, _ *Operand) {
	cpu.setFlag(flagDecimal, false)
}

func cli(cpu *CPU, _ *Operand) {
	cpu.setFlag(flagInterrupt, false)
}

func clv(cpu *CPU, _ *Operand) {
	cpu.setFlag(flagOverflow, false)
}

func sec(cpu *CPU, _ *Operand) {
	cpu.setFlag(flagCarry, true)
}

func sed(cpu *CPU, _ *Operand) {
	cpu.setFlag(flagDecimal, true)
}

func sei(cpu *CPU, _ *Operand) {
	cpu.setFlag(flagInterrupt, true)
}

func dec(cpu *CPU, operand *Operand) {
	val := cpu.Memory.Read(operand.Address) - 1
	cpu.writeOperand(operand, val)

	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func dex(cpu *CPU, _ *Operand) {
	cpu.X -= 1

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func dey(cpu *CPU, _ *Operand) {
	cpu.Y -= 1

	cpu.updateZeroFlag(cpu.Y)
	cpu.updateNegativeFlag(cpu.Y)
}

func inc(cpu *CPU, operand *Operand) {
	val := cpu.Memory.Read(operand.Address) + 1
	cpu.writeOperand(operand, val)

	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func inx(cpu *CPU, _ *Operand) {
	cpu.X += 1

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func iny(cpu *CPU, _ *Operand) {
	cpu.Y += 1

	cpu.updateZeroFlag(cpu.Y)
	cpu.updateNegativeFlag(cpu.Y)
}

func asl(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)
	carry := val&0x80 > 0
	val = val << 1

	cpu.writeOperand(operand, val)
	cpu.setFlag(flagCarry, carry)
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func lsr(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)
	carry := val&0x01 > 0
	val = val >> 1

	cpu.writeOperand(operand, val)
	cpu.setFlag(flagCarry, carry)
	cpu.updateZeroFlag(val)
	cpu.setFlag(flagNegative, false)
}

func rol(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)
	carry := val&0x80 > 0
	val = val<<1 | uint8(cpu.PS&flagCarry)

	cpu.writeOperand(operand, val)
	cpu.setFlag(flagCarry, carry)
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func ror(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)
	carry := val&0x01 > 0
	val = uint8(cpu.PS&flagCarry)<<7 | val>>1

	cpu.writeOperand(operand, val)
	cpu.setFlag(flagCarry, carry)
	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func and(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.A = cpu.A & val

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func ora(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.A = cpu.A | val

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func eor(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.A = cpu.A ^ val

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func bit(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	val = val & cpu.A

	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func adc(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	if cpu.getFlag(flagDecimal) {
		panic("implement decimal adc")
	} else {
		sum := uint16(cpu.A) + uint16(val) + uint16(cpu.getCarry())
		carry := sum > 0xff
		sum8 := uint8(sum & 0xff)

		cpu.updateOverflowFlag(cpu.A, val, sum8)
		cpu.A = sum8
		cpu.setFlag(flagCarry, carry)
		cpu.updateZeroFlag(cpu.A)
		cpu.updateNegativeFlag(cpu.A)
	}
}

// note sbc === adc with ^mem
func sbc(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	if cpu.getFlag(flagDecimal) {
		panic("implement decimal sbc")
	} else {
		sub := uint16(cpu.A) - uint16(val) - uint16(1-cpu.getCarry())
		carry := sub < 0x100
		sub8 := uint8(sub & 0xff)

		cpu.updateOverflowFlag(cpu.A, ^val, sub8)
		cpu.A = sub8
		cpu.setFlag(flagCarry, carry)
		cpu.updateZeroFlag(cpu.A)
		cpu.updateNegativeFlag(cpu.A)
	}
}
