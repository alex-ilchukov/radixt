package analysis

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/null"
)

var (
	empty = generic.New()

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

	methods = generic.New(
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
			Dcfpm: 0,
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
			Dcfpm: 0,
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
			Dcfpm: 0,
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
			Dcfpm: 5,
			Vm:    7,
			N: map[uint]N{
				0: {
					Index:         0,
					Chunk:         "",
					Value:         0,
					HasValue:      false,
					Root:          true,
					Parent:        0,
					ChildrenFirst: 1,
					ChildrenLast:  2,
					ChunkPos:      0,
				},
				1: {
					Index:         1,
					Chunk:         "auth",
					Value:         4,
					HasValue:      true,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 3,
					ChildrenLast:  4,
					ChunkPos:      41,
				},
				2: {
					Index:         2,
					Chunk:         "content-",
					Value:         0,
					HasValue:      false,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 5,
					ChildrenLast:  7,
					ChunkPos:      21,
				},
				3: {
					Index:         3,
					Chunk:         "or",
					Value:         2,
					HasValue:      true,
					Root:          false,
					Parent:        1,
					ChildrenFirst: 8,
					ChildrenLast:  8,
					ChunkPos:      49,
				},
				4: {
					Index:         4,
					Chunk:         "entication",
					Value:         3,
					HasValue:      true,
					Root:          false,
					Parent:        1,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      11,
				},
				5: {
					Index:         5,
					Chunk:         "type",
					Value:         5,
					HasValue:      true,
					Root:          false,
					Parent:        2,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      45,
				},
				6: {
					Index:         6,
					Chunk:         "length",
					Value:         6,
					HasValue:      true,
					Root:          false,
					Parent:        2,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      29,
				},
				7: {
					Index:         7,
					Chunk:         "disposition",
					Value:         7,
					HasValue:      true,
					Root:          false,
					Parent:        2,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      0,
				},
				8: {
					Index:         8,
					Chunk:         "i",
					Value:         0,
					HasValue:      false,
					Root:          false,
					Parent:        3,
					ChildrenFirst: 9,
					ChildrenLast:  10,
					ChunkPos:      1,
				},
				9: {
					Index:         9,
					Chunk:         "ty",
					Value:         0,
					HasValue:      true,
					Root:          false,
					Parent:        8,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      45,
				},
				10: {
					Index:         10,
					Chunk:         "zation",
					Value:         1,
					HasValue:      true,
					Root:          false,
					Parent:        8,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      35,
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
			Dcfpm: 2,
			Vm:    7,
			N: map[uint]N{
				0: {
					Index:         0,
					Chunk:         "",
					Value:         0,
					HasValue:      false,
					Root:          true,
					Parent:        0,
					ChildrenFirst: 1,
					ChildrenLast:  6,
					ChunkPos:      0,
				},
				1: {
					Index:         1,
					Chunk:         "DELETE",
					Value:         0,
					HasValue:      true,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      7,
				},
				2: {
					Index:         2,
					Chunk:         "GET",
					Value:         1,
					HasValue:      true,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      26,
				},
				3: {
					Index:         3,
					Chunk:         "HEAD",
					Value:         2,
					HasValue:      true,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      22,
				},
				4: {
					Index:         4,
					Chunk:         "OPTIONS",
					Value:         3,
					HasValue:      true,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      0,
				},
				5: {
					Index:         5,
					Chunk:         "P",
					Value:         0,
					HasValue:      false,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 7,
					ChildrenLast:  9,
					ChunkPos:      1,
				},
				6: {
					Index:         6,
					Chunk:         "TRACE",
					Value:         7,
					HasValue:      true,
					Root:          false,
					Parent:        0,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      13,
				},
				7: {
					Index:         7,
					Chunk:         "ATCH",
					Value:         4,
					HasValue:      true,
					Root:          false,
					Parent:        5,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      18,
				},
				8: {
					Index:         8,
					Chunk:         "OST",
					Value:         5,
					HasValue:      true,
					Root:          false,
					Parent:        5,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      29,
				},
				9: {
					Index:         9,
					Chunk:         "UT",
					Value:         6,
					HasValue:      true,
					Root:          false,
					Parent:        5,
					ChildrenFirst: 1,
					ChildrenLast:  0,
					ChunkPos:      32,
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

	case a.Dcfpm != b.Dcfpm:
		return "Dcfpm"

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
