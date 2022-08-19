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
			P:   "",
			Pml: 0,
			N:   map[int]N{},
			Nt:  map[int]int{},
			Ca:  map[int]int{},
		},
	},
	{
		tree: null.Tree,
		result: A{
			P:   "",
			Pml: 0,
			N:   map[int]N{},
			Nt:  map[int]int{},
			Ca:  map[int]int{},
		},
	},
	{
		tree: empty,
		result: A{
			P:   "",
			Pml: 0,
			N:   map[int]N{},
			Nt:  map[int]int{},
			Ca:  map[int]int{},
		},
	},
	{
		tree: atree,
		result: A{
			P: "dispositionenticationcontent-lengthzationauth" +
				"typeor",
			Pml: 11,
			N: map[int]N{
				0: {
					Index:    0,
					Pref:     "",
					String:   "",
					Mark:     -1,
					Parent:   NoParent,
					Children: []int{6, 7},
					PrefPos:  0,
				},
				1: {
					Index:    1,
					Pref:     "ty",
					String:   "authority",
					Mark:     0,
					Parent:   3,
					Children: []int{},
					PrefPos:  45,
				},
				2: {
					Index:    2,
					Pref:     "zation",
					String:   "authorization",
					Mark:     1,
					Parent:   3,
					Children: []int{},
					PrefPos:  35,
				},
				3: {
					Index:    3,
					Pref:     "i",
					String:   "",
					Mark:     -1,
					Parent:   4,
					Children: []int{1, 2},
					PrefPos:  1,
				},
				4: {
					Index:    4,
					Pref:     "or",
					String:   "author",
					Mark:     2,
					Parent:   6,
					Children: []int{3},
					PrefPos:  49,
				},
				5: {
					Index:    5,
					Pref:     "entication",
					String:   "authentication",
					Mark:     3,
					Parent:   6,
					Children: []int{},
					PrefPos:  11,
				},
				6: {
					Index:    6,
					Pref:     "auth",
					String:   "auth",
					Mark:     4,
					Parent:   0,
					Children: []int{4, 5},
					PrefPos:  41,
				},
				7: {
					Index:    7,
					Pref:     "content-",
					String:   "",
					Mark:     -1,
					Parent:   0,
					Children: []int{8, 9, 10},
					PrefPos:  21,
				},
				8: {
					Index:    8,
					Pref:     "type",
					String:   "content-type",
					Mark:     5,
					Parent:   7,
					Children: []int{},
					PrefPos:  45,
				},
				9: {
					Index:    9,
					Pref:     "length",
					String:   "content-length",
					Mark:     6,
					Parent:   7,
					Children: []int{},
					PrefPos:  29,
				},
				10: {
					Index:    10,
					Pref:     "disposition",
					String:   "content-disposition",
					Mark:     7,
					Parent:   7,
					Children: []int{},
					PrefPos:  0,
				},
			},
			Nt: map[int]int{
				0: 0,
				1: 9,
				2: 10,
				3: 8,
				4: 3,
				5: 4,
				6: 1,
				7: 2,
				8: 5,
				9: 6,
				10: 7,
			},
			Ca: map[int]int{0: 6, 1: 1, 2: 3, 3: 1},
		},
	},
	{
		tree: methods,
		result: A{
			P:   "OPTIONSDELETETRACEATCHHEADGETOSTUT",
			Pml: 7,
			N: map[int]N{
				0: {
					Index:    0,
					Pref:     "",
					String:   "",
					Mark:     -1,
					Parent:   NoParent,
					Children: []int{1, 2, 3, 4, 5, 9},
					PrefPos:  0,
				},
				1: {
					Index:    1,
					Pref:     "DELETE",
					String:   "DELETE",
					Mark:     0,
					Parent:   0,
					Children: []int{},
					PrefPos:  7,
				},
				2: {
					Index:    2,
					Pref:     "GET",
					String:   "GET",
					Mark:     1,
					Parent:   0,
					Children: []int{},
					PrefPos:  26,
				},
				3: {
					Index:    3,
					Pref:     "HEAD",
					String:   "HEAD",
					Mark:     2,
					Parent:   0,
					Children: []int{},
					PrefPos:  22,
				},
				4: {
					Index:    4,
					Pref:     "OPTIONS",
					String:   "OPTIONS",
					Mark:     3,
					Parent:   0,
					Children: []int{},
					PrefPos:  0,
				},
				5: {
					Index:    5,
					Pref:     "P",
					String:   "",
					Mark:     -1,
					Parent:   0,
					Children: []int{6, 7, 8},
					PrefPos:  1,
				},
				6: {
					Index:    6,
					Pref:     "ATCH",
					String:   "PATCH",
					Mark:     4,
					Parent:   5,
					Children: []int{},
					PrefPos:  18,
				},
				7: {
					Index:    7,
					Pref:     "OST",
					String:   "POST",
					Mark:     5,
					Parent:   5,
					Children: []int{},
					PrefPos:  29,
				},
				8: {
					Index:    8,
					Pref:     "UT",
					String:   "PUT",
					Mark:     6,
					Parent:   5,
					Children: []int{},
					PrefPos:  32,
				},
				9: {
					Index:    9,
					Pref:     "TRACE",
					String:   "TRACE",
					Mark:     7,
					Parent:   0,
					Children: []int{},
					PrefPos:  13,
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
	case a.P != b.P:
		return "P"

	case a.Pml != b.Pml:
		return "Pml"

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
