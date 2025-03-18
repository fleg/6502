package cpu

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var petToAscTable = []byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x14, 0x09, 0x0d, 0x11, 0x93, 0x0a, 0x0e, 0x0f,
	0x10, 0x0b, 0x12, 0x13, 0x08, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
	0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	0x40, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
	0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
	0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf,
	0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf,
	0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
	0x90, 0x91, 0x92, 0x0c, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f,
	0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7, 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf,
	0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7, 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf,
	0x60, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
	0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
	0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7, 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf,
	0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7, 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf,
}

func loadBinaries(t *testing.T) map[string][]byte {
	entries, err := os.ReadDir("../test_suites/wolfgang_lorenz/bin")
	assert.NoError(t, err)
	assert.Len(t, entries, 265)

	result := make(map[string][]byte, 265)

	for _, e := range entries {
		name := e.Name()

		content, err := os.ReadFile("../test_suites/wolfgang_lorenz/bin/" + name)
		assert.NoError(t, err)

		result[name] = content
	}

	return result
}

func TestWolfgangLorenz(t *testing.T) {
	binaries := loadBinaries(t)
	cpu := New()

	var output bytes.Buffer

	load := func(name string) bool {
		if name == "trap17" {
			return true
		}

		bin, ok := binaries[name]
		assert.True(t, ok)

		entryPoint := word(bin[0], bin[1])
		cpu.Memory.writeSlice(entryPoint, bin[2:])

		cpu.Memory.Write(0x0002, 0x00)
		cpu.Memory.Write(0xa002, 0x00)
		cpu.Memory.Write(0xa003, 0x80)
		cpu.Memory.Write(0xfffe, 0x48)
		cpu.Memory.Write(0xffff, 0xff)
		cpu.Memory.Write(0x01fe, 0xff)
		cpu.Memory.Write(0x01ff, 0x7f)

		cpu.Memory.writeSlice(0xff48, []byte{0x48})
		cpu.Memory.writeSlice(0xff49, []byte{0x8a})
		cpu.Memory.writeSlice(0xff4a, []byte{0x48})
		cpu.Memory.writeSlice(0xff4b, []byte{0x98})
		cpu.Memory.writeSlice(0xff4c, []byte{0x48})
		cpu.Memory.writeSlice(0xff4d, []byte{0xba})
		cpu.Memory.writeSlice(0xff4e, []byte{0xbd, 0x04, 0x01})
		cpu.Memory.writeSlice(0xff51, []byte{0x29, 0x10})
		cpu.Memory.writeSlice(0xff53, []byte{0xf0, 0x03})
		cpu.Memory.writeSlice(0xff55, []byte{0x6c, 0x16, 0x03})
		cpu.Memory.writeSlice(0xff58, []byte{0x6c, 0x14, 0x03})

		cpu.SP = 0xfd
		cpu.PS = flagInterrupt
		cpu.PC = 0x0801

		return false
	}

	load("start")

	for {
		cpu.Step()

		// print char
		if cpu.PC == 0xffd2 {
			cpu.Memory.Write(0x030c, 0x00)
			cpu.PC = cpu.popWord() + 1

			c := petToAscTable[cpu.A]
			output.WriteByte(c)

			if c == '\n' {
				line, err := output.ReadString('\n')
				if err == nil {
					if strings.HasPrefix(line, "\x91") {
						if strings.HasSuffix(line, "ok\n") {
							t.Run(line[1:len(line)-6], func(t *testing.T) {})
						} else {
							t.Run(line[1:len(line)-1], func(t *testing.T) {
								t.Fail()
							})
						}
					} else {
						if strings.HasPrefix(line, "before") {
							t.Log("        arg a  x  y  flags    sp")
						}
						if strings.ContainsAny(line, " ") {
							t.Log(line[:len(line)-1])
						}
					}
				}
			}
		}

		// load
		if cpu.PC == 0xe16f {
			addr := word(cpu.Memory.Read(0x00bb), cpu.Memory.Read(0x00bc))
			len := cpu.Memory.Read(0x00b7)
			name := ""

			for i := range uint16(len) {
				name = name + string(petToAscTable[cpu.Memory.Read(addr+i)])
			}

			if load(name) {
				break
			}

			cpu.popWord()
			cpu.PC = 0x0816
		}

		// scan
		if cpu.PC == 0xffe4 {
			cpu.A = 0x03
			cpu.PC = cpu.popWord() + 1
		}

		// exit
		if cpu.PC == 0x8000 || cpu.PC == 0xa474 {
			break
		}
	}
}
