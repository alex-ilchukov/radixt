package lookup_test

import (
	"bufio"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/alex-ilchukov/radixt"
	"github.com/alex-ilchukov/radixt/compact/strg"
	"github.com/alex-ilchukov/radixt/compact/structg"
	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
	"github.com/alex-ilchukov/radixt/lookup"
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

func createGenericTreeFromLines(path string) (lookup.Tritcher, []string) {
	lines := loadLines(path)
	return generic.New(lines...), lines
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

func feedLS(l *lookup.LS, s string) {
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

func benchmarkLookupInTritcher(
	b     *testing.B,
	t     lookup.Tritcher,
	lines []string,
) {
	l := lookup.NewInTritcher(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(lines); j++ {
			l.Reset()
			feedLS(l, lines[j])
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
	g, lines := createGenericTreeFromLines(methods)
	benchmarkLookupInTree(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInGenericTritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	benchmarkLookupInTritcher(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStrgN3(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := strg.New[strg.N3](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStrgN3Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := strg.New[strg.N3](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStrgN4(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := strg.New[strg.N4](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStrgN4Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := strg.New[strg.N4](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStructgUint32(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := structg.New[uint32](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStructgUint32Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := structg.New[uint32](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStructgUint64(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInStructgUint64Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupMethodsInEvident(b *testing.B) {
	g, lines := createGenericTreeFromLines(methods)
	t := evident.New(g)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInMap(b *testing.B) {
	m, lines := createMapFromLines(headers)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInGeneric(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	benchmarkLookupInTree(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInGenericTritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	benchmarkLookupInTritcher(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStrgN4(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	t, err := strg.New[strg.N4](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStrgN4Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	t, err := strg.New[strg.N4](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStructgUint32(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	t, err := structg.New[uint32](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStructgUint32Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	t, err := structg.New[uint32](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStructgUint64(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInStructgUint64Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupHeadersInEvident(b *testing.B) {
	g, lines := createGenericTreeFromLines(headers)
	t := evident.New(g)
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInMap(b *testing.B) {
	m, lines := createMapFromLines(goals)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInGeneric(b *testing.B) {
	g, lines := createGenericTreeFromLines(goals)
	benchmarkLookupInTree(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInGenericTritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(goals)
	benchmarkLookupInTritcher(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInStructgUint64(b *testing.B) {
	g, lines := createGenericTreeFromLines(goals)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupGoalsInStructgUint64Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(goals)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInMap(b *testing.B) {
	m, lines := createMapFromLines(words200k)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInGeneric(b *testing.B) {
	g, lines := createGenericTreeFromLines(words200k)
	benchmarkLookupInTree(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInGenericTritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(words200k)
	benchmarkLookupInTritcher(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInStructgUint64(b *testing.B) {
	g, lines := createGenericTreeFromLines(words200k)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWords200kInStructgUint64Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(words200k)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInMap(b *testing.B) {
	m, lines := createMapFromLines(words)
	benchmarkLookupInMap(b, m, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInGeneric(b *testing.B) {
	g, lines := createGenericTreeFromLines(words)
	benchmarkLookupInTree(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInGenericTritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(words)
	benchmarkLookupInTritcher(b, g, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInStructgUint64(b *testing.B) {
	g, lines := createGenericTreeFromLines(words)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTree(b, t, chooseSomeLines(lines))
}

func BenchmarkLookupWordsInStructgUint64Tritcher(b *testing.B) {
	g, lines := createGenericTreeFromLines(words)
	t, err := structg.New[uint64](g)
	if err != nil {
		panic(err)
	}
	benchmarkLookupInTritcher(b, t, chooseSomeLines(lines))
}
