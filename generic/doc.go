// Package generic contains a generic implementation of radix tree accordingly
// to interface in the parent radixt package.
//
// The implementation is aimed to cover all cases of input data and does not
// care much of consumed memory. The package also provides factory method to
// create an instance from the provided slice of strings, interpretting string
// positions in the slice as values in the resulting tree.
//
// As the tree struct is not exported outside, the implementation assumes that
// instance is never nil. Also, it is totally static and safe to use by
// multiple goroutines concurrently.
package generic
