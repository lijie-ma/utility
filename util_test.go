package utility

import "testing"

func TestIsNumeric(t *testing.T) {
	s := "5"

	t.Log(IsNumeric(s))
	t.Log(IsNumeric("15"))
	t.Log(IsNumeric("12.4"))
	t.Log(IsNumeric("12.2.3"))
}
