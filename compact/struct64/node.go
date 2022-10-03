package struct64

type node uint64

func head(n node, s byte) uint {
	s &= 0x3F
	return uint(n << s >> s)
}

func body(n node, ls, rs byte) uint {
	if rs > 0x3F {
		return 0
	}

	ls &= 0x3F
	rs &= 0x3F
	return uint(n << ls >> rs)
}

func tail(n node, s byte) uint {
	s &= 0x3F
	return uint(n >> s)
}
