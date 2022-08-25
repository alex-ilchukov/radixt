// Package struct4 contains a compactified implementation of radix tree
// accordingly to interface in the parent radixt package.
//
// The implementation is aimed to reduced memory footprint in comparison with
// generic implementation: The most of node information are contained in just 4
// bytes. As it provides only limited abilities to store chunks and values of
// tree nodes, it is not aimed to cover all cases of input data. It has the
// same additional limitations, as the trees in generic package do.

//  1. Non-empty trees always have zero as their root node.
//  2. Node indices of non-empty trees go sequentially from zero til their
//     (Size() - 1) numbers; other numbers are non-node indices.
//
// The package also provides factory method to create a compactified copy of
// the provided tree. As the tree struct is not exported outside, the
// implementation assumes that instance is never nil.
package struct4
