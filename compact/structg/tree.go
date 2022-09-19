package structg

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
	"github.com/alex-ilchukov/radixt/compact/internal/node"
	"github.com/alex-ilchukov/radixt/lookup"
)

type tree[N node.N] struct {
	h      header.A8b
	chunks string
	nodes  []N
}

// Size returns amount of nodes in the tree.
func (t *tree[_]) Size() uint {
	return uint(len(t.nodes))
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t *tree[_]) Value(n uint) (v uint, has bool) {
	if n < t.Size() {
		v, has = header.Value(t.nodes[n], t.h)
	}

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t *tree[_]) Chunk(n uint) (c string) {
	if n < t.Size() {
		l, h := header.ChunkRange(t.nodes[n], t.h)
		c = t.chunks[l:h]
	}

	return
}

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t *tree[_]) EachChild(n uint, e func(uint) bool) {
	if n >= t.Size() {
		return
	}

	for c, h := header.ChildrenRange(n, t.nodes[n], t.h); c < h; c++ {
		if e(c) {
			return
		}
	}
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (t *tree[N]) Hoard() (amount, hint uint) {
	amount = header.Len +
		16 + // tree.chunks
		24 + // tree.nodes
		uint(len(t.chunks)) +
		uint(cap(t.nodes))*(uint(node.BitsLen[N]())/8)

	hint = radixt.HoardExactly

	return
}

// Switch takes node n and byte b. If the node belongs to the tree, it looks
// for a child c of the node with such a chunk, that its first byte coincides
// with b. If such a child is found, it returns the child with its chunk
// without first byte and boolean truth. Otherwise or if the node is not in the
// tree, it returns corresponding default values.
func (t *tree[_]) Switch(n uint, b byte) (c uint, chunk string, found bool) {
	if n >= t.Size() {
		return
	}

	for c, h := header.ChildrenRange(n, t.nodes[n], t.h); c < h; c++ {
		child := t.nodes[c]
		low := header.ChunkLow(child, t.h)
		if t.chunks[low] == b {
			high := low + header.ChunkLen(child, t.h)
			return c, t.chunks[low+1 : high], true
		}
	}

	return
}

var (
	_ radixt.Tree    = (*tree[uint32])(nil)
	_ radixt.Hoarder = (*tree[uint32])(nil)
	_ lookup.Switcher = (*tree[uint32])(nil)
	_ radixt.Tree    = (*tree[uint64])(nil)
	_ radixt.Hoarder = (*tree[uint64])(nil)
	_ lookup.Switcher = (*tree[uint64])(nil)
)
