package analysis

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/null"
	"github.com/alex-ilchukov/radixt/sapling"
)

var (
	empty = sapling.New()

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

	methods = sapling.New(
		"DELETE",
		"GET",
		"HEAD",
		"OPTIONS",
		"PATCH",
		"POST",
		"PUT",
		"TRACE",
	)
)

var doDefaultTests = []struct {
	tree   radixt.Tree
	result A[Default]
}{
	{
		tree: nil,
		result: A[Default]{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     []N[Default]{},
		},
	},
	{
		tree: null.Tree,
		result: A[Default]{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     []N[Default]{},
		},
	},
	{
		tree: empty,
		result: A[Default]{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     []N[Default]{},
		},
	},
	{
		tree: atree,
		result: A[Default]{
			C: "dispositionenticationcontent-lengthzationauth" +
				"typeor",
			Cml:   11,
			Cma:   3,
			Dclpm: 4,
			Vm:    7,
			N: []N[Default]{
				{
					HasValue:     false,
					ChunkFirst:   0,
					ChunkEmpty:   true,
					Index:        0,
					Chunk:        "",
					Value:        0,
					Parent:       0,
					ChildrenLow:  1,
					ChildrenHigh: 3,
					ChunkPos:     0,
				},
				{
					HasValue:     true,
					ChunkFirst:   't',
					ChunkEmpty:   false,
					Index:        9,
					Chunk:        "ty",
					Value:        0,
					Parent:       8,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     45,
				},
				{
					HasValue:     true,
					ChunkFirst:   'z',
					ChunkEmpty:   false,
					Index:        10,
					Chunk:        "zation",
					Value:        1,
					Parent:       8,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     35,
				},
				{
					HasValue:     false,
					ChunkFirst:   'i',
					ChunkEmpty:   false,
					Index:        8,
					Chunk:        "i",
					Value:        0,
					Parent:       4,
					ChildrenLow:  9,
					ChildrenHigh: 11,
					ChunkPos:     1,
				},
				{
					HasValue:     true,
					ChunkFirst:   'o',
					ChunkEmpty:   false,
					Index:        4,
					Chunk:        "or",
					Value:        2,
					Parent:       1,
					ChildrenLow:  8,
					ChildrenHigh: 9,
					ChunkPos:     49,
				},
				{
					HasValue:     true,
					ChunkFirst:   'e',
					ChunkEmpty:   false,
					Index:        3,
					Chunk:        "entication",
					Value:        3,
					Parent:       1,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     11,
				},
				{
					HasValue:     true,
					ChunkFirst:   'a',
					ChunkEmpty:   false,
					Index:        1,
					Chunk:        "auth",
					Value:        4,
					Parent:       0,
					ChildrenLow:  3,
					ChildrenHigh: 5,
					ChunkPos:     41,
				},
				{
					HasValue:     false,
					ChunkFirst:   'c',
					ChunkEmpty:   false,
					Index:        2,
					Chunk:        "content-",
					Value:        0,
					Parent:       0,
					ChildrenLow:  5,
					ChildrenHigh: 8,
					ChunkPos:     21,
				},
				{
					HasValue:     true,
					ChunkFirst:   't',
					ChunkEmpty:   false,
					Index:        7,
					Chunk:        "type",
					Value:        5,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     45,
				},
				{
					HasValue:     true,
					ChunkFirst:   'l',
					ChunkEmpty:   false,
					Index:        6,
					Chunk:        "length",
					Value:        6,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     29,
				},
				{
					HasValue:     true,
					ChunkFirst:   'd',
					ChunkEmpty:   false,
					Index:        5,
					Chunk:        "disposition",
					Value:        7,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     0,
				},
			},
		},
	},
	{
		tree: methods,
		result: A[Default]{
			C:     "OPTIONSDELETETRACEATCHHEADGETOSTUT",
			Cml:   7,
			Cma:   6,
			Dclpm: 2,
			Vm:    7,
			N: []N[Default]{
				{
					HasValue:     false,
					ChunkFirst:   0,
					ChunkEmpty:   true,
					Index:        0,
					Chunk:        "",
					Value:        0,
					Parent:       0,
					ChildrenLow:  1,
					ChildrenHigh: 7,
					ChunkPos:     0,
				},
				{
					HasValue:     true,
					ChunkFirst:   'D',
					ChunkEmpty:   false,
					Index:        1,
					Chunk:        "DELETE",
					Value:        0,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     7,
				},
				{
					HasValue:     true,
					ChunkFirst:   'G',
					ChunkEmpty:   false,
					Index:        2,
					Chunk:        "GET",
					Value:        1,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     26,
				},
				{
					HasValue:     true,
					ChunkFirst:   'H',
					ChunkEmpty:   false,
					Index:        3,
					Chunk:        "HEAD",
					Value:        2,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     22,
				},
				{
					HasValue:     true,
					ChunkFirst:   'O',
					ChunkEmpty:   false,
					Index:        4,
					Chunk:        "OPTIONS",
					Value:        3,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     0,
				},
				{
					HasValue:     false,
					ChunkFirst:   'P',
					ChunkEmpty:   false,
					Index:        5,
					Chunk:        "P",
					Value:        0,
					Parent:       0,
					ChildrenLow:  7,
					ChildrenHigh: 10,
					ChunkPos:     1,
				},
				{
					HasValue:     true,
					ChunkFirst:   'A',
					ChunkEmpty:   false,
					Index:        7,
					Chunk:        "ATCH",
					Value:        4,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     18,
				},
				{
					HasValue:     true,
					ChunkFirst:   'O',
					ChunkEmpty:   false,
					Index:        8,
					Chunk:        "OST",
					Value:        5,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     29,
				},
				{
					HasValue:     true,
					ChunkFirst:   'U',
					ChunkEmpty:   false,
					Index:        9,
					Chunk:        "UT",
					Value:        6,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     32,
				},
				{
					HasValue:     true,
					ChunkFirst:   'T',
					ChunkEmpty:   false,
					Index:        6,
					Chunk:        "TRACE",
					Value:        7,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     13,
				},
			},
		},
	},
}

