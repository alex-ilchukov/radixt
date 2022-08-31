package evident_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
)

var (
	newEmpty = generic.New()

	newATree = generic.New(
		"authority",
		"authorization",
		"author",
		"authentication",
		"auth",
		"content-type",
		"content-length",
		"content-disposition",
	)

	newETree = evident.Tree{
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
	}
)

var newTests = []struct {
	t      radixt.Tree
	result evident.Tree
}{
	{t: nil, result: nil},
	{t: newEmpty, result: nil},
	{t: evident.Tree(nil), result: nil},
	{t: evident.Tree{}, result: evident.Tree{}},
	{t: newATree, result: newETree},
	{t: newETree, result: newETree},
}

const testNewError = "New Test %d: got that New(%v) is\n\n%v\n\nwhich is " +
	"not equal to \n\n%v\n\n (but should be equal)"

func TestNew(t *testing.T) {
	for i, tt := range newTests {
		result := evident.New(tt.t)
		if !result.Eq(tt.result) {
			t.Errorf(testNewError, i, tt.t, result, tt.result)
		}
	}
}
