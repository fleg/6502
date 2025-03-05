package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