func guilty[M Mode](a, b A[M]) string {
	switch {
	case a.C != b.C:
		return "C"

	case a.Cml != b.Cml:
		return "Cml"

	case a.Cma != b.Cma:
		return "Cma"

	case a.Dclpm != b.Dclpm:
		return "Dclpm"

	case a.Vm != b.Vm:
		return "Vm"

	case !reflect.DeepEqual(a.N, b.N):
		if len(a.N) != len(b.N) {
			return "different lengths of N"
		}

		for k, v := range a.N {
			v1 := b.N[k]
			if v != v1 {
				return fmt.Sprintf(
					"N and values at index %d: %v != %v",
					k,
					v,
					v1,
				)
			}
		}
	}

	return ""
}

const doTestDefaultError = "Do[Default] Test %d: got %v for result (should " +
	"be %v, look at %v)"

func TestDoDefault(t *testing.T) {
	for i, tt := range doDefaultTests {
		result := Do[Default](tt.tree)
		if !reflect.DeepEqual(result, tt.result) {
			g := guilty(result, tt.result)
			t.Errorf(doTestDefaultError, i, result, tt.result, g)
		}
	}
}

var doFirstlessTests = []struct {
	tree   radixt.Tree
	result A[Firstless]
}{
	{
		tree: nil,
		result: A[Firstless]{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     []N[Firstless]{},
		},
	},
	{
		tree: null.Tree,
		result: A[Firstless]{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     []N[Firstless]{},
		},
	},
	{
		tree: empty,
		result: A[Firstless]{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     []N[Firstless]{},
		},
	},
	{
		tree: atree,
		result: A[Firstless]{
			C: "ispositionnticationontent-engthuthyper",
			Cml:   10,
			Cma:   3,
			Dclpm: 4,
			Vm:    7,
			N: []N[Firstless]{
				{
					HasValue:     false,
					ChunkFirst:   0,
					ChunkEmpty:   true,
					Index:        0,
					Chunk:        "",
					Value:        0,
					Parent:       0,
					ChildrenLow:  1,
					ChildrenHigh: 3,
					ChunkPos:     0,
				},
				{
					HasValue:     true,
					ChunkFirst:   't',
					ChunkEmpty:   false,
					Index:        9,
					Chunk:        "y",
					Value:        0,
					Parent:       8,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     34,
				},
				{
					HasValue:     true,
					ChunkFirst:   'z',
					ChunkEmpty:   false,
					Index:        10,
					Chunk:        "ation",
					Value:        1,
					Parent:       8,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     14,
				},
				{
					HasValue:     false,
					ChunkFirst:   'i',
					ChunkEmpty:   false,
					Index:        8,
					Chunk:        "",
					Value:        0,
					Parent:       4,
					ChildrenLow:  9,
					ChildrenHigh: 11,
					ChunkPos:     0,
				},
				{
					HasValue:     true,
					ChunkFirst:   'o',
					ChunkEmpty:   false,
					Index:        4,
					Chunk:        "r",
					Value:        2,
					Parent:       1,
					ChildrenLow:  8,
					ChildrenHigh: 9,
					ChunkPos:     37,
				},
				{
					HasValue:     true,
					ChunkFirst:   'e',
					ChunkEmpty:   false,
					Index:        3,
					Chunk:        "ntication",
					Value:        3,
					Parent:       1,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     10,
				},
				{
					HasValue:     true,
					ChunkFirst:   'a',
					ChunkEmpty:   false,
					Index:        1,
					Chunk:        "uth",
					Value:        4,
					Parent:       0,
					ChildrenLow:  3,
					ChildrenHigh: 5,
					ChunkPos:     31,
				},
				{
					HasValue:     false,
					ChunkFirst:   'c',
					ChunkEmpty:   false,
					Index:        2,
					Chunk:        "ontent-",
					Value:        0,
					Parent:       0,
					ChildrenLow:  5,
					ChildrenHigh: 8,
					ChunkPos:     19,
				},
				{
					HasValue:     true,
					ChunkFirst:   't',
					ChunkEmpty:   false,
					Index:        7,
					Chunk:        "ype",
					Value:        5,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     34,
				},
				{
					HasValue:     true,
					ChunkFirst:   'l',
					ChunkEmpty:   false,
					Index:        6,
					Chunk:        "ength",
					Value:        6,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     26,
				},
				{
					HasValue:     true,
					ChunkFirst:   'd',
					ChunkEmpty:   false,
					Index:        5,
					Chunk:        "isposition",
					Value:        7,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     0,
				},
			},
		},
	},
	{
		tree: methods,
		result: A[Firstless]{
			C:     "PTIONSELETERACEEADTCHST",
			Cml:   6,
			Cma:   6,
			Dclpm: 2,
			Vm:    7,
			N: []N[Firstless]{
				{
					HasValue:     false,
					ChunkFirst:   0,
					ChunkEmpty:   true,
					Index:        0,
					Chunk:        "",
					Value:        0,
					Parent:       0,
					ChildrenLow:  1,
					ChildrenHigh: 7,
					ChunkPos:     0,
				},
				{
					HasValue:     true,
					ChunkFirst:   'D',
					ChunkEmpty:   false,
					Index:        1,
					Chunk:        "ELETE",
					Value:        0,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     6,
				},
				{
					HasValue:     true,
					ChunkFirst:   'G',
					ChunkEmpty:   false,
					Index:        2,
					Chunk:        "ET",
					Value:        1,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     8,
				},
				{
					HasValue:     true,
					ChunkFirst:   'H',
					ChunkEmpty:   false,
					Index:        3,
					Chunk:        "EAD",
					Value:        2,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     15,
				},
				{
					HasValue:     true,
					ChunkFirst:   'O',
					ChunkEmpty:   false,
					Index:        4,
					Chunk:        "PTIONS",
					Value:        3,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     0,
				},
				{
					HasValue:     false,
					ChunkFirst:   'P',
					ChunkEmpty:   false,
					Index:        5,
					Chunk:        "",
					Value:        0,
					Parent:       0,
					ChildrenLow:  7,
					ChildrenHigh: 10,
					ChunkPos:     0,
				},
				{
					HasValue:     true,
					ChunkFirst:   'A',
					ChunkEmpty:   false,
					Index:        7,
					Chunk:        "TCH",
					Value:        4,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     18,
				},
				{
					HasValue:     true,
					ChunkFirst:   'O',
					ChunkEmpty:   false,
					Index:        8,
					Chunk:        "ST",
					Value:        5,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     21,
				},
				{
					HasValue:     true,
					ChunkFirst:   'U',
					ChunkEmpty:   false,
					Index:        9,
					Chunk:        "T",
					Value:        6,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     1,
				},
				{
					HasValue:     true,
					ChunkFirst:   'T',
					ChunkEmpty:   false,
					Index:        6,
					Chunk:        "RACE",
					Value:        7,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     11,
				},
			},
		},
	},
}

const doTestFirstlessError = "Do[Firstless] Test %d: got %v for result " +
	"(should be %v, look at %v)"

func TestDoFirstless(t *testing.T) {
	for i, tt := range doFirstlessTests {
		result := Do[Firstless](tt.tree)
		if !reflect.DeepEqual(result, tt.result) {
			g := guilty(result, tt.result)
			t.Errorf(doTestFirstlessError, i, result, tt.result, g)
		}
	}
}
