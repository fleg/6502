package cpu

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTiming(t *testing.T) {
	bin, err := os.ReadFile("../test_suites/timing/timingtest-1.bin")
	assert.NoError(t, err)

	cpu := NewWithRAM()

	cpu.writeSlice(0x1000, bin)
	cpu.SP = 0xfd
	cpu.PS = flagInterrupt
	cpu.PC = 0x1000

	for {
		cpu.Step()

		if cpu.PC == 0x1269 {
			break
		}
	}

	expectedTicks := uint64(1141)
	expectedOps := uint64(299)

	assert.Equal(t, expectedTicks, cpu.totalTicks)
	assert.Equal(t, expectedOps, cpu.totalOps)
}
