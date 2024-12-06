# Day 02

## Boilerplate

I'll do this for each day to get ready to start coding...

First I created a `data/day02` directory with this file and downloaded
my input as `input.txt` and created the `sample.txt` from the puzzle
and `Instructions.md` for the text of the puzzle page at

    https://adventofcode.com/2024/day/2

Then in a terminal from the root I ran this command to create `cmd/day02.go`:

    cobra-cli add day02

And now for the work...

## Solving

In `cmd/day02.go` I cleared the long description and set the short description
to the title for the puzzle.

Looking at the [instructions](Instructions.md), it looks like the
`util.LoadNumbers()` can be used for the input, returning `[][]int`
where each element of the outer array is an array of numbers on that line.

The puzzle is to count the 'safe' lines, where a line is safe if
all the numbers on the line are all either increasing or decreasing (the
entire line must be in the same direction, but the line can be either)
by 1, 2, or 3.   Two identical numbers is unsafe.  If some numbers
are higher than the previous number and some numbers are lower than the
previous numbers, that is unsafe.   If the numbers are separated by 4 or
more, that is unsafe.

Writing pseudocode here, I'm thinking:

```go
func isSafeIncreasing(values []int) {
    for i, value := range values {
        if i > 0 && (value < values[i-1] || value > values[i-1]+3) {
            return false
        }
    }
    return true
}
```

Ok, implementing that I realized along with having to rename `solve()`
I might have other functions I don't want to pollute the `cmd` package
with, so I'll create a separate `day02` directory and package with
the file `cmd/day02/solver.go` and use that.   I can then import
package in `cmd/day02.go` and use `day02.solve()`.

> NOTE: The import actually looks like this and works because `go.mod`
defines all the packages defined in this module as having that as
a parent, and only the final path element is used in code to refer
to the package.

	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day02"

> NOTE: If you have multiple 'day02' packages from different go modules,
you can either just import one.  This is what I do when accessing my
`util` package, even though others come up in:

    import "github.com/JasonGoemaat/go-advent-of-code-2024/util"

> NOTE: If you want to access both, you can alias packages in the import:

    import myutil "github.com/JasonGoemaat/go-advent-of-code-2024/util"

> NOTE: Sometimes you may need to include a package but not reference any
of the code.   This is the case with Sql drivers for example, to make the
driver available for connecting and running commands in go.  In that case
you use the underscore (`_`) character which means it is unused.   Otherwise
go's auto-formatting will detect it is unused and remove the import:

    import _ "github.com/lib/pq"

> NOTE: The underscore (`_`) can also be used to ignore return values you don't
care about.   Fro example the loops using `range` return an index and value,
but a lot of times we don't care about the index:

```go
for _, value := range myArray {
    fmt.Println(value)
}
```

> NOTE: I haven't mentioned it yet, but you can declare variables in go in one
of two ways.   Either use the `var` keyword, or use `:=` if the type can be
infered and you want to assign it right away.  Also you can only use `:=`
if at least one of the variables on the left has not been declared yet.
If assigning only already-declared variables, use `=`:

    // using var
    var value
    value = myFunc()
    value = myFunc()

    // using :=
    value := myFunc()
    value = myFunc()

A note on private/public...   Functions and variables are private to a package
if they are lower-case and public if they are upper-case.   So I had to use
upper-case for `Solve()` so that the command could access it:

```go
func Solve(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	safeCount := 0
	for _, row := range numbers {
		if IsSafeIncreasing(row) || IsSafeDecreasing(row) {
			safeCount++
		}
	}
	return safeCount
}
```

I also decided to make the helper functions public.   I can see them being
used in future days and if they could be I will move them to the 'util'
package.  Also I missed the return type `bool` in my pseudo-code above:

```go
func IsSafeIncreasing(values []int) bool {
	for i, value := range values {
		if i > 0 && (value < values[i-1] || value > values[i-1]+3) {
			return false
		}
	}
	return true
}

func IsSafeDecreasing(values []int) bool {
	for i, value := range values {
		if i > 0 && (value > values[i-1] || value < values[i-1]-3) {
			return false
		}
	}
	return true
}
```

Well this returns the wrong result, it should be 2 instead of 3:

    PS C:\git\go\go-advent-of-code-2024> go run . day02
    day02 called
    sample: 3

So I'll add some output to see what I'm dealing with:

```go
func Solve(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	safeCount := 0
	for _, row := range numbers {
		if IsSafeIncreasing(row) || IsSafeDecreasing(row) {
			safeCount++
			fmt.Printf("Safe  : %v\n", row)
		} else {
			fmt.Printf("Unsafe: %v\n", row)
		}
	}
	return safeCount
}
```

    PS C:\git\go\go-advent-of-code-2024> go run . day02
    day02 called
    Safe  : [7 6 4 2 1]
    Unsafe: [1 2 7 8 9]
    Unsafe: [9 7 6 2 1]
    Unsafe: [1 3 2 4 5]
    Safe  : [8 6 4 4 1]
    Safe  : [1 3 6 7 9]
    sample: 3

