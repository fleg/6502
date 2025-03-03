package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStxAbsoluteStoresX(t *testing.T) {
	cpu := New()
	cpu.X = 0x55

	// $0000 STX $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x8e, 0xcd, 0xab})
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x55), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x55), cpu.Memory.Read(0xabcd))
}
