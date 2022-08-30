package generic

type imago struct {
	chunk    string
	value    uint
	children []uint
}

type shrub struct {
	noValue uint
	imagoes []imago
}

func (s *shrub) transit(n, npos uint, b byte) (bool, uint) {
	imagoes := s.imagoes
	im := imagoes[n]
	chunk := im.chunk
	if npos < uint(len(chunk)) {
		return chunk[npos] == b, n
	}

	for _, c := range im.children {
		if imagoes[c].chunk[0] == b {
			return true, c
		}
	}

	return false, n
}

func (s *shrub) find(str string) (found bool, n, pos, npos uint) {
	l := uint(len(str))
	for ; pos < l; pos++ {
		f, m := s.transit(n, npos, str[pos])
		switch {
		case !f:
			return
		case m == n:
			npos++
		default:
			npos = 1
			n = m
		}
	}

	found = true

	return
}

func (s *shrub) within(n, npos uint) bool {
	return npos < uint(len(s.imagoes[n].chunk))
}

func (s *shrub) insert(str string, value uint) {
	found, n, pos, npos := s.find(str)
	switch {
	case !found:
		if s.within(n, npos) {
			s.splitNode(n, npos, s.noValue)
		}

		s.addChild(n, str[pos:], value)

	case s.within(n, npos):
		s.splitNode(n, npos, value)

	case s.imagoes[n].value == s.noValue:
		s.imagoes[n].value = value
	}
}

func (s *shrub) splitNode(n, npos, value uint) {
	no := s.imagoes[n]
	chunk := no.chunk
	no.chunk = chunk[npos:]
	s.imagoes = append(s.imagoes, no)
	s.imagoes[n] = imago{
		chunk:    chunk[:npos],
		value:    value,
		children: []uint{uint(len(s.imagoes) - 1)},
	}
}

func (s *shrub) addChild(n uint, chunk string, value uint) {
	s.imagoes = append(s.imagoes, imago{chunk: chunk, value: value})
	s.imagoes[n].children = append(
		s.imagoes[n].children,
		uint(len(s.imagoes)-1),
	)
}
