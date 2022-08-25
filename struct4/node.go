package struct4

type node uint32

const sbits = 0x1F

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
