package lookup_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/lookup"
	"github.com/alex-ilchukov/radixt/null"
)

var (
	empty = generic.New()

	atree = generic.New(
		"authorization",
		"content-type",
		"content-length",
		"content-disposition",
	)

	withBlank = generic.New(
		"authorization",
		"content-type",
		"content-length",
		"content-disposition",
		"",
	)
)

var lFeedTests = []struct {
	tree   radixt.Tree
	input  string
	result bool
}{
	{tree: nil, input: "a", result: false},
	{tree: nil, input: "b", result: false},
	{tree: nil, input: "c", result: false},
	{tree: nil, input: "content-type", result: false},
	{tree: empty, input: "a", result: false},
	{tree: empty, input: "b", result: false},
	{tree: empty, input: "c", result: false},
	{tree: empty, input: "content-type", result: false},
	{tree: atree, input: "a", result: true},
	{tree: atree, input: "b", result: false},
	{tree: atree, input: "c", result: true},
	{tree: atree, input: "authorization", result: true},
	{tree: atree, input: "content-type", result: true},
	{tree: atree, input: "content-length", result: true},
	{tree: atree, input: "content-disposition", result: true},
	{tree: atree, input: "content-typ", result: true},
	{tree: atree, input: "content-w", result: false},
	{tree: atree, input: "content-width", result: false},
	{tree: atree, input: "content-", result: true},
	{tree: atree, input: "auth", result: true},
	{tree: atree, input: "authe", result: false},
	{tree: withBlank, input: "a", result: true},
	{tree: withBlank, input: "b", result: false},
	{tree: withBlank, input: "c", result: true},
	{tree: withBlank, input: "authorization", result: true},
	{tree: withBlank, input: "content-type", result: true},
	{tree: withBlank, input: "content-length", result: true},
	{tree: withBlank, input: "content-disposition", result: true},
	{tree: withBlank, input: "content-typ", result: true},
	{tree: withBlank, input: "content-w", result: false},
	{tree: withBlank, input: "content-width", result: false},
	{tree: withBlank, input: "content-", result: true},
	{tree: withBlank, input: "auth", result: true},
	{tree: withBlank, input: "authe", result: false},
}

const testLFeedError = "Test L Feed %d: after input data %s got %t for " +
	"input byte %d (should be %t)"

func TestLFeed(t *testing.T) {
	for i, tt := range lFeedTests {
		tree := tt.tree
		last := len(tt.input) - 1
		input := tt.input[:last]
		b := tt.input[last]
		l := lookup.New(tree)

		for j := 0; j < len(input); j++ {
			l.Feed(input[j])
		}

		result := l.Feed(b)
		if result != tt.result {
			t.Errorf(
				testLFeedError,
				i,
				input,
				result,
				b,
				tt.result,
			)
		}
	}
}

var lsFeedTests = []struct {
	tree   lookup.Tritcher
	input  string
	result bool
}{
	{tree: nil, input: "a", result: false},
	{tree: nil, input: "b", result: false},
	{tree: nil, input: "c", result: false},
	{tree: nil, input: "content-type", result: false},
	{tree: empty, input: "a", result: false},
	{tree: empty, input: "b", result: false},
	{tree: empty, input: "c", result: false},
	{tree: empty, input: "content-type", result: false},
	{tree: atree, input: "a", result: true},
	{tree: atree, input: "b", result: false},
	{tree: atree, input: "c", result: true},
	{tree: atree, input: "authorization", result: true},
	{tree: atree, input: "content-type", result: true},
	{tree: atree, input: "content-length", result: true},
	{tree: atree, input: "content-disposition", result: true},
	{tree: atree, input: "content-typ", result: true},
	{tree: atree, input: "content-w", result: false},
	{tree: atree, input: "content-width", result: false},
	{tree: atree, input: "content-", result: true},
	{tree: atree, input: "auth", result: true},
	{tree: atree, input: "authe", result: false},
	{tree: withBlank, input: "a", result: true},
	{tree: withBlank, input: "b", result: false},
	{tree: withBlank, input: "c", result: true},
	{tree: withBlank, input: "authorization", result: true},
	{tree: withBlank, input: "content-type", result: true},
	{tree: withBlank, input: "content-length", result: true},
	{tree: withBlank, input: "content-disposition", result: true},
	{tree: withBlank, input: "content-typ", result: true},
	{tree: withBlank, input: "content-w", result: false},
	{tree: withBlank, input: "content-width", result: false},
	{tree: withBlank, input: "content-", result: true},
	{tree: withBlank, input: "auth", result: true},
	{tree: withBlank, input: "authe", result: false},
}

