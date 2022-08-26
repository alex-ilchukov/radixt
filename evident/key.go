package evident

const delim = '|'

func findDelim(key string) int {
	for i := len(key) - 1; i >= 0; i-- {
		if key[i] == delim {
			return i
		}
	}

	return -1
}

func extractValue(key string) string {
	i := findDelim(key)
	if i == -1 {
		panic("'|' symbol is not found")
	}

	return key[i+1:]
}

func extractChunk(key string) string {
	i := findDelim(key)
	if i == -1 {
		panic("'|' symbol is not found")
	}

	return key[:i]
}
