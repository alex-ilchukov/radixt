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

// ChildrenRange always returns 1 and 0.
func (tree) ChildrenRange(uint) (uint, uint) {
	return 1, 0
}

// Tree is the only accessible instance of the implementation.
var Tree tree

// To check, if the implementation is compatible with the interface.
var _ radixt.Tree = Tree
