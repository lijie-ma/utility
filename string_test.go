package utility

import "testing"

func TestSubstr(t *testing.T) {
	s := "abcdefghijk"

	t.Log(Substr(s, 1))
	t.Log(Substr(s, -1))
	t.Log(Substr(s, 0, 4))
	t.Log(Substr(s, 3, 100))
}

func TestHttpBuildQuery(t *testing.T) {
	s := map[string]interface{}{
		"a":  "aa",
		"as": []string{"a1", "a2"},
		"b":  10,
		"bs": []int{11, 12},
		"c":  1.1,
		"cs": []float64{1.2, 1.3},
	}

	t.Log(HttpBuildQuery(s))
}
