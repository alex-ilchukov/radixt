package pass

// Yielder represents a processor of a current node during the pass through the
// provided radix tree.
type Yielder interface {
	// Yield is a method of processing current node n on iteration i, which
	// also serves as new node index in the enumeration pass. The pass
	// process also provides tag parameter, which is zero if n is zero (and
	// root if so) or filled during previous iterations. The return value
	// of the method should be tag for children of the node, and it would
	// be passed as tag for them on their iterations.
	Yield(i, n, tag uint) (ctag uint)
}
