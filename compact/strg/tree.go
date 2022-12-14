package strg

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
	"github.com/alex-ilchukov/radixt/lookup"
)

// Tree is radix tree implementation, which support 3-bytes nodes and 4-bytes
// nodes.
type Tree[_ N] string

const (
	hlen   = header.Len
	olen   = 2
	cstart = hlen + olen
)

// ProperLen is minimum length of proper empty tree. Any string with less
// length is _improper_, but is still considered _valid empty_ tree.
const ProperLen = cstart

// Size returns amount of nodes in the tree.
func (t Tree[N]) Size() (result uint) {
	if !t.empty() {
		result = uint((len(t) - t.nOffset()) / bytesLen[N]())
	}

	return
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t Tree[N]) Value(n uint) (v uint, has bool) {
	if valid, limit := t.valid(n); valid {
		v, has = header.Value(t.node(limit), t)
	}

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t Tree[N]) Chunk(n uint) (c string) {
	if valid, limit := t.valid(n); valid {
		l, h := header.ChunkRange(t.node(limit), t)
		c = string(t[cstart:][l:h])
	}

	return
}

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t Tree[N]) EachChild(n uint, e func(uint) bool) {
	valid, limit := t.valid(n)
	if !valid {
		return
	}

	for c, h := header.ChildrenRange(n, t.node(limit), t); c < h; c++ {
		if e(c) {
			return
		}
	}
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (t Tree[_]) Hoard() (uint, uint) {
	return uint(len(t)), radixt.HoardExactly
}

// Switch takes node n and byte b. If the node belongs to the tree, it looks
// for a child c of the node with such a chunk, that its first byte coincides
// with b. If such a child is found, it returns the child with its chunk
// without first byte and boolean truth. Otherwise or if the node is not in the
// tree, it returns corresponding default values.
func (t Tree[_]) Switch(n uint, b byte) (c uint, chunk string, found bool) {
	valid, limit := t.valid(n)
	if !valid {
		return
	}

	l, h := header.ChildrenRange(n, t.node(limit), t)
	if h == 0 {
		return
	}

	chunks := string(t[cstart:])
	offset := t.nOffset()
	// Explicit hiding of c to possible use in return after the loop
	for l < h {
		m := l + (h - l) >> 1
		limit := t.limit(offset, m)
		child := t.node(limit)
		low := header.ChunkLow(child, t)
		b1 := chunks[low]
		switch {
		case b1 == b:
			high := low + header.ChunkLen(child, t)
			return m, chunks[low+1 : high], true

		case b1 > b:
			h = m
		default:
			l = m + 1
		}
	}

	return
}

func (t Tree[_]) empty() bool {
	return len(t) < ProperLen
}

func (t Tree[_]) nOffset() int {
	return int(t[hlen]) | int(t[hlen+1])<<8
}

func (t Tree[N]) limit(offset int, n uint) int {
	return offset + int(n+1)*bytesLen[N]()
}

func (t Tree[_]) valid(n uint) (result bool, limit int) {
	if !t.empty() {
		limit = t.limit(t.nOffset(), n)
		result = limit <= len(t)
	}

	return
}

func (t Tree[N]) node(limit int) (result uint32) {
	i := limit - bytesLen[N]()
	result = uint32(t[i]) | uint32(t[i+1])<<8 | uint32(t[i+2])<<16
	if bytesLen[N]() == 4 {
		result |= uint32(t[i+3]) << 24
	}

	return
}

var (
	_ radixt.Tree     = Tree[N3]("")
	_ radixt.Hoarder  = Tree[N3]("")
	_ lookup.Switcher = Tree[N3]("")
	_ radixt.Tree     = Tree[N4]("")
	_ radixt.Hoarder  = Tree[N4]("")
	_ lookup.Switcher = Tree[N4]("")
)
