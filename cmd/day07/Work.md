# Day 7: Bridge Repair

Getting tired or repeating instructions, just going to put some notes in here on
the puzzles from now on...

## Part 1

This doesn't seem too bad.   We have inputs like this:

    3267: 81 40 27

The left of the colon represents a total value.   We try every combination of 
inserting either `+` or `*` between the numbers on the right and evaluate
left to right (no order of operations).  There are actually two ways to make
that sample work:

* `81 + 40 * 27` = 3267
* `81 * 40 + 27` = 3267

We find the lines where it's possible for one or more combinations to work
and add their totals together (3267 in this case even though there are
two ways to match).

I doubt optimization is required, but:

* If we get too large, quit early.   There's no way to lower the total

Need to parse, I'm thinking two regex.   One for the colon and one for
the numbers like I use in `util.ParseNumbers()`

```go
rxLine := regex.MustCompile("(\\d+): ([^\n\r]+)")
rxNums := regex.MustCompile("(\\d+)")
parts := rxLine.MatchAll()
```

Ignore that, package name is `regexp`.   Had to work out what to do
in the test, in `util.ParseNumbers` I was using split.   Look at this:

```go
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
```

So the result of `FindAllStringSubmatch` is a slice of slices.  The first
element is the entire matched string.   Then there is one element for each
group.   So I think I can just call this on the entire `content` without
parsing lines.   Let me add something...

Ok, couldn't tell a lot with the go `%v` format for arrays when they are 
strings that contain spaces.   There's a `json` package though that is
easy to use.   This converts to json, and falls back to go's `%v` along
with the error message if that fails (don't know why it would, cycles maybe?)

```go
func jsonOrError(a interface{}) string {
	s, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("%v (JSON ERROR: %v)", a, err)
	}
	return string(s)
}
```

Going to check out creating a struct...

```go
type puzzleLine struct {
	Goal   int
	Values []int
}
```

And to parse content:

```go
func parseContent(content string) []puzzleLine {
	rxLine := regexp.MustCompile("(\\d+): ([^\n\r]+)")
	rxSpace := regexp.MustCompile("[ ]+")
	// [][]string{{"190: 10 19", "190", "10 19"}, {"3267: 81 40 27", "3267", "81 40 27"}}
	matches := rxLine.FindAllStringSubmatch(content, -1)
	parsed := make([]puzzleLine, len(matches))
	for i, match := range matches {
		goal, _ := strconv.Atoi(match[1])
		parts := rxSpace.Split(match[2], -1)
		values := make([]int, len(parts))
		for j, s := range parts {
			values[j], _ = strconv.Atoi(s)
		}
		parsed[i] = puzzleLine{goal, values}
	}
	return parsed
}
```

Writing a test to see if that worked...

```go
func TestParseContent(t *testing.T) {
	sample := `190: 10 19
3267: 81 40 27`
	parsed := parseContent(sample)
	util.ExpectJson(t, len(parsed), 2)
	util.ExpectJson(t, parsed[0].Goal, 190)
	util.ExpectSlices(t, parsed[0].Values, []int{10, 19})
	util.ExpectJson(t, parsed[1].Goal, 3267)
	util.ExpectSlices(t, parsed[0].Values, []int{81, 40, 27})
}
```

Can you spot the error?  In the last line I re-use `parsed[0]` instead of
checking `parsed[1]`.  FYI I found a setting in vscode, search for `Go: Test`
in settings and look for test flags.   Select 'edit in settings.json' or
whatever the option is and add "-v", then when you do 'run test' or 'debug test'
from just that test in vscode it will show output that would be from passing
tests also.  Doesn't really help here, but I do that from the command line
pretty often to see `t.Log` output if I'm figuring something out.

    "go.testFlags": [
        "-v"
    ],

Other than that error, it looks good, time to work on the actual solution after 90
minutes :)

So go has an interesting way to add 'methods' to a struct/interface, It reminds
me of extension methods in C#.  You declare a func with the type you want to
put the method on and a name for the instance of that type in the method before
the method name.   Here's what I got without the actual solving logic:

