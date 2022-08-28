package evident

import (
	"sort"
	"strconv"
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

	value := extractValue(key)
	if value != "" {
		v64, err := strconv.ParseUint(value, 0, 0)
		if err != nil {
			panic(err)
		}

		v = uint(v64)
		has = true
	}

	return
}

// EachChild calls func e for every child of node n, if the tree has the node,
// until the func returns boolean true. The order of going over the children is
// fixed for every node, but may not coincide with any natural order.
func (t Tree) EachChild(n uint, e func(uint) bool) {
	for c, l := t.childrenRange(n); c <= l; c++ {
		if e(c) {
			return
		}
	}
}

// ByteAt returns default byte value and boolean false, if npos is outside of
// chunk of the node n, or byte of the chunk at npos and boolean true
// otherwise.
func (t Tree) ByteAt(n uint, npos uint) (b byte, within bool) {
	key := t.key(n)
	if key == "" {
		return
	}

	chunk := extractChunk(key)
	if npos < uint(len(chunk)) {
		b = chunk[npos]
		within = true
	}

	return
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

func (t Tree) childrenRange(n uint) (uint, uint) {
	a, m, q := t.grind(n)
	if a == nil {
		return 1, 0
	}

	keys := a.keys()
	key := keys[m]
	if a[key] == nil {
		return 1, 0
	}

	f := n + uint(len(a)) - m
	for _, c := range q.a {
		f += uint(len(c))
	}

	for _, k := range keys[:m] {
		f += uint(len(a[k]))
	}

	return f, f + uint(len(a[key])) - 1
}