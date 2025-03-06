package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRolAccumulatorRotateA(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 rol a
	cpu.Memory.Write(0x0000, 0x2a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x02), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestRolAccumulatorRotateAClearCarry(t *testing.T) {
	cpu := New()
	cpu.A = 0x01
	cpu.PS = flagCarry

	// $0000 rol a
	cpu.Memory.Write(0x0000, 0x2a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestRolAccumulatorRotateASetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x80

	// $0000 rol a
	cpu.Memory.Write(0x0000, 0x2a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestRolAccumulatorRotateASetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x55

	// $0000 rol a
	cpu.Memory.Write(0x0000, 0x2a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0xaa), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestRolAbsoluteRotateMem(t *testing.T) {
	cpu := New()

	// $0000 rol $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x02), cpu.Memory.Read(0xabcd))
}

func TestRolAbsoluteRotateMemClearCarry(t *testing.T) {
	cpu := New()
	cpu.PS = flagCarry

	// $0000 rol $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x03), cpu.Memory.Read(0xabcd))
}

func TestRolAbsoluteRotateASetsCarryAndZero(t *testing.T) {
	cpu := New()

	// $0000 rol $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x00), cpu.Memory.Read(0xabcd))
}

func TestRolAbsoluteRotateMemSetsNegative(t *testing.T) {
	cpu := New()

	// $0000 rol $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x55)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0xaa), cpu.Memory.Read(0xabcd))
}
