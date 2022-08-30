package struct8

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/null"
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
	{tree: largeValues, result2: ErrorOverflow},
}

const testNewErrorError = "Test New Error %d: got \"%s\" error " +
	"(should be \"%s\")"

func TestNewError(t *testing.T) {
	for i, tt := range newErrorTests {
		_, result2 := New(tt.tree)
		if result2 != tt.result2 {
			t.Errorf(testNewErrorError, i, result2, tt.result2)
		}
	}
}
