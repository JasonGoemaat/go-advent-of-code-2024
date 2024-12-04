package util

import (
	"os"
	"regexp"
	"strconv"
)

func LoadString(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func LoadLines(filePath string) []string {
	all := LoadString(filePath)
	re := regexp.MustCompile("[\r]?[\n]")
	lines := re.Split(all, -1)
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
