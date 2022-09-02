package structg

import (
	"errors"
	"math/bits"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
)

var ErrorOverflow = errors.New("required fields would not fit into node")

// New takes the provided tree t and tries to compactify it. In case of success
// it returns new, compactified representation of the tree and nil for error.
// In case of an error, it returns nil for tree and the error.
func New[N node](t radixt.Tree) (*tree[N], error) {
	a := analysis.Do(t)

	lenChunkPos := bits.Len(uint(len(a.C)))
	// Zero is NoValue, so (a.Vm + 1) values would be in use
	lenValue := bits.Len(a.Vm + 1)

	// Zero is for empty and one-node trees. Any other trees have at least
	// one parent node and, as a corollary, have a.Dcfpm > 0. Indeed, any
	// child's index (including the minimal, the first one) is strictly
	// greater than its parent index, so the difference is always positive.
	lenChildrenStart := 0
	if a.Dcfpm > 0 {
		lenChildrenStart = bits.Len(a.Dcfpm - 1)
	}

	lenChildrenAmount := bits.Len(a.Cma)
	lenChunkLen := bits.Len(a.Cml)

	l := lenChunkPos + lenValue + lenChildrenStart + lenChildrenAmount +
		lenChunkLen

	lenNode := bitslen[N]()

	if l > lenNode {
		return nil, ErrorOverflow
	}

	sChunkPos := byte(lenNode - lenChunkPos)

	sValue := byte(lenChunkPos)
	lsValue := sChunkPos - byte(lenValue)
	rsValue := byte(lenNode - lenValue)

	sChildrenStart := sValue + byte(lenValue)
	lsChildrenStart := lsValue - byte(lenChildrenStart)
	rsChildrenStart := byte(lenNode - lenChildrenStart)

	sChildrenAmount := sChildrenStart + byte(lenChildrenStart)
	lsChildrenAmount := lsChildrenStart - byte(lenChildrenAmount)
	rsChildrenAmount := byte(lenNode - lenChildrenAmount)

	sChunkLen := byte(lenNode - lenChunkLen)

	nodes := make([]N, len(a.N), len(a.N))

	result := &tree[N]{
		sChunkPos:        sChunkPos,
		lsValue:          lsValue,
		rsValue:          rsValue,
		lsChildrenStart:  lsChildrenStart,
		rsChildrenStart:  rsChildrenStart,
		lsChildrenAmount: lsChildrenAmount,
		rsChildrenAmount: rsChildrenAmount,
		sChunkLen:        sChunkLen,
		chunks:           a.C,
		nodes:            nodes,
	}

	for i, n := range a.N {
		value := N(n.Value)
		if n.HasValue {
			value++
		}

		f := n.ChildrenFirst
		l := n.ChildrenLast
		childrenStart := N(0)
		childrenAmount := N(0)
		if f <= l {
			childrenStart = N(f - n.Index - 1)
			childrenAmount = N(l - f + 1)
		}

		nodes[i] = N(n.ChunkPos) |
			value<<sValue |
			childrenStart<<sChildrenStart |
			childrenAmount<<sChildrenAmount |
			N(len(n.Chunk))<<sChunkLen
	}

	return result, nil
}

// MustCreate takes the provided tree t and tries to compactify it. In case of
// success it returns new, compactified representation of the tree. In case of
// an error, it panics.
func MustCreate[N node](t radixt.Tree) *tree[N] {
	result, err := New[N](t)
	if err != nil {
		panic(err)
	}

	return result
}
