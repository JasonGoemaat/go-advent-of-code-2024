# Day 4: Ceres Search

Starting up I create files and download data into `data/day04`, and run this:

    cobra-cli add day04

Got one of these arrays of things you have to parse directionally.  This is a
a block of characters and we need to find how many times 'XMAS' appears in any
way and any direction in a straight line.

Looking [at strings](https://go.dev/blog/strings) in GO, they are a little
different than I'm used to.   Indexing them gives bytes.   Characters are
called 'runes'.   A rune is an int32.   Strings are bytes underneath and
UTF8.   You can use `%+q` to get the 'quoted string', for example a unicode
character will display what you would use in source code to make that
character including double-quotes, i.e. `"\u2318"`.

To check it out I created [day04_test.go] in this directory:

```go
package day04

import (
	"testing"
)

func TestCharacter(t *testing.T) {
	s := "XMASSXSXMASSS"
	a := s[0] // indexing uses the underlying []byte slice and returns a byte

	if a == 'X' { // this seems to work for comparing
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
```

Then changing to the directory and running `go test . -v` look like we can use
that comparison:

    PS C:\git\go\go-advent-of-code-2024\data\day04> go test . -v
    === RUN   TestCharacter
        day04_test.go:16: looking at a which should be X, it IS!
        day04_test.go:24: looking at a which should be X and testing for M, it IS NOT!
    --- PASS: TestCharacter (0.00s)
    PASS
    ok      github.com/JasonGoemaat/go-advent-of-code-2024/data/day04       0.295s

So I'm thinking I'll use `util.LoadLines()` to get the input as an array of
strings.   The length of the array will be height and `len(lines[0])` will
be the width, they should all have the same width.

Then I'll loop through the rows, and loop through the columns for each.
For each row,col I can check in each direction.   To define directions,
maybe I define a direction as how it affects each coordinate:

    [[0 1] [1 1] [1 0] [1 -1] [0 -1] [-1 -1] [-1 0] [-1 1]]

Then I can loop through the directions.   I can multiply the length by the
direction value and test if they are inside the bounds.   Then loop 
starting with the current row and col and adjusting by direction amounts
for each and maek sure the characters match what we want.   For each match
I will increase the count.

This is what I came up with.   I had an error the first time due to index out
of bounds and added logging.   Seeing it was happening with the 'X' at end
of line 4, I quickly traced it to a copy error where I was using 'r' in place
of 'c' in this line:

	cLimit := c + direction[1]*3  // original copy using 'r' in place of 'c' caused error

Here's what I got:

```go
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
```

Part 1 done!

## Part 2

This looks easy.   We're looking for 'MAS' in the shape of an 'X'.   We're
looking for something like this:

    M.S
    .A.
    M.S

Not sure if there's a super-efficient way to define the criteria in code.
Thinking about it, I really only need to check the two diagonals.  If both
have 'M' on one end and 'S' on the other, I'm good.   I can have an outer
loop that ignores the edges because I have to find the 'A' in the center.
Seems easy.   Let's try some code here:

```go
for r := 1; r+1 < rowCount; r++ {
    for c := 1; c+1 < colCount; c++ {
        if lines[r][c] == 'A' && findCross(r, c) {
            total++
        }
    }
}
```

and

```go
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
```

Time to check it.   I like writing it out here first, I think it's helping me
practice without having all the benefits of code completion in the editor.

Worked like a charm, noice!   Done with day 4...


