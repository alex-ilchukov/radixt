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
	indices := make([]uint, len(nodes))
	parents := make(map[uint]uint)
	chunks := []string{}
	cml := uint(0)
	cma := uint(0)
	vm := uint(0)
	ca := make(map[uint]uint)

	for i, n := range nodes {
		index := n.Index
		indices[i] = index

		for c := n.ChildrenFirst; c <= n.ChildrenLast; c++ {
			parents[c] = index
		}

		cl := uint(0)
		if n.ChildrenFirst <= n.ChildrenLast {
			cl = n.ChildrenLast - n.ChildrenFirst + 1
		}
		if cma < cl {
			cma = cl
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
	n := make(map[uint]N)

	for i := range nodes {
		index := nodes[i].Index
		parent, has := parents[index]
		if has {
			nodes[i].Parent = parent
		} else {
			nodes[i].Root = true
		}

		nodes[i].ChunkPos = uint(strings.Index(c, nodes[i].Chunk))

		n[index] = nodes[i]
	}

	return A{C: c, Cml: cml, Cma: cma, Vm: vm, N: n, Ca: ca}
}

func chunk(t radixt.Tree, n uint) string {
	buf := []byte{}
	npos := uint(0)
	for {
		b, within := t.ByteAt(n, npos)
		if !within {
			break
		}

		buf = append(buf, b)
		npos++
	}

	return string(buf)
}

func nodes(t radixt.Tree) []N {
	size := t.Size()
	nodes := make([]N, size, size)

	for n := uint(0); n < size; n++ {
		v, has := t.Value(n)
		f, l := t.ChildrenRange(n)
		nodes[n] = N{
			Index:         n,
			Chunk:         chunk(t, n),
			Value:         v,
			HasValue:      has,
			ChildrenFirst: f,
			ChildrenLast:  l,
		}
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
