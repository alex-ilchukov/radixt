package generic

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
)

// New creates a new generic tree as a copy of the provided tree t and returns
// a pointer on the created tree. It returns empty tree, if t is nil.
func New(t radixt.Tree) *tree {
	a := analysis.Do(t)
	nodes := make([]node, len(a.N), len(a.N))
	for _, n := range a.N {
		nodes[n.Index] = node{
			hasValue:  n.HasValue,
			cAmount:   byte(n.ChildrenHigh - n.ChildrenLow),
			cFirst:    n.ChildrenLow,
			chunkLow:  n.ChunkPos,
			chunkHigh: n.ChunkPos + uint(len(n.Chunk)),
			value:     n.Value,
		}
	}

	return &tree{nodes: nodes, c: a.C}
}
