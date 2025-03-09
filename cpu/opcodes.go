package cpu

import "fmt"

var ops = [256]*Op{
	// opcode: name addrmode size ticks do
	0xea: {"nop", amImp, 1, 2, nop},

	0xa9: {"lda", amImm, 2, 2, lda},
	0xa5: {"lda", amZeP, 2, 3, lda},
	0xb5: {"lda", amZeX, 2, 4, lda},
	0xad: {"lda", amAbs, 3, 4, lda},
	0xbd: {"lda", amAbX, 3, 4, lda},
	0xb9: {"lda", amAbY, 3, 4, lda},
	0xa1: {"lda", amInX, 2, 6, lda},
	0xb1: {"lda", amInY, 2, 5, lda},

	0xa2: {"ldx", amImm, 2, 2, ldx},
	0xa6: {"ldx", amZeP, 2, 3, ldx},
	0xb6: {"ldx", amZeY, 2, 4, ldx},
	0xae: {"ldx", amAbs, 3, 4, ldx},
	0xbe: {"ldx", amAbY, 3, 4, ldx},

	0xa0: {"ldy", amImm, 2, 2, ldy},
	0xa4: {"ldy", amZeP, 2, 3, ldy},
	0xb4: {"ldy", amZeX, 2, 4, ldy},
	0xac: {"ldy", amAbs, 3, 4, ldy},
	0xbc: {"ldy", amAbX, 3, 4, ldy},

	0x85: {"sta", amZeP, 2, 3, sta},
	0x95: {"sta", amZeX, 2, 4, sta},
	0x8d: {"sta", amAbs, 3, 4, sta},
	0x9d: {"sta", amAbX, 3, 5, sta},
	0x99: {"sta", amAbY, 3, 5, sta},
	0x81: {"sta", amInX, 2, 6, sta},
	0x91: {"sta", amInY, 2, 6, sta},

	0x86: {"stx", amZeP, 2, 3, stx},
	0x96: {"stx", amZeY, 2, 4, stx},
	0x8e: {"stx", amAbs, 3, 4, stx},

	0x84: {"sty", amZeP, 2, 3, sty},
	0x94: {"sty", amZeY, 2, 4, sty},
	0x8c: {"sty", amAbs, 3, 4, sty},

	0xaa: {"tax", amImp, 1, 2, tax},
	0xa8: {"tay", amImp, 1, 2, tay},
	0xba: {"tsx", amImp, 1, 2, tsx},
	0x8a: {"txa", amImp, 1, 2, txa},
	0x9a: {"txs", amImp, 1, 2, txs},
	0x98: {"tya", amImp, 1, 2, tya},

	0x48: {"pha", amImp, 1, 3, pha},
	0x68: {"pla", amImp, 1, 3, pla},
	0x08: {"php", amImp, 1, 3, php},
	0x28: {"plp", amImp, 1, 3, plp},

	0x4c: {"jmp", amAbs, 3, 3, jmp},
	0x6c: {"jmp", amInd, 3, 5, jmp},

	0x20: {"jsr", amAbs, 3, 6, jsr},
	0x60: {"rts", amImp, 1, 6, rts},

	0x00: {"brk", amImp, 1, 7, brk},
	0x40: {"rti", amImp, 1, 6, rti},

	0xb0: {"bcs", amRel, 2, 2, bcs},
	0x90: {"bcc", amRel, 2, 2, bcc},
	0xf0: {"beq", amRel, 2, 2, beq},
	0xd0: {"bne", amRel, 2, 2, bne},
	0x30: {"bmi", amRel, 2, 2, bmi},
	0x10: {"bpl", amRel, 2, 2, bpl},
	0x70: {"bvs", amRel, 2, 2, bvs},
	0x50: {"bvc", amRel, 2, 2, bvc},

	0x18: {"clc", amImp, 1, 2, clc},
	0xd8: {"cld", amImp, 1, 2, cld},
	0x58: {"cli", amImp, 1, 2, cli},
	0xb8: {"clv", amImp, 1, 2, clv},

	0x38: {"sec", amImp, 1, 2, sec},
	0xf8: {"sed", amImp, 1, 2, sed},
	0x78: {"sei", amImp, 1, 2, sei},

	0xc6: {"dec", amZeP, 2, 5, dec},
	0xd6: {"dec", amZeX, 2, 6, dec},
	0xce: {"dec", amAbs, 2, 6, dec},
	0xde: {"dec", amAbX, 2, 7, dec},

	0xca: {"dex", amImp, 1, 2, dex},
	0x88: {"dey", amImp, 1, 2, dey},

	0xe6: {"inc", amZeP, 2, 5, inc},
	0xf6: {"inc", amZeX, 2, 6, inc},
	0xee: {"inc", amAbs, 2, 6, inc},
	0xfe: {"inc", amAbX, 2, 5, inc},

	0xe8: {"inx", amImp, 1, 2, inx},
	0xc8: {"iny", amImp, 1, 2, iny},

	0x0a: {"asl", amAcc, 1, 2, asl},
	0x06: {"asl", amZeP, 2, 5, asl},
	0x16: {"asl", amZeX, 2, 6, asl},
	0x0e: {"asl", amAbs, 3, 6, asl},
	0x1e: {"asl", amAbX, 3, 7, asl},

	0x4a: {"lsr", amAcc, 1, 2, lsr},
	0x46: {"lsr", amZeP, 2, 5, lsr},
	0x56: {"lsr", amZeX, 2, 6, lsr},
	0x4e: {"lsr", amAbs, 3, 6, lsr},
	0x5e: {"lsr", amAbX, 3, 7, lsr},

	0x2a: {"rol", amAcc, 1, 2, rol},
	0x26: {"rol", amZeP, 2, 5, rol},
	0x36: {"rol", amZeX, 2, 6, rol},
	0x2e: {"rol", amAbs, 3, 6, rol},
	0x3e: {"rol", amAbX, 3, 7, rol},

	0x6a: {"ror", amAcc, 1, 2, ror},
	0x66: {"ror", amZeP, 2, 5, ror},
	0x76: {"ror", amZeX, 2, 6, ror},
	0x6e: {"ror", amAbs, 3, 6, ror},
	0x7e: {"ror", amAbX, 3, 7, ror},

	0x29: {"and", amImm, 2, 2, and},
	0x25: {"and", amZeP, 2, 3, and},
	0x35: {"and", amZeX, 2, 4, and},
	0x2d: {"and", amAbs, 3, 4, and},
	0x3d: {"and", amAbX, 3, 4, and},
	0x39: {"and", amAbY, 3, 4, and},
	0x21: {"and", amInX, 2, 6, and},
	0x31: {"and", amInY, 2, 5, and},

	0x09: {"ora", amImm, 2, 2, ora},
	0x05: {"ora", amZeP, 2, 3, ora},
	0x15: {"ora", amZeX, 2, 4, ora},
	0x0d: {"ora", amAbs, 3, 4, ora},
	0x1d: {"ora", amAbX, 3, 4, ora},
	0x19: {"ora", amAbY, 3, 4, ora},
	0x01: {"ora", amInX, 2, 6, ora},
	0x11: {"ora", amInY, 2, 5, ora},

	0x49: {"eor", amImm, 2, 2, eor},
	0x45: {"eor", amZeP, 2, 3, eor},
	0x55: {"eor", amZeX, 2, 4, eor},
	0x4d: {"eor", amAbs, 3, 4, eor},
	0x5d: {"eor", amAbX, 3, 4, eor},
	0x59: {"eor", amAbY, 3, 4, eor},
	0x41: {"eor", amInX, 2, 6, eor},
	0x51: {"eor", amInY, 2, 5, eor},

	0x24: {"bit", amZeP, 2, 3, bit},
	0x2c: {"bit", amAbs, 3, 4, bit},

	0x69: {"adc", amImm, 2, 2, adc},
	0x65: {"adc", amZeP, 2, 3, adc},
	0x75: {"adc", amZeX, 2, 4, adc},
	0x6d: {"adc", amAbs, 3, 4, adc},
	0x7d: {"adc", amAbX, 3, 4, adc},
	0x79: {"adc", amAbY, 3, 4, adc},
	0x61: {"adc", amInX, 2, 6, adc},
	0x71: {"adc", amInY, 2, 5, adc},

	0xe9: {"sbc", amImm, 2, 2, sbc},
	0xe5: {"sbc", amZeP, 2, 3, sbc},
	0xf5: {"sbc", amZeX, 2, 4, sbc},
	0xed: {"sbc", amAbs, 3, 4, sbc},
	0xfd: {"sbc", amAbX, 3, 4, sbc},
	0xf9: {"sbc", amAbY, 3, 4, sbc},
	0xe1: {"sbc", amInX, 2, 6, sbc},
	0xf1: {"sbc", amInY, 2, 5, sbc},
}

func opcode2op(opcode uint8) *Op {
	op := ops[opcode]

	if op == nil {
		panic(fmt.Sprintf("Unknown opcode 0x%02x", opcode))
	}

	return op
}
