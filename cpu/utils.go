package cpu

func word(lsb uint8, msb uint8) uint16 {
	return uint16(msb)<<8 | uint16(lsb)
}

func wordLSB(value uint16) uint8 {
	return uint8(value & 0x00ff)
}

func wordMSB(value uint16) uint8 {
	return uint8(value >> 8)
}
