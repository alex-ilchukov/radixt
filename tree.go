package radixt

type Tree interface {
	Size() int
	Has(n int) bool
	Root() int
	NodeMark(n int) int
	NodeString(n int) string
	NodePref(n int) string
	NodeEachChild(n int, e func(int) bool) 
	NodeTransit(n, npos int, b byte) int
}
