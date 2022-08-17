package main

import (
	"fmt"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/lookup"
	"github.com/alex-ilchukov/radixt/null"
)

func feed(l *lookup.L, s string) {
	n := len(s)
	for i := 0; i < n; i++ {
		if !l.Feed(s[i]) {
			return
		}
	}
}

func main() {
	trees := map[string]radixt.Tree{
		"null": null.Tree,
		"headers": generic.New(
			"authorization",
			"content-type",
			"content-disposition",
			"content-length",
		),
	}

	strings := []string{"content-length", "authorization", "auth", "host"}

	for n, tree := range trees {
		l := lookup.New(tree)
		for _, s := range strings {
			l.Reset()
			feed(l, s)
			f := " not "
			if l.Found() {
				f = " "
			}
			fmt.Printf("'%s' is%sfound in %s tree\n", s, f, n)
		}
	}
}
