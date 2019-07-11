package utility

func Substr(s string, start int, length ...int) string {
	fetchLen := len(s)
	if len(length) > 0 {
		fetchLen = length[0]
	}
	runeSlice := []rune(s)
	lenSlice := len(runeSlice)

	if start < 0 {
		start = lenSlice + start
	}
	if start > lenSlice {
		start = lenSlice
	}
	end := start + fetchLen
	if end > lenSlice {
		end = lenSlice
	}
	if fetchLen < 0 {
		end = lenSlice + fetchLen
	}
	if start > end {
		start, end = end, start
	}
	return string(runeSlice[start:end])
}
