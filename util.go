package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func debug(format string, args ...interface{}) {
	if isDebug {
		_, _ = fmt.Fprintf(os.Stderr, "debug: "+format+"\n", args...)
	}
}

func debugPrintTokens(tokens []*Token) {
	if isDebug {
		str := "tokens => [ "

		for _, token := range tokens {
			switch token.ty {
			case tkNum:
				str += strconv.Itoa(token.val) + ", "
			case tkEOF:
				str += "EOF"
			default:
				str += string(token.ty) + ", "
			}
		}
		str += " ]"
		debug(str)
		println()
	}
}

func debugPrintAst(ast *Node) {
	if isDebug {
		debug("===== ast =====")
		debugPrintNode(ast, 0)
		debug("===============")
		println()
	}
}

func debugPrintNode(node *Node, depth int) {
	if node == nil {
		return
	}

	format := "\t" + strings.Repeat("\t", depth)
	format += node.name
	_, _ = fmt.Fprintln(os.Stderr, format)

	debugPrintNode(node.lhs, depth+1)
	debugPrintNode(node.rhs, depth+1)
}
