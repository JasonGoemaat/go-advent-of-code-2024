package util

import (
	"cmp"
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
			t.Logf("Expected[%d]: %v", i, vb)
			t.Logf("--Actual[%d]: %v", i, va)
			t.Fail()
			return false
		}
	}
	return true
}
