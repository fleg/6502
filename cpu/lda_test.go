package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLdaAbsoluteLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00

	// $0000 LDA $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xad, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaAbsoluteLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xff

	// $0000 LDA $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xad, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaZPLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00

	// $0000 LDA $0010
	cpu.Memory.writeSlice(0x0000, []byte{0xa5, 0x10})
	cpu.Memory.Write(0x0010, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaZPLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xFF

	// $0000 LDA $0010
	cpu.Memory.writeSlice(0x0000, []byte{0xa5, 0x10})
	cpu.Memory.Write(0x0010, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaImmediateLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00

	// $0000 LDA #$80
	cpu.Memory.writeSlice(0x0000, []byte{0xa9, 0x80})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaImmediateLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xFF

	// $0000 LDA #$00
	cpu.Memory.writeSlice(0x0000, []byte{0xA9, 0x00})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaAbsXIndexedLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA $abcd,X
	cpu.Memory.writeSlice(0x0000, []byte{0xbd, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.X), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaAbsXIndexedLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xFF
	cpu.X = 0x03

	// $0000 LDA $abcd,X
	cpu.Memory.writeSlice(0x0000, []byte{0xbd, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.X), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaAbsYIndexedLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.Y = 0x03

	// $0000 LDA $abcd,Y
	cpu.Memory.writeSlice(0x0000, []byte{0xb9, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x03), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaAbsYIndexedLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xff
	cpu.Y = 0x03

	// $0000 LDA $abcd,Y
	cpu.Memory.writeSlice(0x0000, []byte{0xb9, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x03), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaIndIndexedXLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA ($0010,X)
	// $0013 vector to $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xa1, 0x10})
	cpu.Memory.writeSlice(0x0013, []byte{0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaIndIndexedXLoadsASetsNFlagIgnoreCarry(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0xe9

	// $0000 LDA ($0051,X)
	// $003a vector to $3104
	cpu.Memory.writeSlice(0x0000, []byte{0xa1, 0x51})
	cpu.Memory.writeSlice(0x003a, []byte{0x04, 0x31})
	cpu.Memory.Write(0x3104, 0x81)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x81), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0xe9), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaIndIndexedXLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA ($0010,X)
	// $0013 vector to $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xa1, 0x10})
	cpu.Memory.writeSlice(0x0013, []byte{0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaIndexedIndYLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.Y = 0x03

	// $0000 LDA ($0010),Y
	// $0010 Vector to $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xb1, 0x10})
	cpu.Memory.writeSlice(0x0010, []byte{0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x03), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaIndexedIndYLoadsASetsNFlagWithCarry(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.Y = 0xe9

	// LDA ($a4),Y
	cpu.Memory.writeSlice(0x0000, []byte{0xb1, 0xa4})
	cpu.Memory.writeSlice(0x00a4, []byte{0x51, 0x3f})
	cpu.Memory.Write(0x403a, 0xbb)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0xbb), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0xe9), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaIndexedIndYLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.Y = 0x03

	// $0000 LDA ($0010),Y
	// $0010 Vector to $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xb1, 0x10})
	cpu.Memory.writeSlice(0x0010, []byte{0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x03), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaZpXIndexedLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA $10,X
	cpu.Memory.writeSlice(0x0000, []byte{0xb5, 0x10})
	cpu.Memory.Write(uint16(0x0010)+uint16(cpu.X), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaZpXIndexedLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xff
	cpu.X = 0x03

	// $0000 LDA $10,X
	cpu.Memory.writeSlice(0x0000, []byte{0xb5, 0x10})
	cpu.Memory.Write(uint16(0x0010)+uint16(cpu.X), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x03), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdaZpXIndexedWrapAroundZp(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0xff

	// $0000 LDA $40,X
	cpu.Memory.writeSlice(0x0000, []byte{0xb5, 0x40})
	cpu.Memory.Write(uint16(0x003f), 0x42)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x42), cpu.A)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0xff), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}
