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
	nodeLt
	nodeLe
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
	lhs, pos := unary(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, '*'); next > pos {
			rhs, pos = unary(tokens, next)
			lhs = &Node{nodeMul, "mul", lhs, rhs, 0}
		} else if next = consume(tokens, pos, '/'); next > pos {
			rhs, pos = unary(tokens, next)
			lhs = &Node{nodeDiv, "div", lhs, rhs, 0}
		} else {
			break
		}
	}

	return lhs, pos
}

// add = mul ("+" mul | "-" mul)*
func add(tokens []*Token, pos int) (*Node, int) {
	lhs, pos := mul(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, '+'); next > pos {
			rhs, pos = mul(tokens, next)
			lhs = &Node{nodeAdd, "add", lhs, rhs, 0}
		} else if next := consume(tokens, pos, '-'); next > pos {
			rhs, pos = mul(tokens, next)
			lhs = &Node{nodeSub, "sub", lhs, rhs, 0}
		} else {
			break
		}
	}

	return lhs, pos
}

// relational = add ("<" add | "<=" add | ">" add | ">=" add)*
func relational(tokens []*Token, pos int) (*Node, int) {
	lhs, pos := add(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, '<'); next > pos {
			rhs, pos = add(tokens, next)
			lhs = &Node{nodeLt, "lt", lhs, rhs, 0}
		} else if next := consume(tokens, pos, tkLe); next > pos {
			rhs, pos = add(tokens, next)
			lhs = &Node{nodeLe, "le", lhs, rhs, 0}
		} else if next := consume(tokens, pos, '>'); next > pos {
			rhs, pos = add(tokens, next)
			lhs = &Node{nodeLt, "lt", rhs, lhs, 0} // swap 'rhs' & 'lhs'
		} else if next := consume(tokens, pos, tkGe); next > pos {
			rhs, pos = add(tokens, next)
			lhs = &Node{nodeLe, "le", rhs, lhs, 0} // swap 'rhs' & 'lhs'
		} else {
			break
		}
	}

	return lhs, pos
}

// equality = relational ("==" relational | "!=" relational)*
func equality(tokens []*Token, pos int) (*Node, int) {
	lhs, pos := relational(tokens, pos)

	var rhs *Node

	for {
		if next := consume(tokens, pos, tkEq); next > pos {
			rhs, pos = relational(tokens, next)
			lhs = &Node{nodeEq, "eq", lhs, rhs, 0}
		} else if next := consume(tokens, pos, tkNe); next > pos {
			rhs, pos = relational(tokens, next)
			lhs = &Node{nodeNe, "ne", lhs, rhs, 0}
		} else {
			break
		}
	}

	return lhs, pos
}

// expr = equality
func expr(tokens []*Token, pos int) (*Node, int) {
	return equality(tokens, pos)
}