Ah, it reports line 5 as safe, but it should be unsafe because there
are two 4's, the first compare needs to be `<=` or `>=` instead of 
`<` or `>`:

    if i > 0 && (value <= values[i-1] || value > values[i-1]+3) {

That worked!

## Part 2

This time we are working with the same data, but we can tolerate
one unsafe value, where if we removed that value from a row's numbers,
it would be safe with the other criteria.

My first thought would be maybe I can just process all of the elements
that are unsafe instead of returning 'false' immediately, but I don't think
it's that easy.  When it's unsafe, I still need to check if the next number
is safe compared to the previous number (as if it were removed).   I'll think
of some cases:

* `1 3 1 5` - Detects unsafe at second '1', '4' needs to be checked against 3, should report safe even though 5 is unsafe with respect to 1 because removing that 1 would make it safe
* `1 2 6 7` - Detects that 6 is unsafe, but still unsafe because `(7-2) > 3`, even though 7 is safe with respect to 6

How do I handle this?   I'm thinking of having a flag on the safe methods for
the new criteria for whether to allow for one unsafe value, then if I find one
I can create a new slice without that number and call it with that flag cleared
so no more can be unsafe.

Another thing I could do would be to split the slice into multiple slices
that are safe.  If a value is detected as unsafe, add a slice of everything
before that value to the list and start checking the remaining slice as if
it started fresh.   In this case for `1 2 6 7` I would end up with 
`[[1 2] [6 7]]`.  If there are more than two slices it would be a failure,
wouldn't it?  I would then check between the slices for the last of the first
and the second in the second, and the first in the second and the second to
last in the first.   Seems like a headache.

I would also have to keep checking the same direction.   For example
`[1 3 7 5 2]` would become `[[1 3] [7 5 2]]`.   The first slice is ascending
and the second is descending even if the 7 were removed, though each would
be safe on their own.   That's not hard, I would call the same ascending or
descending check from that function so the second slice would detect the 5
as bad.

Do I want to keep track of two flags?  `foundBad` and `lastBad`?  If
I hit a bad value I set both and continue.   If `lastBad` is set I
check against the current index minus two instead of the current index
minus one and clear the `lastBad` flag.   If the result would set the
`foundBad` flag and it's already set, return UNSAFE.

No, that won't work if the first element is the bad one.   For example
`[10 2 3 4 5]`, the `2` would set `lastBad` and the `3` would still be
checked against the `10` and report bad, even though removing the `10`
would leave a perfectly valid ascending list.

I could of course call the check for each index and treat it as if the
index didn't exist and or (`||`) all the results.   This seems like it
is very inefficient though.  Short-circuiting when testing removing
the first item would catch all the ones that are already safe though.

Since I only care about removing one number though, maybe I can keep
track of that.   When I'm on the `2` in `[10 2 3 4 5]`, I get a bad result.
So either the `10` or the `2` MUST be removed, but it is still safe if
EITHER is removed.   So I'll have to complicate my checks by testing
each index being compared as if either were removed.   The `2`'s 
previous value would be past the start of the list, so that is fine.
When checking the `3` it is fine if nothing is removed.

Ok, maybe I use a combination.  No need to check if everything is removed.
Maybe when I see the `2` is removed I set `[0 1]` as the possible bad indexes.
Then I have to recursively call the function with two new slices,
removing each of the possible indexes.   That doesn't sound too bad...

Looking up how to split the slice: https://stackoverflow.com/a/37335777/369792

Here's what I came up with:

```go
func IsSafeIncreasingLenient(values []int) bool {
	for i, value := range values {
		if i > 0 && (value >= values[i-1] || value < values[i-1]-3) {
			// try removing index `i` and `i-1` and don't be lenient
			if IsSafeIncreasing(remove(values, i-1)) {
				return true
			}
			if !IsSafeIncreasing(remove(values, i)) {
				return true
			}
		}
	}
	return true
}
```

Well that reports 6 instead of 4 for the sample, time to put the prints in to
see which are causing problems and figure out why.  I see the `remove()`
method is trash, it is over-writing the original slice somehow.  Yeah, stupid
stackoverflow.   Comments and other answers explain why it's trash.  Down-vote it.

I see the slices package has `slices.Delete(s, i, j)` which is also destructive,
but a built-in function.   It has `slices.Clone(s)` also which makes a copy. 
I'll use that.

Now's a good time to talk about testing, I'll write a test to make sure that works.
Testing in go is easy and I think is cool.   Create a new file that ends in `_test.go`
and create functions that start with `Test` and take an argument of type
`*testing.T` like `func TestRemove(t *testing.T)`.   Then run `go test . -v` in the
directory with the test.  The `-v` parameter tells it to output the logs even if
the tests succeed which is what I want for demonstration.   My updated function in
`cmd/day02/solver.go`:

```go
func remove(slice []int, i int) []int {
	var a = slices.Clone(slice)
	a = slices.Delete(a, i, i)
	fmt.Printf(" result: %v\n", a)
	return a
}
```

And the contents of `cmd/day02/solver_test.go`:

```go
package day02

import (
	"fmt"
	"slices"
	"testing"
)

func TestRemove(t *testing.T) {
	t.Log("----- TestRemove() -----")
	s := []int{1, 2, 3, 4, 5, 6}
	t.Logf("s is: %v\n", s)
	a := slices.Clone(s)
	t.Logf("a is: %v\n", s)
	b := slices.Delete(a, 2, 3)
	t.Logf("Called slices.Delete(a, 2, 3): %v\n", b)
	t.Logf("s is: %v\n", s)
	t.Logf("a is: %v\n", a)
}

```

    PS C:\git\go\go-advent-of-code-2024\cmd\day02> go test . -v
    === RUN   TestRemove
        solver_test.go:18: ----- TestRemove() -----
        solver_test.go:20: s is: [1 2 3 4 5 6]
        solver_test.go:22: a is: [1 2 3 4 5 6]
        solver_test.go:24: Called slices.Delete(a, 2, 3): [1 2 4 5 6]
        solver_test.go:25: s is: [1 2 3 4 5 6]
        solver_test.go:26: a is: [1 2 4 5 6 0]
    --- PASS: TestRemove (0.00s)
    PASS
    ok      github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day02        (cached)

That looks better.  I do have to clone it since `slices.Delete()` does modify
the slice you pass.

Ok, had to do some fiddling around, forgot or didn't think that the
second parameter is the index PAST the last one you want to remove.
for example `slices.Delete([1 2 3 4 5], 3, 4)` will remove the 4 at index
3 and not the 5 at index 4.

```go
func remove(slice []int, i int) []int {
	a := slices.Clone(slice)
	return slices.Delete(a, i, i+1)
}
```

Well, this is working for the samples but failing with the main input, trying to find out why...

```go
func Solve2(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	safeCount := 0
	for _, row := range numbers {
		if IsSafeIncreasingLenient(row) || IsSafeDecreasingLenient(row) {
			safeCount++
		}
	}
	return safeCount
}
```

The functions:

```go
func IsSafeIncreasingLenient(values []int) bool {
	for i, value := range values {
		if i > 0 && (value <= values[i-1] || value > values[i-1]-3) {
			// try removing index `i` and `i-1` and don't be lenient
			if IsSafeIncreasing(remove(values, i-1)) {
				return true
			}
			if IsSafeIncreasing(remove(values, i)) {
				return true
			}
			return false
		}
	}
	return true
}

func IsSafeDecreasingLenient(values []int) bool {
	for i, value := range values {
		if i > 0 && (value >= values[i-1] || value < values[i-1]-3) {
			// try removing index `i` and `i-1` and don't be lenient
			if IsSafeDecreasing(remove(values, i-1)) {
				return true
			}
			if IsSafeDecreasing(remove(values, i)) {
				return true
			}
			return false
		}
	}
	return true
}
```

Time for me to create some sample data for edge cases that should pass:

    1 2 3 4 10
    10 1 2 3 4
    1 2 10 3 4
    1 2 2 3 4
    1 1 2 3 4
    1 2 3 4 4

Ah, that only returns `3` when it should be all `6`.   That happens to be the
number of times the bad number isn't on the edge of an array.  I'll add
logging to the functions and see what's up.

Ok, I'm an idiot.   I had the comparison reversed on one I think.   I don't know
how that happened, my code above in this document looks right.

I created a function 'mylog' as a helper so I can disable my debug logging
statements by commenting out the printf:

```go
func mylog(format string, args ...interface{}) {
	// fmt.Printf(format, args...)
}
```

Now I think that would be better in the `util` package, and I could add a flag
using the `cobra` library I'm using for arguments.   Let's see...

I'll add a file `util/logging.go`.   I'll make a boolean for enabling and
capitalize it and the function so it can be used in other packages:

```go
package util

import "fmt"

var MyLogEnabled = false

func MyLog(format string, args ...interface{}) {
	if MyLogEnabled {
		fmt.Printf(format, args...)
	}
}
```

In `root.go` (so it is global) I add the flag in `init()`:

	rootCmd.PersistentFlags().BoolVarP(&util.MyLogEnabled, "log", "l", false, "enables util.MyLog")

And that's it...

    PS C:\git\go\go-advent-of-code-2024> go run . day02
    day02 called
    MyLogEnabled: false
    PS C:\git\go\go-advent-of-code-2024> go run . day02 -l
    day02 called
    MyLogEnabled: true
    PS C:\git\go\go-advent-of-code-2024> go run . day02 --log
    day02 called
    MyLogEnabled: true

All done with day 2!