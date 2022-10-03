package lookup_test

import (
	"bufio"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/str3"
	"github.com/alex-ilchukov/radixt/compact/str4"
	"github.com/alex-ilchukov/radixt/compact/strg"
	"github.com/alex-ilchukov/radixt/compact/structg"
	"github.com/alex-ilchukov/radixt/compact/struct32"
	"github.com/alex-ilchukov/radixt/compact/struct64"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/lookup"
	"github.com/alex-ilchukov/radixt/sapling"
)

func loadLines(path string) (lines []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func createSaplingTreeFromLines(path string) (radixt.Tree, []string) {
	lines := loadLines(path)
	return sapling.New(lines...), lines
}

func createMapFromLines(path string) (map[string]uint, []string) {
	lines := loadLines(path)
	result := make(map[string]uint, len(lines))
	for i, l := range lines {
		result[l] = uint(i)
	}

	return result, lines
}

const linesAmount = 20

func chooseSomeLines(l []string) []string {
	rand.Seed(time.Now().UnixNano())

	amount := linesAmount
	result := make([]string, amount, amount)
	for i := 0; i < amount; i++ {
		result[i] = l[rand.Intn(len(l))]
	}

	return result
}

func feed(l *lookup.L, s string) {
	for i := 0; i < len(s); i++ {
		if !l.Feed(s[i]) {
			return
		}
	}
}

func benchmarkLookupInTree(b *testing.B, t radixt.Tree, lines []string) {
	l := lookup.New(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(lines); j++ {
			l.Reset()
			feed(l, lines[j])
		}
	}
}

var lookupInMapResult uint

func benchmarkLookupInMap(b *testing.B, m map[string]uint, lines []string) {
	var r uint
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(lines); j++ {
			r = m[lines[j]]
		}
	}
	lookupInMapResult = r
}

const (
	methods   = "./testdata/methods.txt"
	headers   = "./testdata/headers.txt"
	goals     = "./testdata/goals.txt"
	words200k = "./testdata/words200k.txt"
	words     = "./testdata/words.txt"
)

func BenchmarkLookupMethodsInMap(b *testing.B) {
	m, lines := createMapFromLines(methods)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInGeneric(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t := generic.New(s)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStrgN3(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := strg.New[strg.N3](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStr3(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := str3.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStrgN4(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := strg.New[strg.N4](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStr4(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := str4.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStructgUint32(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := structg.New[uint32](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStructgUint64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := structg.New[uint64](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStruct32(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := struct32.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStruct64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t, err := struct64.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInEvident(b *testing.B) {
	s, lines := createSaplingTreeFromLines(methods)
	t := evident.New(s)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInMap(b *testing.B) {
	m, lines := createMapFromLines(headers)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInGeneric(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t := generic.New(s)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStrgN4(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t, err := strg.New[strg.N4](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStr4(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t, err := str4.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStructgUint32(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t, err := structg.New[uint32](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStructgUint64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t, err := structg.New[uint64](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStruct32(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t, err := struct32.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStruct64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t, err := struct64.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInEvident(b *testing.B) {
	s, lines := createSaplingTreeFromLines(headers)
	t := evident.New(s)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInMap(b *testing.B) {
	m, lines := createMapFromLines(goals)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInGeneric(b *testing.B) {
	s, lines := createSaplingTreeFromLines(goals)
	t := generic.New(s)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInStructgUint64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(goals)
	t, err := structg.New[uint64](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInStruct64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(goals)
	t, err := struct64.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInMap(b *testing.B) {
	m, lines := createMapFromLines(words200k)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInGeneric(b *testing.B) {
	s, lines := createSaplingTreeFromLines(words200k)
	t := generic.New(s)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInStructgUint64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(words200k)
	t, err := structg.New[uint64](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInStruct64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(words200k)
	t, err := struct64.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInMap(b *testing.B) {
	m, lines := createMapFromLines(words)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInGeneric(b *testing.B) {
	s, lines := createSaplingTreeFromLines(words)
	t := generic.New(s)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInStructgUint64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(words)
	t, err := structg.New[uint64](s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInStruct64(b *testing.B) {
	s, lines := createSaplingTreeFromLines(words)
	t, err := struct64.New(s)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}
