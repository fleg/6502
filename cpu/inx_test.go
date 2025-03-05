package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInxIncrementsX(t *testing.T) {
	cpu := New()
	cpu.X = 0x08

	// $0000 inx
	cpu.Memory.Write(0x0000, 0xe8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x09), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInxIncrementsXSetsZero(t *testing.T) {
	cpu := New()
	cpu.X = 0xff

	// $0000 inx
	cpu.Memory.Write(0x0000, 0xe8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInxIncrementsXSetsNegative(t *testing.T) {
	cpu := New()
	cpu.X = 0x7f

	// $0000 inx
	cpu.Memory.Write(0x0000, 0xe8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x80), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
