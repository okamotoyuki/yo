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
