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

// Has returns if the tree has node n or not.
func (t *tree) Has(n int) bool {
	return 0 <= n && n < len(t.nodes)
}

// Root returns -1 for empty tree and 0 otherwise.
func (t *tree) Root() int {
	if t.Size() > 0 {
		return 0
	}

	return -1
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t *tree) Value(n int) (v uint, has bool) {
	if !t.Has(n) {
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

// EachChild calls func e for every child of node n, if the tree has the node,
// until the func returns boolean true. The order of going over the children is
// fixed for every node, but may not coincide with any natural order.
func (t *tree) EachChild(n int, e func(int) bool) {
	if !t.Has(n) {
		return
	}

	node := t.nodes[n]
	amount := node.body(t.lsChildrenAmount, t.rsChildrenAmount)
	if amount == 0 {
		return
	}

	c := int(node.body(t.lsChildrenStart, t.rsChildrenStart))
	l := c + int(amount)
	for ; c < l; c++ {
		if e(c) {
			return
		}
	}
}

// ByteAt returns default byte value and boolean false, if npos is outside of
// chunk of the node n, or byte of the chunk at npos and boolean true
// otherwise.
func (t *tree) ByteAt(n int, npos uint) (b byte, within bool) {
	if !t.Has(n) {
		return
	}

	node := t.nodes[n]
	if node.tail(t.sChunkLen) <= npos {
		return
	}

	within = true
	pos := node.head(t.sChunkPos)
	b = t.chunks[pos+npos]

	return
}

var _ radixt.Tree = (*tree)(nil)
