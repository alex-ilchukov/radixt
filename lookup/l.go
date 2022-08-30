package lookup

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
)

// L contains information on state of the lookup process.
type L struct {
	t    radixt.Tree
	n    uint
	npos uint
	stop bool
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
	l.npos = 0
	l.stop = l.t.Size() == 0
}

// Feed takes byte b and returns if the byte is found in radix tree accordingly
// to the state or not.
func (l *L) Feed(b byte) bool {
	if l.stop {
		return false
	}

	t := l.t
	n := l.n
	byteAt, within := t.ByteAt(n, l.npos)
	if within {
		if byteAt == b {
			l.npos++
			return true
		}

		l.stop = true
		return false
	}

	for c, f := t.ChildrenRange(n); c <= f; c++ {
		byteAt, _ := t.ByteAt(c, 0)
		if byteAt == b {
			l.n = c
			l.npos = 1
			return true
		}
	}

	l.stop = true
	return false
}

// Found returns if the lookup state points to result string with value in the
// tree or not.
func (l *L) Found() bool {
	if l.stop {
		return false
	}

	t := l.t
	n := l.n
	_, has := t.Value(n)
	if !has {
		return false
	}

	_, within := t.ByteAt(n, l.npos)
	if within {
		return false
	}

	return true
}

// Tree returns radix tree.
func (l *L) Tree() radixt.Tree {
	return l.t
}

// Node returns index of current tree node.
func (l *L) Node() uint {
	return l.n
}
