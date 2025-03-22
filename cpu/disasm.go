package cpu

import "fmt"

func Disasm(bin []uint8) string {
	code := ""
	size := uint(len(bin))
	addr := uint(0)

	for addr < size {
		opcode := bin[addr]
		code += fmt.Sprintf("%04x: %02x", addr, bin[addr])
		addr += 1

		op := opcode2op(opcode)
		arg := uint(0)

		switch op.Size {
		case 1:
			code += "    "
		case 2:
			arg = uint(bin[addr])
			code += fmt.Sprintf("%02x  ", bin[addr])
			addr += 1
		case 3:
			arg = uint(word(bin[addr], bin[addr+1]))
			code += fmt.Sprintf("%02x", bin[addr])
			code += fmt.Sprintf("%02x", bin[addr+1])
			addr += 2
		}

		code += "    " + op.Name

		if op.AddressMode != amImp {
			code += " "
		}

		switch op.AddressMode {
		case amAcc:
			code += "a"
		case amImm:
			code += fmt.Sprintf("#$%02x", arg)
		case amZeP:
			code += fmt.Sprintf("$%02x", arg)
		case amZeX:
			code += fmt.Sprintf("$%02x,x", arg)
		case amZeY:
			code += fmt.Sprintf("$%02x,y", arg)
		case amRel:
			code += fmt.Sprintf("*%+d", int8(arg)+2)
		case amAbs:
			code += fmt.Sprintf("$%04x", arg)
		case amAbX:
			code += fmt.Sprintf("$%04x,x", arg)
		case amAbY:
			code += fmt.Sprintf("$%04x,y", arg)
		case amInd:
			code += fmt.Sprintf("($%04x)", arg)
		case amInX:
			code += fmt.Sprintf("($%02x,x)", arg)
		case amInY:
			code += fmt.Sprintf("($%02x),y", arg)
		}

		code += "\n"
	}

	return code
}
