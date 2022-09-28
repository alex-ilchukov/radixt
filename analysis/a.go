package analysis

// N struct represents data on a node in radix tree in analysis result. All the
// indices presented in instances of the struct are result of renaming process
// (see documentation on [A] struct). The struct depends on type parameter M
// from [Mode] type set, which represents chosen type of processing for node's
// chunk, and that reflects on meaning of some of its fields (see documentation
// on [Firstless] processing type).
type N[M Mode] struct {
	// HasValue reflects the fact if the node has value or not.
	HasValue bool

	// ChunkFirst is the first byte of node's original chunk. It holds
	// default value if the node has empty chunk.
	//
	// Remark: The field does _not_ depend on provided type parameter M.
	ChunkFirst byte

	// ChunkEmpty reflects the fact if the node has empty original chunk or
	// not.
	//
	// Remark: The field does _not_ depend on provided type parameter M.
	ChunkEmpty bool

	// Index is index of the node in the tree.
	Index uint

	// Chunk is chunk of the node. Depending on provided processing mode M,
	// it hold either original chunk ([Default] processing) or original
	// chunk without first byte ([Firstless] processing).
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

	// ChildrenHigh is incremented last (maximum) index of children nodes
	// of the node, if the node has children, or 0 otherwise. "High" here
	// has the same meaning as in slice expression like slice[low : high].
	//
	// Remark: Amount of children of the node can be calculated as
	// ChildrenHigh - [ChildrenLow].
	ChildrenHigh uint

	// ChunkPos is position of [Chunk] in the [A.C] string.
	//
	// Remark: As [Chunk] depends on provided parameter M, ChunkPos depends
	// on it too.
	ChunkPos uint
}

// A struct represents result data of analysis of radix tree.
//
// User of the struct's instance can assume, that the data regarding indices
// reflects results on _renaming_ (reindexing). The renamed indices have the
// following property: all children of any node have their indices in
// sequential order (l, l + 1, l + 2 and so on, that is). The property allows
// to explain the sequences with use of [N.ChildrenLow] and [N.ChildrenHigh]
// fields. The only place, where original indices residue, is indices of [A.N]
// field.
//
// The struct depends on type parameter M from [Mode] type set, which
// represents chosen type of processing for node's chunk, and that reflects on
// meaning of some of its fields (see documentation on [Firstless] processing
// type).
type A[M Mode] struct {
	// N is slice of instances of [analysis.N] struct. Its indices are
	// original node indices in the tree analyzed.
	N []N[M]

	// C is the string, made of all [N.Chunk] "crammed" together.
	//
	// Remark: As [N.Chunk] depends on provided parameter M, C depends on
	// it too.
	C string

	// Cml is the maximum of all lengths of [N.Chunk].

	// Remark: As [N.Chunk] depends on provided parameter M, Cml depends on
	// it too.
	Cml uint

	// Cma is the maximum over children amounts of all nodes.
	Cma uint

	// Dclpm is the maximum over differences between children's low indices
	// and their parent indices.
	Dclpm uint

	// Vm is the maximum over values of all nodes.
	Vm uint
}
