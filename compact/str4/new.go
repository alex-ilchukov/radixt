package str4

import (
	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/internal/header"
)

// New takes the provided tree t and tries to compactify it. In case of success
// it returns new, compactified representation of the tree and nil for error.
// In case of an error, it returns nil for tree and the error.
func New(t radixt.Tree) (Tree, error) {
	a := analysis.Do[analysis.Firstless](t)
	size := len(a.N)
	if size > maskSize {
		return "", compact.ErrorNodesOverflow
	}

	h, nf, err := header.Calc[uint32](nodeLen * 8, a)
	if err != nil {
		return "", err
	}

	bytes := make([]byte, cfstart+(nodeLen+1)*size+len(a.C))
	copy(bytes, h[:])

	emptyRoot := 0
	if size > 0 && a.N[0].ChunkEmpty {
		emptyRoot = 0x80
	}

	bytes[h21] = byte(size & 0xFF)
	bytes[h22] = byte(size >> 8 | emptyRoot)

	for _, n := range a.N {
		node := nf(n)
		i := cfstart + size + int(n.Index)*nodeLen
		bytes[i] = byte(node & 0xFF)
		bytes[i+1] = byte(node >> 8 & 0xFF)
		bytes[i+2] = byte(node >> 16 & 0xFF)
		bytes[i+3] = byte(node >> 24)
		bytes[cfstart+n.Index] = n.ChunkFirst
	}

	copy(bytes[cfstart+(nodeLen+1)*size:], a.C)

	return Tree(bytes), nil
}

// MustCreate takes the provided tree t and tries to compactify it. In case of
// success it returns new, compactified representation of the tree. In case of
// an error, it panics.
func MustCreate(t radixt.Tree) Tree {
	result, err := New(t)
	if err != nil {
		panic(err)
	}

	return result
}
