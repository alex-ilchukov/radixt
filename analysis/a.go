package analysis

// N struct represents data on a node in radix tree in analysis result. All the
// indices presented in instances of the struct are result of renaming process
// (see documentation on [A] struct).
type N struct {
	// HasValue reflects the fact if the node has value or not.
	HasValue bool

	// ChunkFirst is the first byte of node's original chunk. It holds
	// default value if the node has empty chunk.
	ChunkFirst byte

	// Index is index of the node in the tree.
	Index uint

	// Chunk is chunk of the node.
	Chunk string

	// Value is value of the node, if the node has value (see [HasValue]),
	// or default unsigned integer otherwise.
	Value uint

	// Parent is index of parent of the node, if the node is not root (see
	// [Root] flag), or default integer value otherwise.
	Parent uint

	// ChildrenLow is first (minimum) index of children nodes of the
	// node, if the node has children, or 0 otherwise. "Low" here has the
	// same meaning as in slice expression like slice[low : high].
	ChildrenLow uint

	// ChildrenLast is incremented last (maximum) index of children nodes
	// of the node, if the node has children, or 0 otherwise. "High" here
	// has the same meaning as in slice expression like slice[low : high].
	//
	// Remark: Amount of children of the node can be calculated as
	// ChildrenHigh - [ChildrenLow].
	ChildrenHigh uint

	// ChunkPos is position of chunk in the [A.C] string.
	ChunkPos uint
}

// A struct represents result data of analysis of radix tree.
//
// User of the struct's instance can assume, that the data regarding indices
// reflects results on _renaming_ (reindexing). The renamed indices have the
// following property: all children of any node have their indices in
// sequential order (l, l + 1, l + 2 and so on, that is). The property allows
// to explain the sequences with use of [N.ChildrenLow] and [N.ChildrenHigh]
// fields.
//
// The only place, where original indices residue, is indices of [A.N] field.
type A struct {
	// C is the string of all node chunks "crammed" together.
	C string

	// Cml is the maximum over chunk lengths of all nodes.
	Cml uint

	// Cma is the maximum over children amounts of all nodes.
	Cma uint

	// Dclpm is the maximum over differences between children's low indices
	// and their parent indices.
	Dclpm uint

	// Vm is the maximum over values of all nodes.
	Vm uint

	// N is slice of instance of [N] struct. Its indices are original node
	// indices in the tree analyzed.
	N []N
}
