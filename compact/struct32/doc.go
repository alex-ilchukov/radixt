// Package struct32 contains a compactified implementation of radix tree
// accordingly to interface in the parent radixt package.
//
// The implementation is aimed to have reduced memory footprint in comparison
// with generic implementation: The most of node information is contained in
// just 4 bytes. As it provides only limited abilities to store chunks and
// values of tree nodes, it is not aimed to cover all cases of input data. It
// is totally static and safe to use by multiple goroutines concurrently.
//
// The package also provides factory method to create a compactified copy of
// the provided tree. As the tree struct is not exported outside, the
// implementation assumes that instance is never nil.
package struct32
