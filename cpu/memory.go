package cpu

type Memory interface {
	Read(addr uint16) uint8
	Write(addr uint16, value uint8)
}

func (cpu *CPU) read(addr uint16) uint8 {
	return cpu.memory.Read(addr)
}

func (cpu *CPU) write(addr uint16, value uint8) {
	cpu.memory.Write(addr, value)
}

func (cpu *CPU) writeSlice(addr uint16, values []byte) {
	for i, value := range values {
		cpu.memory.Write(addr+uint16(i), value)
	}
}
