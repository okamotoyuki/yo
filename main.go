package main

import (
	"os"
)

const isDebug = true

func main() {
	tokens := tokenize(os.Args[1])
	debugPrintTokens(tokens)

	ast, _ := expr(tokens, 0)
	debugPrintAst(ast)

	generateCode(ast)
}
