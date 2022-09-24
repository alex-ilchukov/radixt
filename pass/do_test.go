package pass_test

import (
	"fmt"
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/pass"
	"github.com/alex-ilchukov/radixt/sapling"
)

var (
	empty = sapling.New()

	atree = sapling.New(
		"authority",
		"authorization",
		"author",
		"authentication",
		"auth",
		"content-type",
		"content-length",
		"content-disposition",
	)
)

type yielder struct {
	a string
}

func (y *yielder) Yield(i, n, tag uint) uint {
	if y.a != "" {
		y.a += ", "
	}
	y.a += fmt.Sprintf("(%d, %d, %d)", i, n, tag)

	return n
}

var doTests = []struct {
	t radixt.Tree
	y *yielder
	a string
}{
	{t: nil, y: nil, a: ""},
	{t: empty, y: nil, a: ""},
	{t: nil, y: &yielder{}, a: ""},
	{t: empty, y: &yielder{}, a: ""},
	{
		t: atree,
		y: &yielder{},
		a: "" +
			"(0, 0, 0), " +
			"(1, 6, 0), " +
			"(2, 7, 0), " +
			"(3, 5, 6), " +
			"(4, 4, 6), " +
			"(5, 10, 7), " +
			"(6, 9, 7), " +
			"(7, 8, 7), " +
			"(8, 3, 4), " +
			"(9, 1, 3), " +
			"(10, 2, 3)",
	},
}

const testDoError = "Do Test %d: got '%s' instead of '%s'"

func TestDo(t *testing.T) {
	for i, tt := range doTests {
		pass.Do(tt.t, tt.y)

		a := ""
		if tt.y != nil {
			a = tt.y.a
		}

		if a != tt.a {
			t.Errorf(testDoError, i, a, tt.a)
		}
	}
}
