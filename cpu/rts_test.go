package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRtsPullsAndUpdatePC(t *testing.T) {
	cpu := New()
	cpu.SP = 0xf0

	// $0000 RTS
	cpu.Memory.Write(0x0000, 0x60)
	cpu.Memory.writeSlice(0x01f1, []byte{0x10, 0xc0})
	cpu.Step()

	assert.Equal(t, uint16(0xc011), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xf2), cpu.SP)
}

func TestRtsPullsAndUpdatePCWrapAround(t *testing.T) {
	cpu := New()
	cpu.PC = 0xf000
	cpu.SP = 0xf0

	// $f000 RTS
	cpu.Memory.Write(0xf000, 0x60)
	cpu.Memory.writeSlice(0x01f1, []byte{0xff, 0xff})
	cpu.Step()

	assert.Equal(t, uint16(0x0000), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xf2), cpu.SP)
}
