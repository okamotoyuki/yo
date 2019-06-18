package main

type Node struct {
	ty  int
	lhs *Node
	rhs *Node
	val int
}

const (
	nodeNum = iota
	nodeAdd
	nodeSub
	nodeMul
	nodeDiv
)

// term = num
func term(tokens []*Token, pos int) (*Node, int) {
	if next := consume(tokens, pos, tkNum); next > pos {
		return &Node{nodeNum, nil, nil, tokens[pos].val}, next
	}

	exitWithError("unexpected token in this context. (token.ty: %d)", tokens[pos].ty)
	return nil, -1
}

// mul  = term ("*" term | "/" term)*
func mul(tokens []*Token, pos int) (*Node, int) {
	node, pos := term(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, tkMul); next > pos {
			rhs, pos = term(tokens, next)
			node = &Node{nodeMul, node, rhs, 0}
		} else if next := consume(tokens, pos, tkDiv); next > pos {
			rhs, pos = term(tokens, next)
			node = &Node{nodeDiv, node, rhs, 0}
		} else {
			break
		}
	}

	return node, pos
}

// expr = mul ("+" mul | "-" mul)*
func expr(tokens []*Token, pos int) (*Node, int) {
	node, pos := mul(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, tkAdd); next > pos {
			rhs, pos = term(tokens, next)
			node = &Node{nodeAdd, node, rhs, 0}
		} else if next := consume(tokens, pos, tkSub); next > pos {
			rhs, pos = term(tokens, next)
			node = &Node{nodeSub, node, rhs, 0}
		} else {
			break
		}
	}

	return node, pos
}
