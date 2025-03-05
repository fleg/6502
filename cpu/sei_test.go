package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeiSetsInterrupt(t *testing.T) {
	cpu := New()

	// $0000 sei
	cpu.Memory.Write(0x0000, 0x78)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagInterrupt, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
