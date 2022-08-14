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

var treeNodeMarkTests = []struct {
	tree   *tree
	n      int
	result int
}{
	{tree: empty, n: -2, result: -1},
	{tree: empty, n: -1, result: -1},
	{tree: empty, n: 0, result: -1},
	{tree: empty, n: 1, result: -1},
	{tree: empty, n: 100, result: -1},
	{tree: atree, n: -2, result: -1},
	{tree: atree, n: -1, result: -1},
	{tree: atree, n: 0, result: -1},
	{tree: atree, n: 1, result: 0},
	{tree: atree, n: 2, result: 1},
	{tree: atree, n: 3, result: -1},
	{tree: atree, n: 4, result: 2},
	{tree: atree, n: 5, result: 3},
	{tree: atree, n: 6, result: 4},
	{tree: atree, n: 7, result: -1},
	{tree: atree, n: 8, result: 5},
	{tree: atree, n: 9, result: 6},
	{tree: atree, n: 10, result: 7},
	{tree: atree, n: 100, result: -1},
}

const testTreeNodeMarkError = "Tree Node Mark Test %d: got %d for mark of " +
	"node %d (should be %d)"

func TestTreeNodeMark(t *testing.T) {
	for i, tt := range treeNodeMarkTests {
		result := tt.tree.NodeMark(tt.n)
		if result != tt.result {
			t.Errorf(
				testTreeNodeMarkError,
				i,
				result,
				tt.n,
				tt.result,
			)
		}
	}
}

var treeNodeStringTests = []struct {
	tree   *tree
	n      int
	result string
}{
	{tree: empty, n: -2, result: ""},
	{tree: empty, n: -1, result: ""},
	{tree: empty, n: 0, result: ""},
	{tree: empty, n: 1, result: ""},
	{tree: empty, n: 100, result: ""},
	{tree: atree, n: -2, result: ""},
	{tree: atree, n: -1, result: ""},
	{tree: atree, n: 0, result: ""},
	{tree: atree, n: 1, result: "authority"},
	{tree: atree, n: 2, result: "authorization"},
	{tree: atree, n: 3, result: ""},
	{tree: atree, n: 4, result: "author"},
	{tree: atree, n: 5, result: "authentication"},
	{tree: atree, n: 6, result: "auth"},
	{tree: atree, n: 7, result: ""},
	{tree: atree, n: 8, result: "content-type"},
	{tree: atree, n: 9, result: "content-length"},
	{tree: atree, n: 10, result: "content-disposition"},
	{tree: atree, n: 100, result: ""},
}

const testTreeNodeStringError = "Tree Node String Test %d: got %s for " +
	"string of node %d (should be %s)"

func TestTreeNodeString(t *testing.T) {
	for i, tt := range treeNodeStringTests {
		result := tt.tree.NodeString(tt.n)
		if result != tt.result {
			t.Errorf(
				testTreeNodeStringError,
				i,
				result,
				tt.n,
				tt.result,
			)
		}
	}
}

