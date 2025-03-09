package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdcAbsoluteClearZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.PS = flagZero

	// $0000 adc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6d, 0xcd, 0xab})
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

func TestAdcAbsoluteAddMemWithCarry(t *testing.T) {
	cpu := New()
	cpu.A = 0x01
	cpu.PS = flagCarry

	// $0000 adc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestAdcAbsoluteAddMemSetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 adc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6d, 0xcd, 0xab})
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

func TestAdcAbsoluteAddMemSetsOverflowBothPositive(t *testing.T) {
	cpu := New()
	cpu.A = 0x7f

	// $0000 adc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x7f)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagOverflow|flagNegative, cpu.PS)
	assert.Equal(t, uint8(0xfe), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x7f), cpu.Memory.Read(0xabcd))
}

func TestAdcAbsoluteAddMemSetsOverflowBothNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x80

	// $0000 adc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagOverflow|flagZero|flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x80), cpu.Memory.Read(0xabcd))
}
