package day06

import "github.com/JasonGoemaat/go-advent-of-code-2024/util"

var up = []int{-1, 0}
var down = []int{1, 0}
var left = []int{0, -1}
var right = []int{0, 1}

// directions in order, start by going up
var directions = [][]int{up, right, down, left}

var rowCount = 0
var colCount = 0
var theMap []byte = nil

func at(r, c int) byte {
	pos, outside := getPos(r, c)
	if outside {
		util.MyLog("OUTSIDE!")
		return 0
	}
	return theMap[pos]
}

func getPos(r, c int) (int, bool) {
	if r < 0 || r >= rowCount || c < 0 || c >= colCount {
		return -1, true
	}
	return r*colCount + c, false
}

func move(r, c, dir int) (int, int, bool) {
	newR := r + directions[dir][0]
	newC := c + directions[dir][1]
	_, outside := getPos(newR, newC)
	return newR, newC, outside
}

func SolvePart1(content string) int {
	var lines = util.ParseLines(content)
	rowCount = len(lines)
	colCount = len(lines[0])
	theMap = make([]byte, rowCount*colCount)
	startRow := 0
	startCol := 0
	r := 0
	c := 0
	for r = 0; r < rowCount; r++ {
		for c = 0; c < colCount; c++ {
			pos, _ := getPos(r, c)
			theMap[pos] = lines[r][c]
			if theMap[pos] == '^' {
				startRow = r
				startCol = c
			}
		}
	}

	visitedCount := 0
	currentDirection := 0
	r = startRow
	c = startCol
	for r >= 0 && r < rowCount && c >= 0 && c < colCount {
		pos, isOutside := getPos(r, c)
		if isOutside {
			break
		}
		if theMap[pos] != 'X' {
			visitedCount++
			theMap[pos] = 'X'
		}
		nextR, nextC, outside := move(r, c, currentDirection)
		if outside {
			break
		}
		if at(nextR, nextC) == '#' {
			currentDirection = (currentDirection + 1) % 4
			continue
		}
		r = nextR
		c = nextC
	}
	return visitedCount
}

func SolvePart2(content string) int {
	var lines = util.ParseLines(content)
	rowCount = len(lines)
	colCount = len(lines[0])
	theMap = make([]byte, rowCount*colCount)
	startRow := 0
	startCol := 0
	r := 0
	c := 0
	for r = 0; r < rowCount; r++ {
		for c = 0; c < colCount; c++ {
			pos, _ := getPos(r, c)
			theMap[pos] = lines[r][c]
			if theMap[pos] == '^' {
				startRow = r
				startCol = c
			}
		}
	}

	visitedCount := 0
	currentDirection := 0
	r = startRow
	c = startCol
	for r >= 0 && r < rowCount && c >= 0 && c < colCount {
		pos, isOutside := getPos(r, c)
		if isOutside {
			break
		}
		if theMap[pos] != 'X' {
			visitedCount++
			theMap[pos] = 'X'
		}
		nextR, nextC, outside := move(r, c, currentDirection)
		if outside {
			break
		}
		if at(nextR, nextC) == '#' {
			currentDirection = (currentDirection + 1) % 4
			continue
		}
		r = nextR
		c = nextC
	}
	return visitedCount
}
