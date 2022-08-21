// Package radixt provides an interface to work with radix trees.
//
// Radix trees are trees of strings, which are interpreted as prefixes or
// chunks. Going through nodes of such a tree from root to some target node,
// a result string can be constructed.
//
// The trees can be considered as a very special implementation of associative
// array with result strings as its keys. The interface in the package provides
// unsigned integer values. The main distinction is efficient lookup of a key,
// as it is not required to be stored (buffered) and can be represented as a
// stream of bytes (which are read, for example, with help of [io.ByteReader]
// implementation). A lookup implementation would need to keep only current
// node and a position within its chunk.
//
// Example:
//
//	            "", Value: (0, false)
//	        /                            \
//	"authorization", Value: (0, true)    "content-", Value: (0, false)
//	                                      /                  \
//	          "length", Value: (1, true)        "type", Value: (2, true)
//
// In the example values are presented as couples of unsigned integer and
// boolean flag of node having value. That allows to distinguish keys
// "authorization", "content-length", "content-type" from non-keys "" and
// "content-".
package radixt
