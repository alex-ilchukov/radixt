package struct8

type node uint64

const sbits = 0x3F

func (n node) head(s byte) uint {
	s &= sbits
	return uint(n << s >> s)
}

func (n node) body(ls, rs byte) uint {
	ls &= sbits
	rs &= sbits
	return uint(n << ls >> rs)
}

func (n node) tail(s byte) uint {
	s &= sbits
	return uint(n >> s)
}
