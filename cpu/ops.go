package cpu

type Op struct {
	Name          string
	AddressMode   AddressMode
	Size          uint16
	Ticks         uint8
	Do            func(*CPU, *Operand)
	PageCrossTick uint8
}

type Operand struct {
	Address     uint16
	AddressMode AddressMode
	PageCrossed bool
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
	cpu.A = cpu.Y

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
	cpu.push(uint8(cpu.PS | flagBreak | flagUnused))
}

func plp(cpu *CPU, _ *Operand) {
	cpu.PS = Flags(cpu.pop()) | flagUnused
	cpu.setFlag(flagBreak, false)
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

		cpu.totalTicks += 1
		if operand.PageCrossed {
			cpu.totalTicks += 1
		}
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
	val := cpu.read(operand.Address) - 1
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
	val := cpu.read(operand.Address) + 1
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

	cpu.updateZeroFlag(val & cpu.A)
	cpu.updateNegativeFlag(val)
	cpu.setFlag(flagOverflow, val&uint8(flagOverflow) > 0)
}

func adc(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	if cpu.getFlag(flagDecimal) {
		loCarry := uint8(0)
		hiCarry := false
		loAdj := uint8(0)
		hiAdj := uint8(0)

		loSum := val&0x0f + cpu.A&0x0f + cpu.getCarry()
		if loSum > 0x09 {
			loAdj = 0x06
			loCarry = 1
		}

		hiSum := (val>>4)&0x0f + (cpu.A>>4)&0x0f + loCarry
		if hiSum > 0x09 {
			hiAdj = 0x06
			hiCarry = true
		}

		sum := ((hiSum & 0x0f) << 4) + loSum&0x0f
		sumAdj := (((hiSum + hiAdj) & 0x0f) << 4) | (loSum+loAdj)&0x0f

		cpu.updateOverflowFlag(cpu.A, val, sum)
		cpu.A = sumAdj
		cpu.setFlag(flagCarry, hiCarry)
		cpu.updateZeroFlag(sum)
		cpu.updateNegativeFlag(sum)
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
		loCarry := uint8(1)
		hiCarry := false
		loAdj := uint8(0)
		hiAdj := uint8(0)

		loSum := (^val)&0x0f + cpu.A&0x0f + cpu.getCarry()
		if loSum <= 0x0f {
			loAdj = 0x0a
			loCarry = 0
		}

		hiSum := ((^val)>>4)&0x0f + (cpu.A>>4)&0x0f + loCarry
		if hiSum <= 0x0f {
			hiAdj = 0xa0
		}

		sum := uint16(cpu.A) + uint16(^val)&0xff + uint16(cpu.getCarry())
		sum8 := uint8(sum & 0xff)
		sumAdj := ((sum8 + hiAdj) & 0xf0) | (sum8+loAdj)&0x0f

		if sum > 0xff {
			hiCarry = true
		}

		cpu.updateOverflowFlag(cpu.A, ^val, sum8)
		cpu.A = sumAdj
		cpu.setFlag(flagCarry, hiCarry)
		cpu.updateZeroFlag(sum8)
		cpu.updateNegativeFlag(sum8)
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

func compare(cpu *CPU, operand *Operand, src uint8) {
	val := cpu.readOperand(operand)

	cpu.setFlag(flagCarry, src >= val)
	cpu.setFlag(flagZero, src == val)
	cpu.updateNegativeFlag(src - val)
}

func cmp(cpu *CPU, operand *Operand) {
	compare(cpu, operand, cpu.A)
}

func cpx(cpu *CPU, operand *Operand) {
	compare(cpu, operand, cpu.X)
}

func cpy(cpu *CPU, operand *Operand) {
	compare(cpu, operand, cpu.Y)
}

func slo(cpu *CPU, operand *Operand) {
	asl(cpu, operand)
	ora(cpu, operand)
}

func rla(cpu *CPU, operand *Operand) {
	rol(cpu, operand)
	and(cpu, operand)
}

func sre(cpu *CPU, operand *Operand) {
	lsr(cpu, operand)
	eor(cpu, operand)
}

func rra(cpu *CPU, operand *Operand) {
	ror(cpu, operand)
	adc(cpu, operand)
}

func dcp(cpu *CPU, operand *Operand) {
	dec(cpu, operand)
	cmp(cpu, operand)
}

func isc(cpu *CPU, operand *Operand) {
	inc(cpu, operand)
	sbc(cpu, operand)
}

func lax(cpu *CPU, operand *Operand) {
	lda(cpu, operand)
	ldx(cpu, operand)
}

func sax(cpu *CPU, operand *Operand) {
	cpu.writeOperand(operand, cpu.A&cpu.X)
}

func anc(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.A = cpu.A & val
	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
	cpu.setFlag(flagCarry, cpu.getFlag(flagNegative))
}

func asr(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	and := cpu.A & val

	cpu.A = and >> 1
	cpu.setFlag(flagNegative, false)
	cpu.updateZeroFlag(cpu.A)
	cpu.setFlag(flagCarry, and&0x01 > 0)
}

func arr(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	if cpu.getFlag(flagDecimal) {
		// undocumented magic
		// https://github.com/vbt1/fba_saturn/blob/d1dade09813b397c6c3abab418842c109c7e2bd2/m6502.new/ill02.h#L71
		res := uint16(val) & uint16(cpu.A)
		lo := res & 0x0f
		hi := res & 0xf0
		and := res

		res = uint16(cpu.PS&flagCarry)<<7 | res>>1
		cpu.updateNegativeFlag(uint8(res))
		cpu.updateZeroFlag(uint8(res))
		cpu.setFlag(flagOverflow, (and^res)&uint16(flagOverflow) > 0)

		if lo+(lo&0x01) > 0x05 {
			res = (res & 0xf0) | ((res + 0x06) & 0x0f)
		}

		if hi+(hi&0x10) > 0x50 {
			cpu.setFlag(flagCarry, true)
			res = res + 0x60
		} else {
			cpu.setFlag(flagCarry, false)
		}

		cpu.A = uint8(res & 0xff)
	} else {
		cpu.A = uint8(cpu.PS&flagCarry)<<7 | (val&cpu.A)>>1
		cpu.updateZeroFlag(cpu.A)
		cpu.updateNegativeFlag(cpu.A)
		cpu.setFlag(flagOverflow, (cpu.A&uint8(flagOverflow) > 0) != (cpu.A&uint8(flagUnused) > 0))
		cpu.setFlag(flagCarry, cpu.A&uint8(flagOverflow) > 0)
	}
}

func xaa(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.A = (cpu.A | magic) & cpu.X & val
	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func lxa(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.A = (cpu.A | magic) & val
	cpu.X = cpu.A
	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func las(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)

	cpu.SP = cpu.SP & val
	cpu.A = cpu.SP
	cpu.X = cpu.SP
	cpu.updateZeroFlag(cpu.SP)
	cpu.updateNegativeFlag(cpu.SP)
}

func sbx(cpu *CPU, operand *Operand) {
	val := cpu.readOperand(operand)
	and := cpu.A & cpu.X

	cpu.X = and - val
	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
	cpu.setFlag(flagCarry, and >= val)
}

func sha(cpu *CPU, operand *Operand) {
	cpu.writeOperand(operand, cpu.A&cpu.X&(1+wordMSB(operand.Address)))
}

func shx(cpu *CPU, operand *Operand) {
	cpu.writeOperand(operand, cpu.X&(1+wordMSB(operand.Address)))
}

func shy(cpu *CPU, operand *Operand) {
	cpu.writeOperand(operand, cpu.Y&(1+wordMSB(operand.Address)))
}

func shs(cpu *CPU, operand *Operand) {
	cpu.SP = cpu.A & cpu.X
	cpu.writeOperand(operand, cpu.SP&(1+wordMSB(operand.Address)))
}

func jam(cpu *CPU, _ *Operand) {
	// TODO halt the cpu?
	cpu.totalTicks += 10
}
