package cpu

func interrupt(cpu *CPU, flags Flags, vec uint16) {
	cpu.pushWord(cpu.PC)
	cpu.push(uint8(flags))
	cpu.setFlag(flagInterrupt, true)
	cpu.PC = cpu.readWord(vec)
}

func nmi(cpu *CPU) {
	interrupt(cpu, cpu.PS|flagUnused, nmiVector)

	cpu.totalTicks += 7
}

func irq(cpu *CPU) {
	interrupt(cpu, cpu.PS|flagUnused, irqVector)

	cpu.totalTicks += 7
}
