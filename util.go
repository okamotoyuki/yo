package main

import (
	"fmt"
	"os"
)

func debug(format string, args ...interface{}) {
	if isDebug {
		_, _ = fmt.Fprintf(os.Stderr, "debug: "+format+"\n", args...)
	}
}

func debugPrintTokens(tokens []*Token) {
	debug("==== tokens ====")

	for _, token := range tokens {
		switch token.ty {
		case tkNum:
			debug("token -> ty: tkNum, val: %d", token.val)
		case tkAdd:
			debug("token -> ty: tkAdd, val: %d", token.val)
		case tkSub:
			debug("token -> ty: tkSub, val: %d", token.val)
		case tkMul:
			debug("token -> ty: tkMul, val: %d", token.val)
		case tkDiv:
			debug("token -> ty: tkDiv, val: %d", token.val)
		case tkEnd:
			debug("token -> ty: tkEnd, val: EOF")
		default:
			debug("token -> unsupported token type. (token.ty: %d)", token.ty)
		}
	}

	debug("================\n")
}

func printAsmLine(inst string, format string, args ...interface{}) {
	fmt.Printf("\t"+inst+"\t"+format+"\n", args...)
}

func printAsmHeader() {
	fmt.Println("#include \"textflag.h\"")
	fmt.Println()
	fmt.Println("TEXT ·run(SB), NOSPLIT, $0")
}
