package main

import (
	"fmt"
)

func printAsmLine(inst string, format string, args ...interface{}) {
	if format == "" {
		fmt.Println("\t" + inst)
		return
	}
	fmt.Printf("\t"+inst+"\t"+format+"\n", args...)
}

func printAsmHeader() {
	fmt.Println("#include \"textflag.h\"")
	fmt.Println()
	fmt.Println("TEXT ·run(SB), NOSPLIT, $0")
}

func printAsmFooter() {
	printAsmLine("MOVQ", "AX, ret+0(FP)")
	printAsmLine("RET", "")
}

func printAsmBody(ast *Node) {
	visit(ast)

	if ast.ty == nodeNum {
		printAsmLine("POPQ", "AX")
	}
}

func generateCode(ast *Node) {
	printAsmHeader()
	printAsmBody(ast)
	printAsmFooter()
}

func visit(node *Node) {
	if node.ty == nodeNum {
		printAsmLine("PUSHQ", "$%d", node.val)
		return
	}

	visit(node.lhs)
	visit(node.rhs)

	printAsmLine("POPQ", "DI")
	printAsmLine("POPQ", "AX")

	switch node.ty {
	case nodeAdd:
		printAsmLine("ADDQ", "DI, AX")
	case nodeSub:
		printAsmLine("SUBQ", "DI, AX")
	case nodeMul:
		printAsmLine("MULQ", "DI")
	case nodeDiv:
		printAsmLine("CQO", "")
		printAsmLine("DIVQ", "DI")
	default:
		exitWithError("unexpected node in this context. (node.ty: %d)", node.ty)
	}
}
