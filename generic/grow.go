package generic

func grow(s *shrub) *tree {
	imagoes := s.imagoes
	l := len(imagoes)
	nt := make(map[uint]uint, l)
	q := make([]uint, 1, l)
	for len(q) > 0 {
		h := q[0]
		q = q[1:]
		for _, c := range imagoes[h].children {
			q = append(q, c)
		}
		nt[h] = uint(len(nt))
	}

	nodes := make([]node, l, l)
	for i, n := range nt {
		im := imagoes[i]
		children := im.children
		cAmount := byte(len(children))
		cFirst := uint(0)
		if cAmount > 0 {
			// shrub.splitNode and shrub.addChild guarantee, that
			// children slice has indices in ascending order, so
			// they go into q in the same order, pop out q in the
			// same order, and get into nt in the same order
			cFirst = nt[children[0]]
		}

		nodes[n] = node{
			cAmount: cAmount,
			cFirst:  cFirst,
			chunk:   im.chunk,
			value:   im.value,
		}
	}

	return &tree{noValue: s.noValue, nodes: nodes}
}
