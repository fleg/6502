package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBvsOverflowSetJumpsForward(t *testing.T) {
	cpu := New()
	cpu.PS = flagOverflow

	// $0000 BVS +6
	cpu.Memory.writeSlice(0x0000, []byte{0x70, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagOverflow, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBvsOverflowSetJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008
	cpu.PS = flagOverflow

	// $0008 BVS -6
	cpu.Memory.writeSlice(0x0008, []byte{0x70, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagOverflow, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBvsOverflowCleanDoesNotJump(t *testing.T) {
	cpu := New()

	// $0000 BVS +6
	cpu.Memory.writeSlice(0x0000, []byte{0x70, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
