package analysis

import (
	"bytes"
	"sort"
	"strings"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
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
	ca := make(map[uint]uint)

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

		ca[cl] += 1
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

	return A{C: c, Cml: cml, Cma: cma, Dclpm: dcplm, Vm: vm, N: n, Ca: ca}
}

func nodes(t radixt.Tree) []N {
	size := t.Size()
	nodes := make([]N, size, size)
	if size == 0 {
		return nodes
	}

	type e struct {
		n uint
		p uint
	}

	for i, q := uint(0), []e{{}}; len(q) > 0; i++ {
		a := q[0]
		q = q[1:]

		n := a.n

		t.EachChild(n, func(c uint) bool {
			q = append(q, e{n: c, p: n})
			return false
		})

		v, has := t.Value(n)
		nodes[n] = N{
			Index:    i,
			Chunk:    t.Chunk(n),
			Value:    v,
			HasValue: has,
		}

		if n == 0 {
			continue
		}

		nodes[n].Parent = nodes[a.p].Index

		if nodes[a.p].ChildrenHigh == 0 {
			nodes[a.p].ChildrenLow = i
		}
		nodes[a.p].ChildrenHigh = i + 1
	}

	return nodes
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
