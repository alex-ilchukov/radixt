package generic

import "github.com/alex-ilchukov/radixt"

type node struct {
	pref     string
	mark     int
	children []int
}

type tree struct {
	strings []string
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

// NodeMark returns mark of node n, if the tree has the node and the node is
// marked, or -1 otherwise.
func (t *tree) NodeMark(n int) int {
	if t.Has(n) {
		return t.mark(n)
	}

	return -1
}

// NodeMark returns string of node n, if the tree has the node and the node is
// marked, or empty string otherwise.
func (t *tree) NodeString(n int) string {
	mark := t.NodeMark(n)
	if mark >= 0 {
		return t.strings[mark]
	}

	return ""
}

// NodePref returns prefix string of node n, if the tree has the node, or empty
// string otherwise.
func (t *tree) NodePref(n int) string {
	if t.Has(n) {
		return t.pref(n)
	}

	return ""
}

// NodeEachChild calls func e for every child of node n, if the tree has the
// node, until the func returns boolean true. The order of going over the
// children is fixed for every node, but may not coincide with any natural
// order.
func (t *tree) NodeEachChild(n int, e func(int) bool) {
	if !t.Has(n) {
		return
	}

	for _, c := range t.children(n) {
		if e(c) {
			return
		}
	}
}

// NodeTransit transits from node n accordingly to the following rules. Let
// start with prefix of the node.
//  1. If npos is less than length of the prefix, and byte of the prefix at
//     npos is b, then n is returned.
//  2. If npos is less than length of the prefix, but byte of the prefix at
//     npos is not b, then -1 (which is non-node) is returned.
//  3. If npos is more or equal to length of the prefix, then children of n are
//     checked; if there is a child with b as first byte of its prefix, then
//     the child is returned; if there is no such a child, -1 is returned.
func (t *tree) NodeTransit(n, npos int, b byte) int {
	if t.Has(n) && npos >= 0 {
		return t.transit(n, npos, b)
	}

	return -1
}

func (t *tree) mark(n int) int {
	return t.nodes[n].mark
}

func (t *tree) pref(n int) string {
	return t.nodes[n].pref
}

func (t *tree) children(n int) []int {
	return t.nodes[n].children
}

func (t *tree) end(n, npos int) bool {
	return len(t.pref(n)) <= npos
}

func (t *tree) transit(n, npos int, b byte) int {
	pref := t.pref(n)
	if npos < len(pref) {
		if pref[npos] == b {
			return n
		}

		return -1
	}

	for _, c := range t.children(n) {
		if t.pref(c)[0] == b {
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

func (t *tree) insert(s string) {
	found, n, pos, npos := t.find(s)
	switch {
	case !found:
		if !t.end(n, npos) {
			t.splitNode(n, npos, -1)
		}

		t.addChild(n, s[pos:])

	case !t.end(n, npos):
		t.splitNode(n, npos, len(t.strings))

	case t.mark(n) < 0:
		t.nodes[n].mark = len(t.strings)

	default:
		return
	}

	t.strings = append(t.strings, s)
}

func (t *tree) splitNode(n, npos, mark int) {
	no := t.nodes[n]
	pref := no.pref
	no.pref = pref[npos:]
	t.nodes = append(t.nodes, no)
	t.nodes[n] = node{pref[:npos], mark, []int{len(t.nodes) - 1}}
}

func (t *tree) addChild(n int, pref string) {
	t.nodes = append(t.nodes, node{pref, len(t.strings), nil})
	t.nodes[n].children = append(t.nodes[n].children, len(t.nodes)-1)
}

// New creates a new generic tree, inserting the provided strings, and returns
// a pointer on the tree
func New(strings ...string) *tree {
	t := new(tree)

	if len(strings) == 0 {
		return t
	}

	s := strings[0]
	t.strings = []string{s}
	t.nodes = []node{{pref: s}}

	for _, s = range strings[1:] {
		t.insert(s)
	}

	return t
}

var _ radixt.Tree = (*tree)(nil)
