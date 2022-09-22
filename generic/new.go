package generic

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/pass"
)

type yielder struct {
	t     radixt.Tree
	nodes []node
}

func (y *yielder) Yield(i, n, tag uint) uint {
	v, has := y.t.Value(n)
	y.nodes[i] = node{hasValue: has, chunk: y.t.Chunk(n), value: v}
	y.processParent(i, tag)

	return i
}

func (y *yielder) processParent(i, p uint) {
	if i == 0 {
		return
	}

	nodes := y.nodes
	if nodes[p].cAmount == 0 {
		nodes[p].cFirst = i
	}
	nodes[p].cAmount = byte(i - nodes[p].cFirst + 1)
}

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

	y := &yielder{t: t, nodes: make([]node, l, l)}
	pass.Do(t, y)
	return &tree{nodes: y.nodes}
}
