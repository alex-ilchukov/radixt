package null

import "testing"

var treeSizeTests = []struct {
	tree   tree
	result int
}{
	{tree: Tree, result: 0},
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
	tree   tree
	n      int
	result bool
}{
	{tree: Tree, n: -2, result: false},
	{tree: Tree, n: -1, result: false},
	{tree: Tree, n: 0, result: false},
	{tree: Tree, n: 1, result: false},
	{tree: Tree, n: 100, result: false},
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
	tree   tree
	result int
}{
	{tree: Tree, result: -1},
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
	tree    tree
	n       int
	result1 uint
	result2 bool
}{
	{tree: Tree, n: -2, result1: 0, result2: false},
	{tree: Tree, n: -1, result1: 0, result2: false},
	{tree: Tree, n: 0, result1: 0, result2: false},
	{tree: Tree, n: 1, result1: 0, result2: false},
	{tree: Tree, n: 100, result1: 0, result2: false},
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
	tree tree
	n    int
}{
	{tree: Tree, n: -2},
	{tree: Tree, n: -1},
	{tree: Tree, n: 0},
	{tree: Tree, n: 1},
	{tree: Tree, n: 100},
}

const testTreeEachChildError = "Tree Each Child Test %d: iterator func got " +
	"called on node %d (shouldn't get called)"

func TestEachChild(t *testing.T) {
	for i, tt := range treeEachChildTests {
		called := false
		e := func(int) bool {
			called = true
			return false
		}
		tt.tree.EachChild(tt.n, e)
		if called {
			t.Errorf(testTreeEachChildError, i, tt.n)
		}
	}
}

var treeByteAtTests = []struct {
	tree    tree
	n       int
	npos    uint
	result1 byte
	result2 bool
}{
	{tree: Tree, n: -2, npos: 0, result1: 0, result2: false},
	{tree: Tree, n: -2, npos: 1, result1: 0, result2: false},
	{tree: Tree, n: -1, npos: 0, result1: 0, result2: false},
	{tree: Tree, n: -1, npos: 1, result1: 0, result2: false},
	{tree: Tree, n: 0, npos: 0, result1: 0, result2: false},
	{tree: Tree, n: 0, npos: 1, result1: 0, result2: false},
	{tree: Tree, n: 1, npos: 0, result1: 0, result2: false},
	{tree: Tree, n: 1, npos: 1, result1: 0, result2: false},
	{tree: Tree, n: 6, npos: 0, result1: 0, result2: false},
	{tree: Tree, n: 6, npos: 1, result1: 0, result2: false},
	{tree: Tree, n: 100, npos: 0, result1: 0, result2: false},
	{tree: Tree, n: 100, npos: 1, result1: 0, result2: false},
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
