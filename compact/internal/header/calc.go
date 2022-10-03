package header

import (
	"math/bits"

	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/internal/node"
)

// NodeFactory represents all factory functions, which take result of node
// analysis and return node in form of type from [internal/node.N] type set.
type NodeFactory[N node.N, M analysis.Mode] func(n analysis.N[M]) N

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
func Calc[N node.N, M analysis.Mode](lenNode int, a analysis.A[M]) (
	h A8b,
	nf NodeFactory[N, M],
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

	h = fillHeader(node.BitsLen[N](), lens)
	nf = createNodeFactory[N, M](lens)

	return
}

func fillLens[M analysis.Mode](a analysis.A[M]) (lens fieldLens) {
	lens[fieldChunkPos] = bits.Len(uint(len(a.C)))
	// Zero is NoValue, so (a.Vm + 1) values would be in use
	lens[fieldValue] = bits.Len(a.Vm + 1)

	// Zero is for empty and one-node trees. Any other trees have at least
	// one parent node and, as a corollary, have a.Dclpm > 0. Indeed, any
	// child's index (including the minimal, the first one) is strictly
	// greater than its parent index, so the difference is always positive.
	if a.Dclpm > 0 {
		lens[fieldChildrenStart] = bits.Len(a.Dclpm - 1)
	}

	lens[fieldChildrenAmount] = bits.Len(a.Cma)
	lens[fieldChunkLen] = bits.Len(a.Cml)

	return
}

func fillHeader(lenNode int, lens fieldLens) (h A8b) {
	h[0] = byte(lenNode - lens[0])
	ls := h[0]
	for i := 1; i < fieldsAmount-1; i++ {
		ls -= byte(lens[i])
		h[2*i-1] = ls                    // left shift
		h[2*i] = byte(lenNode - lens[i]) // right shift
	}
	h[Len-1] = byte(lenNode) - ls

	return
}

func createNodeFactory[N node.N, M analysis.Mode](lens fieldLens) (
	NodeFactory[N, M],
) {
	s := fillShifts(lens)
	return func(n analysis.N[M]) N {
		var f fields

		f[fieldChunkPos] = n.ChunkPos
		f[fieldChunkLen] = uint(len(n.Chunk))

		if n.HasValue {
			f[fieldValue] = n.Value + 1
		}

		f[fieldChildrenAmount] = n.ChildrenHigh - n.ChildrenLow
		if 0 < n.ChildrenHigh {
			f[fieldChildrenStart] = n.ChildrenLow - n.Index - 1
		}

		return N(placeFields(f, s))
	}
}

func fillShifts(lens fieldLens) (s fieldShifts) {
	for i := 1; i < fieldsAmount; i++ {
		s[i] = s[i-1] + byte(lens[i-1])
	}

	return
}

func placeFields(f fields, s fieldShifts) (result uint) {
	for i := 0; i < fieldsAmount; i++ {
		result |= f[i] << s[i]
	}

	return
}
