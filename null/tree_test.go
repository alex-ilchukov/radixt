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

var treeNodeMarkTests = []struct {
	tree   tree
	n      int
	result int
}{
	{tree: Tree, n: -2, result: -1},
	{tree: Tree, n: -1, result: -1},
	{tree: Tree, n: 0, result: -1},
	{tree: Tree, n: 1, result: -1},
	{tree: Tree, n: 100, result: -1},
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
	tree   tree
	n      int
	result string
}{
	{tree: Tree, n: -2, result: ""},
	{tree: Tree, n: -1, result: ""},
	{tree: Tree, n: 0, result: ""},
	{tree: Tree, n: 1, result: ""},
	{tree: Tree, n: 100, result: ""},
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
	tree   tree
	n      int
	result string
}{
	{tree: Tree, n: -2, result: ""},
	{tree: Tree, n: -1, result: ""},
	{tree: Tree, n: 0, result: ""},
	{tree: Tree, n: 1, result: ""},
	{tree: Tree, n: 100, result: ""},
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
	tree tree
	n    int
}{
	{tree: Tree, n: -2},
	{tree: Tree, n: -1},
	{tree: Tree, n: 0},
	{tree: Tree, n: 1},
	{tree: Tree, n: 100},
}

const testTreeNodeEachChildError = "Tree Node Each Child Test %d: iterator " +
	"func got called on node %d (shouldn't get called)"

func TestNodeEachChild(t *testing.T) {
	for i, tt := range treeNodePrefTests {
		called := false
		e := func(int) bool {
			called = true
			return false
		}
		tt.tree.NodeEachChild(tt.n, e)
		if called {
			t.Errorf(testTreeNodeEachChildError, i, tt.n)
		}
	}
}

var treeNodeTransitTests = []struct {
	tree   tree
	n      int
	pos    int
	b      byte
	result int
}{
	{tree: Tree, n: -2, pos: -1, b: 97, result: -1},
	{tree: Tree, n: -2, pos: 0, b: 97, result: -1},
	{tree: Tree, n: -1, pos: -1, b: 97, result: -1},
	{tree: Tree, n: -1, pos: 0, b: 97, result: -1},
	{tree: Tree, n: 0, pos: -1, b: 97, result: -1},
	{tree: Tree, n: 0, pos: 0, b: 97, result: -1},
	{tree: Tree, n: 1, pos: -1, b: 97, result: -1},
	{tree: Tree, n: 1, pos: 0, b: 97, result: -1},
	{tree: Tree, n: 6, pos: -1, b: 97, result: -1},
	{tree: Tree, n: 6, pos: 0, b: 97, result: -1},
	{tree: Tree, n: 100, pos: -1, b: 97, result: -1},
	{tree: Tree, n: 100, pos: 0, b: 97, result: -1},
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
