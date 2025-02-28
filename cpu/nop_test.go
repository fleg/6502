package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNopDoesNothing(t *testing.T) {
	cpu := New()

	// $0000 NOP
	cpu.Memory.Write(0x0000, 0xea)
	cpu.Step()

	assert.Equal(t, uint16(0x0001), cpu.PC)
}
