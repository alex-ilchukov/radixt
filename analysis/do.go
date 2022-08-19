package analysis

import (
	"bytes"
	"math"
	"sort"
	"strings"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
)

func Do(t radixt.Tree) A {
	if t == nil {
		t = null.Tree
	}

	nodes, nt := nodes(t)
	indices := make([]int, len(nodes))
	parents := make(map[int]int)
	prefixes := []string{}
	pml := 0
	ca := make(map[int]int)

	for i, n := range nodes {
		index := n.Index
		indices[i] = index

		for _, c := range n.Children {
			parents[c] = index
		}

		ca[len(n.Children)] += 1
		prefixes = append(prefixes, n.Pref)

		if pml < len(n.Pref) {
			pml = len(n.Pref)
		}
	}

	nonnode := calcNonNode(indices)
	p := cramPrefixes(prefixes)
	n := make(map[int]N)

	for i := range nodes {
		index := nodes[i].Index
		parent, has := parents[index]
		if has {
			nodes[i].Parent = parent
		} else {
			nodes[i].Parent = nonnode
		}

		nodes[i].PrefPos = strings.Index(p, nodes[i].Pref)

		n[index] = nodes[i]
	}

	return A{P: p, Pml: pml, N: n, Nt: nt, Ca: ca}
}

func nodes(t radixt.Tree) ([]N, map[int]int) {
	nodes := []N{}
	nt := map[int]int{}
	queue := []int{}
	root := t.Root()
	if t.Has(root) {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		r := N{
			Index:    n,
			Pref:     t.NodePref(n),
			String:   t.NodeString(n),
			Mark:     t.NodeMark(n),
			Children: []int{},
		}

		t.NodeEachChild(n, func(c int) bool {
			r.Children = append(r.Children, c)
			queue = append(queue, c)
			return false
		})

		sort.Ints(r.Children)
		nodes = append(nodes, r)
		nt[n] = len(nt)
	}

	return nodes, nt
}

func calcNonNode(indices []int) int {
	sort.Ints(indices)
	l := len(indices)
	r := math.MinInt
	for i := 0; i < l; i++ {
		if r < indices[i] {
			return r
		}

		r++
	}

	return r
}

func cramPrefixes(prefixes []string) string {
	sort.SliceStable(prefixes, func(i, j int) bool {
		pi := prefixes[i]
		pj := prefixes[j]
		li := len(pi)
		lj := len(pj)

		return li > lj || (li == lj && pi <= pj)
	})

	l := len(prefixes)
	t := 0
	byteSlices := make([][]byte, l)
	for i, p := range prefixes {
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
