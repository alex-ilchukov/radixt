package header

import "github.com/alex-ilchukov/radixt/compact/internal/node"

// A8b represents a header in a compact implementation with all the bytes
// required for extraction of node's fields.
type A8b [8]byte

// H is type set of header types. Routines, working with H, assume the
// following order of node's fields in node bit string:
//
//  * head — chunk's position;
//  * body 0 — value (mogrified);
//  * body 1 — index of first child (mogrified);
//  * body 2 — amount of children;
//  * tail — chunk's length.
type H interface {
	A8b | ~string
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
