package struct4

import (
	"math"
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
		generic.SV{"GET", math.MaxUint16 - 1},
		generic.SV{"POST", math.MaxUint16 - 1},
		generic.SV{"PATCH", math.MaxUint16 - 1},
		generic.SV{"DELETE", math.MaxUint16 - 1},
		generic.SV{"PUT", math.MaxUint16 - 1},
		generic.SV{"OPTIONS", math.MaxUint16 - 1},
		generic.SV{"CONNECT", math.MaxUint16 - 1},
		generic.SV{"HEAD", math.MaxUint16 - 1},
		generic.SV{"TRACE", math.MaxUint16 - 1},
	)

	largeValues = generic.NewFromSV(
		generic.SV{"GET", math.MaxUint16},
		generic.SV{"POST", math.MaxUint16},
		generic.SV{"PATCH", math.MaxUint16},
		generic.SV{"DELETE", math.MaxUint16},
		generic.SV{"PUT", math.MaxUint16},
		generic.SV{"OPTIONS", math.MaxUint16},
		generic.SV{"CONNECT", math.MaxUint16},
		generic.SV{"HEAD", math.MaxUint16},
		generic.SV{"TRACE", math.MaxUint16},
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
