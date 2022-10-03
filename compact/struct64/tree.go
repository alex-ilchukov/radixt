package struct64

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/lookup"
)

type tree struct {
	emptyRoot bool
	s         shifts
	chunks    string
	cf        string
	nodes     []node
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

	v = body(t.nodes[n], t.s.lsValue, t.s.rsValue)
	if v == 0 {
		return
	}

	has = true
	v--

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t *tree) Chunk(n uint) (c string) {
	if n >= t.Size() || (n == 0 && t.emptyRoot) {
		return
	}

	l := t.chunkPos(n)
	h := l + t.chunkLen(n)
	c = string(append([]byte{t.cf[n]}, t.chunks[l:h]...))

	return
}

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t *tree) EachChild(n uint, e func(uint) bool) {
	if n >= t.Size() {
		return
	}

	ca := t.childrenAmount(n)
	if ca == 0 {
		return
	}

	c := t.childrenStart(n)
	h := c + ca
	for ; c < h; c++ {
		if e(c) {
			return
		}
	}
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (t *tree) Hoard() (amount, hint uint) {
	amount = 8 + // tree.emptyRoot aligned
		8 + //  tree.s
		16 + // tree.chunks
		16 + // tree.cf
		24 + // tree.nodes
		uint(len(t.chunks)) +
		uint(len(t.cf)) +
		uint(cap(t.nodes))*8

	hint = radixt.HoardExactly

	return
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

	ca := t.childrenAmount(n)
	if ca == 0 {
		return
	}

	l := t.childrenStart(n)
	h := l + ca
	for l < h {
		m := (l + h) >> 1
		b1 := t.cf[m]
		switch {
		case b1 == b:
			low := t.chunkPos(m)
			high := low + t.chunkLen(m)
			return m, t.chunks[low:high], true

		case b1 > b:
			h = m
		default:
			l = m + 1
		}
	}

	return
}

func (t *tree) chunkPos(n uint) uint {
	return head(t.nodes[n], t.s.sChunkPos)
}

func (t *tree) chunkLen(n uint) uint {
	return tail(t.nodes[n], t.s.sChunkLen)
}

func (t *tree) childrenAmount(n uint) uint {
	if n >= t.Size() {
		return 0
	}

	return body(t.nodes[n], t.s.lsChildrenAmount, t.s.rsChildrenAmount)
}

func (t *tree) childrenStart(n uint) uint {
	s := body(t.nodes[n], t.s.lsChildrenStart, t.s.rsChildrenStart) + n + 1
	return s
}

var (
	_ radixt.Tree     = (*tree)(nil)
	_ radixt.Hoarder  = (*tree)(nil)
	_ lookup.Switcher = (*tree)(nil)
)
