package analysis

// N struct represents data on a node in radix tree in analysis result. Most of
// the data is return values of Node* methods in [radixt.Tree] interface.
type N struct {
	// Index of the node in the tree.
	Index int

	// Chunk of the node.
	Chunk string

	// String, associated with the node, or empty string if no string is
	// associated.
	String string

	// Mark of the node
	Mark int

	// Index of parent of the node, or non-node index if the node is root.
	Parent int

	// Slice of indices of children of the node. Always is sorted by
	// ascending of the indices.
	Children []int

	// Position of chunk in the [A.P] string.
	ChunkPos int
}

// A struct represents result data of analysis of radix tree.
type A struct {
	// P is the string of all prefixes "crammed" together.
	P string

	// Pml is the maximum over prefix lengths of all nodes.
	Pml int

	// N is map from all node indices to node data in form of [N] structs.
	N map[int]N

	// Nt is map from all node indices to new indices, which would allow to
	// represent all slices of [N.Children] as two numbers: start index and
	// amount of children.
	Nt map[int]int

	// Ca is map from amounts of children to amount of nodes with those
	// children.
	Ca map[int]int
}
