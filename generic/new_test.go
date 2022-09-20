package generic_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/sapling"
)

var newTests = []struct {
	t radixt.Tree
	e evident.Tree
}{
	{t: nil, e: nil},
	{t: sapling.New(), e: nil},
	{
		t: sapling.New(
			"authority",
			"authorization",
			"author",
			"authentication",
			"auth",
			"content-type",
			"content-length",
			"content-disposition",
		),
		e: evident.Tree{
			"|": {
				"auth|4": {
					"entication|3": nil,
					"or|2": {
						"i|": {
							"ty|0":     nil,
							"zation|1": nil,
						},
					},
				},
				"content-|": {
					"disposition|7": nil,
					"length|6":      nil,
					"type|5":        nil,
				},
			},
		},
	},
}

const testNewError = "New Test %d: got that New(%v...) is\n\n%v\n\nwhich is " +
	"not equal to \n\n%v\n\n (but should be equal)"

func TestNew(t *testing.T) {
	for i, tt := range newTests {
		result := generic.New(tt.t)
		if !tt.e.Eq(result) {
			t.Errorf(testNewError, i, tt.t, result, tt.e)
		}
	}
}
