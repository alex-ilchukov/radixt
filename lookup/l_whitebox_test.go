package lookup

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/null"
)

var (
	empty = evident.Tree{}

	atree = evident.Tree{
		"|": {
			"authorization|0": nil,
			"content-|": {
				"type|1": nil,
				"length|2": nil,
				"disposition|3": nil,
			},
		},
	}

	withBlank = evident.Tree{
		"|4": {
			"authorization|0": nil,
			"content-|": {
				"type|1": nil,
				"length|2": nil,
				"disposition|3": nil,
			},
		},
	}
)

var newTests = []struct {
	t      radixt.Tree
	lt     radixt.Tree
	ln     uint
	lchunk string
	lstop  bool
}{
	{t: nil, lt: null.Tree, ln: 0, lchunk: "", lstop: true},
	{t: null.Tree, lt: null.Tree, ln: 0, lchunk: "", lstop: true},
	{t: empty, lt: empty, ln: 0, lchunk: "", lstop: true},
	{t: atree, lt: atree, ln: 0, lchunk: "", lstop: false},
	{t: withBlank, lt: withBlank, ln: 0, lchunk: "", lstop: false},
}

const testNewError = "Test New %d: for tree %v got %v (should be %v)"

func TestNew(t *testing.T) {
	for i, tt := range newTests {
		l := New(tt.t)

		if !evident.New(l.t).Eq(tt.lt) ||
			l.n != tt.ln ||
			l.chunk != tt.lchunk ||
			l.stop != tt.lstop {
			t.Errorf(testNewError, i, tt.t, l, tt)
		}
	}
}

var lResetTests = []struct {
	tree   radixt.Tree
	input  string
	ln     uint
	lchunk string
	lstop  bool
}{
	{tree: nil, input: "", ln: 0, lchunk: "", lstop: true},
	{tree: nil, input: "content-type", ln: 0, lchunk: "", lstop: true},
	{tree: empty, input: "", ln: 0, lchunk: "", lstop: true},
	{tree: empty, input: "content-type", ln: 0, lchunk: "", lstop: true},
	{tree: atree, input: "authorization", ln: 0, lchunk: "", lstop: false},
	{tree: atree, input: "content-type", ln: 0, lchunk: "", lstop: false},
	{
		tree:   atree,
		input:  "content-length",
		ln:     0,
		lchunk: "",
		lstop:  false,
	},
	{
		tree:   atree,
		input:  "content-disposition",
		ln:     0,
		lchunk: "",
		lstop:  false,
	},
	{tree: atree, input: "content-typ", ln: 0, lchunk: "", lstop: false},
	{tree: atree, input: "content-", ln: 0, lchunk: "", lstop: false},
	{tree: atree, input: "auth", ln: 0, lchunk: "", lstop: false},
	{tree: atree, input: "", ln: 0, lchunk: "", lstop: false},

	{
		tree:   withBlank,
		input:  "authorization",
		ln:     0,
		lchunk: "",
		lstop:  false,
	},
	{
		tree:   withBlank,
		input:  "content-type",
		ln:     0,
		lchunk: "",
		lstop:  false,
	},
	{
		tree:   withBlank,
		input:  "content-length",
		ln:     0,
		lchunk: "",
		lstop:  false,
	},
	{
		tree:   withBlank,
		input:  "content-disposition",
		ln:     0,
		lchunk: "",
		lstop:  false,
	},
	{
		tree:   withBlank,
		input:  "content-typ",
		ln:     0,
		lchunk: "",
		lstop:  false,
	},
	{tree: withBlank, input: "content-", ln: 0, lchunk: "", lstop: false},
	{tree: withBlank, input: "auth", ln: 0, lchunk: "", lstop: false},
	{tree: withBlank, input: "", ln: 0, lchunk: "", lstop: false},
}

const testLResetError = "Test L Reset %d: got l.n = %d, l.npos = '%s', " +
	"l.stop = %t (should be %d, '%s' and %t)"

func TestLReset(t *testing.T) {
	for i, tt := range lResetTests {
		tree := tt.tree
		input := tt.input
		l := New(tree)

		for j := 0; j < len(input); j++ {
			l.Feed(input[j])
		}

		l.Reset()
		if l.n != tt.ln || l.chunk != tt.lchunk || l.stop != tt.lstop {
			t.Errorf(
				testLResetError,
				i,
				l.n,
				l.chunk,
				l.stop,
				tt.ln,
				tt.lchunk,
				tt.lstop,
			)
		}
	}
}
