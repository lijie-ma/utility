package utility

func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	runeSlice := []rune(str)
	lenSlice := len(runeSlice)

	if start < 0 {
		start = lenSlice + start
	}
	if start > lenSlice {
		start = lenSlice
	}
	end := start + length
	if end > lenSlice {
		end = lenSlice
	}
	if length < 0 {
		end = lenSlice + length
	}
	if start > end {
		start, end = end, start
	}
	return string(runeSlice[start:end])
}
