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
		debug("type => \"%s\", val => %d", token.name, token.val)
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
	format += "%s"
	debug(format, node.name)

	debugPrintNode(node.lhs, depth+1)
	debugPrintNode(node.rhs, depth+1)
}
