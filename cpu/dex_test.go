package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDexDecrementsX(t *testing.T) {
	cpu := New()
	cpu.X = 0x08

	// $0000 dex
	cpu.Memory.Write(0x0000, 0xca)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x07), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestDexDecrementsXSetsZero(t *testing.T) {
	cpu := New()
	cpu.X = 0x01

	// $0000 dex
	cpu.Memory.Write(0x0000, 0xca)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestDexDecrementsXSetsNegative(t *testing.T) {
	cpu := New()
	cpu.X = 0x00

	// $0000 dex
	cpu.Memory.Write(0x0000, 0xca)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0xff), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
