package null

import "github.com/alex-ilchukov/radixt"

type tree struct{}

// Size always returns zero.
func (tree) Size() uint {
	return 0
}

// Value always returns 0 and boolean false.
func (tree) Value(uint) (v uint, has bool) {
	return
}

// Chunk always returns empty string.
func (tree) Chunk(uint) string {
	return ""
}

// EachChild does nothing.
func (tree) EachChild(uint, func(uint) bool) {}

// Hoard always returns zero amount of bytes, taken by the implementation, with
// [radixt.HoardExactly] as interpretation hint.
func (tree) Hoard() (uint, uint) {
	return 0, radixt.HoardExactly
}

// Tree is the only accessible instance of the implementation.
var Tree tree

// To check, if the implementation is compatible with the interfaces.
var (
	_ radixt.Tree    = Tree
	_ radixt.Hoarder = Tree
)
