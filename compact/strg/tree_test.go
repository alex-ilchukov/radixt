package strg_test

import (
	"testing"

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

var tree3ChildrenRangeTests = []struct {
	tree    strg.Tree[strg.N3]
	n       uint
	result1 uint
	result2 uint
}{
	{tree: blank3, n: 0, result1: 0, result2: 0},
	{tree: blank3, n: 1, result1: 0, result2: 0},
	{tree: blank3, n: 100, result1: 0, result2: 0},
	{tree: tooshort3, n: 0, result1: 0, result2: 0},
	{tree: tooshort3, n: 1, result1: 0, result2: 0},
	{tree: tooshort3, n: 100, result1: 0, result2: 0},
	{tree: empty3, n: 0, result1: 0, result2: 0},
	{tree: empty3, n: 1, result1: 0, result2: 0},
	{tree: empty3, n: 100, result1: 0, result2: 0},
	{tree: atree3, n: 0, result1: 1, result2: 3},
	{tree: atree3, n: 1, result1: 3, result2: 5},
	{tree: atree3, n: 2, result1: 5, result2: 8},
	{tree: atree3, n: 3, result1: 8, result2: 9},
	{tree: atree3, n: 4, result1: 0, result2: 0},
	{tree: atree3, n: 5, result1: 0, result2: 0},
	{tree: atree3, n: 6, result1: 0, result2: 0},
	{tree: atree3, n: 7, result1: 0, result2: 0},
	{tree: atree3, n: 8, result1: 9, result2: 11},
	{tree: atree3, n: 9, result1: 0, result2: 0},
	{tree: atree3, n: 10, result1: 0, result2: 0},
	{tree: atree3, n: 100, result1: 0, result2: 0},
}

const testTree3ChildrenRangeError = "Tree3 Children Range Test %d: got %d " +
	"and %d for low and high indices of children of node %d (should be " +
	"%d and %d)"

func TestTree3ChildrenRange(t *testing.T) {
	for i, tt := range tree3ChildrenRangeTests {
		result1, result2 := tt.tree.ChildrenRange(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree3ChildrenRangeError,
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

var tree4ChildrenRangeTests = []struct {
	tree    strg.Tree[strg.N4]
	n       uint
	result1 uint
	result2 uint
}{
	{tree: blank4, n: 0, result1: 0, result2: 0},
	{tree: blank4, n: 1, result1: 0, result2: 0},
	{tree: blank4, n: 100, result1: 0, result2: 0},
	{tree: tooshort4, n: 0, result1: 0, result2: 0},
	{tree: tooshort4, n: 1, result1: 0, result2: 0},
	{tree: tooshort4, n: 100, result1: 0, result2: 0},
	{tree: empty4, n: 0, result1: 0, result2: 0},
	{tree: empty4, n: 1, result1: 0, result2: 0},
	{tree: empty4, n: 100, result1: 0, result2: 0},
	{tree: atree4, n: 0, result1: 1, result2: 3},
	{tree: atree4, n: 1, result1: 3, result2: 5},
	{tree: atree4, n: 2, result1: 5, result2: 8},
	{tree: atree4, n: 3, result1: 8, result2: 9},
	{tree: atree4, n: 4, result1: 0, result2: 0},
	{tree: atree4, n: 5, result1: 0, result2: 0},
	{tree: atree4, n: 6, result1: 0, result2: 0},
	{tree: atree4, n: 7, result1: 0, result2: 0},
	{tree: atree4, n: 8, result1: 9, result2: 11},
	{tree: atree4, n: 9, result1: 0, result2: 0},
	{tree: atree4, n: 10, result1: 0, result2: 0},
	{tree: atree4, n: 100, result1: 0, result2: 0},
}

const testTree4ChildrenRangeError = "Tree4 Children Range Test %d: " +
	"got %d and %d for low and high indices of children of node %d " +
	"(should be %d and %d)"

func TestTree4ChildrenRange(t *testing.T) {
	for i, tt := range tree4ChildrenRangeTests {
		result1, result2 := tt.tree.ChildrenRange(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree4ChildrenRangeError,
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
