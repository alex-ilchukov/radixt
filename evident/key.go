package evident

import "strconv"

const delim = '|'

const noDelimError = "'|' symbol is not found"

func mustFindDelim(key string) int {
	for i := len(key) - 2; i >= 0; i-- {
		if key[i] == delim {
			return i
		}
	}

	panic(noDelimError)

	return -1
}

func extractValue(key string) (v uint, has bool) {
	// Fast check
	i := len(key) - 1
	if key[i] == delim {
		return
	}

	// Slow check
	i = mustFindDelim(key) + 1
	v64, err := strconv.ParseUint(key[i:], 0, 0)
	if err != nil {
		panic(err)
	}

	v = uint(v64)
	has = true

	return
}

func extractChunk(key string) string {
	// Fast check
	i := len(key) - 1
	if key[i] != delim {
		// Slow check
		i = mustFindDelim(key)
	}

	return key[:i]
}
