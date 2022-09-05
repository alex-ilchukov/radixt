package evident_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/null"
)

var (
	empty = evident.Tree{}

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

	btree = generic.New(
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

var treeChildrenRangeTests = []struct {
	tree    evident.Tree
	n       uint
	result1 uint
	result2 uint
}{
	{tree: nil, n: 0, result1: 0, result2: 0},
	{tree: nil, n: 1, result1: 0, result2: 0},
	{tree: nil, n: 100, result1: 0, result2: 0},
	{tree: empty, n: 0, result1: 0, result2: 0},
	{tree: empty, n: 1, result1: 0, result2: 0},
	{tree: empty, n: 100, result1: 0, result2: 0},
	{tree: etree, n: 0, result1: 1, result2: 3},
	{tree: etree, n: 1, result1: 3, result2: 5},
	{tree: etree, n: 2, result1: 5, result2: 8},
	{tree: etree, n: 3, result1: 0, result2: 0},
	{tree: etree, n: 4, result1: 8, result2: 9},
	{tree: etree, n: 5, result1: 0, result2: 0},
	{tree: etree, n: 6, result1: 0, result2: 0},
	{tree: etree, n: 7, result1: 0, result2: 0},
	{tree: etree, n: 8, result1: 9, result2: 11},
	{tree: etree, n: 9, result1: 0, result2: 0},
	{tree: etree, n: 10, result1: 0, result2: 0},
	{tree: etree, n: 100, result1: 0, result2: 0},
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
