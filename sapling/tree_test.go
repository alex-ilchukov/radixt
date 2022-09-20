package sapling_test

import (
	"strconv"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/sapling"
)

var (
	blank = (*sapling.Tree)(nil)
	empty = sapling.New()

	atree = sapling.New(
		"authority",
		"authorization",
		"author",
		"authentication",
		"auth",
		"content-type",
		"content-length",
		"content-disposition",
	)
)

var treeSizeTests = []struct {
	tree   radixt.Tree
	result uint
}{
	{tree: blank, result: 0},
	{tree: empty, result: 0},
	{tree: atree, result: 11},
}

const testTreeSizeError = "Tree Size Test %d: got %d for size (should be %d)"

func TestTreeSize(t *testing.T) {
	for i, tt := range treeSizeTests {
		result := tt.tree.Size()
		if result != tt.result {
			t.Errorf(testTreeSizeError, i, result, tt.result)
		}
	}
}

var treeValueTests = []struct {
	tree    radixt.Tree
	n       uint
	result1 uint
	result2 bool
}{
	{tree: blank, n: 0, result1: 0, result2: false},
	{tree: blank, n: 1, result1: 0, result2: false},
	{tree: blank, n: 100, result1: 0, result2: false},
	{tree: empty, n: 0, result1: 0, result2: false},
	{tree: empty, n: 1, result1: 0, result2: false},
	{tree: empty, n: 100, result1: 0, result2: false},
	{tree: atree, n: 0, result1: 0, result2: false},
	{tree: atree, n: 1, result1: 0, result2: true},
	{tree: atree, n: 2, result1: 1, result2: true},
	{tree: atree, n: 3, result1: 0, result2: false},
	{tree: atree, n: 4, result1: 2, result2: true},
	{tree: atree, n: 5, result1: 3, result2: true},
	{tree: atree, n: 6, result1: 4, result2: true},
	{tree: atree, n: 7, result1: 0, result2: false},
	{tree: atree, n: 8, result1: 5, result2: true},
	{tree: atree, n: 9, result1: 6, result2: true},
	{tree: atree, n: 10, result1: 7, result2: true},
	{tree: atree, n: 100, result1: 0, result2: false},
}

const testTreeValueError = "Tree Value Test %d: got %d and %t for value of " +
	"node %d (should be %d and %t)"

func TestTreeValue(t *testing.T) {
	for i, tt := range treeValueTests {
		result1, result2 := tt.tree.Value(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTreeValueError,
				i,
				result1,
				result2,
				tt.n,
				tt.result1,
				tt.result2,
			)
		}
	}
}

var treeChunkTests = []struct {
	tree   radixt.Tree
	n      uint
	result string
}{
	{tree: blank, n: 0, result: ""},
	{tree: blank, n: 1, result: ""},
	{tree: blank, n: 100, result: ""},
	{tree: empty, n: 0, result: ""},
	{tree: empty, n: 1, result: ""},
	{tree: empty, n: 100, result: ""},
	{tree: atree, n: 0, result: ""},
	{tree: atree, n: 1, result: "ty"},
	{tree: atree, n: 2, result: "zation"},
	{tree: atree, n: 3, result: "i"},
	{tree: atree, n: 4, result: "or"},
	{tree: atree, n: 5, result: "entication"},
	{tree: atree, n: 6, result: "auth"},
	{tree: atree, n: 7, result: "content-"},
	{tree: atree, n: 8, result: "type"},
	{tree: atree, n: 9, result: "length"},
	{tree: atree, n: 10, result: "disposition"},
	{tree: atree, n: 100, result: ""},
}

const testTreeChunkError = "Tree Chunk Test %d: got '%s' for chunk of node " +
	"%d (should be '%s')"

func TestTreeChunk(t *testing.T) {
	for i, tt := range treeChunkTests {
		result := tt.tree.Chunk(tt.n)
		if result != tt.result {
			t.Errorf(
				testTreeChunkError,
				i,
				result,
				tt.n,
				tt.result,
			)
		}
	}
}

func eachChild(tree radixt.Tree, n uint) (indices string) {
	tree.EachChild(n, func(c uint) bool {
		if len(indices) > 0 {
			indices += ", "
		}

		indices += strconv.FormatUint(uint64(c), 10)

		return false
	})

	return
}

func eachFirstChild(tree radixt.Tree, n uint) (indices string) {
	tree.EachChild(n, func(c uint) bool {
		indices = strconv.FormatUint(uint64(c), 10)

		return true
	})

	return
}

