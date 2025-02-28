package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLdaAbsoluteLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	// $0000 LDA $abcd
	cpu.Memory.WriteSlice(0x0000, []byte{0xad, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaAbsoluteLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xff

	// $0000 LDA $abcd
	cpu.Memory.WriteSlice(0x0000, []byte{0xad, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}

func TestLdaZPLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00

	// $0000 LDA $0010
	cpu.Memory.WriteSlice(0x0000, []byte{0xa5, 0x10})
	cpu.Memory.Write(0x0010, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaZPLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xFF

	// $0000 LDA $0010
	cpu.Memory.WriteSlice(0x0000, []byte{0xa5, 0x10})
	cpu.Memory.Write(0x0010, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}

func TestLdaImmediateLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00

	// $0000 LDA #$80
	cpu.Memory.WriteSlice(0x0000, []byte{0xa9, 0x80})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaImmediateLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xFF

	// $0000 LDA #$00
	cpu.Memory.WriteSlice(0x0000, []byte{0xA9, 0x00})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}

func TestLdaAbsXIndexedLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA $ABCD,X
	cpu.Memory.WriteSlice(0x0000, []byte{0xbd, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xABCD)+uint16(cpu.X), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaAbsXIndexedLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xFF
	cpu.X = 0x03

	// $0000 LDA $ABCD,X
	cpu.Memory.WriteSlice(0x0000, []byte{0xbd, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.X), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}

func TestLdaAbsXIndexedDoesNotPageWrap(t *testing.T) {
	cpu := New()
	cpu.A = 0
	cpu.X = 0xff

	// $0000 LDA $0080,X
	cpu.Memory.WriteSlice(0x0000, []byte{0xbd, 0x80, 0x00})
	cpu.Memory.Write(uint16(0x0080)+uint16(cpu.X), 0x42)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x42), cpu.A)
}

func TestLdaAbsYIndexedLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.Y = 0x03

	// $0000 LDA $abcd,Y
	cpu.Memory.WriteSlice(0x0000, []byte{0xb9, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaAbsYIndexedLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xff
	cpu.Y = 0x03

	// $0000 LDA $abcd,Y
	cpu.Memory.WriteSlice(0x0000, []byte{0xb9, 0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}

func TestLdaAbsYIndexedDoesNotPageWrap(t *testing.T) {
	cpu := New()
	cpu.A = 0
	cpu.Y = 0xff

	// $0000 LDA $0080,Y
	cpu.Memory.WriteSlice(0x0000, []byte{0xb9, 0x80, 0x00})
	cpu.Memory.Write(uint16(0x0080)+uint16(cpu.Y), 0x42)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x42), cpu.A)
}

func TestLdaIndIndexedXLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA ($0010,X)
	// $0013 Vector to $ABCD
	cpu.Memory.WriteSlice(0x0000, []byte{0xa1, 0x10})
	cpu.Memory.WriteSlice(0x0013, []byte{0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaIndIndexedXLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA ($0010,X)
	// $0013 Vector to $ABCD
	cpu.Memory.WriteSlice(0x0000, []byte{0xa1, 0x10})
	cpu.Memory.WriteSlice(0x0013, []byte{0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}

func TestLdaIndexedIndYLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.Y = 0x03

	// $0000 LDA ($0010),Y
	// $0010 Vector to $ABCD
	cpu.Memory.WriteSlice(0x0000, []byte{0xb1, 0x10})
	cpu.Memory.WriteSlice(0x0010, []byte{0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaIndexedIndYLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.Y = 0x03

	// $0000 LDA ($0010),Y
	// $0010 Vector to $ABCD
	cpu.Memory.WriteSlice(0x0000, []byte{0xb1, 0x10})
	cpu.Memory.WriteSlice(0x0010, []byte{0xcd, 0xab})
	cpu.Memory.Write(uint16(0xabcd)+uint16(cpu.Y), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}

func TestLdaZPXIndexedLoadsASetsNFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0x00
	cpu.X = 0x03

	// $0000 LDA $10,X
	cpu.Memory.WriteSlice(0x0000, []byte{0xb5, 0x10})
	cpu.Memory.Write(uint16(0x0010)+uint16(cpu.X), 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagNegative)
	assert.Equal(t, Flags(0), cpu.PS&flagZero)
}

func TestLdaZPXIndexedLoadsASetsZFlag(t *testing.T) {
	cpu := New()
	cpu.A = 0xff
	cpu.X = 0x03

	// $0000 LDA $10,X
	cpu.Memory.WriteSlice(0x0000, []byte{0xb5, 0x10})
	cpu.Memory.Write(uint16(0x0010)+uint16(cpu.X), 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, cpu.PS, cpu.PS&flagZero)
	assert.Equal(t, Flags(0), cpu.PS&flagNegative)
}
