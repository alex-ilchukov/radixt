package strg_test

import (
	"strconv"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/strg"
	"github.com/alex-ilchukov/radixt/generic"
)

var (
	blank3    = strg.Tree[strg.N3]("")
	tooshort3 = strg.Tree[strg.N3]("123")
	empty3    = strg.MustCreate[strg.N3](nil)

	atree3 = strg.MustCreate[strg.N3](
		generic.New(
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

	blank4    = strg.Tree[strg.N4]("")
	tooshort4 = strg.Tree[strg.N4]("123")
	empty4    = strg.MustCreate[strg.N4](nil)

	atree4 = strg.MustCreate[strg.N4](
		generic.New(
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

var tree3SizeTests = []struct {
	tree   strg.Tree[strg.N3]
	result uint
}{
	{tree: blank3, result: 0},
	{tree: tooshort3, result: 0},
	{tree: empty3, result: 0},
	{tree: atree3, result: 11},
}

const testTree3SizeError = "Tree3 Size Test %d: got %d for size (should be %d)"

func TestTree3Size(t *testing.T) {
	for i, tt := range tree3SizeTests {
		result := tt.tree.Size()
		if result != tt.result {
			t.Errorf(testTree3SizeError, i, result, tt.result)
		}
	}
}

var tree3ValueTests = []struct {
	tree    strg.Tree[strg.N3]
	n       uint
	result1 uint
	result2 bool
}{
	{tree: blank3, n: 0, result1: 0, result2: false},
	{tree: blank3, n: 1, result1: 0, result2: false},
	{tree: blank3, n: 100, result1: 0, result2: false},
	{tree: tooshort3, n: 0, result1: 0, result2: false},
	{tree: tooshort3, n: 1, result1: 0, result2: false},
	{tree: tooshort3, n: 100, result1: 0, result2: false},
	{tree: empty3, n: 0, result1: 0, result2: false},
	{tree: empty3, n: 1, result1: 0, result2: false},
	{tree: empty3, n: 100, result1: 0, result2: false},
	{tree: atree3, n: 0, result1: 0, result2: false},
	{tree: atree3, n: 1, result1: 4, result2: true},
	{tree: atree3, n: 2, result1: 0, result2: false},
	{tree: atree3, n: 3, result1: 2, result2: true},
	{tree: atree3, n: 4, result1: 3, result2: true},
	{tree: atree3, n: 5, result1: 5, result2: true},
	{tree: atree3, n: 6, result1: 6, result2: true},
	{tree: atree3, n: 7, result1: 7, result2: true},
	{tree: atree3, n: 8, result1: 0, result2: false},
	{tree: atree3, n: 9, result1: 0, result2: true},
	{tree: atree3, n: 10, result1: 1, result2: true},
	{tree: atree3, n: 100, result1: 0, result2: false},
}

const testTree3ValueError = "Tree3 Value Test %d: got %d and %t for value " +
	"of node %d (should be %d and %t)"

func TestTree3Value(t *testing.T) {
	for i, tt := range tree3ValueTests {
		result1, result2 := tt.tree.Value(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree3ValueError,
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

var tree3ChunkTests = []struct {
	tree   strg.Tree[strg.N3]
	n      uint
	result string
}{
	{tree: blank3, n: 0, result: ""},
	{tree: blank3, n: 1, result: ""},
	{tree: blank3, n: 100, result: ""},
	{tree: tooshort3, n: 0, result: ""},
	{tree: tooshort3, n: 1, result: ""},
	{tree: tooshort3, n: 100, result: ""},
	{tree: empty3, n: 0, result: ""},
	{tree: empty3, n: 1, result: ""},
	{tree: empty3, n: 100, result: ""},
	{tree: atree3, n: 0, result: ""},
	{tree: atree3, n: 1, result: "auth"},
	{tree: atree3, n: 2, result: "content-"},
	{tree: atree3, n: 3, result: "or"},
	{tree: atree3, n: 4, result: "entication"},
	{tree: atree3, n: 5, result: "type"},
	{tree: atree3, n: 6, result: "length"},
	{tree: atree3, n: 7, result: "disposition"},
	{tree: atree3, n: 8, result: "i"},
	{tree: atree3, n: 9, result: "ty"},
	{tree: atree3, n: 10, result: "zation"},
	{tree: atree3, n: 100, result: ""},
}

const testTree3ChunkError = "Tree3 Chunk Test %d: got '%s' for chunk of " +
	"node %d (should be '%s')"

func TestTree3Chunk(t *testing.T) {
	for i, tt := range tree3ChunkTests {
		result := tt.tree.Chunk(tt.n)
		if result != tt.result {
			t.Errorf(
				testTree3ChunkError,
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

var tree3EachChildTests = []struct {
	tree    strg.Tree[strg.N3]
	n       uint
	f       func(radixt.Tree, uint) string
	indices string
}{
	{tree: blank3, n: 0, f: eachChild, indices: ""},
	{tree: blank3, n: 1, f: eachChild, indices: ""},
	{tree: blank3, n: 100, f: eachChild, indices: ""},
	{tree: blank3, n: 0, f: eachFirstChild, indices: ""},
	{tree: blank3, n: 1, f: eachFirstChild, indices: ""},
	{tree: blank3, n: 100, f: eachFirstChild, indices: ""},
	{tree: tooshort3, n: 0, f: eachChild, indices: ""},
	{tree: tooshort3, n: 1, f: eachChild, indices: ""},
	{tree: tooshort3, n: 100, f: eachChild, indices: ""},
	{tree: tooshort3, n: 0, f: eachFirstChild, indices: ""},
	{tree: tooshort3, n: 1, f: eachFirstChild, indices: ""},
	{tree: tooshort3, n: 100, f: eachFirstChild, indices: ""},
	{tree: empty3, n: 0, f: eachChild, indices: ""},
	{tree: empty3, n: 1, f: eachChild, indices: ""},
	{tree: empty3, n: 100, f: eachChild, indices: ""},
	{tree: empty3, n: 0, f: eachFirstChild, indices: ""},
	{tree: empty3, n: 1, f: eachFirstChild, indices: ""},
	{tree: empty3, n: 100, f: eachFirstChild, indices: ""},
	{tree: atree3, n: 0, f: eachChild, indices: "1, 2"},
	{tree: atree3, n: 1, f: eachChild, indices: "3, 4"},
	{tree: atree3, n: 2, f: eachChild, indices: "5, 6, 7"},
	{tree: atree3, n: 3, f: eachChild, indices: "8"},
	{tree: atree3, n: 4, f: eachChild, indices: ""},
	{tree: atree3, n: 5, f: eachChild, indices: ""},
	{tree: atree3, n: 6, f: eachChild, indices: ""},
	{tree: atree3, n: 7, f: eachChild, indices: ""},
	{tree: atree3, n: 8, f: eachChild, indices: "9, 10"},
	{tree: atree3, n: 9, f: eachChild, indices: ""},
	{tree: atree3, n: 10, f: eachChild, indices: ""},
	{tree: atree3, n: 100, f: eachChild, indices: ""},
	{tree: atree3, n: 0, f: eachFirstChild, indices: "1"},
	{tree: atree3, n: 1, f: eachFirstChild, indices: "3"},
	{tree: atree3, n: 2, f: eachFirstChild, indices: "5"},
	{tree: atree3, n: 3, f: eachFirstChild, indices: "8"},
	{tree: atree3, n: 4, f: eachFirstChild, indices: ""},
	{tree: atree3, n: 5, f: eachFirstChild, indices: ""},
	{tree: atree3, n: 6, f: eachFirstChild, indices: ""},
	{tree: atree3, n: 7, f: eachFirstChild, indices: ""},
	{tree: atree3, n: 8, f: eachFirstChild, indices: "9"},
	{tree: atree3, n: 9, f: eachFirstChild, indices: ""},
	{tree: atree3, n: 10, f: eachFirstChild, indices: ""},
	{tree: atree3, n: 100, f: eachFirstChild, indices: ""},
}

const testTree3EachChildError = "Tree3 Each Child Test %d: got %s as result " +
	"indices (should be %s)"

func TestTree3EachChild(t *testing.T) {
	for i, tt := range tree3EachChildTests {
		indices := tt.f(tt.tree, tt.n)
		if indices != tt.indices {
			t.Errorf(
				testTree3EachChildError,
				i,
				indices,
				tt.indices,
			)
		}
	}
}

var tree3HoardTests = []struct {
	tree    strg.Tree[strg.N3]
	result1 uint
	result2 uint
}{
	{tree: blank3, result1: 0, result2: radixt.HoardExactly},
	{tree: tooshort3, result1: 3, result2: radixt.HoardExactly},
	{tree: empty3, result1: 10, result2: radixt.HoardExactly},
	{tree: atree3, result1: 94, result2: radixt.HoardExactly},
}

const testTree3HoardError = "Tree3 Hoard Test %d: got %d and %d (should be " +
	"%d and %d)"

func TestTree3Hoard(t *testing.T) {
	for i, tt := range tree3HoardTests {
		result1, result2 := tt.tree.Hoard()
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree3HoardError,
				i,
				result1,
				result2,
				tt.result1,
				tt.result2,
			)
		}
	}
}

var tree3SwitchTests = []struct {
	switcher strg.Tree[strg.N3]
	n        uint
	b        byte
	result1  uint
	result2  string
	result3  bool
}{
	{
		switcher: blank3,
		n:        0,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: blank3,
		n:        1,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: blank3,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: tooshort3,
		n:        0,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: tooshort3,
		n:        1,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: tooshort3,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty3,
		n:        0,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty3,
		n:        1,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty3,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        0,
		b:        97,
		result1:  1,
		result2:  "uth",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        0,
		b:        99,
		result1:  2,
		result2:  "ontent-",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        0,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        1,
		b:        101,
		result1:  4,
		result2:  "ntication",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        1,
		b:        111,
		result1:  3,
		result2:  "r",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        1,
		b:        112,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        2,
		b:        116,
		result1:  5,
		result2:  "ype",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        2,
		b:        108,
		result1:  6,
		result2:  "ength",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        2,
		b:        100,
		result1:  7,
		result2:  "isposition",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        2,
		b:        99,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        3,
		b:        105,
		result1:  8,
		result2:  "",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        3,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        4,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        4,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        5,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        5,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        6,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        6,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        7,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        7,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        8,
		b:        116,
		result1:  9,
		result2:  "y",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        8,
		b:        122,
		result1:  10,
		result2:  "ation",
		result3:  true,
	},
	{
		switcher: atree3,
		n:        8,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        9,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        9,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        10,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        10,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree3,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
}

const testTree3SwitchError = "Tree3 Switch Test %d: got %d, '%s', and %t, " +
	"trying to switch from node %d by byte %d (should be %d, '%s', and %t)"

func TestTree3Switch(t *testing.T) {
	for i, tt := range tree3SwitchTests {
		result1, result2, result3 := tt.switcher.Switch(tt.n, tt.b)

		e := result1 != tt.result1 ||
			result2 != tt.result2 ||
			result3 != tt.result3

		if e {
			t.Errorf(
				testTree3SwitchError,
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

var tree4SizeTests = []struct {
	tree   strg.Tree[strg.N4]
	result uint
}{
	{tree: blank4, result: 0},
	{tree: tooshort4, result: 0},
	{tree: empty4, result: 0},
	{tree: atree4, result: 11},
}

const testTree4SizeError = "Tree4 Size Test %d: got %d for size (should be %d)"

func TestTree4Size(t *testing.T) {
	for i, tt := range tree4SizeTests {
		result := tt.tree.Size()
		if result != tt.result {
			t.Errorf(testTree4SizeError, i, result, tt.result)
		}
	}
}

var tree4ValueTests = []struct {
	tree    strg.Tree[strg.N4]
	n       uint
	result1 uint
	result2 bool
}{
	{tree: blank4, n: 0, result1: 0, result2: false},
	{tree: blank4, n: 1, result1: 0, result2: false},
	{tree: blank4, n: 100, result1: 0, result2: false},
	{tree: tooshort4, n: 0, result1: 0, result2: false},
	{tree: tooshort4, n: 1, result1: 0, result2: false},
	{tree: tooshort4, n: 100, result1: 0, result2: false},
	{tree: empty4, n: 0, result1: 0, result2: false},
	{tree: empty4, n: 1, result1: 0, result2: false},
	{tree: empty4, n: 100, result1: 0, result2: false},
	{tree: atree4, n: 0, result1: 0, result2: false},
	{tree: atree4, n: 1, result1: 4, result2: true},
	{tree: atree4, n: 2, result1: 0, result2: false},
	{tree: atree4, n: 3, result1: 2, result2: true},
	{tree: atree4, n: 4, result1: 3, result2: true},
	{tree: atree4, n: 5, result1: 5, result2: true},
	{tree: atree4, n: 6, result1: 6, result2: true},
	{tree: atree4, n: 7, result1: 7, result2: true},
	{tree: atree4, n: 8, result1: 0, result2: false},
	{tree: atree4, n: 9, result1: 0, result2: true},
	{tree: atree4, n: 10, result1: 1, result2: true},
	{tree: atree4, n: 100, result1: 0, result2: false},
}

const testTree4ValueError = "Tree4 Value Test %d: got %d and %t for value " +
	"of node %d (should be %d and %t)"

func TestTree4Value(t *testing.T) {
	for i, tt := range tree4ValueTests {
		result1, result2 := tt.tree.Value(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree4ValueError,
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

var tree4ChunkTests = []struct {
	tree   strg.Tree[strg.N4]
	n      uint
	result string
}{
	{tree: blank4, n: 0, result: ""},
	{tree: blank4, n: 1, result: ""},
	{tree: blank4, n: 100, result: ""},
	{tree: tooshort4, n: 0, result: ""},
	{tree: tooshort4, n: 1, result: ""},
	{tree: tooshort4, n: 100, result: ""},
	{tree: empty4, n: 0, result: ""},
	{tree: empty4, n: 1, result: ""},
	{tree: empty4, n: 100, result: ""},
	{tree: atree4, n: 0, result: ""},
	{tree: atree4, n: 1, result: "auth"},
	{tree: atree4, n: 2, result: "content-"},
	{tree: atree4, n: 3, result: "or"},
	{tree: atree4, n: 4, result: "entication"},
	{tree: atree4, n: 5, result: "type"},
	{tree: atree4, n: 6, result: "length"},
	{tree: atree4, n: 7, result: "disposition"},
	{tree: atree4, n: 8, result: "i"},
	{tree: atree4, n: 9, result: "ty"},
	{tree: atree4, n: 10, result: "zation"},
	{tree: atree4, n: 100, result: ""},
}

const testTree4ChunkError = "Tree4 Chunk Test %d: got '%s' for chunk of " +
	"node %d (should be '%s')"

func TestTree4Chunk(t *testing.T) {
	for i, tt := range tree4ChunkTests {
		result := tt.tree.Chunk(tt.n)
		if result != tt.result {
			t.Errorf(
				testTree4ChunkError,
				i,
				result,
				tt.n,
				tt.result,
			)
		}
	}
}

var tree4EachChildTests = []struct {
	tree    strg.Tree[strg.N4]
	n       uint
	f       func(radixt.Tree, uint) string
	indices string
}{
	{tree: blank4, n: 0, f: eachChild, indices: ""},
	{tree: blank4, n: 1, f: eachChild, indices: ""},
	{tree: blank4, n: 100, f: eachChild, indices: ""},
	{tree: blank4, n: 0, f: eachFirstChild, indices: ""},
	{tree: blank4, n: 1, f: eachFirstChild, indices: ""},
	{tree: blank4, n: 100, f: eachFirstChild, indices: ""},
	{tree: tooshort4, n: 0, f: eachChild, indices: ""},
	{tree: tooshort4, n: 1, f: eachChild, indices: ""},
	{tree: tooshort4, n: 100, f: eachChild, indices: ""},
	{tree: tooshort4, n: 0, f: eachFirstChild, indices: ""},
	{tree: tooshort4, n: 1, f: eachFirstChild, indices: ""},
	{tree: tooshort4, n: 100, f: eachFirstChild, indices: ""},
	{tree: empty4, n: 0, f: eachChild, indices: ""},
	{tree: empty4, n: 1, f: eachChild, indices: ""},
	{tree: empty4, n: 100, f: eachChild, indices: ""},
	{tree: empty4, n: 0, f: eachFirstChild, indices: ""},
	{tree: empty4, n: 1, f: eachFirstChild, indices: ""},
	{tree: empty4, n: 100, f: eachFirstChild, indices: ""},
	{tree: atree4, n: 0, f: eachChild, indices: "1, 2"},
	{tree: atree4, n: 1, f: eachChild, indices: "3, 4"},
	{tree: atree4, n: 2, f: eachChild, indices: "5, 6, 7"},
	{tree: atree4, n: 3, f: eachChild, indices: "8"},
	{tree: atree4, n: 4, f: eachChild, indices: ""},
	{tree: atree4, n: 5, f: eachChild, indices: ""},
	{tree: atree4, n: 6, f: eachChild, indices: ""},
	{tree: atree4, n: 7, f: eachChild, indices: ""},
	{tree: atree4, n: 8, f: eachChild, indices: "9, 10"},
	{tree: atree4, n: 9, f: eachChild, indices: ""},
	{tree: atree4, n: 10, f: eachChild, indices: ""},
	{tree: atree4, n: 100, f: eachChild, indices: ""},
	{tree: atree4, n: 0, f: eachFirstChild, indices: "1"},
	{tree: atree4, n: 1, f: eachFirstChild, indices: "3"},
	{tree: atree4, n: 2, f: eachFirstChild, indices: "5"},
	{tree: atree4, n: 3, f: eachFirstChild, indices: "8"},
	{tree: atree4, n: 4, f: eachFirstChild, indices: ""},
	{tree: atree4, n: 5, f: eachFirstChild, indices: ""},
	{tree: atree4, n: 6, f: eachFirstChild, indices: ""},
	{tree: atree4, n: 7, f: eachFirstChild, indices: ""},
	{tree: atree4, n: 8, f: eachFirstChild, indices: "9"},
	{tree: atree4, n: 9, f: eachFirstChild, indices: ""},
	{tree: atree4, n: 10, f: eachFirstChild, indices: ""},
	{tree: atree4, n: 100, f: eachFirstChild, indices: ""},
}

const testTree4EachChildError = "Tree4 Each Child Test %d: got %s as result " +
	"indices (should be %s)"

func TestTree4EachChild(t *testing.T) {
	for i, tt := range tree4EachChildTests {
		indices := tt.f(tt.tree, tt.n)
		if indices != tt.indices {
			t.Errorf(
				testTree4EachChildError,
				i,
				indices,
				tt.indices,
			)
		}
	}
}

var tree4HoardTests = []struct {
	tree    strg.Tree[strg.N4]
	result1 uint
	result2 uint
}{
	{tree: blank4, result1: 0, result2: radixt.HoardExactly},
	{tree: tooshort4, result1: 3, result2: radixt.HoardExactly},
	{tree: empty4, result1: 10, result2: radixt.HoardExactly},
	{tree: atree4, result1: 105, result2: radixt.HoardExactly},
}

const testTree4HoardError = "Tree4 Hoard Test %d: got %d and %d (should be " +
	"%d and %d)"

func TestTree4Hoard(t *testing.T) {
	for i, tt := range tree4HoardTests {
		result1, result2 := tt.tree.Hoard()
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree4HoardError,
				i,
				result1,
				result2,
				tt.result1,
				tt.result2,
			)
		}
	}
}

var tree4SwitchTests = []struct {
	switcher strg.Tree[strg.N4]
	n        uint
	b        byte
	result1  uint
	result2  string
	result3  bool
}{
	{
		switcher: blank4,
		n:        0,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: blank4,
		n:        1,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: blank4,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: tooshort4,
		n:        0,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: tooshort4,
		n:        1,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: tooshort4,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty4,
		n:        0,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty4,
		n:        1,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: empty4,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        0,
		b:        97,
		result1:  1,
		result2:  "uth",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        0,
		b:        99,
		result1:  2,
		result2:  "ontent-",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        0,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        1,
		b:        101,
		result1:  4,
		result2:  "ntication",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        1,
		b:        111,
		result1:  3,
		result2:  "r",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        1,
		b:        112,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        2,
		b:        116,
		result1:  5,
		result2:  "ype",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        2,
		b:        108,
		result1:  6,
		result2:  "ength",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        2,
		b:        100,
		result1:  7,
		result2:  "isposition",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        2,
		b:        99,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        3,
		b:        105,
		result1:  8,
		result2:  "",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        3,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        4,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        4,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        5,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        5,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        6,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        6,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        7,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        7,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        8,
		b:        116,
		result1:  9,
		result2:  "y",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        8,
		b:        122,
		result1:  10,
		result2:  "ation",
		result3:  true,
	},
	{
		switcher: atree4,
		n:        8,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        9,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        9,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        10,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        10,
		b:        98,
		result1:  0,
		result2:  "",
		result3:  false,
	},
	{
		switcher: atree4,
		n:        100,
		b:        97,
		result1:  0,
		result2:  "",
		result3:  false,
	},
}

const testTree4SwitchError = "Tree4 Switch Test %d: got %d, '%s', and %t, " +
	"trying to switch from node %d by byte %d (should be %d, '%s', and %t)"

func TestTree4Switch(t *testing.T) {
	for i, tt := range tree4SwitchTests {
		result1, result2, result3 := tt.switcher.Switch(tt.n, tt.b)

		e := result1 != tt.result1 ||
			result2 != tt.result2 ||
			result3 != tt.result3

		if e {
			t.Errorf(
				testTree4SwitchError,
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
