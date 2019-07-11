package utility

import (
	"testing"
)

func TestMapKeys(t *testing.T) {
	s := map[string]interface{}{
		"a": "b",
		"u": "b",
	}
	c, e := MapKeys(s)
	t.Log(c.([]string), e)

	cf, ee := MapKeys(s, func(key interface{}) bool { return key.(string) == "a" })
	t.Log(cf.([]string), ee)
}
