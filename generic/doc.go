// Package generic contains a generic implementation of radix tree accordingly
// to interface in the parent radixt package.
//
// The implementation is aimed to cover all cases of input data and does not
// care much of consumed memory. It has additional limitations to the listed
// ones in the interface definition.
//
//  1. Non-empty trees always have zero as their root node.
//  2. Node indices of non-empty trees go sequentially from zero til their
//     (Size() - 1) numbers; other numbers are non-node indices.
//  3. Empty prefix is allowed only for root node.
//
// The package also provides factory method to create an instance from the
// provided slice of strings. As the tree struct is not exported outside, the
// implementation assumes that instance is never nil.
package generic
