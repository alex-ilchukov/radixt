package analysis

import (
	"bytes"
	"sort"
	"strings"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
	"github.com/alex-ilchukov/radixt/pass"
)

// Do analyzes radix tree t and returns result of the analysis. It guarantees
// to return the same result for the same tree with the same order of subfield
// slices. It is safe to invoke the function concurrently for the same tree or
// for different trees.
func Do(t radixt.Tree) A {
	if t == nil {
		t = null.Tree
	}

	nodes := nodes(t)
	chunks := []string{}
	cml := uint(0)
	cma := uint(0)
	dcplm := uint(0)
	vm := uint(0)

	for _, n := range nodes {
		cl := n.ChildrenHigh - n.ChildrenLow
		if cma < cl {
			cma = cl
		}

		if 0 < n.ChildrenHigh {
			dclp := n.ChildrenLow - n.Index
			if dcplm < dclp {
				dcplm = dclp
			}
		}

		chunks = append(chunks, n.Chunk)

		cl = uint(len(n.Chunk))
		if cml < cl {
			cml = cl
		}

		if vm < n.Value {
			vm = n.Value
		}
	}

	c := cramChunks(chunks)
	n := make(map[uint]N, len(nodes))

	for i, node := range nodes {
		nodes[i].ChunkPos = uint(strings.Index(c, node.Chunk))
		n[uint(i)] = nodes[i]
	}

	return A{C: c, Cml: cml, Cma: cma, Dclpm: dcplm, Vm: vm, N: n}
}

type yielder struct {
	t     radixt.Tree
	nodes []N
}

func (y *yielder) Yield(i, n, tag uint) uint {
	v, has := y.t.Value(n)
	y.nodes[n] = N{Index: i, Chunk: y.t.Chunk(n), Value: v, HasValue: has}
	y.processParent(i, n, tag)

	return n
}

func (y *yielder) processParent(i, n, p uint) {
	if n == 0 {
		return
	}

	nodes := y.nodes
	nodes[n].Parent = nodes[p].Index

	if nodes[p].ChildrenHigh == 0 {
		nodes[p].ChildrenLow = i
	}
	nodes[p].ChildrenHigh = i + 1
}

func nodes(t radixt.Tree) []N {
	l := t.Size()
	y := &yielder{t: t, nodes: make([]N, l, l)}
	pass.Do(t, y)
	return y.nodes
}

func cramChunks(chunks []string) string {
	sort.SliceStable(chunks, func(i, j int) bool {
		pi := chunks[i]
		pj := chunks[j]
		li := len(pi)
		lj := len(pj)

		return li > lj || (li == lj && pi <= pj)
	})

	l := len(chunks)
	t := 0
	byteSlices := make([][]byte, l)
	for i, p := range chunks {
		byteSlices[i] = []byte(p)
		t += len(p)
	}

	b := make([]byte, t)
	t = 0
	for i := 0; i < l; i++ {
		s := byteSlices[i]
		if !bytes.Contains(b, s) {
			copy(b[t:], s)
			t += len(s)
		}
	}

	return string(b[:t])
}
