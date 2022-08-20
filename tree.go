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
//  3. First bytes of non-empty children prefixes should be unique for every
//     node.
//
// The interface does not put any limitations on order of node indices, leaving
// that detail to implementations. Also, there is no limitation on amount of
// nodes with empty prefixes.
type Tree interface {
	// Size should return amount of nodes in the tree.
	Size() int

	// Has should return if the tree has node n or not.
	Has(n int) bool

	// Root should return index of root node if the tree is not empty, or
	// non-node index otherwise.
	Root() int

	// NodeMark should return ``mark'' of the node n. That is, it should
	// return negative number, if the node is not associated with a string
	// from original string list, or the index of the associated string
	// otherwise.
	Mark(n int) uint

	// NodeEachChild should call func e for every child of the node n,
	// until e returns boolean true.
	NodeEachChild(n int, e func(int) bool)
}
