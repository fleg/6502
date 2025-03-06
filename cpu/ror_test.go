package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRorAccumulatorRotateA(t *testing.T) {
	cpu := New()
	cpu.A = 0x02

	// $0000 ror a
	cpu.Memory.Write(0x0000, 0x6a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x01), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestRorAccumulatorRotateAClearCarrySetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x02
	cpu.PS = flagCarry

	// $0000 ror a
	cpu.Memory.Write(0x0000, 0x6a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x81), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestRorAccumulatorRotateASetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 ror a
	cpu.Memory.Write(0x0000, 0x6a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestRorAbsoluteRotateMem(t *testing.T) {
	cpu := New()

	// $0000 ror $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x02)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestRorAbsoluteRotateMemClearCarrySetsNegative(t *testing.T) {
	cpu := New()
	cpu.PS = flagCarry

	// $0000 ror $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x02)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x81), cpu.Memory.Read(0xabcd))
}

func TestRorAbsoluteRotateASetsCarryAndZero(t *testing.T) {
	cpu := New()

	// $0000 ror $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x6e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x00), cpu.Memory.Read(0xabcd))
}
