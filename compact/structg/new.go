package structg

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
	"github.com/alex-ilchukov/radixt/compact/internal/node"
)

// New takes the provided tree t and tries to compactify it. In case of success
// it returns new, compactified representation of the tree and nil for error.
// In case of an error, it returns nil for tree and the error.
func New[N node.N](t radixt.Tree) (*tree[N], error) {
	a := analysis.Do(t)
	h, nf, err := header.Calc[N](node.BitsLen[N](), a)
	if err != nil {
		return nil, err
	}

	nodes := make([]N, len(a.N), len(a.N))
	for i, n := range a.N {
		nodes[i] = nf(n)
	}

	result := &tree[N]{h: h, chunks: a.C, nodes: nodes}

	return result, nil
}

// MustCreate takes the provided tree t and tries to compactify it. In case of
// success it returns new, compactified representation of the tree. In case of
// an error, it panics.
func MustCreate[N node.N](t radixt.Tree) *tree[N] {
	result, err := New[N](t)
	if err != nil {
		panic(err)
	}

	return result
}
