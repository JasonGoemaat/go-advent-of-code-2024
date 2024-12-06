# Day 5: Print Queue

Starting up I create files and download data into `data/day05`, and run this from the root:

    cobra-cli add day05

And udpate the generated command in `cmd/day05.go`:

```go
var day05Cmd = &cobra.Command{
	Use:   "day05",
	Short: "Day 5: Print Queue",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day05 called")
		fmt.Println("sample part1:", day05.SolveDay05("data/day05/sample.txt"))
		// fmt.Println("input part1 :", day05.SolveDay05("data/day05/input.txt"))
		// fmt.Println("sample part1:", day05.SolveDay05Part2("data/day05/sample.txt"))
		// fmt.Println("input part1 :", day05.SolveDay05Part2("data/day05/input.txt"))
	},
}
```

And my solver `cmd/day05/solve_day05.go`

```go
package day05

import "github.com/JasonGoemaat/go-advent-of-code-2024/util"

func SolveDay05(filePath string) int {
	_ = util.LoadLines(filePath)
	total := 0
	return total
}

func SolveDay05Part2(filePath string) int {
	_ = util.LoadLines(filePath)
	total := 0
	return total
}
```

## Parsing/Testing

This time it looks like we'll need a new loader, the file comes in two
sections separated by a blank line (example, not the full sample):

    47|53
    97|13

    75,47,61,53,29
    97,61,53,29,13

Seems like regex could split the sections with `\n\n`, or maybe just
`strings.Split()`. But windows files will have `\r\n\r\n` which I accounted
for in my original `util.LoadLines()`

I think returning a slice containing groups of lines would work well.
`util.LoadLines()` ignores blank lines so can't use it directly.
I'll create a new function.   Copying some regex parts from
`util.LoadLines()`.   Thinking about it, I may actually want to
change it from taking the path to a file and make it just take a string.
We have `util.LoadString(filePath)` to load the contenst of a file given
the path, and currently the other methods call it themselves.   If they
just take the string and the caller has to actually load the file
that might be better.

For one thing, the command only has to load 'sample.txt' and 'input.txt'
once and can pass the same strings to part 1 and part 2.   For another
think we could write tests for our solvers and pass them strings.

That sounds good, but I don't want to mess up my existing functions.
Great, the name 'Load' doesn't make as much sense anyway, so I'll
take the parsing logic out into new functions called 'ParseXXX'
and change the 'LoadXXX' functions to call those after 'LoadString'.
Sample:

```go
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
```

Now that I think about it, that works doubly-well.  For this puzzle I can just
split the file by empty lines and call parse methods on each section.
Actually here it's pretty simple regex, but I do want to return int arrays,
and based on different separators.  Thinking...

```go
func ParseGroups(content string) [] string {
    re := regexp.MustCompile("[\r]?[\n][\r]?[\n]")
    groups := re.split(content, -1)

// omit last line if blank
    if len(groups[len(groups)-1]) = 0 {
        return groups[:len(groups)-1]
    }
    return groups
}
```

Thinking of conventions, I changed `Loader.go` -> `loader.go`:

    git mv Loader.go loader.go

> NOTE: The language server kept the information under `Loader.go` so I was
getting duplicate definition squiglies.   In vscode I did CTRL+SHIFT+P
and selected `GO: Restart Language Server` and they went away.

And created `util/loader_test.go`:

```go
package util

import "testing"

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
	if groups[0] != expected[0] {
		t.Logf("FAILURE: Expected %q but got %q", expected[0], groups[0])
		t.Fail()
	}
	if groups[1] != expected[1] {
		t.Logf("FAILURE: Expected %q but got %q", expected[0], groups[0])
		t.Fail()
	}
}
```

That fails, but it looks like they're right, what gives?

    PS C:\git\go\go-advent-of-code-2024\util> go test .
    --- FAIL: TestParseGroups (0.00s)
        loader_test.go:22: FAILURE: Expected "47|53\n97|13" but got "47|53\n97|13"
    FAIL
    FAIL    github.com/JasonGoemaat/go-advent-of-code-2024/util     0.223s
    FAIL

I thought maybe it was treating them as `[]byte`, but it seems like you can
do comparisons of different strings like this just fine.   This also fails:

```go
if string(groups[0]) != string(expected[0]) {
```

