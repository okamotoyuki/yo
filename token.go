package main

import (
	"strconv"
)

const (
	tkNum = 256 + iota
	tkEq
	tkNe
	tkLe
	tkGe
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

	var pos int
	var r rune

	// check if token buffer is empty
	bufferIsEmpty := func() bool {
		return start < 0
	}

	// flush token buffer
	flushBuffer := func() string {
		str := source[start:pos]
		start = -1
		return str
	}

	for pos = 0; pos < len(source); pos++ {
		r = rune(source[pos])

		switch r {
		case ' ':
			if !bufferIsEmpty() {
				input := flushBuffer()
				val, _ := strconv.Atoi(input)
				token := Token{tkNum, val, input}
				tokens = append(tokens, &token)
			}
			continue
		case '+', '-', '*', '/', '(', ')':
			if !bufferIsEmpty() {
				input := flushBuffer()
				val, _ := strconv.Atoi(input)
				token := Token{tkNum, val, input}
				tokens = append(tokens, &token)
			}
			token := Token{int(r), int(r), string(r)}
			tokens = append(tokens, &token)
			start = -1
		case '=', '!':
			if !bufferIsEmpty() {
				input := flushBuffer()
				val, _ := strconv.Atoi(input)
				token := Token{tkNum, val, input}
				tokens = append(tokens, &token)
			}
			if source[pos+1] == '=' {
				var token Token
				if r == '=' {
					token = Token{tkEq, 0, "=="}
				} else {

					token = Token{tkNe, 0, "!="}
				}
				tokens = append(tokens, &token)
				start = -1
				pos++
				continue
			}
			exitWithError("unexpected character in this context. ('%s', %d)", string(source[pos+1]), pos+1)
		case '<', '>':
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
