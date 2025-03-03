package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
