package null_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
)

var treeSizeTests = []struct {
	tree   radixt.Tree
	result uint
}{
	{tree: null.Tree, result: 0},
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
	{tree: null.Tree, n: 0, result1: 0, result2: false},
	{tree: null.Tree, n: 1, result1: 0, result2: false},
	{tree: null.Tree, n: 100, result1: 0, result2: false},
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
	{tree: null.Tree, n: 0, result: ""},
	{tree: null.Tree, n: 1, result: ""},
	{tree: null.Tree, n: 100, result: ""},
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
	{tree: null.Tree, n: 0, result1: 0, result2: 0},
	{tree: null.Tree, n: 1, result1: 0, result2: 0},
	{tree: null.Tree, n: 100, result1: 0, result2: 0},
}

const testTreeChildrenRangeError = "Tree Children Range Test %d: got %d " +
	"and %d for low and high indices of children of node %d (should " +
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
