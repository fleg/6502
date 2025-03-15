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

func TestAndAbsoluteAndWithMem(t *testing.T) {
	cpu := New()
	cpu.A = 0x0f

	// $0000 and $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x04)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x04), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x04), cpu.Memory.Read(0xabcd))
}

func TestAndAbsoluteAndWithMemSetsZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x55

	// $0000 and $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2d, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0xaa)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0xaa), cpu.Memory.Read(0xabcd))
}

func TestAndAbsoluteAndWithMemSetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x80

	// $0000 and $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x2d, 0xcd, 0xab})
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

func TestAslAccumulatorShiftsA(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 asl a
	cpu.Memory.Write(0x0000, 0x0a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x02), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestAslAccumulatorShiftsASetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x80

	// $0000 asl a
	cpu.Memory.Write(0x0000, 0x0a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestAslAccumulatorShiftsASetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x55

	// $0000 asl a
	cpu.Memory.Write(0x0000, 0x0a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0xaa), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestAslAbsoluteShiftsMem(t *testing.T) {
	cpu := New()

	// $0000 asl $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x0e, 0xcd, 0xab})
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

func TestAslAbsoluteShiftsASetsCarryAndZero(t *testing.T) {
	cpu := New()

	// $0000 asl $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x0e, 0xcd, 0xab})
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

func TestAslAbsoluteShiftsASetsNegative(t *testing.T) {
	cpu := New()

	// $0000 asl $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x0e, 0xcd, 0xab})
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

func TestBccCarryClearJumpsForward(t *testing.T) {
	cpu := New()

	// $0000 BCC +6
	cpu.Memory.writeSlice(0x0000, []byte{0x90, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBccCarryClearJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008

	// $0008 BCC -6
	cpu.Memory.writeSlice(0x0008, []byte{0x90, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBccCarrySetDoesNotJump(t *testing.T) {
	cpu := New()
	cpu.PS = flagCarry

	// $0000 BCC +6
	cpu.Memory.writeSlice(0x0000, []byte{0x90, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBcsCarrySetJumpsForward(t *testing.T) {
	cpu := New()
	cpu.PS = flagCarry

	// $0000 BCS +6
	cpu.Memory.writeSlice(0x0000, []byte{0xb0, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBccCarrySetJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008
	cpu.PS = flagCarry

	// $0008 BCS -6
	cpu.Memory.writeSlice(0x0008, []byte{0xb0, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBccCarryCleanDoesNotJump(t *testing.T) {
	cpu := New()

	// $0000 BCS +6
	cpu.Memory.writeSlice(0x0000, []byte{0xb0, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBeqZeroSetJumpsForward(t *testing.T) {
	cpu := New()
	cpu.PS = flagZero

	// $0000 BEQ +6
	cpu.Memory.writeSlice(0x0000, []byte{0xf0, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBeqZeroSetJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008
	cpu.PS = flagZero

	// $0008 BEQ -6
	cpu.Memory.writeSlice(0x0008, []byte{0xf0, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBeqZeroCleanDoesNotJump(t *testing.T) {
	cpu := New()

	// $0000 BEQ +6
	cpu.Memory.writeSlice(0x0000, []byte{0xf0, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

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
	assert.Equal(t, flagZero|flagNegative, cpu.PS)
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

func TestBmiNegativeSetJumpsForward(t *testing.T) {
	cpu := New()
	cpu.PS = flagNegative

	// $0000 BMI +6
	cpu.Memory.writeSlice(0x0000, []byte{0x30, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBmiNegativeSetJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008
	cpu.PS = flagNegative

	// $0008 BMI -6
	cpu.Memory.writeSlice(0x0008, []byte{0x30, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBmiNegativeClearDoesNotJump(t *testing.T) {
	cpu := New()

	// $0000 BMI +6
	cpu.Memory.writeSlice(0x0000, []byte{0x30, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBneZeroClearJumpsForward(t *testing.T) {
	cpu := New()

	// $0000 BNE +6
	cpu.Memory.writeSlice(0x0000, []byte{0xd0, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBneZeroClearJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008

	// $0008 BEQ -6
	cpu.Memory.writeSlice(0x0008, []byte{0xd0, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBneZeroSetDoesNotJump(t *testing.T) {
	cpu := New()
	cpu.PS = flagZero

	// $0000 BEQ +6
	cpu.Memory.writeSlice(0x0000, []byte{0xd0, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBplNegativeClearJumpsForward(t *testing.T) {
	cpu := New()

	// $0000 BPL +6
	cpu.Memory.writeSlice(0x0000, []byte{0x10, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBplNegativeClearJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008

	// $0008 BPL -6
	cpu.Memory.writeSlice(0x0008, []byte{0x10, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBplNegativeSetDoesNotJump(t *testing.T) {
	cpu := New()
	cpu.PS = flagNegative

	// $0000 BPL +6
	cpu.Memory.writeSlice(0x0000, []byte{0x10, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBrkPushesPCAndPSAndUpdatePC(t *testing.T) {
	cpu := New()
	cpu.PC = 0xc000
	cpu.SP = 0xfd

	// $c000 BRK
	cpu.Memory.Write(0xc000, 0x00)
	cpu.Memory.writeSlice(0xfffe, []byte{0xcd, 0xab})
	cpu.Step()

	assert.Equal(t, uint16(0xabcd), cpu.PC)
	assert.Equal(t, flagInterrupt, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfa), cpu.SP)
	assert.Equal(t, uint8(0xc0), cpu.Memory.Read(0x01fd))
	assert.Equal(t, uint8(0x02), cpu.Memory.Read(0x01fc))
	assert.Equal(t, uint8(flagUnused|flagBreak), cpu.Memory.Read(0x01fb))
}

func TestBvcOverflowClearJumpsForward(t *testing.T) {
	cpu := New()

	// $0000 BVC +6
	cpu.Memory.writeSlice(0x0000, []byte{0x50, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBvcOverflowClearJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008

	// $0008 BVC -6
	cpu.Memory.writeSlice(0x0008, []byte{0x50, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBvcOverflowSetDoesNotJump(t *testing.T) {
	cpu := New()
	cpu.PS = flagOverflow

	// $0000 BVC +6
	cpu.Memory.writeSlice(0x0000, []byte{0x50, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagOverflow, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBvsOverflowSetJumpsForward(t *testing.T) {
	cpu := New()
	cpu.PS = flagOverflow

	// $0000 BVS +6
	cpu.Memory.writeSlice(0x0000, []byte{0x70, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0008), cpu.PC)
	assert.Equal(t, flagOverflow, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBvsOverflowSetJumpsBackward(t *testing.T) {
	cpu := New()
	cpu.PC = 0x0008
	cpu.PS = flagOverflow

	// $0008 BVS -6
	cpu.Memory.writeSlice(0x0008, []byte{0x70, 0xfa})
	cpu.Step()

	assert.Equal(t, uint16(0x0004), cpu.PC)
	assert.Equal(t, flagOverflow, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestBvsOverflowCleanDoesNotJump(t *testing.T) {
	cpu := New()

	// $0000 BVS +6
	cpu.Memory.writeSlice(0x0000, []byte{0x70, 0x06})
	cpu.Step()

	assert.Equal(t, uint16(0x0002), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestClcClearsCarry(t *testing.T) {
	cpu := New()
	cpu.PS = 0xff

	// $0000 clc
	cpu.Memory.Write(0x0000, 0x18)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xff) & ^flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestCldClearsDecimal(t *testing.T) {
	cpu := New()
	cpu.PS = 0xff

	// $0000 cld
	cpu.Memory.Write(0x0000, 0xd8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xff) & ^flagDecimal, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestCliClearsInterrupt(t *testing.T) {
	cpu := New()
	cpu.PS = 0xff

	// $0000 clc
	cpu.Memory.Write(0x0000, 0x58)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xff) & ^flagInterrupt, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestClvClearsOverflow(t *testing.T) {
	cpu := New()
	cpu.PS = 0xff

	// $0000 clv
	cpu.Memory.Write(0x0000, 0xb8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xff) & ^flagOverflow, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestCmpAbsoluteSetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 cmp $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcd, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x01), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCmpAbsoluteSetsCarry(t *testing.T) {
	cpu := New()
	cpu.A = 0x02

	// $0000 cmp $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcd, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x02), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCmpAbsoluteSetsNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0x00

	// $0000 cmp $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcd, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpxAbsoluteSetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.X = 0x01

	// $0000 cpx $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xec, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x01), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpxAbsoluteSetsCarry(t *testing.T) {
	cpu := New()
	cpu.X = 0x02

	// $0000 cpx $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xec, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x02), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpxAbsoluteSetsNegative(t *testing.T) {
	cpu := New()
	cpu.X = 0x00

	// $0000 cpx $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xec, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpyAbsoluteSetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.Y = 0x01

	// $0000 cpy $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcc, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x01), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpyAbsoluteSetsCarry(t *testing.T) {
	cpu := New()
	cpu.Y = 0x02

	// $0000 cpy $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcc, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x02), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestCpyAbsoluteSetsNegative(t *testing.T) {
	cpu := New()
	cpu.Y = 0x00

	// $0000 cpy $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xcc, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x01), cpu.Memory.Read(0xabcd))
}

func TestDecAbsoluteDecrementsMem(t *testing.T) {
	cpu := New()

	// $0000 dec $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xce, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x08)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x07), cpu.Memory.Read(0xabcd))
}

func TestDecAbsoluteDecrementsMemSetsZero(t *testing.T) {
	cpu := New()

	// $0000 dec $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xce, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x01)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x00), cpu.Memory.Read(0xabcd))
}

func TestDecAbsoluteDecrementsMemSetsNegative(t *testing.T) {
	cpu := New()

	// $0000 dec $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xce, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0xff), cpu.Memory.Read(0xabcd))
}

func TestDexDecrementsX(t *testing.T) {
	cpu := New()
	cpu.X = 0x08

	// $0000 dex
	cpu.Memory.Write(0x0000, 0xca)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x07), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestDexDecrementsXSetsZero(t *testing.T) {
	cpu := New()
	cpu.X = 0x01

	// $0000 dex
	cpu.Memory.Write(0x0000, 0xca)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestDexDecrementsXSetsNegative(t *testing.T) {
	cpu := New()
	cpu.X = 0x00

	// $0000 dex
	cpu.Memory.Write(0x0000, 0xca)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0xff), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestDeyDecrementsY(t *testing.T) {
	cpu := New()
	cpu.Y = 0x08

	// $0000 dey
	cpu.Memory.Write(0x0000, 0x88)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x07), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestDeyDecrementsYSetsZero(t *testing.T) {
	cpu := New()
	cpu.Y = 0x01

	// $0000 dey
	cpu.Memory.Write(0x0000, 0x88)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestDeyDecrementsYSetsNegative(t *testing.T) {
	cpu := New()
	cpu.Y = 0x00

	// $0000 dey
	cpu.Memory.Write(0x0000, 0x88)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0xff), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

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

func TestIncAbsoluteIncrementsMem(t *testing.T) {
	cpu := New()

	// $0000 inc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xee, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x08)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x09), cpu.Memory.Read(0xabcd))
}

func TestIncAbsoluteIncrementsMemSetsZero(t *testing.T) {
	cpu := New()

	// $0000 inc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xee, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0xff)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x00), cpu.Memory.Read(0xabcd))
}

func TestIncAbsoluteIncrementsMemSetsNegative(t *testing.T) {
	cpu := New()

	// $0000 inc $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xee, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x7f)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x80), cpu.Memory.Read(0xabcd))
}

func TestInxIncrementsX(t *testing.T) {
	cpu := New()
	cpu.X = 0x08

	// $0000 inx
	cpu.Memory.Write(0x0000, 0xe8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x09), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInxIncrementsXSetsZero(t *testing.T) {
	cpu := New()
	cpu.X = 0xff

	// $0000 inx
	cpu.Memory.Write(0x0000, 0xe8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInxIncrementsXSetsNegative(t *testing.T) {
	cpu := New()
	cpu.X = 0x7f

	// $0000 inx
	cpu.Memory.Write(0x0000, 0xe8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x80), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInyIncrementsY(t *testing.T) {
	cpu := New()
	cpu.Y = 0x08

	// $0000 iny
	cpu.Memory.Write(0x0000, 0xc8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x09), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInyIncrementsYSetsZero(t *testing.T) {
	cpu := New()
	cpu.Y = 0xff

	// $0000 iny
	cpu.Memory.Write(0x0000, 0xc8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestInyIncrementsYSetsNegative(t *testing.T) {
	cpu := New()
	cpu.Y = 0x7f

	// $0000 iny
	cpu.Memory.Write(0x0000, 0xc8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x80), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestJmpAbsoluteModifyPC(t *testing.T) {
	cpu := New()

	// $0000 JMP $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x4c, 0xcd, 0xab})
	cpu.Step()

	assert.Equal(t, uint16(0xabcd), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestJmpIndirectModifyPC(t *testing.T) {
	cpu := New()

	// $0000 JMP ($abcd)
	cpu.Memory.writeSlice(0x0000, []byte{0x6c, 0x00, 0x02})
	cpu.Memory.writeSlice(0x0200, []byte{0xcd, 0xab})
	cpu.Step()

	assert.Equal(t, uint16(0xabcd), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestJsrPushesAndUpdatePC(t *testing.T) {
	cpu := New()
	cpu.PC = 0xc010
	cpu.SP = 0xfd

	// $c010 JSR $ffd2
	cpu.Memory.writeSlice(0xc010, []byte{0x20, 0xd2, 0xff})
	cpu.Step()

	assert.Equal(t, uint16(0xffd2), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfb), cpu.SP)
	assert.Equal(t, uint8(0xc0), cpu.Memory.Read(0x01fd))
	assert.Equal(t, uint8(0x12), cpu.Memory.Read(0x01fc))
}

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

func TestLdxAbsoluteLoadsXSetsNFlag(t *testing.T) {
	cpu := New()
	cpu.X = 0x00

	// $0000 LDX $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xae, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.X)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdxAbsoluteLoadsXSetsZFlag(t *testing.T) {
	cpu := New()
	cpu.X = 0xff

	// $0000 LDX $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xae, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdyAbsoluteLoadsYSetsNFlag(t *testing.T) {
	cpu := New()
	cpu.Y = 0x00

	// $0000 LDY $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xac, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x80), cpu.Y)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLdyAbsoluteLoadsYSetsZFlag(t *testing.T) {
	cpu := New()
	cpu.Y = 0xff

	// $0000 LDY $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0xac, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLsrAccumulatorShiftsA(t *testing.T) {
	cpu := New()
	cpu.A = 0x02

	// $0000 lsr a
	cpu.Memory.Write(0x0000, 0x4a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x01), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLsrAccumulatorShiftsASetsCarryAndZero(t *testing.T) {
	cpu := New()
	cpu.A = 0x01

	// $0000 lsr a
	cpu.Memory.Write(0x0000, 0x4a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagCarry|flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLsrAccumulatorShiftsAClearNegative(t *testing.T) {
	cpu := New()
	cpu.A = 0xaa
	cpu.PS = flagNegative

	// $0000 lsr a
	cpu.Memory.Write(0x0000, 0x4a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x55), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestLsrAbsoluteShiftsMem(t *testing.T) {
	cpu := New()

	// $0000 lsr $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x4e, 0xcd, 0xab})
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

func TestLsrAbsoluteShiftsASetsCarryAndZero(t *testing.T) {
	cpu := New()

	// $0000 lsr $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x4e, 0xcd, 0xab})
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

func TestLsrAbsoluteShiftsAClearNegative(t *testing.T) {
	cpu := New()
	cpu.PS = flagNegative

	// $0000 lsr $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x4e, 0xcd, 0xab})
	cpu.Memory.Write(0xabcd, 0xaa)
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x55), cpu.Memory.Read(0xabcd))
}

func TestNopDoesNothing(t *testing.T) {
	cpu := New()

	// $0000 NOP
	cpu.Memory.Write(0x0000, 0xea)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

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

func TestPhpPushesPAndUpdatesSP(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.PS = Flags(0x55)

	// $0000 PHP
	cpu.Memory.Write(0x0000, 0x08)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0x55), cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfc), cpu.SP)
	assert.Equal(t, uint8(0x55), cpu.Memory.Read(0x01fd))
}

func TestPlaPullsIntoAAndUpdatesSP(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x01fe, 0x3a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x3a), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}

func TestPlaPullsIntoAAndUpdatesSPOverflow(t *testing.T) {
	cpu := New()
	cpu.SP = 0xff
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x0100, 0x3a)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x3a), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestPlaPullsIntoAAndSetsZFlag(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x01fe, 0x00)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagZero, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}

func TestPlaPullsIntoAAndSetsNFlag(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.A = 0xff

	// $0000 PLA
	cpu.Memory.Write(0x0000, 0x68)
	cpu.Memory.Write(0x01fe, 0x80)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagNegative, cpu.PS)
	assert.Equal(t, uint8(0x80), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}

func TestPlpPullsIntoPAndUpdatesSP(t *testing.T) {
	cpu := New()
	cpu.SP = 0xfd
	cpu.PS = 0xff

	// $0000 PLP
	cpu.Memory.Write(0x0000, 0x28)
	cpu.Memory.Write(0x01fe, 0xaa)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, Flags(0xaa)|flagUnused|flagBreak, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xfe), cpu.SP)
}

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

func TestRtiPullsPCAndPSAndUpdatePC(t *testing.T) {
	cpu := New()
	cpu.SP = 0xf0

	// $0000 RTI
	cpu.Memory.Write(0x0000, 0x40)
	cpu.Memory.writeSlice(0x01f1, []byte{0x01, 0x03, 0xc0})
	cpu.Step()

	assert.Equal(t, uint16(0xc003), cpu.PC)
	assert.Equal(t, flagCarry|flagBreak|flagUnused, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xf3), cpu.SP)
}

func TestRtsPullsAndUpdatePC(t *testing.T) {
	cpu := New()
	cpu.SP = 0xf0

	// $0000 RTS
	cpu.Memory.Write(0x0000, 0x60)
	cpu.Memory.writeSlice(0x01f1, []byte{0x10, 0xc0})
	cpu.Step()

	assert.Equal(t, uint16(0xc011), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xf2), cpu.SP)
}

func TestRtsPullsAndUpdatePCWrapAround(t *testing.T) {
	cpu := New()
	cpu.PC = 0xf000
	cpu.SP = 0xf0

	// $f000 RTS
	cpu.Memory.Write(0xf000, 0x60)
	cpu.Memory.writeSlice(0x01f1, []byte{0xff, 0xff})
	cpu.Step()

	assert.Equal(t, uint16(0x0000), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0xf2), cpu.SP)
}

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

func TestSecSetsCarry(t *testing.T) {
	cpu := New()

	// $0000 sec
	cpu.Memory.Write(0x0000, 0x38)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagCarry, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestSedSetsDecimal(t *testing.T) {
	cpu := New()

	// $0000 sed
	cpu.Memory.Write(0x0000, 0xf8)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagDecimal, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestSeiSetsInterrupt(t *testing.T) {
	cpu := New()

	// $0000 sei
	cpu.Memory.Write(0x0000, 0x78)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
	assert.Equal(t, flagInterrupt, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
}

func TestStaAbsoluteStoresA(t *testing.T) {
	cpu := New()
	cpu.A = 0x55

	// $0000 STA $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x8d, 0xcd, 0xab})
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x55), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x55), cpu.Memory.Read(0xabcd))
}

func TestStxAbsoluteStoresX(t *testing.T) {
	cpu := New()
	cpu.X = 0x55

	// $0000 STX $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x8e, 0xcd, 0xab})
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x55), cpu.X)
	assert.Equal(t, uint8(0x00), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x55), cpu.Memory.Read(0xabcd))
}

func TestStyAbsoluteStoresY(t *testing.T) {
	cpu := New()
	cpu.Y = 0x55

	// $0000 STY $abcd
	cpu.Memory.writeSlice(0x0000, []byte{0x8c, 0xcd, 0xab})
	cpu.Step()

	assert.Equal(t, uint16(0x0003), cpu.PC)
	assert.Equal(t, flagEmpty, cpu.PS)
	assert.Equal(t, uint8(0x00), cpu.A)
	assert.Equal(t, uint8(0x00), cpu.X)
	assert.Equal(t, uint8(0x55), cpu.Y)
	assert.Equal(t, uint8(0x00), cpu.SP)
	assert.Equal(t, uint8(0x55), cpu.Memory.Read(0xabcd))
}
