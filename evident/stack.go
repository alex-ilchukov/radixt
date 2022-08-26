package evident

type stack struct {
	a []Tree
}

func newStack(t Tree) *stack {
	return &stack{[]Tree{t}}
}

func (s *stack) popAndPushChildren() Tree {
	a := s.a
	l := len(a) - 1
	t := a[l]
	s.a = a[:l]
	for _, c := range t {
		if c != nil {
			s.a = append(s.a, c)
		}
	}

	return t
}

func (s *stack) populated() bool {
	return len(s.a) > 0
}
