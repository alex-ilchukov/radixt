package main

import (
	"fmt"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/strg"
	"github.com/alex-ilchukov/radixt/lookup"
	"github.com/alex-ilchukov/radixt/null"
	"github.com/alex-ilchukov/radixt/sapling"
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
	headers := sapling.New(
		"authorization",
		"content-type",
		"content-disposition",
		"content-length",
	)
	trees := map[string]radixt.Tree{
		"null":          null.Tree,
		"headers":       headers,
		"headers-strg3": strg.MustCreate[strg.N3](headers),
		"headers-strg4": strg.MustCreate[strg.N4](headers),
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
