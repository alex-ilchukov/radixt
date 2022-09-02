package structg

import "github.com/alex-ilchukov/radixt"

type tree[N node] struct {
	sChunkPos        byte
	lsValue          byte
	rsValue          byte
	lsChildrenStart  byte
	rsChildrenStart  byte
	lsChildrenAmount byte
	rsChildrenAmount byte
	sChunkLen        byte
	chunks           string
	nodes            []N
}

// Size returns amount of nodes in the tree.
func (t *tree[_]) Size() uint {
	return uint(len(t.nodes))
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t *tree[_]) Value(n uint) (v uint, has bool) {
	if n >= t.Size() {
		return
	}

	v = body(t.nodes[n], t.lsValue, t.rsValue)
	if v == 0 {
		return
	}

	v -= 1
	has = true

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t *tree[_]) Chunk(n uint) string {
	if n >= t.Size() {
		return ""
	}

	node := t.nodes[n]
	l := tail(node, t.sChunkLen)
	pos := head(node, t.sChunkPos)
	return t.chunks[pos : pos+l]
}

// ChildrenRange returns first and last indices of children of node n, if the
// tree has the node and the node has children, or 1 and 0 otherwise.
func (t *tree[_]) ChildrenRange(n uint) (uint, uint) {
	if n >= t.Size() {
		return 1, 0
	}

	node := t.nodes[n]
	amount := body(node, t.lsChildrenAmount, t.rsChildrenAmount)
	if amount == 0 {
		return 1, 0
	}

	f := body(node, t.lsChildrenStart, t.rsChildrenStart) + n + 1
	return f, f + amount - 1
}

var (
	_ radixt.Tree = (*tree[uint32])(nil)
	_ radixt.Tree = (*tree[uint64])(nil)
)
