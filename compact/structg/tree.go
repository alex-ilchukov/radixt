package structg

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
	"github.com/alex-ilchukov/radixt/compact/internal/node"
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

// ChildrenRange returns low and high indices of children of node n, if the
// tree has the node and the node has children, or default unsigned integers
// otherwise.
func (t *tree[_]) ChildrenRange(n uint) (low, high uint) {
	if n < t.Size() {
		low, high = header.ChildrenRange(n, t.nodes[n], t.h)
	}

	return
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

var (
	_ radixt.Tree    = (*tree[uint32])(nil)
	_ radixt.Hoarder = (*tree[uint32])(nil)
	_ radixt.Tree    = (*tree[uint64])(nil)
	_ radixt.Hoarder = (*tree[uint64])(nil)
)
