package sapling

// New creates a new sapling tree, inserting the provided strings, and returns
// a pointer on the tree. Node values are indices of the strings.
func New(strings ...string) *Tree {
	t := new(Tree)
	for v, s := range strings {
		t.Grow(s, uint(v))
	}

	return t
}

// SV represents a couple of string key S and unsigned integer value V to be
// contained in a tree
type SV struct {
	S string
	V uint
}

// NewFromSV creates a new sapling tree, inserting strings with values from the
// provided sv slice, and returns a pointer on the tree.
func NewFromSV(sv ...SV) *Tree {
	t := new(Tree)
	for _, e := range sv {
		t.Grow(e.S, e.V)
	}

	return t
}
