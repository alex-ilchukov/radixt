package lookup

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
)

type l struct {
	t     radixt.Tree
	n     uint
	chunk string
	keep  bool
}

// L contains information on state of the lookup process.
type L l

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
	l.keep = l.t.Size() > 0
}

func (l *L) try(b byte, n uint, chunk string) {
	l.keep = b == chunk[0]
	if l.keep {
		l.n = n
		l.chunk = chunk[1:]
	}
}

// Feed takes byte b and returns if the byte is found in radix tree accordingly
// to the state or not.
func (l *L) Feed(b byte) bool {
	switch {
	case !l.keep:
		// no statement
	case l.chunk != "":
		l.try(b, l.n, l.chunk)
	default:
		l.t.EachChild(l.n, func(c uint) bool {
			l.try(b, c, l.t.Chunk(c))
			return l.keep
		})
	}

	return l.keep
}

// Found returns if the lookup state points to result string with value in the
// tree or not.
func (l *L) Found() (found bool) {
	if l.keep && l.chunk == "" {
		_, found = l.t.Value(l.n)
	}

	return
}

// Tree returns radix tree.
func (l *L) Tree() radixt.Tree {
	return l.t
}

// Node returns index of current tree node.
func (l *L) Node() uint {
	return l.n
}
