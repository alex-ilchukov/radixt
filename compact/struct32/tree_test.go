package struct32_test

import (
	"strconv"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/struct32"
	"github.com/alex-ilchukov/radixt/lookup"
	"github.com/alex-ilchukov/radixt/sapling"
)

var (
	empty = struct32.MustCreate(nil)

	atree = struct32.MustCreate(
		sapling.New(
			"authority",
			"authorization",
			"author",
			"authentication",
			"auth",
			"content-type",
			"content-length",
			"content-disposition",
		),
	)
)

var treeSizeTests = []struct {
	tree   radixt.Tree
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
	tree    radixt.Tree
	n       uint
	result1 uint
	result2 bool
}{
	{tree: empty, n: 0, result1: 0, result2: false},
	{tree: empty, n: 1, result1: 0, result2: false},
	{tree: empty, n: 100, result1: 0, result2: false},
	{tree: atree, n: 0, result1: 0, result2: false},
	{tree: atree, n: 1, result1: 4, result2: true},
	{tree: atree, n: 2, result1: 0, result2: false},
	{tree: atree, n: 3, result1: 3, result2: true},
	{tree: atree, n: 4, result1: 2, result2: true},
	{tree: atree, n: 5, result1: 7, result2: true},
	{tree: atree, n: 6, result1: 6, result2: true},
	{tree: atree, n: 7, result1: 5, result2: true},
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

var treeChunkTests = []struct {
	tree   radixt.Tree
	n      uint
	result string
}{
	{tree: empty, n: 0, result: ""},
	{tree: empty, n: 1, result: ""},
	{tree: empty, n: 100, result: ""},
	{tree: atree, n: 0, result: ""},
	{tree: atree, n: 1, result: "auth"},
	{tree: atree, n: 2, result: "content-"},
	{tree: atree, n: 3, result: "entication"},
	{tree: atree, n: 4, result: "or"},
	{tree: atree, n: 5, result: "disposition"},
	{tree: atree, n: 6, result: "length"},
	{tree: atree, n: 7, result: "type"},
	{tree: atree, n: 8, result: "i"},
	{tree: atree, n: 9, result: "ty"},
	{tree: atree, n: 10, result: "zation"},
	{tree: atree, n: 100, result: ""},
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

func eachChild(tree radixt.Tree, n uint) (indices string) {
	tree.EachChild(n, func(c uint) bool {
		if len(indices) > 0 {
			indices += ", "
		}

		indices += strconv.FormatUint(uint64(c), 10)

		return false
	})

	return
}

func eachFirstChild(tree radixt.Tree, n uint) (indices string) {
	tree.EachChild(n, func(c uint) bool {
		indices = strconv.FormatUint(uint64(c), 10)

		return true
	})

	return
}

var treeEachChildTests = []struct {
	tree    radixt.Tree
	n       uint
	f       func(radixt.Tree, uint) string
	indices string
}{
	{tree: empty, n: 0, f: eachChild, indices: ""},
	{tree: empty, n: 1, f: eachChild, indices: ""},
	{tree: empty, n: 100, f: eachChild, indices: ""},
	{tree: empty, n: 0, f: eachFirstChild, indices: ""},
	{tree: empty, n: 1, f: eachFirstChild, indices: ""},
	{tree: empty, n: 100, f: eachFirstChild, indices: ""},
	{tree: atree, n: 0, f: eachChild, indices: "1, 2"},
	{tree: atree, n: 1, f: eachChild, indices: "3, 4"},
	{tree: atree, n: 2, f: eachChild, indices: "5, 6, 7"},
	{tree: atree, n: 3, f: eachChild, indices: ""},
	{tree: atree, n: 4, f: eachChild, indices: "8"},
	{tree: atree, n: 5, f: eachChild, indices: ""},
	{tree: atree, n: 6, f: eachChild, indices: ""},
	{tree: atree, n: 7, f: eachChild, indices: ""},
	{tree: atree, n: 8, f: eachChild, indices: "9, 10"},
	{tree: atree, n: 9, f: eachChild, indices: ""},
	{tree: atree, n: 10, f: eachChild, indices: ""},
	{tree: atree, n: 100, f: eachChild, indices: ""},
	{tree: atree, n: 0, f: eachFirstChild, indices: "1"},
	{tree: atree, n: 1, f: eachFirstChild, indices: "3"},
	{tree: atree, n: 2, f: eachFirstChild, indices: "5"},
	{tree: atree, n: 3, f: eachFirstChild, indices: ""},
	{tree: atree, n: 4, f: eachFirstChild, indices: "8"},
	{tree: atree, n: 5, f: eachFirstChild, indices: ""},
	{tree: atree, n: 6, f: eachFirstChild, indices: ""},
	{tree: atree, n: 7, f: eachFirstChild, indices: ""},
	{tree: atree, n: 8, f: eachFirstChild, indices: "9"},
	{tree: atree, n: 9, f: eachFirstChild, indices: ""},
	{tree: atree, n: 10, f: eachFirstChild, indices: ""},
	{tree: atree, n: 100, f: eachFirstChild, indices: ""},
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

var treeHoardTests = []struct {
	tree    radixt.Hoarder
	result1 uint
	result2 uint
}{
	{tree: empty, result1: 72, result2: radixt.HoardExactly},
	{tree: atree, result1: 165, result2: radixt.HoardExactly},
}

const testTreeHoardError = "Tree Hoard Test %d: got %d and %d (should be %d " +
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

var treeSwitchTests = []struct {
	switcher lookup.Switcher
	n        uint
	b        byte
	result1  uint
	result2  string
	result3  bool
}{
	{
		switcher: empty,
		n:        0,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty,
		n:        1,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        0,
		b:        97,
		result1:  1,
		result2:  "uth",
		result3:  true,
	},
	{
		switcher: atree,
		n:        0,
		b:        99,
		result1:  2,
		result2:  "ontent-",
		result3:  true,
	},
	{
		switcher: atree,
		n:        0,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        1,
		b:        101,
		result1:  3,
		result2:  "ntication",
		result3:  true,
	},
	{
		switcher: atree,
		n:        1,
		b:        111,
		result1:  4,
		result2:  "r",
		result3:  true,
	},
	{
		switcher: atree,
		n:        1,
		b:        112,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        2,
		b:        116,
		result1:  7,
		result2:  "ype",
		result3:  true,
	},
	{
		switcher: atree,
		n:        2,
		b:        108,
		result1:  6,
		result2:  "ength",
		result3:  true,
	},
	{
		switcher: atree,
		n:        2,
		b:        100,
		result1:  5,
		result2:  "isposition",
		result3:  true,
	},
	{
		switcher: atree,
		n:        2,
		b:        99,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        3,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        3,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        4,
		b:        105,
		result1:  8,
		result2:  "",
		result3:  true,
	},
	{
		switcher: atree,
		n:        4,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        5,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        5,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        6,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        6,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        7,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        7,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        8,
		b:        116,
		result1:  9,
		result2:  "y",
		result3:  true,
	},
	{
		switcher: atree,
		n:        8,
		b:        122,
		result1:  10,
		result2:  "ation",
		result3:  true,
	},
	{
		switcher: atree,
		n:        8,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        9,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        9,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        10,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        10,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
}

const testTreeSwitchError = "Tree Switch Test %d: got %d, '%s', and %t, " +
	"trying to switch from node %d by byte %d (should be %d, '%s', and %t)"

func TestTreeSwitch(t *testing.T) {
	for i, tt := range treeSwitchTests {
		result1, result2, result3 := tt.switcher.Switch(tt.n, tt.b)

		e := result1 != tt.result1 ||
			result2 != tt.result2 ||
			result3 != tt.result3

		if e {
			t.Errorf(
				testTreeSwitchError,
				i,
				result1,
				result2,
				result3,
				tt.n,
				tt.b,
				tt.result1,
				tt.result2,
				tt.result3,
			)
		}
	}
}
