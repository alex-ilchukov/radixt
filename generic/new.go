package generic

import "github.com/alex-ilchukov/radixt"

// New creates a new generic tree as a copy of the provided tree t and returns
// a pointer on the created tree. It returns empty tree, if t is nil.
func New(t radixt.Tree) *tree {
	if t == nil {
		return &tree{}
	}

	l := t.Size()
	if l == 0 {
		return &tree{}
	}

	type e struct {
		n uint
		p uint
	}

	nodes := make([]node, l, l)
	for i, q := uint(0), make([]e, 1, l); len(q) > 0; i++ {
		a := q[0]
		q = q[1:]

		n := a.n

		t.EachChild(n, func(c uint) bool {
			q = append(q, e{n: c, p: i})
			return false
		})

		v, has := t.Value(n)
		nodes[i] = node{hasValue: has, chunk: t.Chunk(n), value: v}

		if n == 0 {
			continue
		}

		p := a.p
		if nodes[p].cAmount == 0 {
			nodes[p].cFirst = i
		}
		nodes[p].cAmount = byte(i - nodes[p].cFirst + 1)
	}

	return &tree{nodes: nodes}
}
