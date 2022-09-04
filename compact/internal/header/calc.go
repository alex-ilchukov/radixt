package header

import (
	"math/bits"

	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/internal/node"
)

// NodeFactory represents all factory functions, which take result of node
// analysis and return node in form of type from [internal/node.N] type set.
type NodeFactory[N node.N] func(n analysis.N) N

type fields [fieldsAmount]uint
type fieldLens [fieldsAmount]int
type fieldShifts [fieldsAmount]byte

// Calc takes maximum used bits length of node, which can be less or equal to
// actual bits length of node, with result of tree analysis. It returns proper
// filled h and node factory function nf with nil err, if no error appearead.
// Otherwise returns the following error values:
//
//  1. [compact.ErrorInvalidLenNode] if provided lenNode is more than actual
//     bits length of node;
//  2. [compact.ErrorOverflow] if node fields can not be fit into node value.
func Calc[N node.N](lenNode int, a analysis.A) (
	h   A8b,
	nf  NodeFactory[N],
	err error,
) {
	if node.BitsLen[N]() < lenNode {
		err = compact.ErrorInvalidLenNode
		return
	}

	lens := fillLens(a)
	l := lenNode
	for _, fieldLen := range lens {
		l -= fieldLen
	}

	if l < 0 {
		err = compact.ErrorOverflow
		return
	}

	h = fillHeader(lenNode, lens)
	nf = createNodeFactory[N](lens)

	return
}

func fillLens(a analysis.A) (lens fieldLens) {
	lens[fieldChunkPos] = bits.Len(uint(len(a.C)))
	// Zero is NoValue, so (a.Vm + 1) values would be in use
	lens[fieldValue] = bits.Len(a.Vm + 1)

	// Zero is for empty and one-node trees. Any other trees have at least
	// one parent node and, as a corollary, have a.Dcfpm > 0. Indeed, any
	// child's index (including the minimal, the first one) is strictly
	// greater than its parent index, so the difference is always positive.
	if a.Dcfpm > 0 {
		lens[fieldChildrenStart] = bits.Len(a.Dcfpm - 1)
	}

	lens[fieldChildrenAmount] = bits.Len(a.Cma)
	lens[fieldChunkLen] = bits.Len(a.Cml)

	return
}

func fillHeader(lenNode int, lens fieldLens) (h A8b) {
	h[0] = byte(lenNode - lens[0])
	ls := h[0]
	for i := 1; i < fieldsAmount - 1; i++ {
		ls -= byte(lens[i])
		h[2*i-1] = ls // left shift
		h[2*i] = byte(lenNode - lens[i]) // right shift
	}
	h[hlen-1] = byte(lenNode) - ls

	return
}

func createNodeFactory[N node.N](lens fieldLens) NodeFactory[N] {
	s := fillShifts(lens)
	return func(n analysis.N) N {
		var f fields

		f[fieldChunkPos] = n.ChunkPos
		f[fieldChunkLen] = uint(len(n.Chunk))

		if n.HasValue {
			f[fieldValue] = n.Value + 1
		}

		a := n.ChildrenFirst
		b := n.ChildrenLast
		if a <= b {
			f[fieldChildrenStart] = a - n.Index - 1
			f[fieldChildrenAmount] = b - a + 1
		}

		return N(placeFields(f, s))
	}
}

func fillShifts(lens fieldLens) (s fieldShifts) {
	for i := 1; i < fieldsAmount; i++ {
		s[i] = s[i - 1] + byte(lens[i - 1])
	}

	return
}

func placeFields(f fields, s fieldShifts) (result uint) {
	for i := 0; i < fieldsAmount; i++ {
		result |= f[i]<<s[i]
	}

	return
}
