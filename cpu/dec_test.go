package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecAbsoluteDecrementsMem(t *testing.T) {
	cpu := New()

	// $0000 dec $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xce, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x08)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x07), cpu.Memory.Read(0xabcd))
}

func TestDecAbsoluteDecrementsMemSetsZero(t *testing.T) {
	cpu := New()

	// $0000 dec $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xce, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x00), cpu.Memory.Read(0xabcd))
}

func TestDecAbsoluteDecrementsMemSetsNegative(t *testing.T) {
	cpu := New()

	// $0000 dec $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xce, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0xff), cpu.Memory.Read(0xabcd))
}
