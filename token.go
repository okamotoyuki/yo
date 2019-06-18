package main

import (
	"strconv"
)

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
	name  string
	val   int
	input string
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
				token := Token{tkNum, "number", val, input}
				tokens = append(tokens, &token)
			}
			start = -1
			continue
		case '+':
			token := Token{tkAdd, "add", tkAdd, string(r)}
			tokens = append(tokens, &token)
			start = -1
		case '-':
			token := Token{tkSub, "sub", tkSub, string(r)}
			tokens = append(tokens, &token)
			start = -1
		case '*':
			token := Token{tkMul, "mul", tkMul, string(r)}
			tokens = append(tokens, &token)
			start = -1
		case '/':
			token := Token{tkDiv, "div", tkDiv, string(r)}
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
		token := Token{tkNum, "num", val, input}
		tokens = append(tokens, &token)
		start = -1
	}

	token := Token{tkEnd, "end", 0, ""}
	tokens = append(tokens, &token)

	debugPrintTokens(tokens)

	return tokens
}

// consume token if the input is expected one
func consume(tokens []*Token, pos int, ty int) int {
	if tokens[pos].ty != ty {
		return pos
	}
	pos++
	return pos
}
