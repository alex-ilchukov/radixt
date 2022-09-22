package analysis

import (
	"bytes"
	"sort"

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

		cl = uint(len(n.Chunk))
		if cml < cl {
			cml = cl
		}

		if vm < n.Value {
			vm = n.Value
		}
	}

	c := cramChunks(nodes)

	return A{C: c, Cml: cml, Cma: cma, Dclpm: dcplm, Vm: vm, N: nodes}
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

func cramChunks(nodes []N) string {
	l := len(nodes)
	pods := make([]*N, l, l)
	t := 0
	for i := range nodes {
		pods[i] = &nodes[i]
		t += len(nodes[i].Chunk)
	}
	b := make([]byte, t)

	sort.SliceStable(pods, func(i, j int) bool {
		pi := pods[i].Chunk
		pj := pods[j].Chunk
		li := len(pi)
		lj := len(pj)

		return li > lj || (li == lj && pi <= pj)
	})

	t = 0
	for _, p := range pods {
		s := []byte(p.Chunk)
		pos := bytes.Index(b, s)
		if pos == -1 {
			pos = t
			t += copy(b[t:], s)
		}
		p.ChunkPos = uint(pos)
	}

	return string(b[:t])
}
