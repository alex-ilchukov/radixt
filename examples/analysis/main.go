package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/generic"
)

func createTreeFromLines(path string) (t radixt.Tree, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	t = generic.New(lines...)

	return
}

const bitsRequired = "\t(%d bits required)\n"

func printWithBitsRequired(template string, value uint) {
	fmt.Printf(template, value)
	fmt.Printf(bitsRequired, bits.Len(value))
}

const (
	lenac = "\tTotal length of chunks crammed together: %d"
	acml  = "\tMaximum chunk length:                    %d"
	avm   = "\tMaximum value:                           %d"
	lenan = "\tAmount of nodes:                         %d"
	acma  = "\tMaximum amount of children:              %d"
)

func printAnalysis(name string, a analysis.A) {
	fmt.Printf("Analysis on %s tree:\n", name)
	printWithBitsRequired(lenac, uint(len(a.C)))
	printWithBitsRequired(acml, a.Cml)
	printWithBitsRequired(avm, a.Vm)
	printWithBitsRequired(lenan, uint(len(a.N)))
	printWithBitsRequired(acma, a.Cma)
	fmt.Printf("\tAmounts of children to amounts of nodes: %v\n", a.Ca)
}

func printErr(path string, err error) {
	fmt.Printf("Error has appeared during loading %s: %s\n", path, err)
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
		t, err := createTreeFromLines(path)
		if err != nil {
			printErr(path, err)
		}

		a := analysis.Do(t)
		printAnalysis(name, a)
		fmt.Println()
	}
}