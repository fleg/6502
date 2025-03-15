package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRtiPullsPCAndPSAndUpdatePC(t *testing.T) {
	cpu := New()
	cpu.SP = 0xf0

	// $0000 RTI
	cpu.Memory.Write(0x0000, 0x40)
	cpu.Memory.writeSlice(0x01f1, []byte{0x01, 0x03, 0xc0})
	cpu.Step()

	assert.Equal(t, uint16(0xc003), cpu.PC)
	assert.Equal(t, flagCarry|flagBreak|flagUnused, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xf3), cpu.SP)
}
