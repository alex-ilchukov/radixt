package generic

import "github.com/alex-ilchukov/radixt"

type node struct {
	cAmount byte
	cFirst  uint
	chunk   string
	value   uint
}

type tree struct {
	noValue uint
	nodes   []node
}

// Size returns amount of nodes in the tree.
func (t *tree) Size() uint {
	return uint(len(t.nodes))
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t *tree) Value(n uint) (v uint, has bool) {
	if n >= t.Size() {
		return
	}

	v = t.nodes[n].value
	if v == t.noValue {
		v = 0
	} else {
		has = true
	}

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t *tree) Chunk(n uint) string {
	if n >= t.Size() {
		return ""
	}

	return t.nodes[n].chunk
}

// ChildrenRange returns first and last indices of children of node n, if the
// tree has the node and the node has children, or 1 and 0 otherwise.
func (t *tree) ChildrenRange(n uint) (uint, uint) {
	if n >= t.Size() {
		return 1, 0
	}

	no := t.nodes[n]
	f := no.cFirst
	return f, f + uint(no.cAmount) - 1
}

var _ radixt.Tree = (*tree)(nil)
