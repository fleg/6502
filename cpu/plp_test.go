package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlpPullsIntoPAndUpdatesSP(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.PS = 0xff

	// $0000 PLP
	cpu.Memory.Write(0x0000, 0x28)
	cpu.Memory.Write(0x01fe, 0xaa)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xaa)|flagUnused|flagBreak, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}
