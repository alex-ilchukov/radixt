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
		generic.SV{"GET", 0xFFFFFFFFFFFC},
		generic.SV{"POST", 0xFFFFFFFFFFFC},
		generic.SV{"PATCH", 0xFFFFFFFFFFFC},
		generic.SV{"DELETE", 0xFFFFFFFFFFFC},
		generic.SV{"PUT", 0xFFFFFFFFFFFC},
		generic.SV{"OPTIONS", 0xFFFFFFFFFFFC},
		generic.SV{"CONNECT", 0xFFFFFFFFFFFC},
		generic.SV{"HEAD", 0xFFFFFFFFFFFC},
		generic.SV{"TRACE", 0xFFFFFFFFFFFC},
	)

	largeValues = generic.NewFromSV(
		generic.SV{"GET", 0xFFFFFFFFFFFF},
		generic.SV{"POST", 0xFFFFFFFFFFFF},
		generic.SV{"PATCH", 0xFFFFFFFFFFFF},
		generic.SV{"DELETE", 0xFFFFFFFFFFFF},
		generic.SV{"PUT", 0xFFFFFFFFFFFF},
		generic.SV{"OPTIONS", 0xFFFFFFFFFFFF},
		generic.SV{"CONNECT", 0xFFFFFFFFFFFF},
		generic.SV{"HEAD", 0xFFFFFFFFFFFF},
		generic.SV{"TRACE", 0xFFFFFFFFFFFF},
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
