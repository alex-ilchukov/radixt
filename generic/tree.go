package generic

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/lookup"
)

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

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t *tree) EachChild(n uint, e func(uint) bool) {
	if n >= t.Size() {
		return
	}

	node := t.nodes[n]
	low := node.cFirst
	high := low + uint(node.cAmount)
	for c := low; c < high; c++ {
		if e(c) {
			return
		}
	}
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (t *tree) Hoard() (uint, uint) {
	amount := uint(24) + // tree
		// node.cAmount gets aligned to 8 bytes
		uint(cap(t.nodes))*(8+8+16+8)

	for _, n := range t.nodes {
		amount += uint(len(n.chunk))
	}

	return amount, radixt.HoardExactly
}

// Switch takes node n and byte b. If the node belongs to the tree, it looks
// for a child c of the node with such a chunk, that its first byte coincides
// with b. If such a child is found, it returns the child with its chunk
// without first byte and boolean truth. Otherwise or if the node is not in the
// tree, it returns corresponding default values.
func (t *tree) Switch(n uint, b byte) (c uint, chunk string, found bool) {
	if n >= t.Size() {
		return
	}

	node := t.nodes[n]
	low := node.cFirst
	high := low + uint(node.cAmount)
	for c := low; c < high; c++ {
		chunk := t.nodes[c].chunk
		if chunk[0] == b {
			return c, chunk[1:], true
		}
	}

	return
}

var (
	_ radixt.Tree     = (*tree)(nil)
	_ radixt.Hoarder  = (*tree)(nil)
	_ lookup.Switcher = (*tree)(nil)
)
