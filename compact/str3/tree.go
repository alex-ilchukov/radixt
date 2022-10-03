package str3

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/lookup"
)

// Tree is radix tree implementation, which support 3-bytes nodes.
type Tree string

const (
	sChunkPos = iota
	lsValue
	rsValue
	lsChildrenStart
	rsChildrenStart
	lsChildrenAmount
	rsChildrenAmount
	sChunkLen
	h21
	h22
	cfstart

	maskSize = 0x7F_FF
	maskEmptyRoot = 0x80_00
)

// ProperLen is minimum length of proper empty tree. Any string with less
// length is _improper_, but is still considered _valid empty_ tree.
const ProperLen = cfstart

// Size returns amount of nodes in the tree.
func (t Tree) Size() (result uint) {
	if !t.empty() {
		result = t.h2() & maskSize
	}

	return
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t Tree) Value(n uint) (v uint, has bool) {
	size := t.Size()
	if n >= size {
		return
	}

	v = body(t.node(n, size), t[lsValue], t[rsValue])
	if v == 0 {
		return
	}

	has = true
	v--

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t Tree) Chunk(n uint) (c string) {
	size := t.Size()
	if n >= size || (n == 0 && t.emptyRoot()) {
		return
	}

	no := t.node(n, size)
	l := t.chunkPos(no)
	h := l + t.chunkLen(no)
	chunks := string(t[cfstart+(nodeLen+1)*size:])
	c = string(append([]byte{t[cfstart+n]}, chunks[l:h]...))

	return
}

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t Tree) EachChild(n uint, e func(uint) bool) {
	size := t.Size()
	if n >= size {
		return
	}

	no := t.node(n, size)
	ca := t.childrenAmount(no)
	if ca == 0 {
		return
	}

	c := t.childrenStart(n, no)
	h := c + ca
	for ; c < h; c++ {
		if e(c) {
			return
		}
	}
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (t Tree) Hoard() (uint, uint) {
	return uint(len(t)), radixt.HoardExactly
}

// Switch takes node n and byte b. If the node belongs to the tree, it looks
// for a child c of the node with such a chunk, that its first byte coincides
// with b. If such a child is found, it returns the child with its chunk
// without first byte and boolean truth. Otherwise or if the node is not in the
// tree, it returns corresponding default values.
func (t Tree) Switch(n uint, b byte) (c uint, chunk string, found bool) {
	size := t.Size()
	if n >= size {
		return
	}

	no := t.node(n, size)
	ca := t.childrenAmount(no)
	if ca == 0 {
		return
	}

	l := t.childrenStart(n, no)
	h := l + ca
	cf := string(t[cfstart:])
	for l < h {
		m := (l + h) >> 1
		b1 := cf[m]
		switch {
		case b1 == b:
			child := t.node(m, size)
			low := t.chunkPos(child)
			high := low + t.chunkLen(child)
			return m, cf[(nodeLen+1)*size:][low:high], true

		case b1 > b:
			h = m
		default:
			l = m + 1
		}
	}

	return
}

func (t Tree) empty() bool {
	return len(t) < ProperLen
}

func (t Tree) h2() uint {
	return uint(t[h21]) | uint(t[h22])<<8
}

func (t Tree) emptyRoot() bool {
	return t.h2() & maskEmptyRoot > 0
}

func (t Tree) node(n, size uint) node {
	i := cfstart + size + n * nodeLen
	return node(t[i]) | node(t[i+1])<<8 | node(t[i+2])<<16
}

func (t Tree) chunkPos(n node) uint {
	return head(n, t[sChunkPos])
}

func (t Tree) chunkLen(n node) uint {
	return tail(n, t[sChunkLen])
}

func (t Tree) childrenAmount(n node) uint {
	return body(n, t[lsChildrenAmount], t[rsChildrenAmount])
}

func (t Tree) childrenStart(n uint, no node) uint {
	return body(no, t[lsChildrenStart], t[rsChildrenStart]) + n + 1
}

var (
	_ radixt.Tree     = Tree("")
	_ radixt.Hoarder  = Tree("")
	_ lookup.Switcher = Tree("")
)
