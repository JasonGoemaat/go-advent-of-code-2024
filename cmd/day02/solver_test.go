package day02

import (
	"slices"
	"testing"
)

// func TestRemove(t *testing.T) {
// 	t.Log("----- TestRemove() -----")
// 	s := []int{1, 2, 3, 4, 5, 6}
// 	t.Logf(fmt.Sprintf("s is: %v\n", s))
// 	b := slices.Delete(s, 2, 3)
// 	t.Logf("Called slices.Delete(s, 2, 3): %v\n", b)
// 	t.Logf("s is: %v\n", s)
// }

func TestRemove(t *testing.T) {
	t.Log("----- TestRemove() -----")
	s := []int{1, 2, 3, 4, 5, 6}
	t.Logf("s is: %v\n", s)
	a := slices.Clone(s)
	t.Logf("a is: %v\n", s)
	b := slices.Delete(a, 2, 3)
	t.Logf("Called slices.Delete(a, 2, 3): %v\n", b)
	t.Logf("s is: %v\n", s)
	t.Logf("a is: %v\n", a)
}
