package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOraAbsoluteOrWithMem(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 ora $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x0d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x02)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x02), cpu.Memory.Read(0xabcd))
}

func TestOraAbsoluteOrWithMemSetsZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x00

	// $0000 ora $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x0d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x00), cpu.Memory.Read(0xabcd))
}

func TestOraAbsoluteOrWithMemSetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x55

	// $0000 ora $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x0d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0xaa)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0xff), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0xaa), cpu.Memory.Read(0xabcd))
}
