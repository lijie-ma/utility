package utility

import "testing"

func TestSubstr(t *testing.T) {
	s := "abcdefghijk"

	t.Log(Substr(s, 1))
	t.Log(Substr(s, -1))
	t.Log(Substr(s, 0, 4))
	t.Log(Substr(s, 3, 100))
}
