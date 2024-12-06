package day05

import "testing"

func TestMissingMapBool(t *testing.T) {
	bs := map[int]bool{}
	bs[0] = false
	bs[1] = true
	t.Logf("bs[0] == %v (false)", bs[0])
	t.Logf("bs[1] == %v (true)", bs[1])
	t.Logf("bs[2] == %v (?)", bs[2])
}

func TestMissingMapInt(t *testing.T) {
	bs := map[int]int{}
	bs[0] = 0
	bs[1] = 1
	t.Logf("bs[0] == %v (0)", bs[0])
	t.Logf("bs[1] == %v (1)", bs[1])
	t.Logf("bs[2] == %v (?)", bs[2])
}
