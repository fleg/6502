package cpu

import "fmt"

var ops = [256]*Op{
	// opcode: name addr mode size ticks do pageCrossTick
	0xea: {"nop", amImp, 1, 2, nop, 0},

	0xa9: {"lda", amImm, 2, 2, lda, 0},
	0xa5: {"lda", amZeP, 2, 3, lda, 0},
	0xb5: {"lda", amZeX, 2, 4, lda, 0},
	0xad: {"lda", amAbs, 3, 4, lda, 0},
	0xbd: {"lda", amAbX, 3, 4, lda, 1},
	0xb9: {"lda", amAbY, 3, 4, lda, 1},
	0xa1: {"lda", amInX, 2, 6, lda, 0},
	0xb1: {"lda", amInY, 2, 5, lda, 1},

	0xa2: {"ldx", amImm, 2, 2, ldx, 0},
	0xa6: {"ldx", amZeP, 2, 3, ldx, 0},
	0xb6: {"ldx", amZeY, 2, 4, ldx, 0},
	0xae: {"ldx", amAbs, 3, 4, ldx, 0},
	0xbe: {"ldx", amAbY, 3, 4, ldx, 1},

	0xa0: {"ldy", amImm, 2, 2, ldy, 0},
	0xa4: {"ldy", amZeP, 2, 3, ldy, 0},
	0xb4: {"ldy", amZeX, 2, 4, ldy, 0},
	0xac: {"ldy", amAbs, 3, 4, ldy, 0},
	0xbc: {"ldy", amAbX, 3, 4, ldy, 1},

	0x85: {"sta", amZeP, 2, 3, sta, 0},
	0x95: {"sta", amZeX, 2, 4, sta, 0},
	0x8d: {"sta", amAbs, 3, 4, sta, 0},
	0x9d: {"sta", amAbX, 3, 5, sta, 0},
	0x99: {"sta", amAbY, 3, 5, sta, 0},
	0x81: {"sta", amInX, 2, 6, sta, 0},
	0x91: {"sta", amInY, 2, 6, sta, 0},

	0x86: {"stx", amZeP, 2, 3, stx, 0},
	0x96: {"stx", amZeY, 2, 4, stx, 0},
	0x8e: {"stx", amAbs, 3, 4, stx, 0},

	0x84: {"sty", amZeP, 2, 3, sty, 0},
	0x94: {"sty", amZeX, 2, 4, sty, 0},
	0x8c: {"sty", amAbs, 3, 4, sty, 0},

	0xaa: {"tax", amImp, 1, 2, tax, 0},
	0xa8: {"tay", amImp, 1, 2, tay, 0},
	0xba: {"tsx", amImp, 1, 2, tsx, 0},
	0x8a: {"txa", amImp, 1, 2, txa, 0},
	0x9a: {"txs", amImp, 1, 2, txs, 0},
	0x98: {"tya", amImp, 1, 2, tya, 0},

	0x48: {"pha", amImp, 1, 3, pha, 0},
	0x68: {"pla", amImp, 1, 3, pla, 0},
	0x08: {"php", amImp, 1, 3, php, 0},
	0x28: {"plp", amImp, 1, 3, plp, 0},

	0x4c: {"jmp", amAbs, 3, 3, jmp, 0},
	0x6c: {"jmp", amInd, 3, 5, jmp, 0},

	0x20: {"jsr", amAbs, 3, 6, jsr, 0},
	0x60: {"rts", amImp, 1, 6, rts, 0},

	0x00: {"brk", amImp, 1, 7, brk, 0},
	0x40: {"rti", amImp, 1, 6, rti, 0},

	// handle extra ticks inside op
	0xb0: {"bcs", amRel, 2, 2, bcs, 0},
	0x90: {"bcc", amRel, 2, 2, bcc, 0},
	0xf0: {"beq", amRel, 2, 2, beq, 0},
	0xd0: {"bne", amRel, 2, 2, bne, 0},
	0x30: {"bmi", amRel, 2, 2, bmi, 0},
	0x10: {"bpl", amRel, 2, 2, bpl, 0},
	0x70: {"bvs", amRel, 2, 2, bvs, 0},
	0x50: {"bvc", amRel, 2, 2, bvc, 0},

	0x18: {"clc", amImp, 1, 2, clc, 0},
	0xd8: {"cld", amImp, 1, 2, cld, 0},
	0x58: {"cli", amImp, 1, 2, cli, 0},
	0xb8: {"clv", amImp, 1, 2, clv, 0},

	0x38: {"sec", amImp, 1, 2, sec, 0},
	0xf8: {"sed", amImp, 1, 2, sed, 0},
	0x78: {"sei", amImp, 1, 2, sei, 0},

	0xc6: {"dec", amZeP, 2, 5, dec, 0},
	0xd6: {"dec", amZeX, 2, 6, dec, 0},
	0xce: {"dec", amAbs, 2, 6, dec, 0},
	0xde: {"dec", amAbX, 2, 7, dec, 0},

	0xca: {"dex", amImp, 1, 2, dex, 0},
	0x88: {"dey", amImp, 1, 2, dey, 0},

	0xe6: {"inc", amZeP, 2, 5, inc, 0},
	0xf6: {"inc", amZeX, 2, 6, inc, 0},
	0xee: {"inc", amAbs, 2, 6, inc, 0},
	0xfe: {"inc", amAbX, 2, 5, inc, 0},

	0xe8: {"inx", amImp, 1, 2, inx, 0},
	0xc8: {"iny", amImp, 1, 2, iny, 0},

	0x0a: {"asl", amAcc, 1, 2, asl, 0},
	0x06: {"asl", amZeP, 2, 5, asl, 0},
	0x16: {"asl", amZeX, 2, 6, asl, 0},
	0x0e: {"asl", amAbs, 3, 6, asl, 0},
	0x1e: {"asl", amAbX, 3, 7, asl, 0},

	0x4a: {"lsr", amAcc, 1, 2, lsr, 0},
	0x46: {"lsr", amZeP, 2, 5, lsr, 0},
	0x56: {"lsr", amZeX, 2, 6, lsr, 0},
	0x4e: {"lsr", amAbs, 3, 6, lsr, 0},
	0x5e: {"lsr", amAbX, 3, 7, lsr, 0},

	0x2a: {"rol", amAcc, 1, 2, rol, 0},
	0x26: {"rol", amZeP, 2, 5, rol, 0},
	0x36: {"rol", amZeX, 2, 6, rol, 0},
	0x2e: {"rol", amAbs, 3, 6, rol, 0},
	0x3e: {"rol", amAbX, 3, 7, rol, 0},

	0x6a: {"ror", amAcc, 1, 2, ror, 0},
	0x66: {"ror", amZeP, 2, 5, ror, 0},
	0x76: {"ror", amZeX, 2, 6, ror, 0},
	0x6e: {"ror", amAbs, 3, 6, ror, 0},
	0x7e: {"ror", amAbX, 3, 7, ror, 0},

	0x29: {"and", amImm, 2, 2, and, 0},
	0x25: {"and", amZeP, 2, 3, and, 0},
	0x35: {"and", amZeX, 2, 4, and, 0},
	0x2d: {"and", amAbs, 3, 4, and, 0},
	0x3d: {"and", amAbX, 3, 4, and, 1},
	0x39: {"and", amAbY, 3, 4, and, 1},
	0x21: {"and", amInX, 2, 6, and, 0},
	0x31: {"and", amInY, 2, 5, and, 1},

	0x09: {"ora", amImm, 2, 2, ora, 0},
	0x05: {"ora", amZeP, 2, 3, ora, 0},
	0x15: {"ora", amZeX, 2, 4, ora, 0},
	0x0d: {"ora", amAbs, 3, 4, ora, 0},
	0x1d: {"ora", amAbX, 3, 4, ora, 1},
	0x19: {"ora", amAbY, 3, 4, ora, 1},
	0x01: {"ora", amInX, 2, 6, ora, 0},
	0x11: {"ora", amInY, 2, 5, ora, 1},

	0x49: {"eor", amImm, 2, 2, eor, 0},
	0x45: {"eor", amZeP, 2, 3, eor, 0},
	0x55: {"eor", amZeX, 2, 4, eor, 0},
	0x4d: {"eor", amAbs, 3, 4, eor, 0},
	0x5d: {"eor", amAbX, 3, 4, eor, 1},
	0x59: {"eor", amAbY, 3, 4, eor, 1},
	0x41: {"eor", amInX, 2, 6, eor, 0},
	0x51: {"eor", amInY, 2, 5, eor, 1},

	0x24: {"bit", amZeP, 2, 3, bit, 0},
	0x2c: {"bit", amAbs, 3, 4, bit, 0},

	0x69: {"adc", amImm, 2, 2, adc, 0},
	0x65: {"adc", amZeP, 2, 3, adc, 0},
	0x75: {"adc", amZeX, 2, 4, adc, 0},
	0x6d: {"adc", amAbs, 3, 4, adc, 0},
	0x7d: {"adc", amAbX, 3, 4, adc, 1},
	0x79: {"adc", amAbY, 3, 4, adc, 1},
	0x61: {"adc", amInX, 2, 6, adc, 0},
	0x71: {"adc", amInY, 2, 5, adc, 1},

	0xe9: {"sbc", amImm, 2, 2, sbc, 0},
	0xe5: {"sbc", amZeP, 2, 3, sbc, 0},
	0xf5: {"sbc", amZeX, 2, 4, sbc, 0},
	0xed: {"sbc", amAbs, 3, 4, sbc, 0},
	0xfd: {"sbc", amAbX, 3, 4, sbc, 1},
	0xf9: {"sbc", amAbY, 3, 4, sbc, 1},
	0xe1: {"sbc", amInX, 2, 6, sbc, 0},
	0xf1: {"sbc", amInY, 2, 5, sbc, 1},

	0xc9: {"cmp", amImm, 2, 2, cmp, 0},
	0xc5: {"cmp", amZeP, 2, 3, cmp, 0},
	0xd5: {"cmp", amZeX, 2, 4, cmp, 0},
	0xcd: {"cmp", amAbs, 3, 4, cmp, 0},
	0xdd: {"cmp", amAbX, 3, 4, cmp, 1},
	0xd9: {"cmp", amAbY, 3, 4, cmp, 1},
	0xc1: {"cmp", amInX, 2, 6, cmp, 0},
	0xd1: {"cmp", amInY, 2, 5, cmp, 1},

	0xe0: {"cpx", amImm, 2, 2, cpx, 0},
	0xe4: {"cpx", amZeP, 2, 3, cpx, 0},
	0xec: {"cpx", amAbs, 3, 4, cpx, 0},

	0xc0: {"cpy", amImm, 2, 2, cpy, 0},
	0xc4: {"cpy", amZeP, 2, 3, cpy, 0},
	0xcc: {"cpy", amAbs, 3, 4, cpy, 0},

	// undocumented, 0
	0x1a: {"nop", amImp, 1, 2, nop, 0},
	0x3a: {"nop", amImp, 1, 2, nop, 0},
	0x5a: {"nop", amImp, 1, 2, nop, 0},
	0x7a: {"nop", amImp, 1, 2, nop, 0},
	0xda: {"nop", amImp, 1, 2, nop, 0},
	0xfa: {"nop", amImp, 1, 2, nop, 0},
	0x80: {"nop", amImm, 2, 2, nop, 0},
	0x82: {"nop", amImm, 2, 2, nop, 0},
	0x89: {"nop", amImm, 2, 2, nop, 0},
	0xc2: {"nop", amImm, 2, 2, nop, 0},
	0xe2: {"nop", amImm, 2, 2, nop, 0},
	0x0c: {"nop", amAbs, 3, 4, nop, 0},
	0x1c: {"nop", amAbX, 3, 4, nop, 1},
	0x3c: {"nop", amAbX, 3, 4, nop, 1},
	0x5c: {"nop", amAbX, 3, 4, nop, 1},
	0x7c: {"nop", amAbX, 3, 4, nop, 1},
	0xdc: {"nop", amAbX, 3, 4, nop, 1},
	0xfc: {"nop", amAbX, 3, 4, nop, 1},
	0x04: {"nop", amZeP, 2, 3, nop, 0},
	0x44: {"nop", amZeP, 2, 3, nop, 0},
	0x64: {"nop", amZeP, 2, 3, nop, 0},
	0x14: {"nop", amZeX, 2, 4, nop, 0},
	0x34: {"nop", amZeX, 2, 4, nop, 0},
	0x54: {"nop", amZeX, 2, 4, nop, 0},
	0x74: {"nop", amZeX, 2, 4, nop, 0},
	0xd4: {"nop", amZeX, 2, 4, nop, 0},
	0xf4: {"nop", amZeX, 2, 4, nop, 0},

	0x0f: {"slo", amAbs, 3, 6, slo, 0},
	0x1f: {"slo", amAbX, 3, 7, slo, 0},
	0x1b: {"slo", amAbY, 3, 7, slo, 0},
	0x07: {"slo", amZeP, 2, 5, slo, 0},
	0x17: {"slo", amZeX, 2, 6, slo, 0},
	0x03: {"slo", amInX, 2, 8, slo, 0},
	0x13: {"slo", amInY, 2, 8, slo, 0},

	0x2f: {"rla", amAbs, 3, 6, rla, 0},
	0x3f: {"rla", amAbX, 3, 7, rla, 0},
	0x3b: {"rla", amAbY, 3, 7, rla, 0},
	0x27: {"rla", amZeP, 2, 5, rla, 0},
	0x37: {"rla", amZeX, 2, 6, rla, 0},
	0x23: {"rla", amInX, 2, 8, rla, 0},
	0x33: {"rla", amInY, 2, 8, rla, 0},

	0x4f: {"sre", amAbs, 3, 6, sre, 0},
	0x5f: {"sre", amAbX, 3, 7, sre, 0},
	0x5b: {"sre", amAbY, 3, 7, sre, 0},
	0x47: {"sre", amZeP, 2, 5, sre, 0},
	0x57: {"sre", amZeX, 2, 6, sre, 0},
	0x43: {"sre", amInX, 2, 8, sre, 0},
	0x53: {"sre", amInY, 2, 8, sre, 0},

	0x6f: {"rra", amAbs, 3, 6, rra, 0},
	0x7f: {"rra", amAbX, 3, 7, rra, 0},
	0x7b: {"rra", amAbY, 3, 7, rra, 0},
	0x67: {"rra", amZeP, 2, 5, rra, 0},
	0x77: {"rra", amZeX, 2, 6, rra, 0},
	0x63: {"rra", amInX, 2, 8, rra, 0},
	0x73: {"rra", amInY, 2, 8, rra, 0},

	0xcf: {"dcp", amAbs, 3, 6, dcp, 0},
	0xdf: {"dcp", amAbX, 3, 7, dcp, 0},
	0xdb: {"dcp", amAbY, 3, 7, dcp, 0},
	0xc7: {"dcp", amZeP, 2, 5, dcp, 0},
	0xd7: {"dcp", amZeX, 2, 6, dcp, 0},
	0xc3: {"dcp", amInX, 2, 8, dcp, 0},
	0xd3: {"dcp", amInY, 2, 8, dcp, 0},

	0xef: {"isc", amAbs, 3, 6, isc, 0},
	0xff: {"isc", amAbX, 3, 7, isc, 0},
	0xfb: {"isc", amAbY, 3, 7, isc, 0},
	0xe7: {"isc", amZeP, 2, 5, isc, 0},
	0xf7: {"isc", amZeX, 2, 6, isc, 0},
	0xe3: {"isc", amInX, 2, 8, isc, 0},
	0xf3: {"isc", amInY, 2, 8, isc, 0},

	0xaf: {"lax", amAbs, 2, 4, lax, 0},
	0xbf: {"lax", amAbY, 3, 4, lax, 1},
	0xa7: {"lax", amZeP, 2, 3, lax, 0},
	0xb7: {"lax", amZeY, 2, 4, lax, 0},
	0xa3: {"lax", amInX, 2, 6, lax, 0},
	0xb3: {"lax", amInY, 2, 5, lax, 1},

	0x8f: {"sax", amAbs, 3, 4, sax, 0},
	0x87: {"sax", amZeP, 2, 3, sax, 0},
	0x97: {"sax", amZeY, 2, 4, sax, 0},
	0x83: {"sax", amInX, 2, 6, sax, 0},

	0x0b: {"anc", amImm, 2, 2, anc, 0},
	0x2b: {"anc", amImm, 2, 2, anc, 0},
	0x4b: {"asr", amImm, 2, 2, asr, 0},
	0x6b: {"arr", amImm, 2, 2, arr, 0},
	0x8b: {"xaa", amImm, 2, 2, xaa, 0},
	0xab: {"lxa", amImm, 2, 2, lxa, 0},
	0xbb: {"las", amAbY, 3, 4, las, 1},
	0xcb: {"sbx", amImm, 2, 2, sbx, 0},
	0xeb: {"sbc", amImm, 2, 2, sbc, 0},

	0x9f: {"sha", amAbY, 3, 5, sha, 0},
	0x93: {"sha", amInY, 2, 6, sha, 0},
	0x9e: {"shx", amAbY, 3, 5, shx, 0},
	0x9c: {"shy", amAbX, 3, 5, shy, 0},
	0x9b: {"shs", amAbY, 3, 5, shs, 0},
}

func opcode2op(opcode uint8) *Op {
	op := ops[opcode]

	if op == nil {
		panic(fmt.Sprintf("Unknown opcode 0x%02x", opcode))
	}

	return op
}
