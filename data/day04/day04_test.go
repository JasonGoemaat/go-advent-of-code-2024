package day04

import (
	"testing"
)

func TestCharacter(t *testing.T) {
	s := "XMASSXSXMASSS"
	a := s[0]

	// ok these are bytes
	if a == 'X' {
		t.Logf("looking at a which should be X, it IS!")
	} else {
		t.Logf("looking at a which should be X, it IS NOT!")
	}

	if a == 'M' {
		t.Logf("looking at a which should be X and testing for M, it IS!")
	} else {
		t.Logf("looking at a which should be X and testing for M, it IS NOT!")
	}
}
