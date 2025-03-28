package cpu

const memSize = 64 * 1024

type RAM struct {
	values [memSize]uint8
}

func (m *RAM) Read(addr uint16) uint8 {
	return m.values[addr]
}

func (m *RAM) Write(addr uint16, value uint8) {
	m.values[addr] = value
}

func NewRAM() *RAM {
	return &RAM{}
}
