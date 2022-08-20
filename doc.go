// Package radixt provides an interface to work with radix trees.
//
// Radix trees are trees of strings, which are interpreted as prefixes or
// chunks. Going through nodes of such a tree from root to some target node,
// a result string can be constructed. The target node can be a leaf or an
// inner node, so there are marks to distinguish if a node points to a result
// string or not. Example:
//
//	                  "", mark: 0
//	        /                            \
//	"authorization", mark: 1       "content-", mark: 0
//	                               /                  \
//	                         "length", mark: 2   "type", mark: 3
//
// In the example strings "authorization", "content-length", "content-type" are
// result strings, while "" or "content-" are not.
package radixt
