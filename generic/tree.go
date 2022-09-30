package generic

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/lookup"
)

type node struct {
	chunkFirst byte
	chunkEmpty bool
	hasValue   bool
	cAmount    byte
	cFirst     uint
	chunkLow   uint
	chunkHigh  uint
	value      uint
}

type tree struct {
	c     string
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
func (t *tree) Chunk(n uint) (chunk string) {
	if n < t.Size() {
		l := t.nodes[n].chunkLow
		h := t.nodes[n].chunkHigh
		chunk = t.c[l:h]
	}

	return
}

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t *tree) EachChild(n uint, e func(uint) bool) {
	for c, high := t.childrenRange(n); c < high; c++ {
		if e(c) {
			return
		}
	}
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (t *tree) Hoard() (uint, uint) {
	amount := uint(40) + // tree
		uint(len(t.c)) +
		// node.cAmount with node.hasValue get aligned together to 8
		// bytes
		uint(len(t.nodes))*(8+8+8+8+8)

	return amount, radixt.HoardExactly
}

// Switch takes node n and byte b. If the node belongs to the tree, it looks
// for a child c of the node with such a chunk, that its first byte coincides
// with b. If such a child is found, it returns the child with its chunk
// without first byte and boolean truth. Otherwise or if the node is not in the
// tree, it returns corresponding default values.
func (t *tree) Switch(n uint, b byte) (c uint, chunk string, found bool) {
	for l, h := t.childrenRange(n); l < h; {
		m := l + (h-l)>>1
		b1 := t.nodes[m].chunkFirst
		switch {
		case b1 == b:
			cl := t.nodes[m].chunkLow + 1
			ch := t.nodes[m].chunkHigh
			return m, t.c[cl:ch], true
		case b1 > b:
			h = m
		default:
			l = m + 1
		}
	}

	return
}

func (t *tree) childrenRange(n uint) (low, high uint) {
	if n >= t.Size() {
		return
	}

	low = t.nodes[n].cFirst
	return low, low + uint(t.nodes[n].cAmount)
}

var (
	_ radixt.Tree     = (*tree)(nil)
	_ radixt.Hoarder  = (*tree)(nil)
	_ lookup.Switcher = (*tree)(nil)
)
