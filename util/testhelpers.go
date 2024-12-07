package util

import (
	"cmp"
	"encoding/json"
	"fmt"
	"testing"
)

func Expect(t testing.TB, a, b interface{}, format string) {
	t.Helper()
	if a == b {
		return
	}
	t.Logf("Expected: "+format, b)
	t.Logf("-----Got: "+format, a)
	t.Fail()
}

func jsonOrError(a interface{}) string {
	s, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("%v (JSON ERROR: %v)", a, err)
	}
	return string(s)
}

func ExpectJson(t testing.TB, a, b interface{}) {
	t.Helper()
	if a == b {
		return
	}

	t.Logf("Expected: %s", jsonOrError(b))
	t.Logf("-----got: %s", jsonOrError(a))
	t.Fail()
}

func ExpectSlices[T cmp.Ordered](t testing.TB, a, b []T) bool {
	t.Helper()
	if len(a) != len(b) {
		t.Logf("Expected slice length: %d", len(b))
		t.Logf("--Actual slice length: %d", len(a))
		t.Fail()
		return false
	}
	for i, va := range a {
		vb := b[i]
		if cmp.Compare(va, vb) != 0 {
			t.Logf("Expected[%d]: %s", i, jsonOrError(b))
			t.Logf("--Actual[%d]: %s", i, jsonOrError(a))
			t.Fail()
			return false
		}
	}
	return true
}
