package structg_test

import (
	"strconv"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/structg"
	"github.com/alex-ilchukov/radixt/generic"
)

var (
	empty32 = structg.MustCreate[uint32](nil)

	atree32 = structg.MustCreate[uint32](
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

	empty64 = structg.MustCreate[uint64](nil)

	atree64 = structg.MustCreate[uint64](
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

var tree32SizeTests = []struct {
	tree   radixt.Tree
	result uint
}{
	{tree: empty32, result: 0},
	{tree: atree32, result: 11},
}

const testTree32SizeError = "Tree[uint32] Size Test %d: got %d for size " +
	"(should be %d)"

func TestTree32Size(t *testing.T) {
	for i, tt := range tree32SizeTests {
		result := tt.tree.Size()
		if result != tt.result {
			t.Errorf(testTree32SizeError, i, result, tt.result)
		}
	}
}

var tree32ValueTests = []struct {
	tree    radixt.Tree
	n       uint
	result1 uint
	result2 bool
}{
	{tree: empty32, n: 0, result1: 0, result2: false},
	{tree: empty32, n: 1, result1: 0, result2: false},
	{tree: empty32, n: 100, result1: 0, result2: false},
	{tree: atree32, n: 0, result1: 0, result2: false},
	{tree: atree32, n: 1, result1: 4, result2: true},
	{tree: atree32, n: 2, result1: 0, result2: false},
	{tree: atree32, n: 3, result1: 2, result2: true},
	{tree: atree32, n: 4, result1: 3, result2: true},
	{tree: atree32, n: 5, result1: 5, result2: true},
	{tree: atree32, n: 6, result1: 6, result2: true},
	{tree: atree32, n: 7, result1: 7, result2: true},
	{tree: atree32, n: 8, result1: 0, result2: false},
	{tree: atree32, n: 9, result1: 0, result2: true},
	{tree: atree32, n: 10, result1: 1, result2: true},
	{tree: atree32, n: 100, result1: 0, result2: false},
}

const testTree32ValueError = "Tree[uint32] Value Test %d: got %d and %t for " +
	"value of node %d (should be %d and %t)"

func TestTree32Value(t *testing.T) {
	for i, tt := range tree32ValueTests {
		result1, result2 := tt.tree.Value(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree32ValueError,
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

var tree32ChunkTests = []struct {
	tree   radixt.Tree
	n      uint
	result string
}{
	{tree: empty32, n: 0, result: ""},
	{tree: empty32, n: 1, result: ""},
	{tree: empty32, n: 100, result: ""},
	{tree: atree32, n: 0, result: ""},
	{tree: atree32, n: 1, result: "auth"},
	{tree: atree32, n: 2, result: "content-"},
	{tree: atree32, n: 3, result: "or"},
	{tree: atree32, n: 4, result: "entication"},
	{tree: atree32, n: 5, result: "type"},
	{tree: atree32, n: 6, result: "length"},
	{tree: atree32, n: 7, result: "disposition"},
	{tree: atree32, n: 8, result: "i"},
	{tree: atree32, n: 9, result: "ty"},
	{tree: atree32, n: 10, result: "zation"},
	{tree: atree32, n: 100, result: ""},
}

const testTree32ChunkError = "Tree[uint32] Chunk Test %d: got '%s' for " +
	"chunk of node %d (should be '%s')"

func TestTree32Chunk(t *testing.T) {
	for i, tt := range tree32ChunkTests {
		result := tt.tree.Chunk(tt.n)
		if result != tt.result {
			t.Errorf(
				testTree32ChunkError,
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

var tree32EachChildTests = []struct {
	tree    radixt.Tree
	n       uint
	f       func(radixt.Tree, uint) string
	indices string
}{
	{tree: empty32, n: 0, f: eachChild, indices: ""},
	{tree: empty32, n: 1, f: eachChild, indices: ""},
	{tree: empty32, n: 100, f: eachChild, indices: ""},
	{tree: empty32, n: 0, f: eachFirstChild, indices: ""},
	{tree: empty32, n: 1, f: eachFirstChild, indices: ""},
	{tree: empty32, n: 100, f: eachFirstChild, indices: ""},
	{tree: atree32, n: 0, f: eachChild, indices: "1, 2"},
	{tree: atree32, n: 1, f: eachChild, indices: "3, 4"},
	{tree: atree32, n: 2, f: eachChild, indices: "5, 6, 7"},
	{tree: atree32, n: 3, f: eachChild, indices: "8"},
	{tree: atree32, n: 4, f: eachChild, indices: ""},
	{tree: atree32, n: 5, f: eachChild, indices: ""},
	{tree: atree32, n: 6, f: eachChild, indices: ""},
	{tree: atree32, n: 7, f: eachChild, indices: ""},
	{tree: atree32, n: 8, f: eachChild, indices: "9, 10"},
	{tree: atree32, n: 9, f: eachChild, indices: ""},
	{tree: atree32, n: 10, f: eachChild, indices: ""},
	{tree: atree32, n: 100, f: eachChild, indices: ""},
	{tree: atree32, n: 0, f: eachFirstChild, indices: "1"},
	{tree: atree32, n: 1, f: eachFirstChild, indices: "3"},
	{tree: atree32, n: 2, f: eachFirstChild, indices: "5"},
	{tree: atree32, n: 3, f: eachFirstChild, indices: "8"},
	{tree: atree32, n: 4, f: eachFirstChild, indices: ""},
	{tree: atree32, n: 5, f: eachFirstChild, indices: ""},
	{tree: atree32, n: 6, f: eachFirstChild, indices: ""},
	{tree: atree32, n: 7, f: eachFirstChild, indices: ""},
	{tree: atree32, n: 8, f: eachFirstChild, indices: "9"},
	{tree: atree32, n: 9, f: eachFirstChild, indices: ""},
	{tree: atree32, n: 10, f: eachFirstChild, indices: ""},
	{tree: atree32, n: 100, f: eachFirstChild, indices: ""},
}

const testTree32EachChildError = "Tree[uint32] Each Child Test %d: got %s " +
	"as result indices (should be %s)"

func TestTree32EachChild(t *testing.T) {
	for i, tt := range tree32EachChildTests {
		indices := tt.f(tt.tree, tt.n)
		if indices != tt.indices {
			t.Errorf(
				testTree32EachChildError,
				i,
				indices,
				tt.indices,
			)
		}
	}
}

var tree32HoardTests = []struct {
	tree    radixt.Hoarder
	result1 uint
	result2 uint
}{
	{tree: empty32, result1: 48, result2: radixt.HoardExactly},
	{tree: atree32, result1: 143, result2: radixt.HoardExactly},
}

const testTree32HoardError = "Tree[uint32] Hoard Test %d: got %d and %d " +
	"(should be %d and %d)"

func TestTree32Hoard(t *testing.T) {
	for i, tt := range tree32HoardTests {
		result1, result2 := tt.tree.Hoard()
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree32HoardError,
				i,
				result1,
				result2,
				tt.result1,
				tt.result2,
			)
		}
	}
}

var tree64SizeTests = []struct {
	tree   radixt.Tree
	result uint
}{
	{tree: empty64, result: 0},
	{tree: atree64, result: 11},
}

const testTree64SizeError = "Tree[uint64] Size Test %d: got %d for size " +
	"(should be %d)"

func TestTree64Size(t *testing.T) {
	for i, tt := range tree64SizeTests {
		result := tt.tree.Size()
		if result != tt.result {
			t.Errorf(testTree64SizeError, i, result, tt.result)
		}
	}
}

var tree64ValueTests = []struct {
	tree    radixt.Tree
	n       uint
	result1 uint
	result2 bool
}{
	{tree: empty64, n: 0, result1: 0, result2: false},
	{tree: empty64, n: 1, result1: 0, result2: false},
	{tree: empty64, n: 100, result1: 0, result2: false},
	{tree: atree64, n: 0, result1: 0, result2: false},
	{tree: atree64, n: 1, result1: 4, result2: true},
	{tree: atree64, n: 2, result1: 0, result2: false},
	{tree: atree64, n: 3, result1: 2, result2: true},
	{tree: atree64, n: 4, result1: 3, result2: true},
	{tree: atree64, n: 5, result1: 5, result2: true},
	{tree: atree64, n: 6, result1: 6, result2: true},
	{tree: atree64, n: 7, result1: 7, result2: true},
	{tree: atree64, n: 8, result1: 0, result2: false},
	{tree: atree64, n: 9, result1: 0, result2: true},
	{tree: atree64, n: 10, result1: 1, result2: true},
	{tree: atree64, n: 100, result1: 0, result2: false},
}

const testTree64ValueError = "Tree[uint64] Value Test %d: got %d and %t for " +
	"value of node %d (should be %d and %t)"

func TestTree64Value(t *testing.T) {
	for i, tt := range tree64ValueTests {
		result1, result2 := tt.tree.Value(tt.n)
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree64ValueError,
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

var tree64ChunkTests = []struct {
	tree   radixt.Tree
	n      uint
	result string
}{
	{tree: empty64, n: 0, result: ""},
	{tree: empty64, n: 1, result: ""},
	{tree: empty64, n: 100, result: ""},
	{tree: atree64, n: 0, result: ""},
	{tree: atree64, n: 1, result: "auth"},
	{tree: atree64, n: 2, result: "content-"},
	{tree: atree64, n: 3, result: "or"},
	{tree: atree64, n: 4, result: "entication"},
	{tree: atree64, n: 5, result: "type"},
	{tree: atree64, n: 6, result: "length"},
	{tree: atree64, n: 7, result: "disposition"},
	{tree: atree64, n: 8, result: "i"},
	{tree: atree64, n: 9, result: "ty"},
	{tree: atree64, n: 10, result: "zation"},
	{tree: atree64, n: 100, result: ""},
}

const testTree64ChunkError = "Tree[uint64] Chunk Test %d: got '%s' for " +
	"chunk of node %d (should be '%s')"

func TestTree64Chunk(t *testing.T) {
	for i, tt := range tree64ChunkTests {
		result := tt.tree.Chunk(tt.n)
		if result != tt.result {
			t.Errorf(
				testTree64ChunkError,
				i,
				result,
				tt.n,
				tt.result,
			)
		}
	}
}

var tree64EachChildTests = []struct {
	tree    radixt.Tree
	n       uint
	f       func(radixt.Tree, uint) string
	indices string
}{
	{tree: empty64, n: 0, f: eachChild, indices: ""},
	{tree: empty64, n: 1, f: eachChild, indices: ""},
	{tree: empty64, n: 100, f: eachChild, indices: ""},
	{tree: empty64, n: 0, f: eachFirstChild, indices: ""},
	{tree: empty64, n: 1, f: eachFirstChild, indices: ""},
	{tree: empty64, n: 100, f: eachFirstChild, indices: ""},
	{tree: atree64, n: 0, f: eachChild, indices: "1, 2"},
	{tree: atree64, n: 1, f: eachChild, indices: "3, 4"},
	{tree: atree64, n: 2, f: eachChild, indices: "5, 6, 7"},
	{tree: atree64, n: 3, f: eachChild, indices: "8"},
	{tree: atree64, n: 4, f: eachChild, indices: ""},
	{tree: atree64, n: 5, f: eachChild, indices: ""},
	{tree: atree64, n: 6, f: eachChild, indices: ""},
	{tree: atree64, n: 7, f: eachChild, indices: ""},
	{tree: atree64, n: 8, f: eachChild, indices: "9, 10"},
	{tree: atree64, n: 9, f: eachChild, indices: ""},
	{tree: atree64, n: 10, f: eachChild, indices: ""},
	{tree: atree64, n: 100, f: eachChild, indices: ""},
	{tree: atree64, n: 0, f: eachFirstChild, indices: "1"},
	{tree: atree64, n: 1, f: eachFirstChild, indices: "3"},
	{tree: atree64, n: 2, f: eachFirstChild, indices: "5"},
	{tree: atree64, n: 3, f: eachFirstChild, indices: "8"},
	{tree: atree64, n: 4, f: eachFirstChild, indices: ""},
	{tree: atree64, n: 5, f: eachFirstChild, indices: ""},
	{tree: atree64, n: 6, f: eachFirstChild, indices: ""},
	{tree: atree64, n: 7, f: eachFirstChild, indices: ""},
	{tree: atree64, n: 8, f: eachFirstChild, indices: "9"},
	{tree: atree64, n: 9, f: eachFirstChild, indices: ""},
	{tree: atree64, n: 10, f: eachFirstChild, indices: ""},
	{tree: atree64, n: 100, f: eachFirstChild, indices: ""},
}

const testTree64EachChildError = "Tree[uint64] Each Child Test %d: got %s " +
	"as result indices (should be %s)"

func TestTree64EachChild(t *testing.T) {
	for i, tt := range tree64EachChildTests {
		indices := tt.f(tt.tree, tt.n)
		if indices != tt.indices {
			t.Errorf(
				testTree64EachChildError,
				i,
				indices,
				tt.indices,
			)
		}
	}
}

var tree64HoardTests = []struct {
	tree    radixt.Hoarder
	result1 uint
	result2 uint
}{
	{tree: empty64, result1: 48, result2: radixt.HoardExactly},
	{tree: atree64, result1: 187, result2: radixt.HoardExactly},
}

const testTree64HoardError = "Tree[uint64] Hoard Test %d: got %d and %d " +
	"(should be %d and %d)"

func TestTree64Hoard(t *testing.T) {
	for i, tt := range tree64HoardTests {
		result1, result2 := tt.tree.Hoard()
		if result1 != tt.result1 || result2 != tt.result2 {
			t.Errorf(
				testTree64HoardError,
				i,
				result1,
				result2,
				tt.result1,
				tt.result2,
			)
		}
	}
}
