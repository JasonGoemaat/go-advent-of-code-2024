# Day 5: Print Queue

## Getting ready

First I'll create the command:

    cobra-cli add day06

Then I create a new `cmd/day06` directory and move `cmd/day06.go` there,
and add this `Work.md` and `Instructions.md`.   In `cmd/day06/day06.go`
I change the command to upper-case `Day06Cmd` and remove the `init()`
function, then add the command in `cmd/root.go`:

```go
rootCmd.AddCommand(day06.Day06Cmd)
```

And created an initial `cmd/day06/solve_day06.go`:

```go
package day06

func SolvePart1(content string) int {
    return 0
}

func SolvePart2(content string) int {
    return 0
}
```

And create the directory `cmd/day06/data` and create a `sample.txt` file from
the instructions and download the input to `input.txt`.

Then modify `cmd/day06/day06.go` to call them:

```go
// day06Cmd represents the day06 command
var Day06Cmd = &cobra.Command{
	Use:   "day06",
	Short: "Day 6: Guard Gallivant",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day06 called")
		content := util.GetContent()
		if content != "" {
			fmt.Println("Part 1:", SolvePart1(content))
			fmt.Println("Part 2:", SolvePart2(content))
			return
		}
		sample := util.LoadString("cmd/day06/data/sample.txt")
		input := util.LoadString("cmd/day06/data/input.txt")
		fmt.Println("Part 1 Sample:", SolvePart1(sample))
		fmt.Println("Part 1 Input :", SolvePart1(input))
		fmt.Println("Part 2 Sample:", SolvePart2(sample))
		fmt.Println("Part 2 Input :", SolvePart2(input))
	},
}
```

This is quite a bit of boilerplate, but it does make it nice to run.   I 
can run any day's solver from the command line using arguments and it operates
on the inputs or can take the input from `--file <name>` in the arguments or
standard input with `--stdin`.

I could probably make it work without the `day06.go` command file though.
I could create a function that would let me replace the line I add in
`cmd/root.go`:

```go
rootCmd.AddCommand("day06", "Day 6: Guard Gallivant", day06.SolvePart1, day06.SolvePart2)
```

## Part 1

I wanted to see how fast I could do this so I didn't write this as I was
coding.  One thing I did was setup directions again, turning right
represents moving one further in the array or back to the start after the
end using `%`:

```go
var up = []int{-1, 0}
var down = []int{1, 0}
var left = []int{0, -1}
var right = []int{0, 1}

// directions in order, start by going up
var directions = [][]int{up, right, down, left}
```

I use package variables for the map and counts, thinking I can use them in
other functions in the package:

```go
var rowCount = 0
var colCount = 0
var theMap []byte = nil
```

I have some helper functions I thought would be useful.   Some make use of
go's ability to return multiple values.  For instance `getPos` returns an
int representing the index into `theMap` as well as a boolean representing
whether the position is outside the map.   The position might be inside
the slice/array but wrap around from one row to the next.

```go
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
```

The solver starts by initializing globals

```go
func SolvePart1(content string) int {
	var lines = util.ParseLines(content)
	rowCount = len(lines)
	colCount = len(lines[0])
	theMap = make([]byte, rowCount*colCount)
	startRow := 0
	startCol := 0
	r := 0
	c := 0
```

And I loop through the lines to populate `theMap` with data and set the
starting position:

```go
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
```

Then I set some initial values and start looping while we're inside the bounds:

```go
    visitedCount := 0
	currentDirection := 0
	r = startRow
	c = startCol
	for r >= 0 && r < rowCount && c >= 0 && c < colCount {
```

First thing in the loop is getting the position and marking it as visited by
placing an `X` in the map and increasing the `visistedCount` we report as the
result if we haven't already visited it.

```go
pos, _ := getPos(r, c)
if theMap[pos] != 'X' {
    visitedCount++
    theMap[pos] = 'X'
}
```

Then I get the next position and break out of the loop if it will move outside
the map:

```go
nextR, nextC, outside := move(r, c, currentDirection)
if outside {
    break
}
```

If the next location contains a blocker, we turn right.  Otherwise we move
to the new position and continue the loop

```go
    if at(nextR, nextC) == '#' {
        currentDirection = (currentDirection + 1) % 4
    } else {
        r = nextR
        c = nextC
    }
}
return visitedCount
```

That wasn't too bad, took me about 40 minutes though.

## Part 2

This was rough because I over-thought it.   I spent forever trying to figure
out an efficient way to track everything all at once.   I wanted to avoid a
brute-force attack because I thought it would take too long.

I thought about just look for crossing a path, if turning right would put me
back on it then placing a blocker after the intersection would cause a loop.
that wouldn't work though if it crossed behind the start it wouldn't know that
turning to hit the start in the same direction again would cause a loop.

Then I thought maybe I could backup from the start in the oppostite direction.
That way it would catch any path that would lead back to the start and cause
a loop.  That turned out to be way to complicated.   I also had to take into
account when moving if I had placed a blocker next to the path to make it
turn.

Then I thought about using a graph to map all paths.   Or starting at the edges
and tracking how many steps for each direction to get to the end in a map.

All of these ran into problems.   Adding a blocker could alter the previous
paths and data I had already set.  Finally I decided just to brute-force it.

### Soltion Part 2

So I created some extra package variables.   `beenMoving` contains a bitmask
of directions we've been travelling at that spot (1=top, 2=right, 4=down,
8=left).   I have copies of that and the map because after placing an
obstacle I want to get back.   


```go
var mapCopy []byte = nil
var beenMoving []byte = nil
var beenMovingCopy []byte = nil
```
Then I have a map of blockers so that I only add a blocker to the same position.
once.

```go
var blockers map[int]bool = nil
```

Thinking about this more now, I don't really need it.  Since I'm going in order
and setting the map to 'X' I could just check the map.   I need to anyway
since I can't put a blocker at the start.

So the main loop is pretty similar with a few things added.  Here I set
`beenMoving` before turning.   I don't think this is strictly necessary
because the next iteration will set it for the direction you turn to
anyway (unless it also has a blocker I guess) but I already put it
before thinking about this and it can only help.

```go
nextPos, _ := getPos(nextR, nextC)
nextItem := theMap[nextPos]
if nextItem == '#' {
    beenMoving[pos] |= (1 << currentDirection)
    currentDirection = (currentDirection + 1) % 4
    continue
}
```

Then the special thing.   If the next map space is blank and we haven't
already tried placing a blocker there from running into it earlier,
we check if adding a blocker would cause a loop.  If so, we mark that
as a blocker and indrement our count.  Originally I was checking that it
wasn't the starting position before I settled on using the actual map for
checking if the space was open.

```go
// if theMap[nextPos] == '.' && !blockers[nextPos] && (nextR != startRow || nextC != startCol) {
if theMap[nextPos] == '.' && !blockers[nextPos] {
    if hasLoop(r, c, currentDirection) {
        blockers[nextPos] = true
        blockerCount++
    }
}
```

The `hasLoop()` function makes copies of the global data and basically does
the original solution loop but checks for a 'loop' in the path:

```go
func hasLoop(r, c, dir int) bool {
	// save var theMap []byte = nil
	copy(mapCopy, theMap)
	copy(beenMovingCopy, beenMoving)
	nr, nc, _ := move(r, c, dir)
	np, _ := getPos(nr, nc)
	mapCopy[np] = 'O'
```

Checking for a loop:

```go
if (beenMovingCopy[pos] & (1 << dir)) > 0 {
    // found our loop!
    return true
}
```
