package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInyIncrementsY(t *testing.T) {
	cpu := New()
	cpu.Y = 0x08

	// $0000 iny
	cpu.Memory.Write(0x0000, 0xc8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x09), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInyIncrementsYSetsZero(t *testing.T) {
	cpu := New()
	cpu.Y = 0xff

	// $0000 iny
	cpu.Memory.Write(0x0000, 0xc8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInyIncrementsYSetsNegative(t *testing.T) {
	cpu := New()
	cpu.Y = 0x7f

	// $0000 iny
	cpu.Memory.Write(0x0000, 0xc8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x80), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
