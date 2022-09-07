package strg

import (
	"math"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
)

const maxChunksLen = math.MaxUint16 - cstart

// New takes the provided tree t and tries to compactify it. In case of success
// it returns new, compactified representation of the tree and nil for error.
// In case of an error, it returns nil for tree and the error.
func New[NX N](t radixt.Tree) (Tree[NX], error) {
	a := analysis.Do(t)
	if len(a.C) > maxChunksLen {
		return "", compact.ErrorChunksOverflow
	}

	h, nf, err := header.Calc[uint32](8 * bytesLen[NX](), a)
	if err != nil {
		return "", err
	}

	noffset := cstart + len(a.C)
	bytes := make([]byte, noffset + bytesLen[NX]() * len(a.N))
	copy(bytes, h[:])

	bytes[hlen] = byte(noffset&0xFF)
	bytes[hlen+1] = byte(noffset>>8)

	copy(bytes[cstart:], a.C)

	for i, n := range a.N {
		o := noffset + int(i) * bytesLen[NX]()
		node := nf(n)
		bytes[o] = byte(node&0xFF)
		bytes[o+1] = byte(node>>8&0xFF)
		switch bytesLen[NX]() {
		case 3:
			bytes[o+2] = byte(node>>16)
		case 4:
			bytes[o+2] = byte(node>>16&0xFF)
			bytes[o+3] = byte(node>>24)
		}
	}

	return Tree[NX](bytes), nil
}

// MustCreate takes the provided tree t and tries to compactify it. In case of
// success it returns new, compactified representation of the tree. In case of
// an error, it panics.
func MustCreate[NX N](t radixt.Tree) Tree[NX] {
	result, err := New[NX](t)
	if err != nil {
		panic(err)
	}

	return result
}
