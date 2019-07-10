package utility

import "testing"

func filterS(v interface{}) bool {
	return v.(string) == "b"
}

func filterI(v interface{}) bool {
	return v.(int) % 2 == 0
}

func TestArrayFilter(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	t.Log(ArrayFilter(s, filterS))
	i := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(ArrayFilter(i, filterI))
}
