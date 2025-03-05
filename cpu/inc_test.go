package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncAbsoluteIncrementsMem(t *testing.T) {
	cpu := New()

	// $0000 inc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xee, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x08)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x09), cpu.Memory.Read(0xabcd))
}

func TestIncAbsoluteIncrementsMemSetsZero(t *testing.T) {
	cpu := New()

	// $0000 inc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xee, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0xff)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x00), cpu.Memory.Read(0xabcd))
}

func TestIncAbsoluteIncrementsMemSetsNegative(t *testing.T) {
	cpu := New()

	// $0000 inc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xee, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x7f)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x80), cpu.Memory.Read(0xabcd))
}
