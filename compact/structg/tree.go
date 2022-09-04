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

// ChildrenRange returns first and last indices of children of node n, if the
// tree has the node and the node has children, or 1 and 0 otherwise.
func (t *tree[_]) ChildrenRange(n uint) (uint, uint) {
	if n >= t.Size() {
		return 1, 0
	}

	return header.ChildrenRange(n, t.nodes[n], t.h)
}

var (
	_ radixt.Tree = (*tree[uint32])(nil)
	_ radixt.Tree = (*tree[uint64])(nil)
)
