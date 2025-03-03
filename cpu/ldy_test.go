package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLdyAbsoluteLoadsYSetsNFlag(t *testing.T) {
	cpu := New()
	cpu.Y = 0x00

	// $0000 LDY $abcd
	cpu.Memory.WriteSlice(0x0000, []byte{0xac, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.Y)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdyAbsoluteLoadsYSetsZFlag(t *testing.T) {
	cpu := New()
	cpu.Y = 0xff

	// $0000 LDY $abcd
	cpu.Memory.WriteSlice(0x0000, []byte{0xac, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
