package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecSetsCarry(t *testing.T) {
	cpu := New()

	// $0000 sec
	cpu.Memory.Write(0x0000, 0x38)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
