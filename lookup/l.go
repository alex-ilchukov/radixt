package lookup

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
)

// L contains information on state of the lookup process.
type L struct {
	t    radixt.Tree
	n    int
	npos int
}

// New creates and initializes new lookup state accordingly to the provided
// radix tree t, and returns a pointer the state. Nil values of t are supported
// and interpreted as empty tree.
func New(t radixt.Tree) *L {
	if t == nil {
		t = null.Tree
	}

	l := &L{t: t}
	l.Reset()

	return l
}

// Reset resets the lookup state.
func (l *L) Reset() {
	l.n = l.t.Root()
	l.npos = 0
}

// Feed takes byte b and returns if the byte is found in radix tree accordingly
// to the state or not.
func (l *L) Feed(b byte) bool {
	t := l.t
	n := l.n
	l.n = t.NodeTransit(n, l.npos, b)
	if !t.Has(l.n) {
		return false
	}
	if l.n == n {
		l.npos++
	} else {
		l.npos = 1
	}

	return true
}

// Found returns if the lookup state points to string in the tree or not.
func (l *L) Found() bool {
	t := l.t
	n := l.n
	return t.NodeMark(n) >= 0 && len(t.NodePref(n)) <= l.npos
}

// Tree returns radix tree.
func (l *L) Tree() radixt.Tree {
	return l.t
}

// Node returns index of current tree node or non-node index if the lookup
// process has already finished with failure.
func (l *L) Node() int {
	return l.n
}
