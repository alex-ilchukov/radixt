package generic

import "github.com/alex-ilchukov/radixt"

type node struct {
	hasValue bool
	cAmount  byte
	cFirst   uint
	chunk    string
	value    uint
}

type tree struct {
	nodes []node
}

// Size returns amount of nodes in the tree.
func (t *tree) Size() uint {
	return uint(len(t.nodes))
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t *tree) Value(n uint) (v uint, has bool) {
	if n < t.Size() {
		node := t.nodes[n]
		v = node.value
		has = node.hasValue
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

// ChildrenRange returns low and high indices of children of node n, if the
// tree has the node and the node has children, or default unsigned integers
// otherwise.
func (t *tree) ChildrenRange(n uint) (low, high uint) {
	if n < t.Size() {
		node := t.nodes[n]
		low = node.cFirst
		high = low + uint(node.cAmount)
	}

	return
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (t *tree) Hoard() (uint, uint) {
	amount := uint(32) + // tree
		// node.cAmount gets aligned to 8 bytes
		uint(cap(t.nodes))*(8+8+16+8)

	for _, n := range t.nodes {
		amount += uint(len(n.chunk))
	}

	return amount, radixt.HoardExactly
}

var (
	_ radixt.Tree    = (*tree)(nil)
	_ radixt.Hoarder = (*tree)(nil)
)
