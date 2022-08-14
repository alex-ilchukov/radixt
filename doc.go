// Package radixt provides an interface to work with radix trees.
//
// Radix trees are trees of strings, which are interpreted as prefixes or
// chunks. Going through nodes of such a tree from root to some target node,
// a result string can be constructed. The target node can be a leaf or an
// inner node, so there are marks to distinguish if a node points to a result
// string or not. Example:
//
//	                        ""
//	        /                            \
//	"authorization", mark: 0       "content-", mark: -1
//	                               /                  \
//	                         "length", mark: 1   "type", mark: 2
//
// Here in the example strings "authorization", "length", "type" are result
// strings, while "content-" is not.
package radixt
