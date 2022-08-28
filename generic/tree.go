package generic

import "github.com/alex-ilchukov/radixt"

type node struct {
	chunk    string
	value    uint
	children []uint
}

type tree struct {
	noValue uint
	nodes   []node
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
func (t *tree) EachChild(n uint, e func(uint) bool) {
	if n >= t.Size() {
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
func (t *tree) ByteAt(n, npos uint) (b byte, within bool) {
	if n >= t.Size() {
		return
	}

	return t.byteAt(n, npos)
}

func (t *tree) value(n uint) uint {
	return t.nodes[n].value
}

func (t *tree) chunk(n uint) string {
	return t.nodes[n].chunk
}

func (t *tree) children(n uint) []uint {
	return t.nodes[n].children
}

func (t *tree) end(n, npos uint) bool {
	return uint(len(t.chunk(n))) <= npos
}

func (t *tree) byteAt(n, npos uint) (b byte, within bool) {
	chunk := t.chunk(n)
	if uint(len(chunk)) <= npos {
		return
	}

	within = true
	b = chunk[npos]

	return
}

func (t *tree) transit(n, npos uint, b byte) (found bool, m uint) {
	byteAt, within := t.byteAt(n, npos)
	if within {
		found = byteAt == b
		m = n

		return
	}

	for _, c := range t.children(n) {
		byteAt, _ = t.byteAt(c, 0)
		found = byteAt == b
		if found {
			m = c
			return
		}
	}

	return
}

func (t *tree) find(s string) (found bool, n, pos, npos uint) {
	l := uint(len(s))
	for ; pos < l; pos++ {
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

func (t *tree) splitNode(n, npos, value uint) {
	no := t.nodes[n]
	chunk := no.chunk
	no.chunk = chunk[npos:]
	t.nodes = append(t.nodes, no)
	t.nodes[n] = node{chunk[:npos], value, []uint{uint(len(t.nodes)-1)}}
}

func (t *tree) addChild(n uint, chunk string, value uint) {
	t.nodes = append(t.nodes, node{chunk, value, nil})
	t.nodes[n].children = append(t.nodes[n].children, uint(len(t.nodes)-1))
}

// New creates a new generic tree, inserting the provided strings, and returns
// a pointer on the tree. Node values are indices of the strings.
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

// SV represents a couple of string key S and unsigned integer value V to be
// contained in a tree
type SV struct {
	S string
	V uint
}

// NewFromSV creates a new generic tree, inserting strings with values from the
// provided sv slice, and returns a pointer on the tree.
func NewFromSV(sv ...SV) *tree {
	t := new(tree)

	if len(sv) == 0 {
		return t
	}

	t.noValue = uint(len(sv))
	t.nodes = []node{{chunk: sv[0].S, value: sv[0].V}}

	for _, e := range sv[1:] {
		t.insert(e.S, e.V)
	}

	return t
}

var _ radixt.Tree = (*tree)(nil)