I'll try `strings.Compare()`, but it gives the same error:

	if strings.Compare(groups[0], expected[0]) != 0 {

The `%q` should print quoted strings, I'll try to use the `%v`
which handles any type.  But then I realized the error was coming
from the wrong line to be the first comparison.   It's actually
the second string that's different, but I output the first ones.
When I fix that I see the actual has an extra `\n` at the end.

Cleaning up that part a bit:

```go
if groups[1] != expected[1] {
    t.Logf("FAILURE1: Expected %q\n", expected[1])
    t.Logf("          Got      %q\n", groups[1])
    t.Fail()
}
```

The problem is much easier to see:

    PS C:\git\go\go-advent-of-code-2024\util> go test .
    --- FAIL: TestParseGroups (0.00s)
        loader_test.go:28: FAILURE1: Expected "75,47,61,53,29\n97,61,53,29,13"
        loader_test.go:29:           Got      "75,47,61,53,29\n97,61,53,29,13\n"
    FAIL
    FAIL    github.com/JasonGoemaat/go-advent-of-code-2024/util     0.320s
    FAIL

I changed `ParseGroups()` to remove the final `\n` or `\r\n` if they exist.
I'm glad I wrote the test so I now know exactly what I'm getting.

There seems like a lot of code for the test, so I did some research and wrote
a helper.  The Go testing package doesn't have `assert` or `expect` built-in.
For helper functions, make the first argument of type `testing.TB` and pass in
the `t *testing.T` argument of the test.   Then call `t.Helper()` to mark the
function as a helper.   This will make the testing framework ignore the
function when computing line numbers.

```go
func Expect(t testing.TB, a, b interface{}, format string) {
	t.Helper()
	if a == b {
		return
	}
    if format == nil {
        format = "%v"
    } 
	t.Logf("Expected: "+format, b)
	t.Logf("     Got: "+format, a)
	t.Fail()
}
```

So for example I call `t.Fail()` in the `Expect()` function, but because it
is marked as a helper the failure show the line that called it in the
test function:

    PS C:\git\go\go-advent-of-code-2024\util> go test .
    --- FAIL: TestParseLines (0.00s)
        loader_test.go:53: Expected: 97|133
        loader_test.go:53:      Got: 97|13
        loader_test.go:54: Expected: 97|133
        loader_test.go:54:      Got: 97|13
    FAIL
    FAIL    github.com/JasonGoemaat/go-advent-of-code-2024/util     0.285s
    FAIL

That's pretty nifty.   Assert would be even easier:

```go
func Assert(t testing.TB, result bool) {
	t.Helper()
	if !result {
        t.Fail()
    }
}
```

Ok, since Go 1.18+ you can have generics, makes for a nice ExpectSlices()
helper:

```go
func ExpectSlices[T cmp.Ordered](t testing.TB, a, b []T) bool {
	t.Helper()
	if len(a) != len(b) {
		t.Logf("Expected slice length: %d", len(b))
		t.Logf("--Actual slice length: %d", len(a))
		t.Fail()
		return false
	}
	for i, va := range a {
		vb := b[i]
		if cmp.Compare(va, vb) != 0 {
			t.Logf("Expected[%d]: %v", i, vb)
			t.Logf("--Actual[%d]: %v", i, va)
			t.Fail()
			return false
		}
	}
	return true
}
```

Makes my `ParseNumbers()` function easy to test:

```go
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
```

## Part 1

Ok, on to the puzzle...

We have rules with a list of pairs of pages that must appear in that order.
Then we have multiple lists of 'udpates' which are arrays of page numbers
which may or may not be in the rules.  If any are in the rules, they must
obey them.  The warning in the instructions seems important:

> Of course, you'll need to be careful: the actual list of page ordering
rules is bigger and more complicated than the above example.

I was expecting that, we need to find a very efficient way to test the rules.

My first thought is a list of the pages that must appear before others.
Then we process each update in-order and flag the numbers we've seen before.
We then check the list of pages that have to appear after that number,
if any have been seen before it is invalid.

Now the instructions don't specify a max page number, but looking at my
input, they are all 1-2 digits.  This is handy.   We can use an slice
consisting of 100 slices.  Say we look at the first 4 lines:

    47|53
    97|13
    97|61
    97|47

Our rules slice would have these:

    rules[47] = []int{53}
    rules[97] = []int{13,47,61}

Then when we come across input like this:

    47,53,13,97,14

We have a 'seen' slice with 100 bools that are start as false.

When we see 47, we set `seen[47] = true`.  We check `rules[47]`
and check if `seen[53]` is true.   If it is, we have numbers
that break the rules and we return false.  It isn't, so we continue.

We add set `seen[53] = true`.   `rules[53]` will be an empty
slice, no pages have to appear after it, so we move on.

We get to 13 and set `seen[13] = true`.  `rules[13]` is an empty slice so
we move on.

We get to 97 and set `seen[97] = true`.  `rules[97]` is `{13,47,61}`
so we check and see that `seen[13] == true`.  Since 97 must appear
after 13 according to the rules we mark this update as bad.

I'm wondering what Part 2 will be.  I'm thinking we might as well
use a maps (hashtables) and wondering if we can just use strings.
A map would allow for any size page number.   I think I'll use that.

Ok, I developed this one in vscode first and had some iterations.
First trying `map[int]map[int]bool` for the rules.  This

```go
package day05

import "github.com/JasonGoemaat/go-advent-of-code-2024/util"

func SolveDay05(content string) int {
	groups := util.ParseGroups(content)
	rulesInput := util.ParseNumbers(util.ParseLines(groups[0]), "|")
	updates := util.ParseNumbers(util.ParseLines(groups[1]), ",")
	total := 0

	// rules is map of page number to list of pages that cannot appear
	// before it in an update
	rules := map[int][]int{}
	for _, rule := range rulesInput {
		if rules[rule[0]] == nil {
			rules[rule[0]] = []int{}
		}
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	isOk := func(pages []int) bool {
		seen := map[int]bool{}
		for _, page := range pages {
			seen[page] = true
			for _, after := range rules[page] {
				if seen[after] {
					return false
				}
			}
		}
		return true
	}

	for _, update := range updates {
		if isOk(update) {
			total += update[len(update)/2] // add middle page
		}
	}

	return total
}
```

This worked the first time and gave me the expected 5108 for part 1.
However, now that I think about it and looking at the input data,
it might be better to make the rules a `map[int]map[int]bool` after
all.   There seem to be a LOT of rules (about 1200), and the updates
aren't too big (largest is 23 pages in an update).  It would probably be
faster if I went through all the previous elements in an update
and checked that map of maps.

If all 23 pages in an update have 100 rules, I'll be checking 2200 times.
The other way I would only be checking 231 times.

It seems like it could be an issue either way if the data changes.
If there are updates that have 10,000 pages, that could be a lot worse
the other way.   But if there are 1,000 pages with an average of 500 rules,
that would be much worse the first way.

## Part 2

Now I see why the inputs are so small:

> After taking only the incorrectly-ordered updates and ordering them
correctly, their middle page numbers are 47, 29, and 47. Adding these
together produces 123.

Time for a quick test.   What does a map give if the value isn't there?
Is it the default, i.e. 'false' for bool and '0' for int?

```go
func TestMissingMapInt(t *testing.T) {
	bs := map[int]int{}
	bs[0] = 0
	bs[1] = 1
	t.Logf("bs[0] == %v (0)", bs[0])
	t.Logf("bs[1] == %v (1)", bs[1])
	t.Logf("bs[2] == %v (?)", bs[2])
}
```

Yeah, bools default to false and ints default to 0.  I think that's ok for
int as there doesn't seem to be a page 0.   Otherwise I would have to
initialize values to -1 or MAX_INT (whatever that is in go) to mark
it as not set yet.

Ok, looking up some stuff on part 5 on reddit I assume that every combination
of pages will appear in some rule.  I'll go on that assumption.  By the
instructions I assumed there could be pages in the updates that weren't in the
rules.   In that case I didn't see a way that we could produce a definitive order.

First to note in the rules.   We have to report the middle number after
putting the pages in the correct order, but only for the ones not already in
the correct order.   So we'll have to keep track of if we made any changes.

This makes it seem like doing our sort would be as easy as using a built-in
go sort method, in which case we'd have to compare the slices to make sure
our sorted result wasn't the same as the input.

I think I can also assume that the same page won't be listed twice.   That
makes *common* sense, but isn't specified in the puzzle.

So I'll change `seen map[int]bool{}` to `seen map[int]int{}` to keep track of
the position it was seen in.  Wait, that won't work by itself, or will it?
If I loop through the rules array for pages it must appear before and find
the minimum `seen` location, I can move it before that.   The rules
should guarantee already the pages up to that point are in the correct order.
In this case I would have to initialize all the values to -1, or keep track
of the indexes.  I would also have to change those indexes when inserting
items later.  Not good.

So let's look at sorting.   I can pass a 'less' function to one of the sort
methods in the go standard library I think.  That takes two values and returns
true if the first is less than the second.  That seems perfect, I just check
the rules.   Let me code...

Well, that worked great.  I renamed 'rules' to 'isBeforeRules' to make it
easier to understand, `isBeforeRules[97][13]` will return true if there is
a rule `97|13`.

```go
isBeforeRules := map[int]map[int]bool{}
for _, rule := range rulesInput {
    before := rule[0]
    after := rule[1]
    if isBeforeRules[before] == nil {
        isBeforeRules[before] = map[int]bool{}
    }
    if isBeforeRules[after] == nil {
        // setting this for that one page that doesn't
        // exist as a before rule so we don't get errors
        isBeforeRules[after] = map[int]bool{}
    }

    isBeforeRules[before][after] = true
}
```

I left my `isOk(update)` method rather than do a comparison after sorting.
This might save some time if it's already sorted and we will ignore it,
so there's no reason to write a function to compare the slices.

Then I created a `sortUpdate()` function that takes the pages in an update
and sorts them according to the rules.   That uses the built-in go function
`slices.SortFunc()` along with my `sortFunc()` function that just uses the
`isBeforeRules` map. 

```go
sortFunc := func(a, b int) int {
    if isBeforeRules[a][b] {
        return -1
    }
    return 1 // shouldn't ever be equal in our case
}

sortUpdate := func(pages []int) []int {
    sorted := slices.Clone(pages)
    slices.SortFunc(sorted, sortFunc)
    return sorted
}
```

Then the code to come up with the solution looks like this:

```
for _, update := range updates {
    if !isOk(update) {
        sorted := sortUpdate(update)
        total += sorted[len(sorted)/2] // add middle page number
    }
}
```

Done with day 5!
