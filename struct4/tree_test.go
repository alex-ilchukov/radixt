package struct4

import (
	"testing"

	"github.com/alex-ilchukov/radixt/generic"
)

var (
	empty = MustCreate(nil)

	atree = MustCreate(generic.New(
		"authority",
		"authorization",
		"author",
		"authentication",
		"auth",
		"content-type",
		"content-length",
		"content-disposition",
	))
)

var treeSizeTests = []struct {
	tree   *tree
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

var treeHasTests = []struct {
	tree   *tree
	n      int
	result bool
}{
	{tree: empty, n: -2, result: false},
	{tree: empty, n: -1, result: false},
	{tree: empty, n: 0, result: false},
	{tree: empty, n: 1, result: false},
	{tree: empty, n: 100, result: false},
	{tree: atree, n: -2, result: false},
	{tree: atree, n: -1, result: false},
	{tree: atree, n: 0, result: true},
	{tree: atree, n: 1, result: true},
	{tree: atree, n: 2, result: true},
	{tree: atree, n: 3, result: true},
	{tree: atree, n: 4, result: true},
	{tree: atree, n: 5, result: true},
	{tree: atree, n: 6, result: true},
	{tree: atree, n: 7, result: true},
	{tree: atree, n: 8, result: true},
	{tree: atree, n: 9, result: true},
	{tree: atree, n: 10, result: true},
	{tree: atree, n: 100, result: false},
}

const testTreeHasError = "Tree Has Test %d: got %t for if the tree has node " +
	"node %d (should be %t)"

func TestTreeHas(t *testing.T) {
	for i, tt := range treeHasTests {
		result := tt.tree.Has(tt.n)
		if result != tt.result {
			t.Errorf(testTreeHasError, i, result, tt.n, tt.result)
		}
	}
}

var treeRootTests = []struct {
	tree   *tree
	result int
}{
	{tree: empty, result: -1},
	{tree: atree, result: 0},
}

const testTreeRootError = "Tree Root Test %d: got %d for root of the tree " +
	"(should be %d)"

func TestTreeRoot(t *testing.T) {
	for i, tt := range treeRootTests {
		result := tt.tree.Root()
		if result != tt.result {
			t.Errorf(testTreeRootError, i, result, tt.result)
		}
	}
}

var treeValueTests = []struct {
	tree    *tree
	n       int
	result1 uint
	result2 bool
}{
	{tree: empty, n: -2, result1: 0, result2: false},
	{tree: empty, n: 0, result1: 0, result2: false},
	{tree: empty, n: 1, result1: 0, result2: false},
	{tree: empty, n: 100, result1: 0, result2: false},
	{tree: atree, n: -2, result1: 0, result2: false},
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

var treeEachChildTests = []struct {
	tree *tree
	n    int
	sum  int
}{
	{tree: empty, n: -2, sum: 0},
	{tree: empty, n: -1, sum: 0},
	{tree: empty, n: 0, sum: 0},
	{tree: empty, n: 1, sum: 0},
	{tree: empty, n: 100, sum: 0},
	{tree: atree, n: -2, sum: 0},
	{tree: atree, n: -1, sum: 0},
	{tree: atree, n: 0, sum: 3},
	{tree: atree, n: 1, sum: 7},
	{tree: atree, n: 2, sum: 11},
	{tree: atree, n: 3, sum: 8},
	{tree: atree, n: 4, sum: 0},
	{tree: atree, n: 5, sum: 0},
	{tree: atree, n: 6, sum: 0},
	{tree: atree, n: 7, sum: 0},
	{tree: atree, n: 8, sum: 19},
	{tree: atree, n: 9, sum: 0},
	{tree: atree, n: 10, sum: 0},
	{tree: atree, n: 100, sum: 0},
}

const testTreeEachChildError = "Tree Each Child Test %d: got %d for sum of " +
	"indices of the first two children of node %d (should be %d)"

func TestTreeEachChild(t *testing.T) {
	for i, tt := range treeEachChildTests {
		sum := 0
		counter := 0

		e := func(c int) bool {
			counter++
			if counter <= 2 {
				sum += c
			}

			return counter >= 2
		}

		tt.tree.EachChild(tt.n, e)
		if sum != tt.sum {
			t.Errorf(testTreeEachChildError, i, sum, tt.n, tt.sum)
		}
	}
}

var treeByteAtTests = []struct {
	tree    *tree
	n       int
	npos    uint
	result1 byte
	result2 bool
}{
	{tree: empty, n: -2, npos: 0, result1: 0, result2: false},
	{tree: empty, n: -2, npos: 1, result1: 0, result2: false},
	{tree: empty, n: -1, npos: 0, result1: 0, result2: false},
	{tree: empty, n: -1, npos: 1, result1: 0, result2: false},
	{tree: empty, n: 0, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 0, npos: 1, result1: 0, result2: false},
	{tree: empty, n: 1, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 1, npos: 1, result1: 0, result2: false},
	{tree: empty, n: 6, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 6, npos: 1, result1: 0, result2: false},
	{tree: empty, n: 100, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 100, npos: 1, result1: 0, result2: false},
	{tree: atree, n: -2, npos: 0, result1: 0, result2: false},
	{tree: atree, n: -2, npos: 1, result1: 0, result2: false},
	{tree: atree, n: -1, npos: 0, result1: 0, result2: false},
	{tree: atree, n: -1, npos: 1, result1: 0, result2: false},
	{tree: atree, n: 0, npos: 0, result1: 0, result2: false},
	{tree: atree, n: 0, npos: 1, result1: 0, result2: false},
	{tree: atree, n: 1, npos: 0, result1: 97, result2: true},
	{tree: atree, n: 1, npos: 1, result1: 117, result2: true},
	{tree: atree, n: 1, npos: 2, result1: 116, result2: true},
	{tree: atree, n: 1, npos: 3, result1: 104, result2: true},
	{tree: atree, n: 1, npos: 4, result1: 0, result2: false},
	{tree: atree, n: 1, npos: 5, result1: 0, result2: false},
	{tree: atree, n: 9, npos: 0, result1: 116, result2: true},
	{tree: atree, n: 9, npos: 1, result1: 121, result2: true},
	{tree: atree, n: 9, npos: 2, result1: 0, result2: false},
	{tree: atree, n: 9, npos: 3, result1: 0, result2: false},
	{tree: atree, n: 100, npos: 0, result1: 0, result2: false},
	{tree: atree, n: 100, npos: 1, result1: 0, result2: false},
}

const testTreeByteAtError = "Tree ByteAt Test %d: got %d and %t for byte at " +
	"position %d of chunk of node %d (should be %d and %t)"

func TestTreeByteAt(t *testing.T) {
	for i, tt := range treeByteAtTests {
		result1, result2 := tt.tree.ByteAt(tt.n, tt.npos)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTreeByteAtError,
				i,
				result1,
				result2,
				tt.npos,
				tt.n,
				tt.result1,
				tt.result2,
			)
		}
	}
}
