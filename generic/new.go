package generic

// New creates a new generic tree, inserting the provided strings, and returns
// a pointer on the tree. Node values are indices of the strings.
func New(strings ...string) *tree {
	if len(strings) == 0 {
		return new(tree)
	}

	s := new(shrub)
	s.noValue = uint(len(strings))
	s.imagoes = []imago{{chunk: strings[0], value: 0, hasValue: true}}

	for v, str := range strings[1:] {
		s.insert(str, uint(v+1))
	}

	return grow(s)
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
	if len(sv) == 0 {
		return new(tree)
	}

	s := new(shrub)
	s.noValue = uint(len(sv))
	s.imagoes = []imago{{chunk: sv[0].S, value: sv[0].V, hasValue: true}}

	for _, e := range sv[1:] {
		s.insert(e.S, e.V)
	}

	return grow(s)
}
