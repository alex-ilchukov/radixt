package header

import (
	"math/bits"

	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/internal/node"
)

// Calc takes maximum used bits length of node, which can be less or equal to
// actual bits length of node, with result of tree analysis. It returns proper
// filled h and s structures with nil err, if no error appearead. Otherwise
// returns the following error values:
//
//  1. [compact.ErrorInvalidLenNode] if provided lenNode is more than actual
//     bits length of node;
//  2. [compact.ErrorOverflow] if node fields can not be fit into node value.
func Calc[N node.N](lenNode int, a analysis.A) (h H[N], s S, err error) {
	if node.BitsLen[N]() < lenNode {
		err = compact.ErrorInvalidLenNode
		return
	}

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

	if l > lenNode {
		err = compact.ErrorOverflow
		return
	}

	s.ChunkPos = 0
	h.sChunkPos = byte(lenNode - lenChunkPos)

	s.Value = byte(lenChunkPos)
	h.lsValue = h.sChunkPos - byte(lenValue)
	h.rsValue = byte(lenNode - lenValue)

	s.ChildrenStart = s.Value + byte(lenValue)
	h.lsChildrenStart = h.lsValue - byte(lenChildrenStart)
	h.rsChildrenStart = byte(lenNode - lenChildrenStart)

	s.ChildrenAmount = s.ChildrenStart + byte(lenChildrenStart)
	h.lsChildrenAmount = h.lsChildrenStart - byte(lenChildrenAmount)
	h.rsChildrenAmount = byte(lenNode - lenChildrenAmount)

	s.ChunkLen = byte(lenNode - lenChunkLen)
	h.sChunkLen = s.ChunkLen

	return
}
