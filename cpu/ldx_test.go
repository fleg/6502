package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLdxAbsoluteLoadsXSetsNFlag(t *testing.T) {
	cpu := New()
	cpu.X = 0x00

	// $0000 LDX $abcd
	cpu.Memory.WriteSlice(0x0000, []byte{0xae, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.X)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdxAbsoluteLoadsXSetsZFlag(t *testing.T) {
	cpu := New()
	cpu.X = 0xff

	// $0000 LDX $abcd
	cpu.Memory.WriteSlice(0x0000, []byte{0xae, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
