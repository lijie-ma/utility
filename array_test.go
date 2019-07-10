package utility

import "testing"

func filterS(v interface{}) bool {
	return v.(string) == "b"
}

func filterI(v interface{}) bool {
	return v.(int) % 2 == 0
}

func TestSliceFilter(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	t.Log(SliceFilter(s, filterS))
	i := []int{1, 2, 3, 4, 5, 6, 7}
	cc := SliceFilter(i, filterI)
	t.Log(cc.([]int))
	t.Log(cc)
}

func TestSlicePop(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	SlicePop(&s)
	t.Log(s)
}

func TestSliceIntersect(t *testing.T) {
	s1 := []string{"a", "b", "c", "d", "e"}
	s2 := []string{"f", "b", "c", "d", "e"}
	ti := SliceIntersect(s1, s2)
	t.Log(ti.([]string))
	t.Log(ti)
}

func TestSliceDiff(t *testing.T) {
	s1 := []string{"a", "b", "c", "d", "e"}
	s2 := []string{"f", "b", "c", "d", "e"}
	ti := SliceDiff(s1, s2)
	t.Log(ti.([]string))
	t.Log(ti)
}