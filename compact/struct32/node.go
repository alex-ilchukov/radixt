package struct32

type node uint32

func head(n node, s byte) uint {
	s &= 0x1F
	return uint(n << s >> s)
}

func body(n node, ls, rs byte) uint {
	if rs > 0x1F {
		return 0
	}

	ls &= 0x1F
	rs &= 0x1F
	return uint(n << ls >> rs)
}

func tail(n node, s byte) uint {
	s &= 0x1F
	return uint(n >> s)
}
