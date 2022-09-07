package strg

// N3 is 3-bytes array and used to select how many bytes per node is used in
// tree implementation.
type N3 [3]byte

// N4 is 4-bytes array and used to select how many bytes per node is used in
// tree implementation.
type N4 [3]byte

// N represents a set of number of bytes for nodes in tree implementation. As
// Go doesn't support constants as generic parameters, the types are used
// instead.
type N interface {
	N3 | N4
}

func bytesLen[n N]() int {
	var a n
	return len(a)
}
