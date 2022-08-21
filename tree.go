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
//  4. Empty chunk is allowed only for root nodes of the trees.
//
// The interface does not put any limitations on order of node indices, leaving
// that detail to implementations. It also does not put any limitations on
// values.
type Tree interface {
	// Size should return amount of nodes in the tree.
	Size() uint

	// Has should return if the tree has node n or not.
	Has(n int) bool

	// Root should return index of root node if the tree is not empty, or
	// non-node index otherwise.
	Root() int

	// Value should return value v of node n with boolean true flag, if the
	// tree has the node and the node has value, or default unsigned
	// integer with boolean false otherwise.
	Value(n int) (v uint, has bool)

	// EachChild should call func e for every child of the node n, until e
	// returns boolean true.
	EachChild(n int, e func(int) bool)

	// ByteAt should return byte b at npos of chunk of node n with boolean
	// true flag, if the tree has the node and npos is within the chunk, or
	// default byte value and boolean false otherwise.
	ByteAt(n int, npos uint) (b byte, within bool)
}
