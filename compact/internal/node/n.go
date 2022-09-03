package node

// N is union type set for types of nodes of compact radix tree
// implementations.
type N interface {
	~uint32 | ~uint64
}

// BitsLen returns 32 for types with underlying type uint32 and 64 for types
// with underlying type uint64.
func BitsLen[T N]() int {
	if ^T(0)>>32 > 0 {
		return 64
	}

	return 32
}

// Head returns the lowest bits of n in amount of (BitsLen[T]() - s) in form of
// unsigned integer. If s is more or equal to BitsLen[T](), then result is 0.
//
// For example, Head[uint32](0xABCDEF76, 27) would return the lowest 5 bits in
// form of unsigned integer, that is, 0x16.
func Head[T N](n T, s byte) uint {
	return uint(n << s >> s)
}

// Body returns the bits of n, which are in the following range:
//
//  [rs - ls, BitsLen[T]() - ls - 1]
//
// (here bits are numbered from lowest 0 to highest (BitsLen[T] - 1)). The bits
// are returned in form of unsigned integer. If ls is more or equal to
// BitsLen[T](), then result is 0. If rs is less than ls, then the range starts
// with 0 and result is shifted to right by (ls - rs) zero bits. If rs is more
// or equal to BitsLen[T](), then result is 0.
//
// For example, Body[uint32](0xABCDEF76, 22, 27) would return the bits from
// [5, 9] range, that is 0x1B.
func Body[T N](n T, ls, rs byte) uint {
	return uint(n << ls >> rs)
}

// Tail returns the highest bits of n, starting from bit s, in form of unsigned
// integer. If s is more or equal to BitsLen[T](), then result is 0.
//
// For example, Tail[uint32](0xABCDEF76, 27) would return the highest 5 bits in
// form of unsigned integer, that is, 0x15.
func Tail[T N](n T, s byte) uint {
	return uint(n >> s)
}
