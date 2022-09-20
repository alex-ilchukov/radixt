package lookup

// Switcher is auxiliary interface for radix trees implementation. Lookup
// process in trees, which realize the interface, can go faster.
type Switcher interface {
	// Switch should take node n and byte b. If the node belongs to the
	// tree, it should look for a child c of the node with such a chunk,
	// that its first byte coincides with b. If such a child is found, it
	// should return the child with its chunk _without first byte_ (chunk
	// is c.chunk[1:], that is) and boolean truth. Otherwise or in the case
	// the node is not in the tree, the method should return corresponding
	// default values (zero, empty string, and boolean false that is).
	Switch(n uint, b byte) (c uint, chunk string, found bool)
}
