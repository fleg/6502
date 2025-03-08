package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEorAbsoluteXorWithMem(t *testing.T) {
	cpu := New()
	cpu.A = 0x05

	// $0000 eor $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x4d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x02)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x07), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x02), cpu.Memory.Read(0xabcd))
}

func TestEorAbsoluteXorWithMemSetsZero(t *testing.T) {
	cpu := New()
	cpu.A = 0xff

	// $0000 eor $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x4d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0xff)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0xff), cpu.Memory.Read(0xabcd))
}

func TestEorAbsoluteXorWithMemSetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 eor $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x4d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x81), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x80), cpu.Memory.Read(0xabcd))
}