var treeNodePrefTests = []struct {
	tree   *tree
	n      int
	result string
}{
	{tree: empty, n: -2, result: ""},
	{tree: empty, n: -1, result: ""},
	{tree: empty, n: 0, result: ""},
	{tree: empty, n: 1, result: ""},
	{tree: empty, n: 100, result: ""},
	{tree: atree, n: -2, result: ""},
	{tree: atree, n: -1, result: ""},
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

const testTreeNodePrefError = "Tree Node Pref Test %d: got %s for string of " +
	"node %d (should be %s)"

func TestTreeNodePref(t *testing.T) {
	for i, tt := range treeNodePrefTests {
		result := tt.tree.NodePref(tt.n)
		if result != tt.result {
			t.Errorf(
				testTreeNodePrefError,
				i,
				result,
				tt.n,
				tt.result,
			)
		}
	}
}

var treeNodeEachChildTests = []struct {
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

const testTreeNodeEachChildError = "Tree Node Each Child Test %d: got %d " +
	"for sum of indices of the first two children of node %d (should be " +
	"%d)"

func TestTreeNodeEachChild(t *testing.T) {
	for i, tt := range treeNodeEachChildTests {
		sum := 0
		counter := 0

		e := func(c int) bool {
			counter++
			if counter <= 2 {
				sum += c
			}

			return counter >= 2
		}

		tt.tree.NodeEachChild(tt.n, e)
		if sum != tt.sum {
			t.Errorf(
				testTreeNodeEachChildError,
				i,
				sum,
				tt.n,
				tt.sum,
			)
		}
	}
}

var treeNodeTransitTests = []struct {
	tree   *tree
	n      int
	pos    int
	b      byte
	result int
}{
	{tree: empty, n: -2, pos: -1, b: 97, result: -1},
	{tree: empty, n: -2, pos: 0, b: 97, result: -1},
	{tree: empty, n: -1, pos: -1, b: 97, result: -1},
	{tree: empty, n: -1, pos: 0, b: 97, result: -1},
	{tree: empty, n: 0, pos: -1, b: 97, result: -1},
	{tree: empty, n: 0, pos: 0, b: 97, result: -1},
	{tree: empty, n: 1, pos: -1, b: 97, result: -1},
	{tree: empty, n: 1, pos: 0, b: 97, result: -1},
	{tree: empty, n: 6, pos: -1, b: 97, result: -1},
	{tree: empty, n: 6, pos: 0, b: 97, result: -1},
	{tree: empty, n: 100, pos: -1, b: 97, result: -1},
	{tree: empty, n: 100, pos: 0, b: 97, result: -1},
	{tree: atree, n: -2, pos: -1, b: 97, result: -1},
	{tree: atree, n: -2, pos: 0, b: 97, result: -1},
	{tree: atree, n: -1, pos: -1, b: 97, result: -1},
	{tree: atree, n: -1, pos: 0, b: 97, result: -1},
	{tree: atree, n: 0, pos: -1, b: 97, result: -1},
	{tree: atree, n: 0, pos: 0, b: 97, result: 6},
	{tree: atree, n: 0, pos: 1, b: 97, result: 6},
	{tree: atree, n: 0, pos: -1, b: 99, result: -1},
	{tree: atree, n: 0, pos: 0, b: 99, result: 7},
	{tree: atree, n: 0, pos: 1, b: 99, result: 7},
	{tree: atree, n: 0, pos: 0, b: 117, result: -1},
	{tree: atree, n: 0, pos: 1, b: 117, result: -1},
	{tree: atree, n: 1, pos: -1, b: 97, result: -1},
	{tree: atree, n: 1, pos: 0, b: 97, result: -1},
	{tree: atree, n: 6, pos: -1, b: 97, result: -1},
	{tree: atree, n: 6, pos: 0, b: 97, result: 6},
	{tree: atree, n: 6, pos: 0, b: 117, result: -1},
	{tree: atree, n: 6, pos: 1, b: 117, result: 6},
	{tree: atree, n: 6, pos: 2, b: 116, result: 6},
	{tree: atree, n: 6, pos: 3, b: 104, result: 6},
	{tree: atree, n: 6, pos: 4, b: 111, result: 4},
	{tree: atree, n: 6, pos: 5, b: 111, result: 4},
	{tree: atree, n: 6, pos: 4, b: 101, result: 5},
	{tree: atree, n: 6, pos: 5, b: 101, result: 5},
	{tree: atree, n: 100, pos: -1, b: 97, result: -1},
	{tree: atree, n: 100, pos: 0, b: 97, result: -1},
}

const testTreeNodeTransitError = "Tree Node Transit Test %d: got %d for " +
	"transition of node %d by byte %d at position %d (should be %d)"

func TestTreeNodeTransit(t *testing.T) {
	for i, tt := range treeNodeTransitTests {
		result := tt.tree.NodeTransit(tt.n, tt.pos, tt.b)
		if result != tt.result {
			t.Errorf(
				testTreeNodeTransitError,
				i,
				result,
				tt.n,
				tt.b,
				tt.pos,
				tt.result,
			)
		}
	}
}
