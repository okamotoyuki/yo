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
	nodeEq
	nodeNe
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

	exitWithError("unexpected token in this context. (token.ty: '%s')", string(tokens[pos].ty))
	return nil, -1
}

// unary = ("+" | "-")? term
func unary(tokens []*Token, pos int) (*Node, int) {
	if next := consume(tokens, pos, '+'); next > pos {
		return term(tokens, next)
	} else if next = consume(tokens, pos, '-'); next > pos {
		lhs := &Node{nodeNum, "number", nil, nil, 0}
		rhs, pos := term(tokens, next)
		return &Node{nodeSub, "sub", lhs, rhs, 0}, pos
	}

	return term(tokens, pos)
}

// mul  = unary ("*" unary | "/" unary)*
func mul(tokens []*Token, pos int) (*Node, int) {
	node, pos := unary(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, '*'); next > pos {
			rhs, pos = unary(tokens, next)
			node = &Node{nodeMul, "mul", node, rhs, 0}
		} else if next = consume(tokens, pos, '/'); next > pos {
			rhs, pos = unary(tokens, next)
			node = &Node{nodeDiv, "div", node, rhs, 0}
		} else {
			break
		}
	}

	return node, pos
}

// add = mul ("+" mul | "-" mul)*
func add(tokens []*Token, pos int) (*Node, int) {
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

// relational = add ("<" add | "<=" add | ">" add | ">=" add)*
func relational(tokens []*Token, pos int) (*Node, int) {
	return add(tokens, pos)
}

// equality = relational ("==" relational | "!=" relational)*
func equality(tokens []*Token, pos int) (*Node, int) {
	node, pos := relational(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, tkEq); next > pos {
			rhs, pos = relational(tokens, next)
			node = &Node{nodeEq, "eq", node, rhs, 0}
		} else if next := consume(tokens, pos, tkNe); next > pos {
			rhs, pos = relational(tokens, next)
			node = &Node{nodeNe, "ne", node, rhs, 0}
		} else {
			break
		}
	}

	return node, pos
}

// expr = equality
func expr(tokens []*Token, pos int) (*Node, int) {
	return equality(tokens, pos)
}
