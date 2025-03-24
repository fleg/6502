package cpu

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKlausDormannFunctional(t *testing.T) {
	bin, err := os.ReadFile("../test_suites/klaus_dormann/bin/6502_functional_test.bin")
	assert.NoError(t, err)

	cpu := NewWithRAM()

	cpu.writeSlice(0x000a, bin)
	cpu.SP = 0xfd
	cpu.PS = flagInterrupt
	cpu.PC = 0x0400

	prevPC := cpu.PC
	fail := false

	for {
		cpu.Step()

		// print char
		if cpu.PC == 0x455c {
			fmt.Printf("%c", cpu.A)
		}

		// get char
		if cpu.PC == 0x453a {
			// just break here without getting char
			break
		}

		// fail
		if cpu.PC == 0x445b {
			fail = true
		}

		if cpu.PC == prevPC {
			break
		}

		prevPC = cpu.PC
	}

	if fail {
		t.Fail()
	}
}

func TestKlausDormannBCD(t *testing.T) {
	bin, err := os.ReadFile("../test_suites/klaus_dormann/bin/6502_decimal_test.bin")
	assert.NoError(t, err)

	cpu := NewWithRAM()

	cpu.writeSlice(0x0200, bin)
	cpu.SP = 0xfd
	cpu.PS = flagInterrupt
	cpu.PC = 0x0200

	for {
		cpu.Step()

		if cpu.PC == 0x024b {
			break
		}
	}

	if cpu.read(0x000b) == 0x01 {
		t.Fail()
	}
}

func TestKlausDormannInterrupt(t *testing.T) {
	bin, err := os.ReadFile("../test_suites/klaus_dormann/bin/6502_interrupt_test.bin")
	assert.NoError(t, err)

	cpu := NewWithRAM()

	cpu.writeSlice(0x000a, bin)
	cpu.SP = 0xfd
	cpu.PS = flagInterrupt
	cpu.PC = 0x0400

	prevPC := cpu.PC
	fail := false

	const nmiMask = 0x02
	const irqMask = 0x01
	const intPortAddr = 0xbffc

	intState := uint8(0)
	cpu.write(intPortAddr, 0x00)

	for {
		cpu.Step()

		intState = cpu.read(intPortAddr)
		if intState&nmiMask == nmiMask {
			if cpu.TriggerNMI() {
				intState = intState & ^uint8(nmiMask)
			}
		}
		if intState&irqMask == irqMask {
			if cpu.TriggerIRQ() {
				intState = intState & ^uint8(irqMask)
			}
		}
		cpu.write(intPortAddr, intState)

		// print char
		if cpu.PC == 0x09e0 {
			fmt.Printf("%c", cpu.A)
		}

		// get char
		if cpu.PC == 0x09be {
			// just break here without getting char
			break
		}

		// fail
		if cpu.PC == 0x08b9 {
			fail = true
		}

		if cpu.PC == prevPC {
			break
		}

		prevPC = cpu.PC
	}

	if fail {
		t.Fail()
	}
}
