package structg_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/structg"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/null"
	"github.com/alex-ilchukov/radixt/sapling"
)

const (
	border32 = 0x7_FF_FF - 1
	large32  = 0x7_FF_FF
	border64 = 0x7_FF_FF_FF_FF_FF_FF - 1
	large64  = 0x7_FF_FF_FF_FF_FF_FF
)

var (
	emptyOriginal = sapling.New()

	regularValues = sapling.New(
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
		"OPTIONS",
		"CONNECT",
		"HEAD",
		"TRACE",
	)

	borderValues32 = sapling.NewFromSV(
		sapling.SV{S: "GET", V: border32},
		sapling.SV{S: "POST", V: border32},
		sapling.SV{S: "PATCH", V: border32},
		sapling.SV{S: "DELETE", V: border32},
		sapling.SV{S: "PUT", V: border32},
		sapling.SV{S: "OPTIONS", V: border32},
		sapling.SV{S: "CONNECT", V: border32},
		sapling.SV{S: "HEAD", V: border32},
		sapling.SV{S: "TRACE", V: border32},
	)

	largeValues32 = sapling.NewFromSV(
		sapling.SV{S: "GET", V: large32},
		sapling.SV{S: "POST", V: large32},
		sapling.SV{S: "PATCH", V: large32},
		sapling.SV{S: "DELETE", V: large32},
		sapling.SV{S: "PUT", V: large32},
		sapling.SV{S: "OPTIONS", V: large32},
		sapling.SV{S: "CONNECT", V: large32},
		sapling.SV{S: "HEAD", V: large32},
		sapling.SV{S: "TRACE", V: large32},
	)

	borderValues64 = sapling.NewFromSV(
		sapling.SV{S: "GET", V: border64},
		sapling.SV{S: "POST", V: border64},
		sapling.SV{S: "PATCH", V: border64},
		sapling.SV{S: "DELETE", V: border64},
		sapling.SV{S: "PUT", V: border64},
		sapling.SV{S: "OPTIONS", V: border64},
		sapling.SV{S: "CONNECT", V: border64},
		sapling.SV{S: "HEAD", V: border64},
		sapling.SV{S: "TRACE", V: border64},
	)

	largeValues64 = sapling.NewFromSV(
		sapling.SV{S: "GET", V: large64},
		sapling.SV{S: "POST", V: large64},
		sapling.SV{S: "PATCH", V: large64},
		sapling.SV{S: "DELETE", V: large64},
		sapling.SV{S: "PUT", V: large64},
		sapling.SV{S: "OPTIONS", V: large64},
		sapling.SV{S: "CONNECT", V: large64},
		sapling.SV{S: "HEAD", V: large64},
		sapling.SV{S: "TRACE", V: large64},
	)
)

var newError32Tests = []struct {
	tree    radixt.Tree
	result2 error
}{
	{tree: nil, result2: nil},
	{tree: emptyOriginal, result2: nil},
	{tree: null.Tree, result2: nil},
	{tree: regularValues, result2: nil},
	{tree: borderValues32, result2: nil},
	{tree: largeValues32, result2: compact.ErrorOverflow},
}

const testNewError32Error = "Test New[uint32] Error %d: got \"%s\" error " +
	"(should be \"%s\")"

func TestNewError32(t *testing.T) {
	for i, tt := range newError32Tests {
		_, result2 := structg.New[uint32](tt.tree)
		if result2 != tt.result2 {
			t.Errorf(testNewError32Error, i, result2, tt.result2)
		}
	}
}

var new32Tests = []struct {
	tree radixt.Tree
	e    evident.Tree
}{
	{tree: nil, e: nil},
	{tree: emptyOriginal, e: nil},
	{tree: null.Tree, e: nil},
	{
		tree: regularValues,
		e: evident.Tree{
			"|": {
				"GET|0": nil,
				"P|": {
					"OST|1":  nil,
					"ATCH|2": nil,
					"UT|4":   nil,
				},
				"DELETE|3":  nil,
				"OPTIONS|5": nil,
				"CONNECT|6": nil,
				"HEAD|7":    nil,
				"TRACE|8":   nil,
			},
		},
	},
}

const testNew32Error = "Test New[uint32] %d: got that New(%v) is\n\n%v\n\n" +
	"which is not equal to\n\n%v\n\n(but should be equal)"

func TestNew32(t *testing.T) {
	for i, tt := range new32Tests {
		result, _ := structg.New[uint32](tt.tree)
		if !tt.e.Eq(result) {
			t.Errorf(testNew32Error, i, tt.tree, result, tt.e)
		}
	}
}

var newError64Tests = []struct {
	tree    radixt.Tree
	result2 error
}{
	{tree: nil, result2: nil},
	{tree: emptyOriginal, result2: nil},
	{tree: null.Tree, result2: nil},
	{tree: regularValues, result2: nil},
	{tree: borderValues64, result2: nil},
	{tree: largeValues64, result2: compact.ErrorOverflow},
}

const testNewError64Error = "Test New[uint64] Error %d: got \"%s\" error " +
	"(should be \"%s\")"

func TestNewError64(t *testing.T) {
	for i, tt := range newError64Tests {
		_, result2 := structg.New[uint64](tt.tree)
		if result2 != tt.result2 {
			t.Errorf(testNewError64Error, i, result2, tt.result2)
		}
	}
}

var new64Tests = []struct {
	tree radixt.Tree
	e    evident.Tree
}{
	{tree: nil, e: nil},
	{tree: emptyOriginal, e: nil},
	{tree: null.Tree, e: nil},
	{
		tree: regularValues,
		e: evident.Tree{
			"|": {
				"GET|0": nil,
				"P|": {
					"OST|1":  nil,
					"ATCH|2": nil,
					"UT|4":   nil,
				},
				"DELETE|3":  nil,
				"OPTIONS|5": nil,
				"CONNECT|6": nil,
				"HEAD|7":    nil,
				"TRACE|8":   nil,
			},
		},
	},
}

const testNew64Error = "Test New[uint64] %d: got that New(%v) is\n\n%v\n\n" +
	"which is not equal to\n\n%v\n\n(but should be equal)"

func TestNew64(t *testing.T) {
	for i, tt := range new64Tests {
		result, _ := structg.New[uint64](tt.tree)
		if !tt.e.Eq(result) {
			t.Errorf(testNew64Error, i, tt.tree, result, tt.e)
		}
	}
}
