// Package header provides means and routines for compact radix tree
// implementations to calculate and use information on how to extract required
// for the trees data from instances of types from [internal/node.N] type set.
//
// The package assumes, that the nodes are bit strings, where the following
// data is packed to:
//
//  * chunk's position in string of chunks combined;
//  * chunk's length;
//  * value;
//  * index of first child;
//  * amount of children.
//
// Some data is packed as is, other one is mogrified. So, the bit string
// divided into five continious regions accordingly to names of functions in
// [internal/node]: head, body 0, body 1, body 2, tail. As the head and the
// tail would need a byte parameter each to extract, and the bodies would need
// two bytes each, the total is eight bytes. That's the length of a header.
package header
