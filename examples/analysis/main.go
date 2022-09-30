package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/analysis"
	"github.com/alex-ilchukov/radixt/sapling"
)

func createTreeFromLines(path string) (t radixt.Tree, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	s := sapling.New()
	for i := uint(0); scanner.Scan(); i++ {
		s.Grow(scanner.Text(), i)
	}
	t = s

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

func modeString[M analysis.Mode]() string {
	var m M
	if len(m) == 0 {
		return "in default mode"
	}

	return "in firstless mode"
}

func printAnalysis[M analysis.Mode](name string, a analysis.A[M]) {
	fmt.Printf("Analysis %s on %s tree:\n", modeString[M](), name)
	printWithBitsRequired(lenac, uint(len(a.C)))
	printWithBitsRequired(acml, a.Cml)
	printWithBitsRequired(avm, a.Vm)
	printWithBitsRequired(lenan, uint(len(a.N)))
	printWithBitsRequired(acma, a.Cma)
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

		printAnalysis(name, analysis.Do[analysis.Default](t))
		fmt.Println()
		printAnalysis(name, analysis.Do[analysis.Firstless](t))
		fmt.Println()
	}
}
