package cpu

const memCapacity = 64 * 1024

type Memory struct {
	values [memCapacity]byte
}

func (m *Memory) Read(addr uint16) uint8 {
	return m.values[addr]
}

func (m *Memory) ReadWord(addr uint16) uint16 {
	return uint16(m.Read(addr+1))<<8 | uint16(m.Read(addr))
}

func (m *Memory) Write(addr uint16, value uint8) {
	m.values[addr] = value
}

func (m *Memory) WriteSlice(addr uint16, values []byte) {
	for i, value := range values {
		m.values[addr+uint16(i)] = value
	}
}

func (m *Memory) Clear() {
	for i := range memCapacity {
		m.values[i] = 0
	}
}
