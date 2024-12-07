package day07

import (
	"regexp"
	"testing"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
)

func TestRegexFindAllStringSubmatch(t *testing.T) {
	rxLine := regexp.MustCompile("(\\d+): ([^\n\r]+)")
	sample := `3267: 81 40 27`
	expected := []string{sample, "3267", "81 40 27"}
	result := rxLine.FindAllStringSubmatch(sample, -1)
	t.Log("sample:", sample)
	t.Log("expected:", expected)
	t.Log("result:", result[0]) //
	util.ExpectSlices(t, result[0], expected)
}

// type mytest testing.TB

// func (t mytest) SayHi() {
// 	fmt.Println("Hello, world!")
// }

func TestRegexFindAllStringSubmatchMultiple(t *testing.T) {
	rxLine := regexp.MustCompile("(\\d+): ([^\n\r]+)")
	sample := `190: 10 19
3267: 81 40 27`
	expected := [][]string{{"190: 10 19", "190", "10 19"}, {"3267: 81 40 27", "3267", "81 40 27"}}
	result := rxLine.FindAllStringSubmatch(sample, -1)
	util.Expect(t, len(result), len(expected), "%d")
	for i, _ := range result {
		util.ExpectSlices(t, result[i], expected[i])
	}
}

func TestParseContent(t *testing.T) {
	sample := `190: 10 19
3267: 81 40 27`
	parsed := parseContent(sample)
	util.ExpectJson(t, len(parsed), 2)
	util.ExpectJson(t, parsed[0].Goal, int64(190))
	util.ExpectSlices(t, parsed[0].Values, []int{10, 19})
	util.ExpectJson(t, parsed[1].Goal, int64(3267))
	util.ExpectSlices(t, parsed[1].Values, []int{81, 40, 27})
}

func TestPart1Sample(t *testing.T) {
	pl := puzzleLine{Goal: 3267, Values: []int{81, 40, 27}}
	result := pl.doesWork(0, 0)
	util.ExpectJson(t, result, true)
}
