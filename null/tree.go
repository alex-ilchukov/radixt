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

// EachChild does nothing.
func (tree) EachChild(uint, func(uint) bool) {
	return
}

// ByteAt always returns 0 and boolean false.
func (tree) ByteAt(uint, uint) (b byte, within bool) {
	return
}

// Tree is the only accessible instance of the implementation.
var Tree tree

// To check, if the implementation is compatible with the interface.
var _ radixt.Tree = Tree
