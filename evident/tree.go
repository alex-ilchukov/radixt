package evident

import (
	"sort"

	"github.com/alex-ilchukov/radixt"
)

// Tree is an implementation of [radixt.Tree] interface, based on maps.
type Tree map[string]Tree

// Size returns amount of nodes in the tree.
func (t Tree) Size() uint {
	size := uint(0)
	if t == nil {
		return size
	}

	s := newStack(t)
	for s.populated() {
		size += uint(len(s.popAndPushChildren()))
	}

	return size
}

// Value returns value v of node n with boolean true flag, if the tree has the
// node and the node has value, or default unsigned integer with boolean false
// otherwise.
func (t Tree) Value(n uint) (v uint, has bool) {
	key := t.key(n)
	if key == "" {
		return
	}

	return extractValue(key)
}

// Chunk returns chunk of node n, if the tree has the node, or empty string
// otherwise.
func (t Tree) Chunk(n uint) string {
	key := t.key(n)
	if key == "" {
		return ""
	}

	return extractChunk(key)
}

// ChildrenRange returns low and high indices of children of node n, if the
// tree has the node and the node has children, or default unsigned integers
// otherwise.
func (t Tree) ChildrenRange(n uint) (low, high uint) {
	a, m, q := t.grind(n)
	if a == nil {
		return
	}

	keys := a.keys()
	key := keys[m]
	if a[key] == nil {
		return
	}

	low = n + uint(len(a)) - m
	for _, c := range q.a {
		low += uint(len(c))
	}

	for _, k := range keys[:m] {
		low += uint(len(a[k]))
	}

	high = low + uint(len(a[key]))

	return
}

// Eq returns true, if the provided tree u has the same structure, node chunks,
// and node values as the original tree t. It supposes that empty and nil trees
// are equal.
func (t Tree) Eq(u radixt.Tree) bool {
	o := New(u)
	if len(t) == 0 && len(o) == 0 {
		return true
	}

	type se struct {
		tv Tree
		ov Tree
	}

	s := []se{{tv: t, ov: o}}
	for len(s) > 0 {
		l := len(s) - 1
		e := s[l]
		s = s[:l]
		tv := e.tv
		ov := e.ov

		if len(tv) != len(ov) {
			return false
		}

		for k, v := range tv {
			u, has := ov[k]
			if !has {
				return false
			}

			bv := len(u) > 0
			bu := len(v) > 0

			if bv != bu {
				return false
			}

			if bv && bu {
				s = append(s, se{tv: v, ov: u})
			}
		}
	}

	return true
}

func (t Tree) keys() []string {
	result := make([]string, len(t), len(t))
	i := 0
	for k := range t {
		result[i] = k
		i++
	}

	sort.SliceStable(result, func(i, j int) bool {
		return extractChunk(result[i]) < extractChunk(result[j])
	})

	return result
}

func (t Tree) grind(n uint) (a Tree, m uint, q *queue) {
	if t == nil {
		return
	}

	q = newQueue(t)
	m = n
	for q.populated() {
		a = q.pop()
		l := uint(len(a))
		if m < l {
			return
		}

		m -= l
		q.pushChildren(a)
	}

	a = nil
	return
}

func (t Tree) key(n uint) string {
	a, m, _ := t.grind(n)
	if a == nil {
		return ""
	}

	return a.keys()[m]
}

var _ radixt.Tree = Tree(nil)
