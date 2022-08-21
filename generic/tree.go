package generic

import "github.com/alex-ilchukov/radixt"

type node struct {
	chunk    string
	value    uint
	children []int
}

type tree struct {
	noValue uint
	nodes   []node
}

// Size returns amount of nodes in the tree.
func (t *tree) Size() int {
	return len(t.nodes)
}

// Has returns if the tree has node n or not.
func (t *tree) Has(n int) bool {
	return 0 <= n && n < t.Size()
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

	v = t.value(n)
	if v == t.noValue {
		v = 0
	} else {
		has = true
	}

	return
}

// EachChild calls func e for every child of node n, if the tree has the node,
// until the func returns boolean true. The order of going over the children is
// fixed for every node, but may not coincide with any natural order.
func (t *tree) EachChild(n int, e func(int) bool) {
	if !t.Has(n) {
		return
	}

	for _, c := range t.children(n) {
		if e(c) {
			return
		}
	}
}

// ByteAt returns default byte value and boolean false, if npos is outside of
// chunk of the node n, or byte of the chunk at npos and boolean true
// otherwise.
func (t *tree) ByteAt(n int, npos uint) (b byte, within bool) {
	if t.Has(n) {
		return t.byteAt(n, npos)
	}

	return
}

func (t *tree) value(n int) uint {
	return t.nodes[n].value
}

func (t *tree) chunk(n int) string {
	return t.nodes[n].chunk
}

func (t *tree) children(n int) []int {
	return t.nodes[n].children
}

func (t *tree) end(n int, npos uint) bool {
	return uint(len(t.chunk(n))) <= npos
}

func (t *tree) byteAt(n int, npos uint) (b byte, within bool) {
	chunk := t.chunk(n)
	if uint(len(chunk)) <= npos {
		return
	}

	within = true
	b = chunk[npos]

	return
}

func (t *tree) transit(n int, npos uint, b byte) int {
	byteAt, within := t.byteAt(n, npos)
	if within {
		if byteAt == b {
			return n
		}

		return -1
	}

	for _, c := range t.children(n) {
		byteAt, _ = t.byteAt(c, 0)
		if byteAt == b {
			return c
		}
	}

	return -1
}

func (t *tree) find(s string) (found bool, n, pos int, npos uint) {
	l := len(s)
	for ; pos < l; pos++ {
		m := t.transit(n, npos, s[pos])
		switch m {
		case -1:
			return
		case n:
			npos++
		default:
			npos = 1
			n = m
		}
	}

	found = true

	return
}

func (t *tree) insert(s string, value uint) {
	found, n, pos, npos := t.find(s)
	switch {
	case !found:
		if !t.end(n, npos) {
			t.splitNode(n, npos, t.noValue)
		}

		t.addChild(n, s[pos:], value)

	case !t.end(n, npos):
		t.splitNode(n, npos, value)

	case t.value(n) == t.noValue:
		t.nodes[n].value = value

	default:
		return
	}
}

func (t *tree) splitNode(n int, npos uint, value uint) {
	no := t.nodes[n]
	chunk := no.chunk
	no.chunk = chunk[npos:]
	t.nodes = append(t.nodes, no)
	t.nodes[n] = node{chunk[:npos], value, []int{len(t.nodes) - 1}}
}

func (t *tree) addChild(n int, chunk string, value uint) {
	t.nodes = append(t.nodes, node{chunk, value, nil})
	t.nodes[n].children = append(t.nodes[n].children, len(t.nodes)-1)
}

// New creates a new generic tree, inserting the provided strings, and returns
// a pointer on the tree. Node values are indicies of the strings, incremented
// by one.
func New(strings ...string) *tree {
	t := new(tree)

	if len(strings) == 0 {
		return t
	}

	t.noValue = uint(len(strings))
	t.nodes = []node{{chunk: strings[0], value: 0}}

	for v, s := range strings[1:] {
		t.insert(s, uint(v+1))
	}

	return t
}

var _ radixt.Tree = (*tree)(nil)
