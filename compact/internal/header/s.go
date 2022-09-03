package header

// S represents a set of shifts to organize a node for compactified
// implementation of radix tree.
type S struct {
	Value          byte
	ChunkPos       byte
	ChunkLen       byte
	ChildrenStart  byte
	ChildrenAmount byte
}
