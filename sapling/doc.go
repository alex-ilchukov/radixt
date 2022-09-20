// Package sapling contains an implementation of radix tree accordingly
// to interface in the parent radixt package. It also provides a ways to _grow_
// a tree, adding new strings and values into it.
//
// The implementation is aimed to cover all cases of input data and does not
// care much of consumed memory. The package also provides factory methods to
// create an instance from the provided slice of strings, interpretting string
// positions in the slice as values in the resulting tree, or from the provided
// slice of couples of strings and their values.
//
// The tree struct is exported outside, and the implementation supports nil
// pointers to the struct. As the implementation is _dynamic_, it does _not_
// guarantee safety over concurrent reading and writing.
package sapling
