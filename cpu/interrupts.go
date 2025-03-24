package cpu

func nmi(cpu *CPU) {
	cpu.pushWord(cpu.PC)
	cpu.push(uint8(cpu.PS | flagUnused))
	cpu.setFlag(flagInterrupt, true)
	cpu.PC = cpu.readWord(nmiVector)

	cpu.totalTicks += 7
}

func irq(cpu *CPU) {
	cpu.pushWord(cpu.PC)
	cpu.push(uint8(cpu.PS | flagUnused))
	cpu.setFlag(flagInterrupt, true)
	cpu.PC = cpu.readWord(irqVector)

	cpu.totalTicks += 7
}
