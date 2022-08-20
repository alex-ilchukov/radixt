package generic

import "github.com/alex-ilchukov/radixt"

type node struct {
	chunk    string
	mark     uint
	children []int
}

type tree struct {
	nodes []node
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

// Mark returns mark of node n, if the tree has the node and the node is
// marked, or zero otherwise.
func (t *tree) Mark(n int) uint {
	if t.Has(n) {
		return t.mark(n)
	}

	return 0
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
func (t *tree) ByteAt(n, npos int) (b byte, within bool) {
	if !t.Has(n) || npos < 0 {
		return
	}

	return t.byteAt(n, npos)
}

func (t *tree) mark(n int) uint {
	return t.nodes[n].mark
}

func (t *tree) chunk(n int) string {
	return t.nodes[n].chunk
}

func (t *tree) children(n int) []int {
	return t.nodes[n].children
}

func (t *tree) end(n, npos int) bool {
	return len(t.chunk(n)) <= npos
}

func (t *tree) byteAt(n, npos int) (b byte, within bool) {
	chunk := t.chunk(n)
	if len(chunk) <= npos {
		return
	}

	within = true
	b = chunk[npos]

	return
}

func (t *tree) transit(n, npos int, b byte) int {
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

func (t *tree) find(s string) (found bool, n, pos, npos int) {
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

func (t *tree) insert(s string, mark uint) {
	found, n, pos, npos := t.find(s)
	switch {
	case !found:
		if !t.end(n, npos) {
			t.splitNode(n, npos, 0)
		}

		t.addChild(n, s[pos:], mark)

	case !t.end(n, npos):
		t.splitNode(n, npos, mark)

	case t.mark(n) == 0:
		t.nodes[n].mark = mark

	default:
		return
	}
}

func (t *tree) splitNode(n, npos int, mark uint) {
	no := t.nodes[n]
	chunk := no.chunk
	no.chunk = chunk[npos:]
	t.nodes = append(t.nodes, no)
	t.nodes[n] = node{chunk[:npos], mark, []int{len(t.nodes) - 1}}
}

func (t *tree) addChild(n int, chunk string, mark uint) {
	t.nodes = append(t.nodes, node{chunk, mark, nil})
	t.nodes[n].children = append(t.nodes[n].children, len(t.nodes)-1)
}

// New creates a new generic tree, inserting the provided strings, and returns
// a pointer on the tree. Node marks are indicies of the strings, incremented
// by one.
func New(strings ...string) *tree {
	t := new(tree)

	if len(strings) == 0 {
		return t
	}

	t.nodes = []node{{chunk: strings[0], mark: 1}}
	for m, s := range strings[1:] {
		t.insert(s, uint(m+2))
	}

	return t
}

var _ radixt.Tree = (*tree)(nil)
