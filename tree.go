package radixt

// Tree is the interface for radix tree implementations. Besides the methods
// below, the implementations should also serve the following contract lines.
//
//  1. Nodes of a tree are named by unique indices. In the documentation below
//     the indices (named by n parameter) are identified with nodes, that is,
//     instead of "node with index n" just "node n" expression is used.
//  2. Node indices of non-empty tree go from zero to (tree size - 1) number in
//     sequential order: 0, 1, ….
//  3. Non-empty tree always rocks zero as its root.
//  4. Any child's index is always strictly greater than index of its parent.
//  5. Empty chunk is allowed only for root nodes of the trees (so any child
//     node must have non-empty chunk).
//  6. First bytes of children chunks should be unique over every parent node.
//
// The interface does not put any limitations on values besides its domain of
// unsigned integers.
type Tree interface {
	// Size should return amount of nodes in the tree.
	Size() uint

	// Value should return value v of node n with boolean true flag, if the
	// tree has the node and the node has value, or default unsigned
	// integer with boolean false otherwise.
	Value(n uint) (v uint, has bool)

	// Chunk should return chunk of node n, if the tree has the node, or
	// empty string otherwise.
	Chunk(n uint) string

	// ChildrenRange should return low and high indices of children of
	// node n, if the tree has the node and the node has children, or
	// default unsigned integers otherwise. (The terms "low" and "high"
	// have the same meaning as in the [specification] of Go slice
	// expressions.)
	//
	// [specification]: https://go.dev/ref/spec#Slice_expressions
	ChildrenRange(n uint) (low, high uint)
}
