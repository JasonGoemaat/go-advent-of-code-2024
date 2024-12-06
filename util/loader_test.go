package util

import (
	"testing"
)

func TestParseGroups(t *testing.T) {
	sample := `47|53
97|13

75,47,61,53,29
97,61,53,29,13
`
	expected := []string{`47|53
97|13`, `75,47,61,53,29
97,61,53,29,13`}

	groups := ParseGroups(sample)
	// if groups[0] != expected[0] {
	// if string(groups[0]) != string(expected[0]) {
	if groups[0] != expected[0] {
		t.Logf("FAILURE1: Expected %q\n", expected[0])
		t.Logf("          Got      %q\n", groups[0])
		t.Fail()
	}
	if groups[1] != expected[1] {
		t.Logf("FAILURE1: Expected %q\n", expected[1])
		t.Logf("          Got      %q\n", groups[1])
		t.Fail()
	}
}

func TestParseLines(t *testing.T) {
	sample1 := `47|53
97|13`
	sample2 := `47|53
97|13
`
	expected := []string{"47|53", "97|13"}
	lines := [][]string{ParseLines(sample1), ParseLines(sample2)}
	Expect(t, lines[0][0], expected[0], "%v")
	Expect(t, lines[1][0], expected[0], "%v")
	Expect(t, lines[0][1], expected[1], "%v")
	Expect(t, lines[1][1], expected[1], "%v")
}

func TestParseNumbers(t *testing.T) {
	sample := []string{"47|53", "97|13"}
	expected := [][]int{{47, 53}, {97, 13}}
	result := ParseNumbers(sample, "|")
	Expect(t, len(result), len(expected), "%v")
	for i := 0; i < len(result); i++ {
		if !ExpectSlices(t, result[i], expected[i]) {
			t.Logf("Problem with sample[%d]\n", i)
			t.FailNow()
		}
	}
}
