package header

import "github.com/alex-ilchukov/radixt/compact/internal/node"

// Routines, working with headers, assume the following order of node's fields
// in node bit string:
//
//   - head — chunk's position;
//   - body 1 — value (mogrified);
//   - body 2 — index of first child (mogrified);
//   - body 3 — amount of children;
//   - tail — chunk's length.
const (
	fieldChunkPos = iota
	fieldValue
	fieldChildrenStart
	fieldChildrenAmount
	fieldChunkLen
	fieldsAmount
)

// Len is amount of bytes, required to form the header
const Len = 2 + (fieldsAmount-2)*2

// A8b represents a header in a compact implementation with all the bytes
// required for extraction of node's fields.
type A8b [Len]byte

// H is type set of header types.
type H interface {
	A8b | ~string
}

func head[N node.N, Header H](n N, h Header) uint {
	return node.Head(n, h[0])
}

func body[N node.N, Header H](n N, h Header, i int) uint {
	return node.Body(n, h[2*i-1], h[2*i])
}

func tail[N node.N, Header H](n N, h Header) uint {
	return node.Tail(n, h[Len-1])
}

// Value takes node n with header h and returns value v of the node with
// boolean true flag, if the node has value, or default unsigned integer with
// boolean false otherwise.
func Value[N node.N, Header H](n N, h Header) (v uint, has bool) {
	v = body(n, h, fieldValue)
	if v == 0 {
		return
	}

	has = true
	v -= 1

	return
}

// ChunkLow takes node n with header h and returns low index to select the
// node's chunk from string of all chunks combined.
func ChunkLow[N node.N, Header H](n N, h Header) uint {
	return head(n, h)
}

// ChunkLen takes node n with header h and returns length of the node's chunk.
func ChunkLen[N node.N, Header H](n N, h Header) uint {
	return tail(n, h)
}

// ChunkRange takes node n with header h and returns low and high indices to
// select the node's chunk from string of all chunks combined.
func ChunkRange[N node.N, Header H](n N, h Header) (low, high uint) {
	low = ChunkLow(n, h)
	high = low + ChunkLen(n, h)
	return
}

// ChildrenRange takes node n with its index i and header h, and returns low
// and high indices of children of the node, if the node has children, or
// default unsigned integer values otherwise.
func ChildrenRange[N node.N, Header H](i uint, n N, h Header) (lo, hi uint) {
	amount := body(n, h, fieldChildrenAmount)
	if amount > 0 {
		lo = body(n, h, fieldChildrenStart) + i + 1
		hi = lo + amount
	}

	return
}
