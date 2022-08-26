// Package evident provides an implementation of radix tree interface from the
// parent radixt package.
//
// The implementation is aimed at _evident_ representation of a tree, and
// allows to define one in the following way:
//
//	var atree = evident.Tree{
//		"|": {
//			"auth|4": {
//				"or|2": {
//					"i|": {
//						"ty|0": nil,
//						"zation|1": nil,
//					},
//				},
//				"entication|3": nil,
//			},
//			"content-|": {
//				"type|5": nil,
//				"length|6": nil,
//				"disposition|7": nil,
//			},
//		},
//	}
//
// Here keys in the maps encapsulate both node chunk and node value with use of
// '|' delimiter.
//
// The implementation is not aimed at low memory consumption or high
// performance. As the maps are unordered, stabilization is achieved via
// sorting of chunks. If a key is found, which doesn't satisfy format with '|'
// above, the implementation panics. Technically, keys with the same chunk are
// allowed in the same map and wouldn't bring a panic, but that can bring chaos
// to lookup process. The implementation does not provide any factory function
// and because of that never checks the tree during construction.
package evident
