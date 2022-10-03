package str3

type node uint32

const (
	maskShift = 0x1F
	nodeLen = 3
)

func head(n node, s byte) uint {
	s &= maskShift
	return uint(n << s >> s)
}

func body(n node, ls, rs byte) uint {
	if rs > maskShift {
		return 0
	}

	ls &= maskShift
	rs &= maskShift
	return uint(n << ls >> rs)
}

func tail(n node, s byte) uint {
	s &= maskShift
	return uint(n >> s)
}
