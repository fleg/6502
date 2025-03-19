package main

import "github.com/fleg/6502/cpu"

func main() {
	ram := cpu.NewRAM()
	cpu := cpu.New(ram)

	ram.Write(0xfffd, 0x60)
	ram.Write(0xfffc, 0x00)
	ram.Write(0x6000, 0xea)

	cpu.Reset()

	cpu.Step()
}
