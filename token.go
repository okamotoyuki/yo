package main

import (
	"strconv"
)

const (
	tkNum = 256
	tkEOF = -1
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
		case '+', '-', '*', '/', '(', ')':
			token := Token{int(r), int(r), string(r)}
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

	token := Token{tkEOF, 0, ""}
	tokens = append(tokens, &token)
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
