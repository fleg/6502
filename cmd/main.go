package main

import "github.com/fleg/6502/cpu"

func main() {
	cpu.InitOpcodes()

	cpu := cpu.New()

	cpu.Memory.Write(0xfffd, 0x60)
	cpu.Memory.Write(0xfffc, 0x00)
	cpu.Memory.Write(0x6000, 0xea)

	cpu.Reset()

	cpu.Tick()
	cpu.Tick()
	cpu.Tick()
}
