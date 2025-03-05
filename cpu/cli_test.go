package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliClearsInterrupt(t *testing.T) {
	cpu := New()
	cpu.PS = 0xff

	// $0000 clc
	cpu.Memory.Write(0x0000, 0x58)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xff) & ^flagInterrupt, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
