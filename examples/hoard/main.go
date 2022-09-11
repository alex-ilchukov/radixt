package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/strg"
	"github.com/alex-ilchukov/radixt/compact/structg"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
)

func loadLines(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func alloc() uint {
	var stats runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats)
	return uint(stats.Alloc)
}

const amount = 300

var maps [amount]map[string]uint

func estimateMapHoard(lines []string) (hoard uint) {
	for i := 0; i < amount; i++ {
		maps[i] = nil
	}

	before := alloc()
	for i := 0; i < amount; i++ {
		m := make(map[string]uint)
		for j := 0; j < len(lines); j++ {
			m[lines[j]] = uint(j)
		}

		maps[i] = m
	}

	after := alloc()
	hoard = uint((float64(after) - float64(before)) / float64(amount))

	for j := 0; j < len(lines); j++ {
		hoard += 16 + uint(len(lines[j]))
	}

	return
}

func printGenericErr(path string, err error) {
	fmt.Printf("Error has appeared during loading %s: %s\n", path, err)
}

func printHoard(h radixt.Hoarder) {
	var hintStr string

	amount, hint := h.Hoard()
	switch hint {
	case radixt.HoardExactly:
		hintStr = "exactly"
	case radixt.HoardAtLeast:
		hintStr = "at least"
	}

	fmt.Printf("\tHoard is %s %d bytes\n\n", hintStr, amount)
}

func processGeneric(name string, t radixt.Tree) {
	fmt.Printf("Tree of %s has been successfully loaded\n", name)
	h := t.(radixt.Hoarder)
	printHoard(h)

	factories := []struct {
		name    string
		factory func(radixt.Tree) (radixt.Tree, error)
		err     error
	}{
		{
			name: "strg.New[strg.N3]",
			factory: func(t radixt.Tree) (radixt.Tree, error) {
				return strg.New[strg.N3](t)
			},
		},
		{
			name: "strg.New[strg.N4]",
			factory: func(t radixt.Tree) (radixt.Tree, error) {
				return strg.New[strg.N4](t)
			},
		},
		{
			name: "structg.New[uint32]",
			factory: func(t radixt.Tree) (radixt.Tree, error) {
				return structg.New[uint32](t)
			},
		},
		{
			name: "structg.New[uint64]",
			factory: func(t radixt.Tree) (radixt.Tree, error) {
				return structg.New[uint64](t)
			},
		},
	}

	for i, f := range factories {
		fmt.Printf("\tTrying to compactify with %s...\n", f.name)

		c, err := f.factory(t)
		if err != nil {
			fmt.Printf("\tFailure! Error: %s\n\n", err)
			factories[i].err = err
			continue
		}

		fmt.Printf("\tSuccess!\n")
		h := c.(radixt.Hoarder)
		printHoard(h)
	}

	errs := 0
	for _, f := range factories {
		if f.err != nil {
			errs++
		}
	}

	if errs > 2 {
		return
	}

	fmt.Printf("\tThere also is evident representation!\n")
	e := evident.New(t)
	printHoard(e)
}

func main() {
	paths := map[string]string{
		"HTTP methods":               "./methods.txt",
		"HTTP request headers":       "./headers.txt",
		"words from goals documents": "./goals.txt",
		"200k of English words":      "./words200k.txt",
		"all English words":          "./words.txt",
	}

	for name, path := range paths {
		lines, err := loadLines(path)
		if err != nil {
			printGenericErr(path, err)
		}

		processGeneric(name, generic.New(lines...))

		fmt.Printf(
			"\tA map from the %s to unsigned integers would "+
				"estimately hoard %d bytes\n\n",
			name,
			estimateMapHoard(lines),
		)
	}
}
