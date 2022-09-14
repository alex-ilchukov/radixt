// Package null contains an implementation of radix tree accordingly to
// interface in the parent radixt package. The implementation represents an
// empty tree and provides an instance of the tree with no factory method. All
// integer numbers are non-node indices for the tree. The tree is totally
// static and safe for use by multiple goroutines concurrently.
package null
