# Day 3: Mull It Over

Starting up I create files and download data into `data/day03`, and run this:

    cobra-cli add day03

Got a change I want to make to the organization.   I'm renaming the code
file like `solve_day03.go` and naming the function 
`SolveDay03(filePath string)` instead of using the same name for each
day.   The command file that calls it is different anyway, this will not
caus any extra work and will let me combine everything into one directory
later if I wanted to.

This puzzle seems a little odd, there seems to be only one line.
Looks like a good test for regex, which it would be nice to know about in
GO.

Thinking, it seems like this would be a good regular expression
to capture each 'mul' instruction:

    mul\((\d+),(\d+)\)

This has two capture groups and represents `mul(a,b)` with no possible
alterations.  We capture all matches and look through each with a global
regex and parse each of the two groups and multiply, then total. My guess
is that part 2 will require us to allow for whitespace in the instruction,
we'll see.  Maybe it is to allow hex (`0xfe` for instance) or something
else.

There's a `regexp` package it seems, let's try this...   The docs don't
seem great here, I don't know what the `n` parameter means...

```go
rx := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
results := rx.FindAllStringSubmatch(line, -1)
```

So this is what the results look like for the sample input:

    PS C:\git\go\go-advent-of-code-2024> go run . day03 --log
    day03 called
    xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
    [[mul(2,4) 2 4] [mul(5,5) 5 5] [mul(11,8) 11 8] [mul(8,5) 8 5]]

That seems perfect!   It returns a slice of slices.   Each child slice
is an instance of matching the regex.   Each child has three elements: the
matched string and the two matched groups.   So let's solve it!   I have to
parse ints and have to look that up...  Looks like it's `strconv.Atoi` to
convert Ascii string to integer.  Don't even need a helper function, this
works for sample data:

```go
func SolveDay03(filePath string) int {
	line := util.LoadLines(filePath)[0]
	rx := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	results := rx.FindAllStringSubmatch(line, -1)
	total := 0
	for _, match := range results {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		total += a * b
	}
	return total
}
```

Ok, that didn't work for the input, re-reading the instructions it specifies
that the numbers are 1-3 digits and I allow any number of digits.   Let's
change the regex and try this:

```go
rx := regexp.MustCompile("mul\\((\\d\\d?\\d?),(\\d\\d?\\d?)\\)")
```

No deal, I get the same wrong answer.   Before logging matches to see
what's going on, I looked up `?`: https://stackoverflow.com/questions/5583579/question-marks-in-regular-expressions

So let's try `??` instead of `?` to be greedy...

Same answer, ok, going to log all matches along with calcualted
values for each number, multiplication, and running totals.   The input isn't
toooo long...

```go
func SolveDay03(filePath string) int {
	line := util.LoadLines(filePath)[0]
	// rx := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	rx := regexp.MustCompile("mul\\((\\d\\d??\\d??),(\\d\\d??\\d??)\\)")
	results := rx.FindAllStringSubmatch(line, -1)
	total := 0
	for _, match := range results {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		product := a * b
		total += a * b
		util.MyLog("%s: %d * %d = %d  (running total %d)\n", match[0], a, b, product, total)
	}
	return total
}
```

Everything I do catch looks good, all have digits and are converted to integers.
Doing a quick search through the text I don't find any instances of "`(-`" or "`,-`"
so I don't think it has to do with there being negative numbers.

Ok, I'm an idiot.   I loaded the input in vscode and did regexp replace of `mul`
with `mul\n` so I could view each possible multiply instruction on a line.
Spot-checking looked sane to me.   I couldn't figure it out then realized the 
input has 6 lines.   I changed this line and it worked:

	// line := util.LoadLines(filePath)[0]
	line := util.LoadString(filePath)

## Part 2

Now we have to deal with a few other instructions.  There are `do()` and `don't()`
instructions that enable and disable whether `mul()` instructions are executed
after them until the next one.   The text starts with them enabled.

Doesn't seem too bad, I can use the existing code for each enabled section.
What I think I want to do is split the string based on `do()`.   For each of these
I want to remove anything after a `don't()`.   Then I can combine them all
back into one string and run the existing code.

First I'll separate out the code to calculate totals for a string.  Now the
solver for part 1 looks like this:

```go
func SolveDay03(filePath string) int {
	line := util.LoadString(filePath)
	return calculate(line)
}
```

There is a package `strings` and `strings.Split(s string) []string` method that
looke like what we need.   Split main string by "`do()`".   For each part,
split by "`don't()`" and process only the first.   If there is no "`don't()`",
there should be one element in the slice.   Trying this:

```go
func SolveDay03Part2(filePath string) int {
	line := util.LoadString(filePath)
	dos := strings.Split(line, "do()")
	total := 0
	for _, s := range(dos) {
		split := strings.Split(s, "don't()")
		total += calculate(split[0])
	}
	return calculate(line)
}
```

Well, I got the exact same results, that's not good :)  I'll add some logs and
run for only the sample:

```go
func SolveDay03Part2(filePath string) int {
	line := util.LoadString(filePath)
	dos := strings.Split(line, "do()")
	util.MyLog("dos: %v\n", dos)
	total := 0
	for _, s := range dos {
		split := strings.Split(s, "don't()")
		util.MyLog("split: %v\n", split)
		total += calculate(split[0])
	}
	return calculate(line)
}
```

Making sense of the input:

    xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))

I get these for `dos`:

    xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)un
    ?mul(8,5))

The first split has these two parts:

    xmul(2,4)&mul[3,7]!^
    _mul(5,5)+mul(32,64](mul(11,8)un

STFU.   I'm retarded.   I don't return the total, I return
`calculate(line)` at the end which ignores everything I did up until then
and returns the same as part1 for the entire string...

smh...

Returning 'total' instead of that fixed it and Day 3 is done!

