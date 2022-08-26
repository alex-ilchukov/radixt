package evident

type queue struct {
	a []Tree
}

func newQueue(t Tree) *queue {
	return &queue{[]Tree{t}}
}

func (q *queue) pop() Tree {
	a := q.a
	t := a[0]
	q.a = a[1:]
	return t
}

func (q *queue) pushChildren(t Tree) {
	for _, k := range t.keys() {
		c := t[k]
		if c != nil {
			q.a = append(q.a, c)
		}
	}
}

func (q *queue) populated() bool {
	return len(q.a) > 0
}
