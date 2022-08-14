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

// NodeMark always returns -1.
func (tree) NodeMark(int) int {
	return -1
}

// NodeString always returns empty string.
func (tree) NodeString(int) string {
	return ""
}

// NodePref always returns empty string.
func (tree) NodePref(int) string {
	return ""
}

// NodeEachChild does nothing.
func (tree) NodeEachChild(int, func(int) bool) {
	return
}

// NodeTransit always returns -1.
func (tree) NodeTransit(int, int, byte) int {
	return -1
}

// Tree is the only accessible instance of the implementation.
var Tree tree

// To check, if the implementation is compatible with the interface.
var _ radixt.Tree = Tree
