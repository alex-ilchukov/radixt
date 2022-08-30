package struct8

import "github.com/alex-ilchukov/radixt"

type tree struct {
	sChunkPos        byte
	lsValue          byte
	rsValue          byte
	lsChildrenStart  byte
	rsChildrenStart  byte
	lsChildrenAmount byte
	rsChildrenAmount byte
	sChunkLen        byte
	chunks           string
	nodes            []node
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

	v = t.nodes[n].body(t.lsValue, t.rsValue)
	if v == 0 {
		return
	}

	v -= 1
	has = true

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t *tree) Chunk(n uint) string {
	if n >= t.Size() {
		return ""
	}

	node := t.nodes[n]
	l := node.tail(t.sChunkLen)
	pos := node.head(t.sChunkPos)
	return t.chunks[pos : pos+l]
}

// ChildrenRange returns first and last indices of children of node n, if the
// tree has the node and the node has children, or 1 and 0 otherwise.
func (t *tree) ChildrenRange(n uint) (uint, uint) {
	if n >= t.Size() {
		return 1, 0
	}

	node := t.nodes[n]
	amount := node.body(t.lsChildrenAmount, t.rsChildrenAmount)
	if amount == 0 {
		return 1, 0
	}

	f := node.body(t.lsChildrenStart, t.rsChildrenStart)
	return f, f + amount - 1
}

var _ radixt.Tree = (*tree)(nil)
