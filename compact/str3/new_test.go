package str3_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact"
	"github.com/alex-ilchukov/radixt/compact/str3"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/null"
	"github.com/alex-ilchukov/radixt/sapling"
)

const (
	border = 0xF_FF - 1
	large  = 0xF_FF
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

	borderValues = sapling.NewFromSV(
		sapling.SV{S: "GET", V: border},
		sapling.SV{S: "POST", V: border},
		sapling.SV{S: "PATCH", V: border},
		sapling.SV{S: "DELETE", V: border},
		sapling.SV{S: "PUT", V: border},
		sapling.SV{S: "OPTIONS", V: border},
		sapling.SV{S: "CONNECT", V: border},
		sapling.SV{S: "HEAD", V: border},
		sapling.SV{S: "TRACE", V: border},
	)

	largeValues = sapling.NewFromSV(
		sapling.SV{S: "GET", V: large},
		sapling.SV{S: "POST", V: large},
		sapling.SV{S: "PATCH", V: large},
		sapling.SV{S: "DELETE", V: large},
		sapling.SV{S: "PUT", V: large},
		sapling.SV{S: "OPTIONS", V: large},
		sapling.SV{S: "CONNECT", V: large},
		sapling.SV{S: "HEAD", V: large},
		sapling.SV{S: "TRACE", V: large},
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
	{tree: largeValues, result2: compact.ErrorOverflow},
}

const testNewErrorError = "Test New Error %d: got \"%s\" error " +
	"(should be \"%s\")"

func TestNewError(t *testing.T) {
	for i, tt := range newErrorTests {
		_, result2 := str3.New(tt.tree)
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

const testNewError = "Test New %d: got that New(%v) is\n\n%v\n\n" +
	"which is not equal to\n\n%v\n\n(but should be equal)"

func TestNew(t *testing.T) {
	for i, tt := range newTests {
		result, _ := str3.New(tt.tree)
		if !tt.e.Eq(result) {
			t.Errorf(testNewError, i, tt.tree, result, tt.e)
		}
	}
}
