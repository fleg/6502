package cpu

type Flags uint8

const (
	flagEmpty     Flags = 0
	flagCarry     Flags = 1 << 0
	flagZero      Flags = 1 << 1
	flagInterrupt Flags = 1 << 2
	flagDecimal   Flags = 1 << 3
	flagBreak     Flags = 1 << 4
	flagUnused    Flags = 1 << 5
	flagOverflow  Flags = 1 << 6
	flagNegative  Flags = 1 << 7
)

func (cpu *CPU) setFlag(mask Flags, value bool) {
	if value {
		cpu.PS = cpu.PS | mask
	} else {
		cpu.PS = cpu.PS & ^mask
	}
}

func (cpu *CPU) getFlag(mask Flags) bool {
	return cpu.PS&mask > 0
}

func (cpu *CPU) updateZeroFlag(val uint8) {
	cpu.setFlag(flagZero, val == 0)
}

func (cpu *CPU) updateNegativeFlag(val uint8) {
	cpu.setFlag(flagNegative, val&0x80 > 0)
}