var treeEachChildTests = []struct {
	tree    radixt.Tree
	n       uint
	f       func(radixt.Tree, uint) string
	indices string
}{
	{tree: blank, n: 0, f: eachChild, indices: ""},
	{tree: blank, n: 1, f: eachChild, indices: ""},
	{tree: blank, n: 100, f: eachChild, indices: ""},
	{tree: blank, n: 0, f: eachFirstChild, indices: ""},
	{tree: blank, n: 1, f: eachFirstChild, indices: ""},
	{tree: blank, n: 100, f: eachFirstChild, indices: ""},
	{tree: empty, n: 0, f: eachChild, indices: ""},
	{tree: empty, n: 1, f: eachChild, indices: ""},
	{tree: empty, n: 100, f: eachChild, indices: ""},
	{tree: empty, n: 0, f: eachFirstChild, indices: ""},
	{tree: empty, n: 1, f: eachFirstChild, indices: ""},
	{tree: empty, n: 100, f: eachFirstChild, indices: ""},
	{tree: atree, n: 0, f: eachChild, indices: "6, 7"},
	{tree: atree, n: 1, f: eachChild, indices: ""},
	{tree: atree, n: 2, f: eachChild, indices: ""},
	{tree: atree, n: 3, f: eachChild, indices: "1, 2"},
	{tree: atree, n: 4, f: eachChild, indices: "3"},
	{tree: atree, n: 5, f: eachChild, indices: ""},
	{tree: atree, n: 6, f: eachChild, indices: "4, 5"},
	{tree: atree, n: 7, f: eachChild, indices: "8, 9, 10"},
	{tree: atree, n: 8, f: eachChild, indices: ""},
	{tree: atree, n: 9, f: eachChild, indices: ""},
	{tree: atree, n: 10, f: eachChild, indices: ""},
	{tree: atree, n: 100, f: eachChild, indices: ""},
	{tree: atree, n: 0, f: eachFirstChild, indices: "6"},
	{tree: atree, n: 1, f: eachFirstChild, indices: ""},
	{tree: atree, n: 2, f: eachFirstChild, indices: ""},
	{tree: atree, n: 3, f: eachFirstChild, indices: "1"},
	{tree: atree, n: 4, f: eachFirstChild, indices: "3"},
	{tree: atree, n: 5, f: eachFirstChild, indices: ""},
	{tree: atree, n: 6, f: eachFirstChild, indices: "4"},
	{tree: atree, n: 7, f: eachFirstChild, indices: "8"},
	{tree: atree, n: 8, f: eachFirstChild, indices: ""},
	{tree: atree, n: 9, f: eachFirstChild, indices: ""},
	{tree: atree, n: 10, f: eachFirstChild, indices: ""},
	{tree: atree, n: 100, f: eachFirstChild, indices: ""},
}

const testTreeEachChildError = "Tree Each Child Test %d: got %s as result " +
	"indices (should be %s)"

func TestTreeEachChild(t *testing.T) {
	for i, tt := range treeEachChildTests {
		indices := tt.f(tt.tree, tt.n)
		if indices != tt.indices {
			t.Errorf(
				testTreeEachChildError,
				i,
				indices,
				tt.indices,
			)
		}
	}
}

var treeHoardTests = []struct {
	tree    radixt.Hoarder
	result1 uint
	result2 uint
}{
	{tree: empty, result1: 24, result2: radixt.HoardExactly},
	{
		tree: atree,
		result1: 24 +
			56*11 +
			0 + 2*8 +
			2 + 0*8 +
			6 + 0*8 +
			1 + 2*8 +
			2 + 1*8 +
			10 + 0*8 +
			4 + 2*8 +
			8 + 3*8 +
			4 + 0*8 +
			6 + 0*8 +
			11 + 0*8,
		result2: radixt.HoardExactly,
	},
}

const testTreeHoardError = "Tree Hoard Test %d: got %d and %d (should be %d " +
	"and %d)"

func TestTreeHoard(t *testing.T) {
	for i, tt := range treeHoardTests {
		result1, result2 := tt.tree.Hoard()
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTreeHoardError,
				i,
				result1,
				result2,
				tt.result1,
				tt.result2,
			)
		}
	}
}

var treeGrowTests = []struct {
	tree     *sapling.Tree
	s        string
	v        uint
	size     uint
	n        uint
	value    uint
	hasValue bool
	chunk    string
	children string
}{
	{ // Creation of root node
		tree:     sapling.New(),
		s:        "authority",
		v:        100,
		size:     1,
		n:        0,
		value:    100,
		hasValue: true,
		chunk:    "authority",
		children: "",
	},
	{ // String is not found, but it has common prefix within a node
		tree:     sapling.New("authority"),
		s:        "authorization",
		v:        200,
		size:     3,
		n:        0,
		value:    0,
		hasValue: false,
		chunk:    "authori",
		children: "1, 2",
	},
	{ // String is not found, but it has common prefix til a node
		tree:     sapling.New("auth"),
		s:        "authorization",
		v:        200,
		size:     2,
		n:        0,
		value:    0,
		hasValue: true,
		chunk:    "auth",
		children: "1",
	},
	{ // String is found within a node
		tree:     sapling.New("authorization"),
		s:        "auth",
		v:        200,
		size:     2,
		n:        0,
		value:    200,
		hasValue: true,
		chunk:    "auth",
		children: "1",
	},
	{ // String is found til a node
		tree:     sapling.New("authorization"),
		s:        "authorization",
		v:        200,
		size:     1,
		n:        0,
		value:    200,
		hasValue: true,
		chunk:    "authorization",
		children: "",
	},
}

const testTreeGrowError = "Tree Grow Test %d: grown the tree with '%s' and " +
	"value %d, got tree.Size() = %d, tree.Value(%d) = (%d, %t), " +
	"tree.Chunk(%d) = '%s', and '%s' for indices of children of %d node " +
	"(should be %d, (%d, %t), '%s', and '%s')"

func TestTreeGrow(t *testing.T) {
	for i, tt := range treeGrowTests {
		tree := tt.tree
		tree.Grow(tt.s, tt.v)
		size := tree.Size()
		value, hasValue := tree.Value(tt.n)
		chunk := tree.Chunk(tt.n)
		children := eachChild(tree, tt.n)

		e := size != tt.size ||
			value != tt.value ||
			hasValue != tt.hasValue ||
			chunk != tt.chunk ||
			children != tt.children

		if e {
			t.Errorf(
				testTreeGrowError,
				i,
				tt.s,
				tt.v,
				size,
				tt.n,
				value,
				hasValue,
				tt.n,
				chunk,
				children,
				tt.n,
				tt.size,
				tt.value,
				tt.hasValue,
				tt.chunk,
				tt.children,
			)
		}
	}
}
