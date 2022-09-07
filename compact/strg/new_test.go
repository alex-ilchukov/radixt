package strg_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/strg"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/null"
)

const (
	border3 = 0x1_FF - 1
	large3  = 0x1_FF
	border4 = 0x1_FF_FF - 1
	large4  = 0x1_FF_FF
)

var (
	emptyOriginal = generic.New()

	regularValues = generic.New(
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

	borderValues3 = generic.NewFromSV(
		generic.SV{S: "GET", V: border3},
		generic.SV{S: "POST", V: border3},
		generic.SV{S: "PATCH", V: border3},
		generic.SV{S: "DELETE", V: border3},
		generic.SV{S: "PUT", V: border3},
		generic.SV{S: "OPTIONS", V: border3},
		generic.SV{S: "CONNECT", V: border3},
		generic.SV{S: "HEAD", V: border3},
		generic.SV{S: "TRACE", V: border3},
	)

	largeValues3 = generic.NewFromSV(
		generic.SV{S: "GET", V: large3},
		generic.SV{S: "POST", V: large3},
		generic.SV{S: "PATCH", V: large3},
		generic.SV{S: "DELETE", V: large3},
		generic.SV{S: "PUT", V: large3},
		generic.SV{S: "OPTIONS", V: large3},
		generic.SV{S: "CONNECT", V: large3},
		generic.SV{S: "HEAD", V: large3},
		generic.SV{S: "TRACE", V: large3},
	)

	borderValues4 = generic.NewFromSV(
		generic.SV{S: "GET", V: border4},
		generic.SV{S: "POST", V: border4},
		generic.SV{S: "PATCH", V: border4},
		generic.SV{S: "DELETE", V: border4},
		generic.SV{S: "PUT", V: border4},
		generic.SV{S: "OPTIONS", V: border4},
		generic.SV{S: "CONNECT", V: border4},
		generic.SV{S: "HEAD", V: border4},
		generic.SV{S: "TRACE", V: border4},
	)

	largeValues4 = generic.NewFromSV(
		generic.SV{S: "GET", V: large4},
		generic.SV{S: "POST", V: large4},
		generic.SV{S: "PATCH", V: large4},
		generic.SV{S: "DELETE", V: large4},
		generic.SV{S: "PUT", V: large4},
		generic.SV{S: "OPTIONS", V: large4},
		generic.SV{S: "CONNECT", V: large4},
		generic.SV{S: "HEAD", V: large4},
		generic.SV{S: "TRACE", V: large4},
	)
)

var newError3Tests = []struct {
	tree    radixt.Tree
	result2 error
}{
	{tree: nil, result2: nil},
	{tree: emptyOriginal, result2: nil},
	{tree: null.Tree, result2: nil},
	{tree: regularValues, result2: nil},
	{tree: borderValues3, result2: nil},
	{tree: largeValues3, result2: compact.ErrorOverflow},
}

const testNewError3Error = "Test New[N3] Error %d: got \"%s\" error " +
	"(should be \"%s\")"

func TestNewError3(t *testing.T) {
	for i, tt := range newError3Tests {
		_, result2 := strg.New[strg.N3](tt.tree)
		if result2 != tt.result2 {
			t.Errorf(testNewError3Error, i, result2, tt.result2)
		}
	}
}

var new3Tests = []struct {
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

const testNew3Error = "Test New[N3] %d: got that New(%v) is\n\n%v\n\n" +
	"which is not equal to\n\n%v\n\n(but should be equal)"

func TestNew3(t *testing.T) {
	for i, tt := range new3Tests {
		result, _ := strg.New[strg.N3](tt.tree)
		if !tt.e.Eq(result) {
			t.Errorf(testNew3Error, i, tt.tree, result, tt.e)
		}
	}
}

var newError4Tests = []struct {
	tree    radixt.Tree
	result2 error
}{
	{tree: nil, result2: nil},
	{tree: emptyOriginal, result2: nil},
	{tree: null.Tree, result2: nil},
	{tree: regularValues, result2: nil},
	{tree: borderValues4, result2: nil},
	{tree: largeValues4, result2: compact.ErrorOverflow},
}

const testNewError4Error = "Test New[N4] Error %d: got \"%s\" error " +
	"(should be \"%s\")"

func TestNewError4(t *testing.T) {
	for i, tt := range newError4Tests {
		_, result2 := strg.New[strg.N4](tt.tree)
		if result2 != tt.result2 {
			t.Errorf(testNewError4Error, i, result2, tt.result2)
		}
	}
}

var new4Tests = []struct {
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

const testNew4Error = "Test New[N4] %d: got that New(%v) is\n\n%v\n\n" +
	"which is not equal to\n\n%v\n\n(but should be equal)"

func TestNew4(t *testing.T) {
	for i, tt := range new4Tests {
		result, _ := strg.New[strg.N4](tt.tree)
		if !tt.e.Eq(result) {
			t.Errorf(testNew4Error, i, tt.tree, result, tt.e)
		}
	}
}
