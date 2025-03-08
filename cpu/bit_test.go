package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitAbsoluteClearZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x01
	cpu.PS = flagZero

	// $0000 bit $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2c, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x01), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestBitAbsoluteSetsZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x55

	// $0000 bit $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2c, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0xaa)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x55), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0xaa), cpu.Memory.Read(0xabcd))
}

func TestBitAbsoluteClearNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x81
	cpu.PS = flagNegative

	// $0000 bit $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2c, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x81), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestBitAbsoluteSetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x80

	// $0000 bit $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2c, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x81)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x81), cpu.Memory.Read(0xabcd))
}
