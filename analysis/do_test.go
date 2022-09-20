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

var doTests = []struct {
	tree   radixt.Tree
	result A
}{
	{
		tree: nil,
		result: A{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     map[uint]N{},
			Ca:    map[uint]uint{},
		},
	},
	{
		tree: null.Tree,
		result: A{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     map[uint]N{},
			Ca:    map[uint]uint{},
		},
	},
	{
		tree: empty,
		result: A{
			C:     "",
			Cml:   0,
			Cma:   0,
			Dclpm: 0,
			Vm:    0,
			N:     map[uint]N{},
			Ca:    map[uint]uint{},
		},
	},
	{
		tree: atree,
		result: A{
			C: "dispositionenticationcontent-lengthzationauth" +
				"typeor",
			Cml:   11,
			Cma:   3,
			Dclpm: 5,
			Vm:    7,
			N: map[uint]N{
				0: {
					Index:        0,
					Chunk:        "",
					Value:        0,
					HasValue:     false,
					Parent:       0,
					ChildrenLow:  1,
					ChildrenHigh: 3,
					ChunkPos:     0,
				},
				1: {
					Index:        9,
					Chunk:        "ty",
					Value:        0,
					HasValue:     true,
					Parent:       8,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     45,
				},
				2: {
					Index:        10,
					Chunk:        "zation",
					Value:        1,
					HasValue:     true,
					Parent:       8,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     35,
				},
				3: {
					Index:        8,
					Chunk:        "i",
					Value:        0,
					HasValue:     false,
					Parent:       3,
					ChildrenLow:  9,
					ChildrenHigh: 11,
					ChunkPos:     1,
				},
				4: {
					Index:        3,
					Chunk:        "or",
					Value:        2,
					HasValue:     true,
					Parent:       1,
					ChildrenLow:  8,
					ChildrenHigh: 9,
					ChunkPos:     49,
				},
				5: {
					Index:        4,
					Chunk:        "entication",
					Value:        3,
					HasValue:     true,
					Parent:       1,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     11,
				},
				6: {
					Index:        1,
					Chunk:        "auth",
					Value:        4,
					HasValue:     true,
					Parent:       0,
					ChildrenLow:  3,
					ChildrenHigh: 5,
					ChunkPos:     41,
				},
				7: {
					Index:        2,
					Chunk:        "content-",
					Value:        0,
					HasValue:     false,
					Parent:       0,
					ChildrenLow:  5,
					ChildrenHigh: 8,
					ChunkPos:     21,
				},
				8: {
					Index:        5,
					Chunk:        "type",
					Value:        5,
					HasValue:     true,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     45,
				},
				9: {
					Index:        6,
					Chunk:        "length",
					Value:        6,
					HasValue:     true,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     29,
				},
				10: {
					Index:        7,
					Chunk:        "disposition",
					Value:        7,
					HasValue:     true,
					Parent:       2,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     0,
				},
			},
			Ca: map[uint]uint{0: 6, 1: 1, 2: 3, 3: 1},
		},
	},
	{
		tree: methods,
		result: A{
			C:     "OPTIONSDELETETRACEATCHHEADGETOSTUT",
			Cml:   7,
			Cma:   6,
			Dclpm: 2,
			Vm:    7,
			N: map[uint]N{
				0: {
					Index:        0,
					Chunk:        "",
					Value:        0,
					HasValue:     false,
					Parent:       0,
					ChildrenLow:  1,
					ChildrenHigh: 7,
					ChunkPos:     0,
				},
				1: {
					Index:        1,
					Chunk:        "DELETE",
					Value:        0,
					HasValue:     true,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     7,
				},
				2: {
					Index:        2,
					Chunk:        "GET",
					Value:        1,
					HasValue:     true,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     26,
				},
				3: {
					Index:        3,
					Chunk:        "HEAD",
					Value:        2,
					HasValue:     true,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     22,
				},
				4: {
					Index:        4,
					Chunk:        "OPTIONS",
					Value:        3,
					HasValue:     true,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     0,
				},
				5: {
					Index:        5,
					Chunk:        "P",
					Value:        0,
					HasValue:     false,
					Parent:       0,
					ChildrenLow:  7,
					ChildrenHigh: 10,
					ChunkPos:     1,
				},
				6: {
					Index:        7,
					Chunk:        "ATCH",
					Value:        4,
					HasValue:     true,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     18,
				},
				7: {
					Index:        8,
					Chunk:        "OST",
					Value:        5,
					HasValue:     true,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     29,
				},
				8: {
					Index:        9,
					Chunk:        "UT",
					Value:        6,
					HasValue:     true,
					Parent:       5,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     32,
				},
				9: {
					Index:        6,
					Chunk:        "TRACE",
					Value:        7,
					HasValue:     true,
					Parent:       0,
					ChildrenLow:  0,
					ChildrenHigh: 0,
					ChunkPos:     13,
				},
			},
			Ca: map[uint]uint{0: 8, 3: 1, 6: 1},
		},
	},
}

func guilty(a, b A) string {
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

	case !reflect.DeepEqual(a.Ca, b.Ca):
		for k, v := range a.Ca {
			v1, has := b.Ca[k]
			switch {
			case !has:
				return fmt.Sprintf(
					"Ca and unnessesary key %d",
					k,
				)
			case v != v1:
				return fmt.Sprintf(
					"Ca and values at key %d: %d != %d",
					k,
					v,
					v1,
				)
			}
		}

		for k, v1 := range b.Ca {
			v, has := a.Ca[k]
			switch {
			case !has:
				return fmt.Sprintf(
					"Ca and required key %d",
					k,
				)
			case v != v1:
				return fmt.Sprintf(
					"Ca and values at key %d: %d != %d",
					k,
					v,
					v1,
				)
			}
		}

	case !reflect.DeepEqual(a.N, b.N):
		for k, v := range a.N {
			v1, has := b.N[k]
			switch {
			case !has:
				return fmt.Sprintf(
					"N and unnessesary key %d",
					k,
				)
			case !reflect.DeepEqual(v, v1):
				return fmt.Sprintf(
					"N and values at key %d: %v != %v",
					k,
					v,
					v1,
				)
			}
		}

		for k, v1 := range b.N {
			v, has := a.N[k]
			switch {
			case !has:
				return fmt.Sprintf(
					"N and required key %d",
					k,
				)
			case !reflect.DeepEqual(v, v1):
				return fmt.Sprintf(
					"N and values at key %d: %v != %v",
					k,
					v,
					v1,
				)
			}
		}
	}

	return ""
}

const doTestError = "Do Test %d: got %v for result (should be %v, look at %v)"

func TestDo(t *testing.T) {
	for i, tt := range doTests {
		result := Do(tt.tree)
		if !reflect.DeepEqual(result, tt.result) {
			g := guilty(result, tt.result)
			t.Errorf(doTestError, i, result, tt.result, g)
		}
	}
}
