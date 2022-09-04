package header

import "github.com/alex-ilchukov/radixt/compact/internal/node"

const hlen = 8

// A8b represents a header in a compact implementation with all the bytes
// required for extraction of node's fields.
type A8b [hlen]byte

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

func head[N node.N, Header H](n N, h Header) uint {
	return node.Head(n, h[0])
}

func body[N node.N, Header H](n N, h Header, i int) uint {
	return node.Body(n, h[2*i+1], h[2*i+2])
}

func tail[N node.N, Header H](n N, h Header) uint {
	return node.Tail(n, h[hlen-1])
}

const (
	bodyValue = iota
	bodyStart
	bodyAmount
)

// Value takes node n with header h and returns value v of the node with
// boolean true flag, if the node has value, or default unsigned integer with
// boolean false otherwise.
func Value[N node.N, Header H](n N, h Header) (v uint, has bool) {
	v = body(n, h, bodyValue)
	if v == 0 {
		return
	}

	has = true
	v -= 1

	return
}

// ChunkRange takes node n with header h and returns low and high indices to
// select the node's chunk from string of all chunks combined.
func ChunkRange[N node.N, Header H](n N, h Header) (low, high uint) {
	low = head(n, h)
	high = low + tail(n, h)
	return
}

// ChildrenRange takes node n with its index i and header h, and returns first
// and last indices of children of the node, if the node has children, or 1 and
// 0 otherwise.
func ChildrenRange[N node.N, Header H](i uint, n N, h Header) (uint, uint) {
	amount := body(n, h, bodyAmount)
	if amount == 0 {
		return 1, 0
	}

	f := body(n, h, bodyStart) + i + 1
	return f, f + amount  - 1
}
