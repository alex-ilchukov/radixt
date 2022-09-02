package analysis

// N struct represents data on a node in radix tree in analysis result. Most of
// the data is return values of Node* methods in [radixt.Tree] interface.
type N struct {
	// Index is index of the node in the tree.
	Index uint

	// Chunk is chunk of the node.
	Chunk string

	// HasValue reflects the fact if the node has value or not.
	HasValue bool

	// Value is value of the node, if the node has value (see [HasValue]),
	// or default unsigned integer otherwise.
	Value uint

	// Root is boolean flag, which indicates if the node is root or not.
	Root bool

	// Parent is index of parent of the node, if the node is not root (see
	// [Root] flag), or default integer value otherwise.
	Parent uint

	// ChildrenFirst is first (minimum) index of children nodes of the
	// node, if the node has children, or 1 otherwise
	ChildrenFirst uint

	// ChildrenLast is last (maximum) index of children nodes of the node,
	// if the node has children, or 0 otherwise
	ChildrenLast uint

	// ChunkPos is position of chunk in the [A.P] string.
	ChunkPos uint
}

// A struct represents result data of analysis of radix tree.
type A struct {
	// C is the string of all node chunks "crammed" together.
	C string

	// Cml is the maximum over chunk lengths of all nodes.
	Cml uint

	// Cma is the maximum over children amounts of all nodes.
	Cma uint

	// Dcfpm is the maximum over differences between children's first
	// indices and their parent indices
	Dcfpm uint

	// Vm is the maximum over values of all nodes.
	Vm uint

	// N is map from all node indices to node data in form of [N] structs.
	N map[uint]N

	// Ca is map from amounts of children to amount of nodes with those
	// children.
	Ca map[uint]uint
}
