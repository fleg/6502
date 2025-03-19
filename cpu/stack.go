package cpu

func (cpu *CPU) push(val uint8) {
	cpu.write(stackBase|uint16(cpu.SP), val)
	cpu.SP -= 1
}

func (cpu *CPU) pop() uint8 {
	cpu.SP += 1
	val := cpu.read(stackBase | uint16(cpu.SP))

	return val
}

func (cpu *CPU) pushWord(val uint16) {
	cpu.push(uint8((val & 0xff00) >> 8))
	cpu.push(uint8(val & 0x00ff))
}

func (cpu *CPU) popWord() uint16 {
	return word(cpu.pop(), cpu.pop())
}
