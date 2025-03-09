package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSbcAbsoluteClearZeroSetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.PS = flagZero

	// $0000 sbc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xed, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0xfe), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestSbcAbsoluteSubMemWithCarry(t *testing.T) {
	cpu := New()
	cpu.A = 0x02
	cpu.PS = flagCarry

	// $0000 sbc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xed, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x01), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestSbcAbsoluteSubMemSetsOverflowNegSubPos(t *testing.T) {
	cpu := New()
	cpu.A = 0x81
	cpu.PS = flagNegative

	// $0000 sbc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xed, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x7f)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagOverflow|flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x01), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x7f), cpu.Memory.Read(0xabcd))
}

func TestSbcAbsoluteSubMemSetsOverflowPosSubNeg(t *testing.T) {
	cpu := New()
	cpu.A = 0x7f

	// $0000 sbc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xed, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x81)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagOverflow|flagNegative, cpu.PS)
	assert.Equal(t, uint8(0xfd), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x81), cpu.Memory.Read(0xabcd))
}
