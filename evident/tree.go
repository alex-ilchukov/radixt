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

// EachChild calls function e just once for every child of node n in ascending
// order, if the tree has the node, until the function returns boolean truth.
// The method does nothing if the tree does not have the node.
func (t Tree) EachChild(n uint, e func(uint) bool) {
	a, m, q := t.grind(n)
	if a == nil {
		return
	}

	keys := a.keys()
	key := keys[m]
	if a[key] == nil {
		return
	}

	low := n + uint(len(a)) - m
	for _, c := range q.a {
		low += uint(len(c))
	}

	for _, k := range keys[:m] {
		low += uint(len(a[k]))
	}

	high := low + uint(len(a[key]))

	for c := low; c < high; c++ {
		if e(c) {
			return
		}
	}
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

// Hoard returns amount of bytes, taken by the implementation t, with an
// interpretation hint. It checks the following cases.
//
//  1. Nil tree. The result is zero and [radixt.HoardExactly].
//  2. Empty tree. The result is amount of bytes, required for internal hmap
//     struct instance with no buckets, and [radixt.HoardExactly].
//  3. Non-empty tree. The result is an estimation with [radixt.HoardAtLeast]
//     as interpretation hint.
func (t Tree) Hoard() (uint, uint) {
	if t == nil {
		return 0, radixt.HoardExactly
	}

	if len(t) == 0 {
		return 48, radixt.HoardExactly
	}

	amount := uint(0)
	s := []Tree{t}
	for len(s) > 0 {
		l := len(s) - 1
		a := s[l]
		s = s[:l]
		amount += 48 + // size of instance of hmap struct
			128 // amount of bytes in at least one bucket of hmap

		for k, c := range a {
			if c != nil {
				s = append(s, c)
			}

			amount += 16 + // size of string header
				uint(len(k))
		}
	}

	return amount, radixt.HoardAtLeast
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

var (
	_ radixt.Tree    = Tree(nil)
	_ radixt.Hoarder = Tree(nil)
)
