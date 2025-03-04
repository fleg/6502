package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsrPushesAndUpdatePC(t *testing.T) {
	cpu := New()
	cpu.PC = 0xc010
	cpu.SP = 0xfd

	// $c010 JSR $ffd2
	cpu.Memory.writeSlice(0xc010, []byte{0x20, 0xd2, 0xff})
	cpu.Step()

	assert.Equal(t, uint16(0xffd2), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfb), cpu.SP)
	assert.Equal(t, uint8(0xc0), cpu.Memory.Read(0x01fd))
	assert.Equal(t, uint8(0x12), cpu.Memory.Read(0x01fc))
}
