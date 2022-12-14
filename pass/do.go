package pass

import (
	"sort"

	"github.com/alex-ilchukov/radixt"
)

// Do iterates over nodes of the provided radix tree t in breadth-wide search
// manner and yields to y just once for every node. If the tree is not empty,
// it starts with zeroth node (root, that is). It also provides user-filled
// tags with the nodes, gathered from their parents. As the root has no parent,
// it supposes it has zero tag. The function does nothing if either of t and y
// is nil.
//
// Additionally, Do sorts children of a node by ascending order of first bytes
// of their chunks. That would allow to implement lookup in more efficien way.
//
// Example. For the following instance of grown [sapling.Tree]
//
//	                             0: ""
//	                     /                  \
//	            6: "auth"                    7: "content-"
//	               / \                      /        |       \
//	      4: "or" 5: "entication"   8: "type" 9: "length" 10: "disposition"
//	         |
//	      3: "i"
//	     /      \
//	1: "ty"     2: "zation"
//
// the nodes would be enumerated and yielded in the following order: 0, 6, 7,
// 5, 4, 10, 9, 8, 3, 1, 2.
func Do(t radixt.Tree, y Yielder) {
	if t == nil || y == nil || t.Size() == 0 {
		return
	}

	type e struct {
		n   uint
		tag uint
	}

	children := []uint{}
	for i, q := uint(0), []e{{}}; len(q) > 0; i, q = i+1, q[1:] {
		a := q[0]
		ctag := y.Yield(i, a.n, a.tag)

		t.EachChild(a.n, func(c uint) bool {
			children = append(children, c)
			return false
		})

		sort.Slice(children, func(i, j int) bool {
			ci := t.Chunk(children[i])
			cj := t.Chunk(children[j])
			return ci[0] < cj[0]
		})

		for _, c := range children {
			q = append(q, e{n: c, tag: ctag})
		}

		children = children[:0]
	}
}
