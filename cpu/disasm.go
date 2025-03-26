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

		if op.AddressMode != AmImp {
			code += " "
		}

		switch op.AddressMode {
		case AmAcc:
			code += "a"
		case AmImm:
			code += fmt.Sprintf("#$%02x", arg)
		case AmZeP:
			code += fmt.Sprintf("$%02x", arg)
		case AmZeX:
			code += fmt.Sprintf("$%02x,x", arg)
		case AmZeY:
			code += fmt.Sprintf("$%02x,y", arg)
		case AmRel:
			code += fmt.Sprintf("*%+d", int8(arg)+2)
		case AmAbs:
			code += fmt.Sprintf("$%04x", arg)
		case AmAbX:
			code += fmt.Sprintf("$%04x,x", arg)
		case AmAbY:
			code += fmt.Sprintf("$%04x,y", arg)
		case AmInd:
			code += fmt.Sprintf("($%04x)", arg)
		case AmInX:
			code += fmt.Sprintf("($%02x,x)", arg)
		case AmInY:
			code += fmt.Sprintf("($%02x),y", arg)
		}

		code += "\n"
	}

	return code
}
