package generic

import "testing"

var (
	empty = New()

	atree = New(
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
	tree   *tree
	result int
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

var treeMarkTests = []struct {
	tree   *tree
	n      int
	result uint
}{
	{tree: empty, n: -2, result: 0},
	{tree: empty, n: 0, result: 0},
	{tree: empty, n: 0, result: 0},
	{tree: empty, n: 1, result: 0},
	{tree: empty, n: 100, result: 0},
	{tree: atree, n: -2, result: 0},
	{tree: atree, n: 0, result: 0},
	{tree: atree, n: 0, result: 0},
	{tree: atree, n: 1, result: 1},
	{tree: atree, n: 2, result: 2},
	{tree: atree, n: 3, result: 0},
	{tree: atree, n: 4, result: 3},
	{tree: atree, n: 5, result: 4},
	{tree: atree, n: 6, result: 5},
	{tree: atree, n: 7, result: 0},
	{tree: atree, n: 8, result: 6},
	{tree: atree, n: 9, result: 7},
	{tree: atree, n: 10, result: 8},
	{tree: atree, n: 100, result: 0},
}

const testTreeMarkError = "Tree Mark Test %d: got %d for mark of node %d " +
	"(should be %d)"

func TestTreeMark(t *testing.T) {
	for i, tt := range treeMarkTests {
		result := tt.tree.Mark(tt.n)
		if result != tt.result {
			t.Errorf(testTreeMarkError, i, result, tt.n, tt.result)
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
	{tree: atree, n: 0, sum: 13},
	{tree: atree, n: 1, sum: 0},
	{tree: atree, n: 2, sum: 0},
	{tree: atree, n: 3, sum: 3},
	{tree: atree, n: 4, sum: 3},
	{tree: atree, n: 5, sum: 0},
	{tree: atree, n: 6, sum: 9},
	{tree: atree, n: 7, sum: 17},
	{tree: atree, n: 8, sum: 0},
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
	{tree: atree, n: 1, npos: 0, result1: 116, result2: true},
	{tree: atree, n: 1, npos: 1, result1: 121, result2: true},
	{tree: atree, n: 1, npos: 2, result1: 0, result2: false},
	{tree: atree, n: 1, npos: 3, result1: 0, result2: false},
	{tree: atree, n: 6, npos: 0, result1: 97, result2: true},
	{tree: atree, n: 6, npos: 1, result1: 117, result2: true},
	{tree: atree, n: 6, npos: 2, result1: 116, result2: true},
	{tree: atree, n: 6, npos: 3, result1: 104, result2: true},
	{tree: atree, n: 6, npos: 4, result1: 0, result2: false},
	{tree: atree, n: 6, npos: 5, result1: 0, result2: false},
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
