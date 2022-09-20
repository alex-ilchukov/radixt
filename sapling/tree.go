package sapling

import "github.com/alex-ilchukov/radixt"

type node struct {
	chunk    string
	value    uint
	hasValue bool
	children []uint
}

// Tree represents radix tree implementation.
type Tree struct {
	nodes []node
}

// Size returns amount of nodes in the tree. It returns zero for if t is nil.
func (t *Tree) Size() uint {
	if t == nil {
		return 0
	}

	return uint(len(t.nodes))
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t *Tree) Value(n uint) (v uint, has bool) {
	if n < t.Size() {
		node := t.nodes[n]
		v = node.value
		has = node.hasValue
	}

	return
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t *Tree) Chunk(n uint) (chunk string) {
	if n < t.Size() {
		chunk = t.nodes[n].chunk
	}

	return
}

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t *Tree) EachChild(n uint, e func(uint) bool) {
	if n >= t.Size() {
		return
	}

	for _, c := range t.nodes[n].children {
		if e(c) {
			return
		}
	}
}

// Hoard returns amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint. It returns zero if t is nil.
func (t *Tree) Hoard() (uint, uint) {
	if t == nil {
		return 0, radixt.HoardExactly
	}

	amount := uint(24) + // Tree
		(16+8+8+24)*uint(len(t.nodes)) // hasBool is aligned to 8 bytes

	for _, n := range t.nodes {
		amount += uint(len(n.chunk)) + uint(len(n.children)*8)
	}

	return amount, radixt.HoardExactly
}

// Grow adds string s to the tree, associating it with the provided value and
// appending new nodes if they're required. It panics if t is nil. If the tree
// already has the string with a value, it overwrites the value with the
// provided one.
func (t *Tree) Grow(s string, v uint) {
	switch {
	case t == nil:
		panic("can not add a string into nil tree")

	case t.Size() == 0:
		t.nodes = []node{{chunk: s, value: v, hasValue: true}}

	default:
		found, n, pos, npos := t.find(s)
		switch {
		case !found:
			if t.within(n, npos) {
				t.splitNode(n, npos, 0, false)
			}

			t.addChild(n, s[pos:], v)

		case t.within(n, npos):
			t.splitNode(n, npos, v, true)

		default:
			t.nodes[n].value = v
			t.nodes[n].hasValue = true
		}
	}
}

func (t *Tree) find(s string) (found bool, n, pos, npos uint) {
	for l := uint(len(s)); pos < l; pos++ {
		f, m := t.transit(n, npos, s[pos])
		switch {
		case !f:
			return
		case m == n:
			npos++
		default:
			npos = 1
			n = m
		}
	}

	found = true

	return
}

func (t *Tree) transit(n, npos uint, b byte) (bool, uint) {
	nodes := t.nodes
	no := nodes[n]
	chunk := no.chunk
	if npos < uint(len(chunk)) {
		return chunk[npos] == b, n
	}

	for _, c := range no.children {
		if nodes[c].chunk[0] == b {
			return true, c
		}
	}

	return false, n
}

func (t *Tree) within(n, npos uint) bool {
	return npos < uint(len(t.nodes[n].chunk))
}

func (t *Tree) splitNode(n, npos, value uint, hasValue bool) {
	no := t.nodes[n]
	chunk := no.chunk
	no.chunk = chunk[npos:]
	t.nodes = append(t.nodes, no)
	t.nodes[n] = node{
		chunk:    chunk[:npos],
		value:    value,
		hasValue: hasValue,
		children: []uint{uint(len(t.nodes) - 1)},
	}
}

func (t *Tree) addChild(n uint, chunk string, value uint) {
	no := node{chunk: chunk, value: value, hasValue: true}
	t.nodes = append(t.nodes, no)
	t.nodes[n].children = append(t.nodes[n].children, uint(len(t.nodes)-1))
}

var (
	_ radixt.Tree    = (*Tree)(nil)
	_ radixt.Hoarder = (*Tree)(nil)
)
