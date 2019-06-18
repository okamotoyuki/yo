package main

import (
	"fmt"
	"os"
	"strings"
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

	debug("================")
	println()
}

func debugPrintAst(ast *Node) {
	debug("===== ast =====")
	debugPrintNode(ast, 0)
	debug("===============")
	println()
}

func debugPrintNode(node *Node, depth int) {
	if node == nil {
		return
	}

	format := strings.Repeat("\t", depth)
	format += "%d"
	debug(format, node.ty)

	debugPrintNode(node.lhs, depth+1)
	debugPrintNode(node.rhs, depth+1)
}
