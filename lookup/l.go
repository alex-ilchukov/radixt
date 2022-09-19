package lookup

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
)

type l[W Way] struct {
	t     radixt.Tree
	w     W
	n     uint
	chunk string
	keep  bool
}

// L contains information on state of the lookup process.
type L struct {
	l[Default]
}

// New creates and initializes new lookup state accordingly to the provided
// radix tree t, and returns a pointer the state. Nil values of t are supported
// and interpreted as empty tree.
func New(t radixt.Tree) *L {
	if t == nil {
		t = null.Tree
	}

	l := &L{l: l[Default]{t: t}}
	l.Reset()

	return l
}

// Reset resets the lookup state.
func (lkp *l[_]) Reset() {
	lkp.n = 0
	lkp.chunk = lkp.t.Chunk(0)
	lkp.keep = lkp.t.Size() > 0
}

func (lkp *l[_]) try(b byte, n uint, chunk string) {
	lkp.keep = b == chunk[0]
	if lkp.keep {
		lkp.n = n
		lkp.chunk = chunk[1:]
	}
}

// Feed takes byte b and returns if the byte is found in radix tree accordingly
// to the state or not.
func (lkp *l[W]) Feed(b byte) bool {
	switch {
	case !lkp.keep:
		// no statement
	case lkp.chunk != "":
		lkp.try(b, lkp.n, lkp.chunk)
	case IsSwitch[W]():
		// The trick with len is required to coax the compiler: putting
		// just 0 would bring its wrath on the case, where the array
		// has zero length (Default way that is), despite the guardian
		// IsSwitch check.
		w := lkp.w[len(lkp.w)-1]
		lkp.n, lkp.chunk, lkp.keep = w.Switch(lkp.n, b)
	default:
		lkp.t.EachChild(lkp.n, func(c uint) bool {
			lkp.try(b, c, lkp.t.Chunk(c))
			return lkp.keep
		})
	}

	return lkp.keep
}

// Found returns if the lookup state points to result string with value in the
// tree or not.
func (lkp *l[_]) Found() (found bool) {
	if lkp.keep && lkp.chunk == "" {
		_, found = lkp.t.Value(lkp.n)
	}

	return
}

// Tree returns radix tree.
func (lkp *l[_]) Tree() radixt.Tree {
	return lkp.t
}

// Node returns index of current tree node.
func (lkp *l[_]) Node() uint {
	return lkp.n
}
