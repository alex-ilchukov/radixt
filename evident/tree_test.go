package evident_test

import (
	"strconv"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/null"
	"github.com/alex-ilchukov/radixt/sapling"
)

var (
	empty = evident.Tree{}

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

	btree = sapling.New(
		"authority",
		"authorization",
		"author",
		"authentication",
		"content-type",
		"content-length",
		"content-disposition",
	)

	etree = evident.Tree{
		"|": { //                                           0
			"auth|4": { //                               1
				"entication|3": nil, //              .3
				"or|2": { //                         .4
					"i|": { //                   ..8
						"ty|0":     nil, //  .. 9
						"zation|1": nil, //  .. 10
					}, //                        ..
				}, //                                ..
			}, //                                        ..
			"content-|": { //                            2.
				"disposition|7": nil, //              5
				"length|6":      nil, //              6
				"type|5":        nil, //              7
			},
		},
	}
)

var treeSizeTests = []struct {
	tree   evident.Tree
	result uint
}{
	{tree: nil, result: 0},
	{tree: empty, result: 0},
	{tree: etree, result: 11},
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
	tree    evident.Tree
	n       uint
	result1 uint
	result2 bool
}{
	{tree: nil, n: 0, result1: 0, result2: false},
	{tree: nil, n: 1, result1: 0, result2: false},
	{tree: nil, n: 100, result1: 0, result2: false},
	{tree: empty, n: 0, result1: 0, result2: false},
	{tree: empty, n: 1, result1: 0, result2: false},
	{tree: empty, n: 100, result1: 0, result2: false},
	{tree: etree, n: 0, result1: 0, result2: false},
	{tree: etree, n: 1, result1: 4, result2: true},
	{tree: etree, n: 2, result1: 0, result2: false},
	{tree: etree, n: 3, result1: 3, result2: true},
	{tree: etree, n: 4, result1: 2, result2: true},
	{tree: etree, n: 5, result1: 7, result2: true},
	{tree: etree, n: 6, result1: 6, result2: true},
	{tree: etree, n: 7, result1: 5, result2: true},
	{tree: etree, n: 8, result1: 0, result2: false},
	{tree: etree, n: 9, result1: 0, result2: true},
	{tree: etree, n: 10, result1: 1, result2: true},
	{tree: etree, n: 100, result1: 0, result2: false},
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
	tree   evident.Tree
	n      uint
	result string
}{
	{tree: nil, n: 0, result: ""},
	{tree: nil, n: 1, result: ""},
	{tree: nil, n: 100, result: ""},
	{tree: empty, n: 0, result: ""},
	{tree: empty, n: 1, result: ""},
	{tree: empty, n: 100, result: ""},
	{tree: etree, n: 0, result: ""},
	{tree: etree, n: 1, result: "auth"},
	{tree: etree, n: 2, result: "content-"},
	{tree: etree, n: 3, result: "entication"},
	{tree: etree, n: 4, result: "or"},
	{tree: etree, n: 5, result: "disposition"},
	{tree: etree, n: 6, result: "length"},
	{tree: etree, n: 7, result: "type"},
	{tree: etree, n: 8, result: "i"},
	{tree: etree, n: 9, result: "ty"},
	{tree: etree, n: 10, result: "zation"},
	{tree: etree, n: 100, result: ""},
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

func eachChild(tree evident.Tree, n uint) (indices string) {
	tree.EachChild(n, func(c uint) bool {
		if len(indices) > 0 {
			indices += ", "
		}

		indices += strconv.FormatUint(uint64(c), 10)

		return false
	})

	return
}

func eachFirstChild(tree evident.Tree, n uint) (indices string) {
	tree.EachChild(n, func(c uint) bool {
		indices = strconv.FormatUint(uint64(c), 10)

		return true
	})

	return
}

var treeEachChildTests = []struct {
	tree    evident.Tree
	n       uint
	f       func(evident.Tree, uint) string
	indices string
}{
	{tree: nil, n: 0, f: eachChild, indices: ""},
	{tree: nil, n: 1, f: eachChild, indices: ""},
	{tree: nil, n: 100, f: eachChild, indices: ""},
	{tree: nil, n: 0, f: eachFirstChild, indices: ""},
	{tree: nil, n: 1, f: eachFirstChild, indices: ""},
	{tree: nil, n: 100, f: eachFirstChild, indices: ""},
	{tree: empty, n: 0, f: eachChild, indices: ""},
	{tree: empty, n: 1, f: eachChild, indices: ""},
	{tree: empty, n: 100, f: eachChild, indices: ""},
	{tree: empty, n: 0, f: eachFirstChild, indices: ""},
	{tree: empty, n: 1, f: eachFirstChild, indices: ""},
	{tree: empty, n: 100, f: eachFirstChild, indices: ""},
	{tree: etree, n: 0, f: eachChild, indices: "1, 2"},
	{tree: etree, n: 1, f: eachChild, indices: "3, 4"},
	{tree: etree, n: 2, f: eachChild, indices: "5, 6, 7"},
	{tree: etree, n: 3, f: eachChild, indices: ""},
	{tree: etree, n: 4, f: eachChild, indices: "8"},
	{tree: etree, n: 5, f: eachChild, indices: ""},
	{tree: etree, n: 6, f: eachChild, indices: ""},
	{tree: etree, n: 7, f: eachChild, indices: ""},
	{tree: etree, n: 8, f: eachChild, indices: "9, 10"},
	{tree: etree, n: 9, f: eachChild, indices: ""},
	{tree: etree, n: 10, f: eachChild, indices: ""},
	{tree: etree, n: 100, f: eachChild, indices: ""},
	{tree: etree, n: 0, f: eachFirstChild, indices: "1"},
	{tree: etree, n: 1, f: eachFirstChild, indices: "3"},
	{tree: etree, n: 2, f: eachFirstChild, indices: "5"},
	{tree: etree, n: 3, f: eachFirstChild, indices: ""},
	{tree: etree, n: 4, f: eachFirstChild, indices: "8"},
	{tree: etree, n: 5, f: eachFirstChild, indices: ""},
	{tree: etree, n: 6, f: eachFirstChild, indices: ""},
	{tree: etree, n: 7, f: eachFirstChild, indices: ""},
	{tree: etree, n: 8, f: eachFirstChild, indices: "9"},
	{tree: etree, n: 9, f: eachFirstChild, indices: ""},
	{tree: etree, n: 10, f: eachFirstChild, indices: ""},
	{tree: etree, n: 100, f: eachFirstChild, indices: ""},
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

var treeEqTests = []struct {
	t      evident.Tree
	u      radixt.Tree
	result bool
}{
	{t: nil, u: nil, result: true},
	{t: nil, u: empty, result: true},
	{t: empty, u: empty, result: true},
	{t: empty, u: nil, result: true},
	{t: nil, u: null.Tree, result: true},
	{t: empty, u: null.Tree, result: true},
	{t: empty, u: etree, result: false},
	{t: etree, u: empty, result: false},
	{t: etree, u: etree, result: true},
	{
		t: etree,
		u: evident.Tree{
			"|": {
				"content-|": {
					"length|6":      nil,
					"disposition|7": nil,
					"type|5":        nil,
				},
				"auth|4": {
					"or|2": {
						"i|": {
							"ty|0":     nil,
							"zation|1": nil,
						},
					},
					"entication|3": nil,
				},
			},
		},
		result: true, // exactly the same tree but different node order
	},
	{
		t: etree,
		u: evident.Tree{
			"|": {
				"auth|4": {
					"entication|3": nil,
					"or|2": {
						"i|": {
							"ty|0":     nil,
							"zation|2": nil,
						},
					},
				},
				"content-|": {
					"disposition|7": nil,
					"length|6":      nil,
					"type|5":        nil,
				},
			},
		},
		result: false, // "zation" has value 2 instead of 1
	},
	{
		t: etree,
		u: evident.Tree{
			"|": {
				"auth|4": {
					"entication|3": nil,
					"or|2": {
						"i|": {
							"ty|0":     nil,
							"zation|1": nil,
						},
					},
				},
				"content-|": {
					"disposition|7": nil,
					"length|6":      nil,
					"type|5":        nil,
					"rage|":         nil,
				},
			},
		},
		result: false, // additional node under "content-|"
	},
	{
		t: etree,
		u: evident.Tree{
			"|": {
				"content-|": {
					"length|6":      nil,
					"disposition|7": nil,
					"type|5":        nil,
				},
				"auth|4": {
					"or|2": {
						"i|": {
							"ty|0": {
								"ty|": nil,
							},
							"zation|1": nil,
						},
					},
					"entication|3": nil,
				},
			},
		},
		result: false, // additional node under "ty|0"
	},
	{
		t: etree,
		u: evident.Tree{
			"|": {
				"content-|": {
					"length|6":      nil,
					"disposition|7": nil,
				},
				"auth|4": {
					"or|2": {
						"i|": {
							"ty|0":     nil,
							"zation|1": nil,
						},
					},
					"entication|3": nil,
				},
			},
		},
		result: false, // "type|5" node is absent
	},
	{t: etree, u: atree, result: true},
	{t: etree, u: btree, result: false},
}

const testTreeEqError = "Tree Eq Test %d: got that %v.Eq(%v) = %t (should " +
	"be %t)"

func TestTreeEq(t *testing.T) {
	for i, tt := range treeEqTests {
		result := tt.t.Eq(tt.u)
		if result != tt.result {
			t.Errorf(
				testTreeEqError,
				i,
				tt.t,
				tt.u,
				result,
				tt.result,
			)
		}
	}
}

var treeHoardTests = []struct {
	tree    evident.Tree
	result1 uint
	result2 uint
}{
	{tree: nil, result1: 0, result2: radixt.HoardExactly},
	{tree: empty, result1: 48, result2: radixt.HoardExactly},
	{
		tree: etree,
		result1: 48 + 128 + // {"|": …}
			16 + 1 + //    "|"
			48 + 128 + //  {"auth|4": …, "content-|": …}
			16 + 6 + //    "auth|4"
			16 + 9 + //    "content-|"
			48 + 128 + //  {"entication|3": …, "or|2": …}
			16 + 12 + //   "entication|3"
			16 + 4 + //    "or|2"
			48 + 128 + //  {"disposition|7": …, "length|6": …, …}
			16 + 13 + //   "disposition|7"
			16 + 8 + //    "length|6"
			16 + 6 + //    "type|5"
			48 + 128 + //  {"i|": …}
			16 + 2 + //    "i|"
			48 + 128 + //  {"ty|0": …, "zation|1": …}
			16 + 4 + //    "ty|0"
			16 + 8, //     "zation|1"
		result2: radixt.HoardAtLeast,
	},
}

const testTreeHoardError = "Tree Hoard Test %d: got %d and %d (should be " +
	"%d and %d)"

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
