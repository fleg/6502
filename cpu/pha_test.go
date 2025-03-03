package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhaPushesAAndUpdatesSP(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.A = 0xab

	// $0000 PHA
	cpu.Memory.Write(0x0000, 0x48)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0xab), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfc), cpu.SP)
	assert.Equal(t, uint8(0xab), cpu.Memory.Read(0x01fd))
}

func TestPhaPushesAAndUpdatesSPUnderflow(t *testing.T) {
	cpu := New()
	cpu.SP = 0x00
	cpu.A = 0xab

	// $0000 PHA
	cpu.Memory.Write(0x0000, 0x48)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0xab), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xff), cpu.SP)
	assert.Equal(t, uint8(0xab), cpu.Memory.Read(0x0100))
}
