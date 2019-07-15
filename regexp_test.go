package utility

import (
	"fmt"
	"testing"
)

func TestRegReplaceCallback(t *testing.T) {
	repl := func(v string) string {
		ss := RegMatchAll(v, `\((\d*-\d*-\d*)\)([\x{4e00}-\x{9fa5}]?)`)
		days := (Date2Unix(`2019-07-15`) - Date2Unix(ss[0][1])) / 86400

		return fmt.Sprintf(`%d天`, days)
	}
	s := `2019年7月15日来北京4(2016-04-22)年多了`
	r := RegReplaceCallback(s, repl, `\d\((\d*-\d*-\d*)\)([\x{4e00}-\x{9fa5}]?)`)
	t.Log(r)
}
