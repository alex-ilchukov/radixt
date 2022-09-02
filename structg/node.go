package structg

type node interface {
	~uint32 | ~uint64
}

func bitslen[N node]() int {
	if ^N(0)>>32 > 0 {
		return 64
	}

	return 32
}

func head[N node](n N, s byte) uint {
	return uint(n << s >> s)
}

func body[N node](n N, ls, rs byte) uint {
	return uint(n << ls >> rs)
}

func tail[N node](n N, s byte) uint {
	return uint(n >> s)
}
