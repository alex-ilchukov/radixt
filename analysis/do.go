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

	l := t.Size()
	y := &yielder{t: t, a: A{N: make([]N, l, l)}, p: make([]*N, l, l)}
	pass.Do(t, y)
	y.cramChunks()

	return y.a
}

type yielder struct {
	t  radixt.Tree
	a  A
	p  []*N
	cl uint
}

func (y *yielder) Yield(i, n, tag uint) uint {
	y.processNode(i, n)
	y.processParent(i, n, tag)

	return n
}

func (y *yielder) processNode(i, n uint) {
	t := y.t

	chunk := t.Chunk(n)
	v, has := t.Value(n)
	y.a.N[n] = N{Index: i, Chunk: chunk, Value: v, HasValue: has}
	y.p[n] = &y.a.N[n]

	cl := uint(len(chunk))
	y.cl += cl
	if y.a.Cml < cl {
		y.a.Cml = cl
	}

	if y.a.Vm < v {
		y.a.Vm = v
	}
}

func (y *yielder) processParent(i, n, p uint) {
	if n == 0 {
		return
	}

	nodes := y.a.N
	nodes[n].Parent = nodes[p].Index

	if nodes[p].ChildrenHigh == 0 {
		nodes[p].ChildrenLow = i

		dclp := i - nodes[p].Index
		if y.a.Dclpm < dclp {
			y.a.Dclpm = dclp
		}
	}
	nodes[p].ChildrenHigh = i + 1

	ca := i + 1 - nodes[p].ChildrenLow
	if y.a.Cma < ca {
		y.a.Cma = ca
	}
}

func (y *yielder) cramChunks() {
	sort.SliceStable(y.p, func(i, j int) bool {
		pi := y.p[i].Chunk
		pj := y.p[j].Chunk
		li := len(pi)
		lj := len(pj)

		return li > lj || (li == lj && pi <= pj)
	})

	t := 0
	b := make([]byte, y.cl)
	for _, p := range y.p {
		s := []byte(p.Chunk)
		pos := bytes.Index(b, s)
		if pos == -1 {
			pos = t
			t += copy(b[t:], s)
		}
		p.ChunkPos = uint(pos)
	}

	y.a.C = string(b[:t])
}
