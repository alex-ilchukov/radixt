package analysis

type N struct {
	Index    int
	Pref     string
	String   string
	Mark     int
	Parent   int
	Children []int
	PrefPos  int
}

type A struct {
	P   string
	Pml int
	N   map[int]N
	Nt  map[int]int
	Ca  map[int]int
}
