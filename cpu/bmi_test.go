package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBmiNegativeSetJumpsForward(t *testing.T) {
	cpu := New()
	cpu.PS = flagNegative

	// $0000 BMI +6
	cpu.Memory.writeSlice(0x0000, []byte{0x30, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBmiNegativeSetJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008
	cpu.PS = flagNegative

	// $0008 BMI -6
	cpu.Memory.writeSlice(0x0008, []byte{0x30, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBmiNegativeClearDoesNotJump(t *testing.T) {
	cpu := New()

	// $0000 BMI +6
	cpu.Memory.writeSlice(0x0000, []byte{0x30, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
