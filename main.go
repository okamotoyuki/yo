package main

import (
	"fmt"
	"os"
	"strconv"
)

const isDebug = true

const (
	tkNum = iota
	tkAdd
	tkSub
	tkMul
	tkDiv
	tkEnd
)

type Token struct {
	ty    int
	val   int
	input string
}

func exitWithError(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}

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

func tokenize(source string) []*Token {
	tokens := []*Token{}
	start := -1 // start index of token buffer

	for pos, r := range source {
		switch r {
		case ' ':
			if start >= 0 {
				input := source[start:pos]
				val, _ := strconv.Atoi(input)
				token := Token{tkNum, val, input}
				tokens = append(tokens, &token)
			}
			start = -1
			continue
		case '+':
			token := Token{tkAdd, tkAdd, string(r)}
			tokens = append(tokens, &token)
			start = -1
		case '-':
			token := Token{tkSub, tkSub, string(r)}
			tokens = append(tokens, &token)
			start = -1
		case '*':
			token := Token{tkMul, tkMul, string(r)}
			tokens = append(tokens, &token)
			start = -1
		case '/':
			token := Token{tkDiv, tkDiv, string(r)}
			tokens = append(tokens, &token)
			start = -1
		default:
			if start < 0 {
				start = pos
				continue
			}
		}
	}

	// if any string is stored in the token buffer, create a token from that
	if start >= 0 {
		input := source[start:]
		val, _ := strconv.Atoi(input)
		token := Token{tkNum, val, input}
		tokens = append(tokens, &token)
		start = -1
	}

	token := Token{tkEnd, -1, ""}
	tokens = append(tokens, &token)

	debugPrintTokens(tokens)

	return tokens
}

func printAsmLine(inst string, format string, args ...interface{}) {
	fmt.Printf("\t"+inst+"\t"+format+"\n", args...)
}

func printAsmHeader() {
	fmt.Println("#include \"textflag.h\"")
	fmt.Println()
	fmt.Println("TEXT ·run(SB), NOSPLIT, $0")
}

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
