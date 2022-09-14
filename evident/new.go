package evident

import "github.com/alex-ilchukov/radixt"

// New creates evident representation of the provided radix tree t and returns
// it. The function checks for the following special cases.
//
//  1. If t is nil, then nil is returned.
//  2. If t is an evident tree, then it is returned.
//  3. If t is empty, then nil is returned.
func New(t radixt.Tree) (result Tree) {
	if t == nil {
		return
	}

	result, success := t.(Tree)
	if success || t.Size() == 0 {
		return
	}

	result = make(Tree)

	type es struct {
		n uint
		e Tree
	}

	s := []es{{n: 0, e: result}}
	for len(s) > 0 {
		l := len(s) - 1
		a := s[l]
		s = s[:l]

		n := a.n
		e := a.e

		child := Tree(nil)
		t.EachChild(n, func(c uint) bool {
			if child == nil {
				child = make(Tree)
			}
			s = append(s, es{c, child})

			return false
		})

		chunk := t.Chunk(n)
		value, has := t.Value(n)
		key := newKey(chunk, value, has)
		e[key] = child
	}

	return
}
