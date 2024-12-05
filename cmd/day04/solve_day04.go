package day04

import "github.com/JasonGoemaat/go-advent-of-code-2024/util"

var directions = [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
var lines []string
var rowCount = 0
var colCount = 0

func findDirection(r, c int, direction [2]int) bool {
	rLimit := r + direction[0]*3 // we'll move 3 in this direction
	cLimit := c + direction[1]*3 // and 3 in this direction, original copy using 'r' in place of 'c' caused error
	if rLimit < 0 || rLimit >= rowCount || cLimit < 0 || cLimit >= colCount {
		return false
	}
	chars := "MAS"
	for i := 0; i < 3; i++ {
		r += direction[0]
		c += direction[1]
		if lines[r][c] != chars[i] {
			return false
		}
	}
	return true
}

func find(r, c int) int {
	count := 0
	for _, direction := range directions {
		if findDirection(r, c, direction) {
			count++
		}
	}
	return count
}

func SolveDay04(filePath string) int {
	lines = util.LoadLines(filePath)
	rowCount = len(lines)
	colCount = len(lines[0])
	total := 0
	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			// util.MyLog("(%d, %d) is %c\n", r, c, lines[r][c])
			// optimization, no need to check directions if char we're on isn't 'X'
			if lines[r][c] == 'X' {
				total += find(r, c)
			}
		}
	}
	return total
}

func findCross(r, c int) bool {
	isMatch := func(a, b byte) bool {
		return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
	}
	// to make it more readable
	topLeft := lines[r-1][c-1]
	topRight := lines[r-1][c+1]
	bottomLeft := lines[r+1][c-1]
	bottomRight := lines[r+1][c+1]
	return isMatch(topLeft, bottomRight) && isMatch(topRight, bottomLeft)
}

func SolveDay04Part2(filePath string) int {
	lines = util.LoadLines(filePath)
	rowCount = len(lines)
	colCount = len(lines[0])
	total := 0
	for r := 1; r+1 < rowCount; r++ {
		for c := 1; c+1 < colCount; c++ {
			if lines[r][c] == 'A' && findCross(r, c) {
				total++
			}
		}
	}
	return total
}
