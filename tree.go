package radixt

// Tree is the interface for radix tree implementations. Besides the methods
// below, the implementations should also serve the following contract lines.
//
//  1. The trees are static and read-only. That means that their users suppose
//     that no new nodes are added or deleted after the users got an instance
//     of tree. That also means, that trees are safe for concurrent use.
//  2. Nodes of a tree are named by unique indices. In the documentation below
//     the indices (named by n parameter) are identified with nodes, that is,
//     instead of "node with index n" just "node n" expression is used.
//  3. Node indices of non-empty tree go from zero to (tree size - 1) number in
//     sequential order: 0, 1, ….
//  4. Non-empty tree always rocks zero as its root.
//  5. Children indices go in sequential order for any parent node: f, (f + 1),
//     (f + 2), ….
//  6. Any child's index is always strictly greater than index of its parent.
//  7. Empty chunk is allowed only for root nodes of the trees (so any child
//     node must have non-empty chunk).
//  8. First bytes of children chunks should be unique over every parent node.
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

	// ChildrenRange should return first and last indices of children of
	// node n, if the tree has the node and the node has children, or 1 and
	// 0 otherwise.
	ChildrenRange(n uint) (f, l uint)
}
