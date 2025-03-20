package cpu

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tomHarteState struct {
	PC  uint16     `json:"pc"`
	SP  uint8      `json:"s"`
	A   uint8      `json:"a"`
	X   uint8      `json:"x"`
	Y   uint8      `json:"y"`
	PS  uint8      `json:"p"`
	RAM [][]uint16 `json:"ram"`
}

type tomHarteTestCase struct {
	Name    string        `json:"name"`
	Initial tomHarteState `json:"initial"`
	Final   tomHarteState `json:"final"`
	Cycles  [][]any       `json:"cycles"`
}

type tomHarteTestSuite []tomHarteTestCase

func TestTomHarte(t *testing.T) {
	for i := range 256 {
		switch i {
		case 0x93, 0x9f, 0x9e, 0x9c, 0x9b:
			// something wrong with operand address here
			// skip it for now
			continue
		}

		p := fmt.Sprintf("../test_suites/tom_harte/6502/v1/%02x.json", i)
		data, err := os.ReadFile(p)
		assert.NoError(t, err)

		var testSuite tomHarteTestSuite
		err = json.Unmarshal(data, &testSuite)
		assert.NoError(t, err)

		for _, tc := range testSuite {
			if tc.Name == "20 55 13" {
				// really weird self modifying instruction edge case
				// skip it for now
				// more details below
				// https://github.com/SingleStepTests/ProcessorTests/issues/65
				// https://github.com/NationalSecurityAgency/ghidra/issues/5871
				continue
			}

			cpu := NewWithRAM()

			cpu.PC = tc.Initial.PC
			cpu.SP = tc.Initial.SP
			cpu.A = tc.Initial.A
			cpu.X = tc.Initial.X
			cpu.Y = tc.Initial.Y
			cpu.PS = Flags(tc.Initial.PS)

			for _, cell := range tc.Initial.RAM {
				cpu.write(cell[0], uint8(cell[1]))
			}

			cpu.Step()

			assert.Equal(t, uint64(1), cpu.totalOps, tc.Name)
			assert.Equal(t, uint64(len(tc.Cycles)), cpu.totalTicks, tc.Name)
			assert.Equal(t, tc.Final.PC, cpu.PC, tc.Name)
			assert.Equal(t, tc.Final.SP, cpu.SP, tc.Name)
			assert.Equal(t, tc.Final.A, cpu.A, tc.Name)
			assert.Equal(t, tc.Final.X, cpu.X, tc.Name)
			assert.Equal(t, tc.Final.Y, cpu.Y, tc.Name)
			assert.Equal(t, Flags(tc.Final.PS), cpu.PS, tc.Name)

			for _, cell := range tc.Final.RAM {
				addr := cell[0]
				expected := uint8(cell[1])
				actual := cpu.read(addr)
				assert.Equal(t, expected, actual, "%s, ram at 0x%04x", tc.Name, addr)
			}
		}

		t.Logf("done %02x (%s)", i, ops[i].Name)
	}
}
