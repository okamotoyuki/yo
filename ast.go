package main

type Node struct {
	ty   int
	name string
	lhs  *Node
	rhs  *Node
	val  int
}

const (
	nodeNum = iota
	nodeAdd
	nodeSub
	nodeMul
	nodeDiv
)

// term = "(" expr ")" | number
func term(tokens []*Token, pos int) (*Node, int) {
	if next := consume(tokens, pos, '('); next > pos {
		node, pos := expr(tokens, next)

		if next = consume(tokens, pos, ')'); next == pos {
			exitWithError("')' is expected.")
		}

		return node, next
	} else if next = consume(tokens, pos, tkNum); next > pos {
		return &Node{nodeNum, "number", nil, nil, tokens[pos].val}, next
	}

	exitWithError("unexpected token in this context. (token.ty: %d)", tokens[pos].ty)
	return nil, -1
}

// mul  = term ("*" term | "/" term)*
func mul(tokens []*Token, pos int) (*Node, int) {
	node, pos := term(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, '*'); next > pos {
			rhs, pos = term(tokens, next)
			node = &Node{nodeMul, "mul", node, rhs, 0}
		} else if next = consume(tokens, pos, '/'); next > pos {
			rhs, pos = term(tokens, next)
			node = &Node{nodeDiv, "div", node, rhs, 0}
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
		if next := consume(tokens, pos, '+'); next > pos {
			rhs, pos = mul(tokens, next)
			node = &Node{nodeAdd, "add", node, rhs, 0}
		} else if next := consume(tokens, pos, '-'); next > pos {
			rhs, pos = mul(tokens, next)
			node = &Node{nodeSub, "sub", node, rhs, 0}
		} else {
			break
		}
	}

	return node, pos
}
