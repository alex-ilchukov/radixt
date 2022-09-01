package generic_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/generic"
)

var (
	empty = generic.New()

	atree = generic.New(
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
	{tree: empty, n: 0, result1: 0, result2: false},
	{tree: empty, n: 1, result1: 0, result2: false},
	{tree: empty, n: 100, result1: 0, result2: false},
	{tree: atree, n: 0, result1: 0, result2: false},
	{tree: atree, n: 1, result1: 4, result2: true},
	{tree: atree, n: 2, result1: 0, result2: false},
	{tree: atree, n: 3, result1: 2, result2: true},
	{tree: atree, n: 4, result1: 3, result2: true},
	{tree: atree, n: 5, result1: 5, result2: true},
	{tree: atree, n: 6, result1: 6, result2: true},
	{tree: atree, n: 7, result1: 7, result2: true},
	{tree: atree, n: 8, result1: 0, result2: false},
	{tree: atree, n: 9, result1: 0, result2: true},
	{tree: atree, n: 10, result1: 1, result2: true},
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
	{tree: empty, n: 0, result: ""},
	{tree: empty, n: 1, result: ""},
	{tree: empty, n: 100, result: ""},
	{tree: atree, n: 0, result: ""},
	{tree: atree, n: 1, result: "auth"},
	{tree: atree, n: 2, result: "content-"},
	{tree: atree, n: 3, result: "or"},
	{tree: atree, n: 4, result: "entication"},
	{tree: atree, n: 5, result: "type"},
	{tree: atree, n: 6, result: "length"},
	{tree: atree, n: 7, result: "disposition"},
	{tree: atree, n: 8, result: "i"},
	{tree: atree, n: 9, result: "ty"},
	{tree: atree, n: 10, result: "zation"},
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

var treeChildrenRangeTests = []struct {
	tree    radixt.Tree
	n       uint
	result1 uint
	result2 uint
}{
	{tree: empty, n: 0, result1: 1, result2: 0},
	{tree: empty, n: 1, result1: 1, result2: 0},
	{tree: empty, n: 100, result1: 1, result2: 0},
	{tree: atree, n: 0, result1: 1, result2: 2},
	{tree: atree, n: 1, result1: 3, result2: 4},
	{tree: atree, n: 2, result1: 5, result2: 7},
	{tree: atree, n: 3, result1: 8, result2: 8},
	{tree: atree, n: 4, result1: 1, result2: 0},
	{tree: atree, n: 5, result1: 1, result2: 0},
	{tree: atree, n: 6, result1: 1, result2: 0},
	{tree: atree, n: 7, result1: 1, result2: 0},
	{tree: atree, n: 8, result1: 9, result2: 10},
	{tree: atree, n: 9, result1: 1, result2: 0},
	{tree: atree, n: 10, result1: 1, result2: 0},
	{tree: atree, n: 100, result1: 1, result2: 0},
}

const testTreeChildrenRangeError = "Tree Children Range Test %d: got %d " +
	"and %d for first and last indices of children of node %d (should " +
	"be %d and %d)"

func TestTreeChildrenRange(t *testing.T) {
	for i, tt := range treeChildrenRangeTests {
		result1, result2 := tt.tree.ChildrenRange(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTreeChildrenRangeError,
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
