package utils

func ReplaceAtIndex(in string, replacementChar rune, i int) string {
	out := []rune(in)
	out[i] = replacementChar
	return string(out)
}
