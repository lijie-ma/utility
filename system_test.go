package utility

import "testing"

func testDebugBacktrace() []interface{} {
	return DebugBacktrace(2)
}
func TestDebugBacktrace(t *testing.T) {
	t.Log(testDebugBacktrace())
}
