package struct8_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/null"
	"github.com/alex-ilchukov/radixt/struct8"
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

	borderValues = generic.NewFromSV(
		generic.SV{S: "GET", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "POST", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "PATCH", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "DELETE", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "PUT", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "OPTIONS", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "CONNECT", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "HEAD", V: 0xFFFFFFFFFFFC},
		generic.SV{S: "TRACE", V: 0xFFFFFFFFFFFC},
	)

	largeValues = generic.NewFromSV(
		generic.SV{S: "GET", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "POST", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "PATCH", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "DELETE", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "PUT", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "OPTIONS", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "CONNECT", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "HEAD", V: 0xFFFFFFFFFFFF},
		generic.SV{S: "TRACE", V: 0xFFFFFFFFFFFF},
	)
)

var newErrorTests = []struct {
	tree    radixt.Tree
	result2 error
}{
	{tree: nil, result2: nil},
	{tree: emptyOriginal, result2: nil},
	{tree: null.Tree, result2: nil},
	{tree: regularValues, result2: nil},
	{tree: borderValues, result2: nil},
	{tree: largeValues, result2: struct8.ErrorOverflow},
}

const testNewErrorError = "Test New Error %d: got \"%s\" error " +
	"(should be \"%s\")"

func TestNewError(t *testing.T) {
	for i, tt := range newErrorTests {
		_, result2 := struct8.New(tt.tree)
		if result2 != tt.result2 {
			t.Errorf(testNewErrorError, i, result2, tt.result2)
		}
	}
}

var newTests = []struct {
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
					"OST|1": nil,
					"ATCH|2": nil,
					"UT|4": nil,
				},
				"DELETE|3": nil,
				"OPTIONS|5": nil,
				"CONNECT|6": nil,
				"HEAD|7": nil,
				"TRACE|8": nil,
			},
		},
	},
}

const testNewError = "Test New %d: got that New(%v) is\n\n%v\n\nwhich is " +
	"not equal to\n\n%v\n\n(but should be equal)"

func TestNew(t *testing.T) {
	for i, tt := range newTests {
		result, _ := struct8.New(tt.tree)
		if !tt.e.Eq(result) {
			t.Errorf(testNewError, i, tt.tree, result, tt.e)
		}
	}
}
