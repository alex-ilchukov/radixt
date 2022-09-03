package header

import "github.com/alex-ilchukov/radixt/compact/internal/node"

// H represents a header in a compact implementation of radix tree, which packs
// its nodes into uint32 or uint64 (see [compact/internal/node.N]). It assumes,
// that bits in the uintXX value represent the following unsigned intgers,
// going from lowest at the top of the diagram to highest ones at the bottom:
//
//    ------------------------------------------------------------------------
//  / — \
//  | — - Position of node's chunk in the string of all chunks concatenated
//  | — - (see [analytics.A.C] string) — the lowest bits, which have been
//  | — - extracted via node.Head(n, H.sChunkPos).
//  | — /
//  | ------------------------------------------------------------------------
//  | — \
//  | — - Incremented node's value, if the node has value, or zero otherwise.
//  | — - The value has been extracted via node.Body(n, H.lsValue, H.rsValue).
//  | — /
//  u ------------------------------------------------------------------------
//  i — \
//  n — - Decremented difference between minimal child's index and the node's
//  t — - index (see [analytics.A.Dcfpm] value), if the node has children, or
//  X — - zero otherwise. The value has been extracted via
//  X — - node.Body(n, H.lsChildrenStart, H.rsChildrenStart).
//  | — /
//  | ------------------------------------------------------------------------
//  | — \
//  | — - Amount of children, which has been extracted via
//  | — - node.Body(n, H.lsChildrenAmount, H.rsChildrenAmount).
//  | — /
//  | ------------------------------------------------------------------------
//  | — \
//  | — - Length of node's chunk — the highest bits, which have been extracted
//  | — - via node.Tail(n, H.sChunkLen).
//  \ — /
//    ------------------------------------------------------------------------
type H[N node.N] struct {
	sChunkPos        byte
	lsValue          byte
	rsValue          byte
	lsChildrenStart  byte
	rsChildrenStart  byte
	lsChildrenAmount byte
	rsChildrenAmount byte
	sChunkLen        byte
}

// Value returns value v of node n with boolean true flag, if the node has
// value, or default unsigned integer with boolean false otherwise.
func (h H[N]) Value(n N) (v uint, has bool) {
	v = node.Body(n, h.lsValue, h.rsValue)
	if v == 0 {
		return
	}

	has = true
	v -= 1

	return
}

// ChunkPos returns position of chunk of node n in the string of all chunks
// concatenated.
func (h H[N]) ChunkPos(n N) uint {
	return node.Head(n, h.sChunkPos)
}

// ChunkLen returns length of chunk of node n.
func (h H[N]) ChunkLen(n N) uint {
	return node.Tail(n, h.sChunkLen)
}

// ChildrenRange returns first and last indices of children of node n with
// index i, if the node has children, or 1 and 0 otherwise.
func (h H[N]) ChildrenRange(i uint, n N) (uint, uint) {
	amount := node.Body(n, h.lsChildrenAmount, h.rsChildrenAmount)
	if amount == 0 {
		return 1, 0
	}

	f := node.Body(n, h.lsChildrenStart, h.rsChildrenStart) + i + 1
	return f, f + amount  - 1
}
