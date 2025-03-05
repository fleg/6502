package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSedSetsDecimal(t *testing.T) {
	cpu := New()

	// $0000 sed
	cpu.Memory.Write(0x0000, 0xf8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagDecimal, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
