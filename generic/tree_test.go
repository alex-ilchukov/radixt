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
	tree    *tree
	n       uint
	result1 uint
	result2 bool
}{
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

var treeEachChildTests = []struct {
	tree *tree
	n    uint
	sum  uint
}{
	{tree: empty, n: 0, sum: 0},
	{tree: empty, n: 1, sum: 0},
	{tree: empty, n: 100, sum: 0},
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
		sum := uint(0)
		counter := 0

		e := func(c uint) bool {
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
	n       uint
	npos    uint
	result1 byte
	result2 bool
}{
	{tree: empty, n: 0, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 0, npos: 1, result1: 0, result2: false},
	{tree: empty, n: 1, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 1, npos: 1, result1: 0, result2: false},
	{tree: empty, n: 6, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 6, npos: 1, result1: 0, result2: false},
	{tree: empty, n: 100, npos: 0, result1: 0, result2: false},
	{tree: empty, n: 100, npos: 1, result1: 0, result2: false},
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
