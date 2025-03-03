package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaPullsIntoAAndUpdatesSP(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x01fe, 0x3a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x3a), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}

func TestPlaPullsIntoAAndUpdatesSPOverflow(t *testing.T) {
	cpu := New()
	cpu.SP = 0xff
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x0100, 0x3a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x3a), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestPlaPullsIntoAAndSetsZFlag(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x01fe, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}

func TestPlaPullsIntoAAndSetsNFlag(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x01fe, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}
