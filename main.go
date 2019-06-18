package main

import (
	"os"
)

const isDebug = true

func main() {
	tokens := tokenize(os.Args[1])

	printAsmHeader()

	for i := 0; tokens[i].ty != tkEnd; i++ {
		if i == 0 {
			if tokens[i].ty != tkNum {
				exitWithError("the first token should be \"tkNum\".")
			}
			printAsmLine("MOVQ", "$%d, AX", tokens[i].val)
			continue
		}

		switch tokens[i].ty {
		case tkAdd:
			i++
			printAsmLine("ADDQ", "$%d, AX", tokens[i].val)
			continue
		case tkSub:
			i++
			printAsmLine("SUBQ", "$%d, AX", tokens[i].val)
			continue
		default:
			exitWithError("unexpected token in this context. (token.ty: %d)", tokens[i].ty)
		}
	}

	printAsmLine("MOVQ", "AX, ret+0(FP)")
	printAsmLine("RET", "")
}
