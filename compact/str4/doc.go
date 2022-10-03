// Package str4 contains a compactified implementation of radix tree
// accordingly to interface in the parent radixt packagei, based on regular Go
// strings.
//
// The implementation is aimed to have reduced memory footprint in comparison
// with generic implementation: The most of node information are contained in
// just 4 bytes. As it provides only limited abilities to store chunks and
// values of tree nodes, it is not aimed to cover all cases of input data. It
// is totally static and safe to use by multiple goroutines concurrently.
//
// The package also provides factory method to create a compactified copy of
// the provided tree. Also, the copies could be saved as regular Go strings
// (even constant ones).
//
// The implementation has the following feature: all strings with length less
// than [ProperLen] constant are considered valid empty trees. Empty string, of
// course, is in the category too. Value of the [ProperLen] depends on length
// of headers, but it must be strictly more than 2 and less than 18.
package str4