```go
func (pl puzzleLine) doesWork() bool {
	return false
}

func SolvePart1(content string) int {
	total := 0
	parsed := parseContent(content)
	for _, line := range parsed {
		if line.doesWork() {
			total += line.Goal
		}
	}
	return 0
}
```

Ok, now the logic at 12:36 am...  Ah, good thing I checked the input.  Some
of the numbers are larger than a 32 bit int would fit.   It looks like the
largest I have is 15 digits, so int64 should work and I don't need some
big number library.  It looks like int64 * int32 will give int64 so no
need to rewrite the struct, and tests for the values, just need to
update the goal.  Well, that part at least I guess...

Reading the pop-up helper for `strconv.Atoi` it mentions it is an alias
for a call to `strconf.ParseInt(s, 10, 0)`.   So I can parse as 64 bit
decimal like so:

```go
goal, _ := strconv.ParseInt(match[1], 10, 64)
```

Have to cast my numbers to int64 in my tests too:

```go
util.ExpectJson(t, parsed[0].Goal, int64(190))
```

This is what I came up with, see actual code for more comments and showing
a mistake I made by trying to optimize when not needed...

```go
func (pl puzzleLine) doesWork(sum int64, index int) bool {
	// if index has passed the end, return true only if matched exactly
	if index >= len(pl.Values) {
		return sum == pl.Goal
	}

	// only add first number, but try multiply with others
	if index > 0 {
		multiplied := sum * int64(pl.Values[index])
		return pl.doesWork(multiplied, index+1)
	}

	// try to add and call recursively
	added := sum + int64(pl.Values[index])
	return pl.doesWork(added, index+1)
}
```

This gives me '190' for my test, so only the first line of the file
is working.   Maybe it's a parsing issue, but my tests should have
found that.   Let me write a test for the second sample and use
vscode to debug the test and step into my solver:

```go
func TestPart1Sample(t *testing.T) {
	pl := puzzleLine{Goal: 3267, Values: []int{81, 40, 27}}
	result := pl.doesWork(0, 0)
	util.ExpectJson(t, result, true)
}
```

Ah, the multiply used to be last and I rearranged the code and was just
returning here, now it will continue and try the add with this change:

```go
// was return pl.doesWork(multiplied, index+1) 
if pl.doesWork(multiplied, index+1) {
    return true
}
```

That worked, done with part 1!  Learned some about testing and structs.
I tried to see about adding my 'Expect' methods to the `testing.T` type,
but it says their's an error passing a lock by value or something.  Trying
to make the 'extension' type a pointer gave an error, it didn't like that.

## Part 2

Ok, time for me to guess how they can make this harder.  First guess is just
adding `-` and `/` to the mix.    Second guess is that they DO want to use
normal math order of operations, so you'd have to do any "*"s in your
equation before doing the "+"s.  All I can think of in a minute, let's see...

Ah, fiendish.   Add some string manipulation and parsing into the mix...
I don't see how this will be very difficult though.   Certainly seems easier
than what I was thinking.

> The concatenation operator (||) combines the digits from its left and right
inputs into a single number. For example, 12 || 345 would become 12345. All
operators are still evaluated left-to-right.

While adding a `checkDone2` method to the `puzzleLine` struct, I though for
a sec that might make my checks overflow even if the results didn't, but
that would about the same for multiply by all 9s.

I tried adding this for my `checkDone2` where the multiply check is:

```go
concatenated, _ := strconv.ParseInt(fmt.Sprintf("%d%d", sum, pl.Values[index]), 10, 64)
if pl.doesWork2(concatenated, index+1) {
    return true
}
```

Initially I was getting the same results as part 1, but I realized I forgot to
change all the calls to recursivly call `doesWork2()` and some were sill calling
`doesWork()`.   Fixing those it worked!

Done with Day 7!

2.5 hours, but a lot of messing around with testing and writing this...

