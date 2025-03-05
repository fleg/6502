package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClvClearsOverflow(t *testing.T) {
	cpu := New()
	cpu.PS = 0xff

	// $0000 clv
	cpu.Memory.Write(0x0000, 0xb8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xff) & ^flagOverflow, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
