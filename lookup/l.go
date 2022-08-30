package lookup

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
)

// L contains information on state of the lookup process.
type L struct {
	t     radixt.Tree
	n     uint
	chunk string
	stop  bool
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
	l.n = 0
	l.chunk = l.t.Chunk(0)
	l.stop = l.t.Size() == 0
}

func (l *L) try(b byte, n uint, chunk string) bool {
	l.stop = b != chunk[0]
	if !l.stop {
		l.n = n
		l.chunk = chunk[1:]
	}

	return l.stop
}

// Feed takes byte b and returns if the byte is found in radix tree accordingly
// to the state or not.
func (l *L) Feed(b byte) bool {
	if l.stop {
		return false
	}

	if l.chunk != "" {
		l.try(b, l.n, l.chunk)
	} else {
		c, f := l.t.ChildrenRange(l.n)
		for ; c <= f && l.try(b, c, l.t.Chunk(c)); c++ {
		}
	}

	return !l.stop
}

// Found returns if the lookup state points to result string with value in the
// tree or not.
func (l *L) Found() bool {
	if l.stop || l.chunk != "" {
		return false
	}

	_, has := l.t.Value(l.n)
	return has
}

// Tree returns radix tree.
func (l *L) Tree() radixt.Tree {
	return l.t
}

// Node returns index of current tree node.
func (l *L) Node() uint {
	return l.n
}
