package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpyAbsoluteSetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.Y = 0x01

	// $0000 cpy $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcc, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x01), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpyAbsoluteSetsCarry(t *testing.T) {
	cpu := New()
	cpu.Y = 0x02

	// $0000 cpy $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcc, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x02), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpyAbsoluteSetsNegative(t *testing.T) {
	cpu := New()
	cpu.Y = 0x00

	// $0000 cpy $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcc, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}
