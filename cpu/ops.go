package cpu

type Op struct {
	Name        string
	AddressMode AddressMode
	Size        uint16
	Ticks       uint8
	Do          func(*CPU)
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
	cpu.push(cpu.A)
}

func pla(cpu *CPU) {
	cpu.A = cpu.pop()

	cpu.updateZeroFlag(cpu.A)
	cpu.updateNegativeFlag(cpu.A)
}

func php(cpu *CPU) {
	cpu.push(uint8(cpu.PS))
}

func plp(cpu *CPU) {
	cpu.PS = Flags(cpu.pop())
}

func jmp(cpu *CPU) {
	addr := cpu.fetchOpAddress()

	cpu.PC = addr
}

func jsr(cpu *CPU) {
	ret := cpu.PC + 1

	cpu.PC = cpu.fetchOpAddress()
	cpu.pushWord(ret)
}

func rts(cpu *CPU) {
	cpu.PC = 1 + cpu.popWord()
}

func brk(cpu *CPU) {
	// note: if an IRQ happens at the same time as a BRK instruction,
	// the BRK instruction is ignored

	cpu.nextPC()
	cpu.pushWord(cpu.PC)
	cpu.push(uint8(cpu.PS | flagBreak | flagUnused))
	cpu.setFlag(flagInterrupt, true)
	cpu.PC = cpu.readWord(irqVector)
}

func rti(cpu *CPU) {
	cpu.PS = Flags(cpu.pop())
	cpu.PC = cpu.popWord()

	cpu.setFlag(flagUnused, true)
	cpu.setFlag(flagBreak, false)
}

func branch(cpu *CPU, flag Flags, isSet bool) {
	addr := cpu.fetchOpAddress()

	if isSet == cpu.getFlag(flag) {
		cpu.PC = addr
	}
}

func bcs(cpu *CPU) {
	branch(cpu, flagCarry, true)
}

func bcc(cpu *CPU) {
	branch(cpu, flagCarry, false)
}

func beq(cpu *CPU) {
	branch(cpu, flagZero, true)
}

func bne(cpu *CPU) {
	branch(cpu, flagZero, false)
}

func bmi(cpu *CPU) {
	branch(cpu, flagNegative, true)
}

func bpl(cpu *CPU) {
	branch(cpu, flagNegative, false)
}

func bvs(cpu *CPU) {
	branch(cpu, flagOverflow, true)
}

func bvc(cpu *CPU) {
	branch(cpu, flagOverflow, false)
}

func clc(cpu *CPU) {
	cpu.setFlag(flagCarry, false)
}

func cld(cpu *CPU) {
	cpu.setFlag(flagDecimal, false)
}

func cli(cpu *CPU) {
	cpu.setFlag(flagInterrupt, false)
}

func clv(cpu *CPU) {
	cpu.setFlag(flagOverflow, false)
}

func sec(cpu *CPU) {
	cpu.setFlag(flagCarry, true)
}

func sed(cpu *CPU) {
	cpu.setFlag(flagDecimal, true)
}

func sei(cpu *CPU) {
	cpu.setFlag(flagInterrupt, true)
}

func dec(cpu *CPU) {
	addr := cpu.fetchOpAddress()
	val := cpu.Memory.Read(addr) - 1
	cpu.Memory.Write(addr, val)

	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func dex(cpu *CPU) {
	cpu.X -= 1

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func dey(cpu *CPU) {
	cpu.Y -= 1

	cpu.updateZeroFlag(cpu.Y)
	cpu.updateNegativeFlag(cpu.Y)
}

func inc(cpu *CPU) {
	addr := cpu.fetchOpAddress()
	val := cpu.Memory.Read(addr) + 1
	cpu.Memory.Write(addr, val)

	cpu.updateZeroFlag(val)
	cpu.updateNegativeFlag(val)
}

func inx(cpu *CPU) {
	cpu.X += 1

	cpu.updateZeroFlag(cpu.X)
	cpu.updateNegativeFlag(cpu.X)
}

func iny(cpu *CPU) {
	cpu.Y += 1

	cpu.updateZeroFlag(cpu.Y)
	cpu.updateNegativeFlag(cpu.Y)
}
