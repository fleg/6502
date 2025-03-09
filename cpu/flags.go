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

func (cpu *CPU) getCarry() uint8 {
	if cpu.getFlag(flagCarry) {
		return 1
	}

	return 0
}

func (cpu *CPU) updateOverflowFlag(a uint8, b uint8, c uint8) {
	// c = a + b
	signA := a & 0x80
	signB := b & 0x80
	signC := c & 0x80

	// if the sign(c) is different from both sign(a) and sign(b)

	// if both inputs are positive and the result is negative,
	// or both are negative and the result is positive
	//overflow := (signA^signC)&(signB^signC) > 0
	overflow := (signA == 0 && signB == 0 && signC > 0) || (signA > 0 && signB > 0 && signC == 0)

	cpu.setFlag(flagOverflow, overflow)
}
