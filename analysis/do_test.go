package analysis

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/null"
)

const NoParent = math.MinInt

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
			C:   "",
			Cml: 0,
			Vm:  0,
			N:   map[int]N{},
			Nt:  map[int]int{},
			Ca:  map[int]int{},
		},
	},
	{
		tree: null.Tree,
		result: A{
			C:   "",
			Cml: 0,
			Vm:  0,
			N:   map[int]N{},
			Nt:  map[int]int{},
			Ca:  map[int]int{},
		},
	},
	{
		tree: empty,
		result: A{
			C:   "",
			Cml: 0,
			Vm:  0,
			N:   map[int]N{},
			Nt:  map[int]int{},
			Ca:  map[int]int{},
		},
	},
	{
		tree: atree,
		result: A{
			C: "dispositionenticationcontent-lengthzationauth" +
				"typeor",
			Cml: 11,
			Vm:  7,
			N: map[int]N{
				0: {
					Index:    0,
					Chunk:    "",
					Value:    0,
					HasValue: false,
					Parent:   NoParent,
					Children: []int{6, 7},
					ChunkPos: 0,
				},
				1: {
					Index:    1,
					Chunk:    "ty",
					Value:    0,
					HasValue: true,
					Parent:   3,
					Children: []int{},
					ChunkPos: 45,
				},
				2: {
					Index:    2,
					Chunk:    "zation",
					Value:    1,
					HasValue: true,
					Parent:   3,
					Children: []int{},
					ChunkPos: 35,
				},
				3: {
					Index:    3,
					Chunk:    "i",
					Value:    0,
					HasValue: false,
					Parent:   4,
					Children: []int{1, 2},
					ChunkPos: 1,
				},
				4: {
					Index:    4,
					Chunk:    "or",
					Value:    2,
					HasValue: true,
					Parent:   6,
					Children: []int{3},
					ChunkPos: 49,
				},
				5: {
					Index:    5,
					Chunk:    "entication",
					Value:    3,
					HasValue: true,
					Parent:   6,
					Children: []int{},
					ChunkPos: 11,
				},
				6: {
					Index:    6,
					Chunk:    "auth",
					Value:    4,
					HasValue: true,
					Parent:   0,
					Children: []int{4, 5},
					ChunkPos: 41,
				},
				7: {
					Index:    7,
					Chunk:    "content-",
					Value:    0,
					HasValue: false,
					Parent:   0,
					Children: []int{8, 9, 10},
					ChunkPos: 21,
				},
				8: {
					Index:    8,
					Chunk:    "type",
					Value:    5,
					HasValue: true,
					Parent:   7,
					Children: []int{},
					ChunkPos: 45,
				},
				9: {
					Index:    9,
					Chunk:    "length",
					Value:    6,
					HasValue: true,
					Parent:   7,
					Children: []int{},
					ChunkPos: 29,
				},
				10: {
					Index:    10,
					Chunk:    "disposition",
					Value:    7,
					HasValue: true,
					Parent:   7,
					Children: []int{},
					ChunkPos: 0,
				},
			},
			Nt: map[int]int{
				0:  0,
				1:  9,
				2:  10,
				3:  8,
				4:  3,
				5:  4,
				6:  1,
				7:  2,
				8:  5,
				9:  6,
				10: 7,
			},
			Ca: map[int]int{0: 6, 1: 1, 2: 3, 3: 1},
		},
	},
	{
		tree: methods,
		result: A{
			C:   "OPTIONSDELETETRACEATCHHEADGETOSTUT",
			Cml: 7,
			Vm:  7,
			N: map[int]N{
				0: {
					Index:    0,
					Chunk:    "",
					Value:    0,
					HasValue: false,
					Parent:   NoParent,
					Children: []int{1, 2, 3, 4, 5, 9},
					ChunkPos: 0,
				},
				1: {
					Index:    1,
					Chunk:    "DELETE",
					Value:    0,
					HasValue: true,
					Parent:   0,
					Children: []int{},
					ChunkPos: 7,
				},
				2: {
					Index:    2,
					Chunk:    "GET",
					Value:    1,
					HasValue: true,
					Parent:   0,
					Children: []int{},
					ChunkPos: 26,
				},
				3: {
					Index:    3,
					Chunk:    "HEAD",
					Value:    2,
					HasValue: true,
					Parent:   0,
					Children: []int{},
					ChunkPos: 22,
				},
				4: {
					Index:    4,
					Chunk:    "OPTIONS",
					Value:    3,
					HasValue: true,
					Parent:   0,
					Children: []int{},
					ChunkPos: 0,
				},
				5: {
					Index:    5,
					Chunk:    "P",
					Value:    0,
					HasValue: false,
					Parent:   0,
					Children: []int{6, 7, 8},
					ChunkPos: 1,
				},
				6: {
					Index:    6,
					Chunk:    "ATCH",
					Value:    4,
					HasValue: true,
					Parent:   5,
					Children: []int{},
					ChunkPos: 18,
				},
				7: {
					Index:    7,
					Chunk:    "OST",
					Value:    5,
					HasValue: true,
					Parent:   5,
					Children: []int{},
					ChunkPos: 29,
				},
				8: {
					Index:    8,
					Chunk:    "UT",
					Value:    6,
					HasValue: true,
					Parent:   5,
					Children: []int{},
					ChunkPos: 32,
				},
				9: {
					Index:    9,
					Chunk:    "TRACE",
					Value:    7,
					HasValue: true,
					Parent:   0,
					Children: []int{},
					ChunkPos: 13,
				},
			},
			Nt: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				3: 3,
				4: 4,
				5: 5,
				6: 7,
				7: 8,
				8: 9,
				9: 6,
			},
			Ca: map[int]int{0: 8, 3: 1, 6: 1},
		},
	},
}

func guilty(a, b A) string {
	switch {
	case a.C != b.C:
		return "C"

	case a.Cml != b.Cml:
		return "Cml"

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

	case !reflect.DeepEqual(a.Nt, b.Nt):
		for k, v := range a.Nt {
			v1, has := b.Nt[k]
			switch {
			case !has:
				return fmt.Sprintf(
					"Nt and unnessesary key %d",
					k,
				)
			case v != v1:
				return fmt.Sprintf(
					"Nt and values at key %d: %d != %d",
					k,
					v,
					v1,
				)
			}
		}

		for k, v1 := range b.Nt {
			v, has := a.Nt[k]
			switch {
			case !has:
				return fmt.Sprintf(
					"Nt and required key %d",
					k,
				)
			case v != v1:
				return fmt.Sprintf(
					"Nt and values at key %d: %d != %d",
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
