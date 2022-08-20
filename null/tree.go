package null

import "github.com/alex-ilchukov/radixt"

type tree struct{}

// Size always returns zero.
func (tree) Size() int {
	return 0
}

// Has always returns boolean false.
func (tree) Has(int) bool {
	return false
}

// Root always returns -1.
func (tree) Root() int {
	return -1
}

// Mark always returns 0.
func (tree) Mark(int) uint {
	return 0
}

// EachChild does nothing.
func (tree) EachChild(int, func(int) bool) {
	return
}

// ByteAt always returns 0 and boolean false.
func (tree) ByteAt(int, int) (byte, bool) {
	return 0, false
}

// Tree is the only accessible instance of the implementation.
var Tree tree

// To check, if the implementation is compatible with the interface.
var _ radixt.Tree = Tree
