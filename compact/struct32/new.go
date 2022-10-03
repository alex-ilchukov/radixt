package struct32

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
)

// New takes the provided tree t and tries to compactify it. In case of success
// it returns new, compactified representation of the tree and nil for error.
// In case of an error, it returns nil for tree and the error.
func New(t radixt.Tree) (*tree, error) {
	a := analysis.Do[analysis.Firstless](t)
	h, nf, err := header.Calc[uint32](32, a)
	if err != nil {
		return nil, err
	}

	s := shifts{
		sChunkPos:        h[0],
		lsValue:          h[1],
		rsValue:          h[2],
		lsChildrenStart:  h[3],
		rsChildrenStart:  h[4],
		lsChildrenAmount: h[5],
		rsChildrenAmount: h[6],
		sChunkLen:        h[7],
	}

	l := len(a.N)
	nodes := make([]node, l, l)
	cf := make([]byte, l, l)
	for _, n := range a.N {
		nodes[n.Index] = node(nf(n))
		cf[n.Index] = n.ChunkFirst
	}

	result := &tree{
		emptyRoot: l > 0 && a.N[0].ChunkEmpty,
		s:         s,
		chunks:    a.C,
		cf:        string(cf),
		nodes:     nodes,
	}

	return result, nil
}

// MustCreate takes the provided tree t and tries to compactify it. In case of
// success it returns new, compactified representation of the tree. In case of
// an error, it panics.
func MustCreate(t radixt.Tree) *tree {
	result, err := New(t)
	if err != nil {
		panic(err)
	}

	return result
}
