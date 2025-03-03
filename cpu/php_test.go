package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhpPushesPAndUpdatesSP(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.PS = Flags(0x55)

	// $0000 PHP
	cpu.Memory.Write(0x0000, 0x08)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0x55), cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfc), cpu.SP)
	assert.Equal(t, uint8(0x55), cpu.Memory.Read(0x01fd))
}
