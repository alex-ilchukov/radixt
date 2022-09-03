package structg

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
	"github.com/alex-ilchukov/radixt/compact/internal/node"
)

type tree[N node.N] struct {
	h      header.H[N]
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
		v, has = t.h.Value(t.nodes[n])
	}

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t *tree[_]) Chunk(n uint) string {
	if n >= t.Size() {
		return ""
	}

	node := t.nodes[n]
	l := t.h.ChunkLen(node)
	pos := t.h.ChunkPos(node)
	return t.chunks[pos : pos+l]
}

// ChildrenRange returns first and last indices of children of node n, if the
// tree has the node and the node has children, or 1 and 0 otherwise.
func (t *tree[_]) ChildrenRange(n uint) (uint, uint) {
	if n >= t.Size() {
		return 1, 0
	}

	return t.h.ChildrenRange(n, t.nodes[n])
}

var (
	_ radixt.Tree = (*tree[uint32])(nil)
	_ radixt.Tree = (*tree[uint64])(nil)
)
