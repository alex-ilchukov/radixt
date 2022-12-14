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

var treeEachChildTests = []struct {
	tree radixt.Tree
	n    uint
}{
	{tree: null.Tree, n: 0},
	{tree: null.Tree, n: 1},
	{tree: null.Tree, n: 100},
}

const testTreeEachChildError = "Tree Each Child Test %d: provided function " +
	"got called (and it should not)"

func TestTreeEachChild(t *testing.T) {
	for i, tt := range treeEachChildTests {
		called := false
		tt.tree.EachChild(tt.n, func(uint) bool {
			called = true
			return false
		})
		if called {
			t.Errorf(testTreeEachChildError, i)
		}
	}
}

var treeHoardTests = []struct {
	tree    radixt.Hoarder
	result1 uint
	result2 uint
}{
	{tree: null.Tree, result1: 0, result2: radixt.HoardExactly},
}

const testTreeHoardError = "Tree Hoard Test %d: got %d and %d (should be %d" +
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