const testLSFeedError = "Test LS Feed %d: after input data %s got %t for " +
	"input byte %d (should be %t)"

func TestLSFeed(t *testing.T) {
	for i, tt := range lsFeedTests {
		tree := tt.tree
		last := len(tt.input) - 1
		input := tt.input[:last]
		b := tt.input[last]
		l := lookup.NewInTritcher(tree)

		for j := 0; j < len(input); j++ {
			l.Feed(input[j])
		}

		result := l.Feed(b)
		if result != tt.result {
			t.Errorf(
				testLSFeedError,
				i,
				input,
				result,
				b,
				tt.result,
			)
		}
	}
}

var lFoundTests = []struct {
	tree   radixt.Tree
	input  string
	result bool
}{
	{tree: nil, input: "", result: false},
	{tree: nil, input: "content-type", result: false},
	{tree: empty, input: "", result: false},
	{tree: empty, input: "content-type", result: false},
	{tree: atree, input: "authorization", result: true},
	{tree: atree, input: "content-type", result: true},
	{tree: atree, input: "content-length", result: true},
	{tree: atree, input: "content-disposition", result: true},
	{tree: atree, input: "content-typ", result: false},
	{tree: atree, input: "content-", result: false},
	{tree: atree, input: "auth", result: false},
	{tree: atree, input: "", result: false},
	{tree: withBlank, input: "authorization", result: true},
	{tree: withBlank, input: "content-type", result: true},
	{tree: withBlank, input: "content-length", result: true},
	{tree: withBlank, input: "content-disposition", result: true},
	{tree: withBlank, input: "content-typ", result: false},
	{tree: withBlank, input: "content-", result: false},
	{tree: withBlank, input: "auth", result: false},
	{tree: withBlank, input: "", result: true},
}

const testLFoundError = "Test L Found %d: for input data %s got %t (should " +
	"be %t)"

func TestLFound(t *testing.T) {
	for i, tt := range lFoundTests {
		tree := tt.tree
		input := tt.input
		l := lookup.New(tree)

		for j := 0; j < len(input); j++ {
			l.Feed(input[j])
		}

		result := l.Found()
		if result != tt.result {
			t.Errorf(
				testLFoundError,
				i,
				tt.input,
				result,
				tt.result,
			)
		}
	}
}

var lTreeTests = []struct {
	tree   radixt.Tree
	result radixt.Tree
}{
	{tree: nil, result: null.Tree},
	{tree: empty, result: empty},
	{tree: atree, result: atree},
	{tree: withBlank, result: withBlank},
}

const testLTreeError = "Test L Tree %d: got %v (should be %v)"

func TestLTree(t *testing.T) {
	for i, tt := range lTreeTests {
		tree := tt.tree
		l := lookup.New(tree)
		result := l.Tree()
		if result != tt.result {
			t.Errorf(testLTreeError, i, result, tt.result)
		}
	}
}

var lNodeTests = []struct {
	tree   radixt.Tree
	input  string
	result uint
}{
	{tree: nil, input: "", result: 0},
	{tree: nil, input: "content-type", result: 0},
	{tree: empty, input: "", result: 0},
	{tree: empty, input: "content-type", result: 0},
	{tree: atree, input: "authorization", result: 1},
	{tree: atree, input: "content-type", result: 3},
	{tree: atree, input: "content-length", result: 4},
	{tree: atree, input: "content-disposition", result: 5},
	{tree: atree, input: "content-typ", result: 3},
	{tree: atree, input: "content-", result: 2},
	{tree: atree, input: "auth", result: 1},
	{tree: atree, input: "", result: 0},
	{tree: withBlank, input: "authorization", result: 1},
	{tree: withBlank, input: "content-type", result: 3},
	{tree: withBlank, input: "content-length", result: 4},
	{tree: withBlank, input: "content-disposition", result: 5},
	{tree: withBlank, input: "content-typ", result: 3},
	{tree: withBlank, input: "content-", result: 2},
	{tree: withBlank, input: "auth", result: 1},
	{tree: withBlank, input: "", result: 0},
}

const testLNodeError = "Test L Node %d: for input data %s got " +
	"%d (should be %d)"

func TestLNode(t *testing.T) {
	for i, tt := range lNodeTests {
		tree := tt.tree
		input := tt.input
		l := lookup.New(tree)

		for j := 0; j < len(input); j++ {
			l.Feed(input[j])
		}

		result := l.Node()
		if result != tt.result {
			t.Errorf(
				testLNodeError,
				i,
				tt.input,
				result,
				tt.result,
			)
		}
	}
}
