package util

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LoadString(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func LoadLines(filePath string) []string {
	return ParseLines(LoadString(filePath))
}

func ParseLines(content string) []string {
	re := regexp.MustCompile("[\r]?[\n]")
	lines := re.Split(content, -1)
	if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[0 : len(lines)-1]
	}
	return lines
}

func LoadNumbers(filePath string) [][]int {
	x := [][]int{}
	re := regexp.MustCompile("[ ]+") // multiple spaces
	for _, line := range LoadLines(filePath) {
		y := []int{}
		// parts := strings.Split(line, " ")
		parts := re.Split(line, -1)
		for _, part := range parts {
			n, _ := strconv.Atoi(part)
			y = append(y, n)
		}
		x = append(x, y)
	}
	return x
}

var reDoubleCRLF = regexp.MustCompile("[\r]?[\n][\r]?[\n]")

func ParseGroups(content string) []string {
	// remove final crlf if there is one, files usually
	// end in crlf and would have a blank line.   Splitting
	// on two here means other groups won't have them
	if content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}
	if content[len(content)-1] == '\r' {
		content = content[:len(content)-1]
	}
	groups := reDoubleCRLF.Split(content, -1)

	// omit last group if blank
	if len(groups[len(groups)-1]) == 0 {
		groups = groups[:len(groups)-1]
	}

	return groups
}

func ParseNumbers(lines []string, separator string) [][]int {
	var results = make([][]int, len(lines))
	for j, s := range lines {
		parts := strings.Split(s, separator)
		result := make([]int, len(parts))
		for i, part := range parts {
			result[i], _ = strconv.Atoi(part)
		}
		results[j] = result
	}
	return results
}
