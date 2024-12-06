# Day 01 Work

## Part 1

This one seems pretty easy.

1. Create two arrays of numbers from the input
2. Sort arrays
3. Process arrays, taking difference of each set and totalling them
4. Report the total of all differences

First I created some util functions and sample code to check them out.
I created a util package (see blow) with some helper functions to load
the file.   The run command in `day01.go` now looks like this:

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day01 called")
		// solve(filepath.Join("cmd", "day01", "data", "sample.txt"))
		solve(filepath.Join("cmd/day01/data/sample.txt"))
	},

I wasn't sure you could use '/' as a path separator on windows, but it seems
to work.   `filepath.Join` will use a separator that is proper for the os and
actually returns `cmd/day01/data/sample.txt`, but I'm glad to see that I can use
forward slashes on Windows also because it's easier to read.

I decided to use a 'solve' function taking the path to the input file.
This way I can easily change out the name from `sample.txt` to `in.txt`,
or call them both.  Actually maybe the `solve()` function should take
the actual data so I can write tests easier and add a `solveFile` function
to load and parse the file and then call solve, but this seems good for now.

---

Now I want to create two slices (arrays), re-formatting the current structure
which has one array element containing both numbers on each line:

    [[3 4] [4 3] [2 5] [1 3] [3 9] [3 3]]

To two arrays, one containing the first number for each line and one
containing the second number for each line:

    [3 4 2 1 3 3]
    [4 3 5 3 9 3]

In javascript/typescript I'd probably do something like this:

    const parts = arr.reduce((p, c) => {
        p.a.push(c[0])
        p.b.push(c[1])
        return p
    }, {a: [], b: []})

I don't know of an equivalent in go, so let's write a function:

```go
func takeElement(numbers int[][], index int): int[] {
    // I could start with an empty slice and use append, but
    // if you know the size ahead of time I think there's a
    // function called make(), looking that up...
    // didn't find a good explanation quickly, but let's try it
    var a = make([]int, len(numbers))
    for i,n := range(numbers) {
        a[i] = numbers[i][index]
    }
    return a
}
```

I'm not sure that will be used elsewhere, so putting it in `day01.go` for now.  Seems to work:

```go
func solve(filePath string) string {
	numbers := util.LoadNumbers(filePath)
	a := takeElement(numbers, 0)
	b := takeElement(numbers, 1)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	return "answer"
}
```

    a: [3 4 2 1 3 3]
    b: [4 3 5 3 9 3]

I looked up and saw there was a package for sorting...  Let me look
that up again.  Ok, there's a `sort` package with the `Sort(slice)`
function and it sorts in-place.  Trying that...

Ok, there's several helpers in the `sort` package.   This seems easiest
since we're dealing with a popular type:

    sort.Ints(a)

Looks good to me:

    a: [1 2 3 3 3 4]
    b: [3 3 3 4 5 9]

I want to check out the other methods though.   There's a `sort.Slice()`
function.  That seems to take the slice and a function as arguments.   The
function takes two integer arguments representing indexes into the slice
and should return true if the element at the first index is less than the
element at the second index.   This seems to work too.

```go
sort.Slice(a, func(i int, j int) bool {
    return a[i] < a[j]
})
sort.Slice(b, func(i int, j int) bool {
    return b[i] < b[j]
})
```

Looking at the docs, there's also a `sort.SliceStable()` function that will
preserve ordering of equal elements FYI.  There are also `Search()` and 
`Find()` methods for finding elements in a sorted slice doing a binary
search, that's handy.

Another method is to call `sort.Sort()` or `sort.Stable()` which take an
interface that should have a `Len()` function that returns the length of the
slice, a `Less()` function that works like the function above returning
true if the value at the first index is less than the second, and ...

Well, maybe I have that wrong...   Looking at [go by example](https://gobyexample.com/sorting),
it looks like we can just call `Sort(slice)`...  Nope, well I should
have looked closer, that is in the `slice` package and this works just fine:

```go
slices.Sort(a)
slices.Sort(b)
```

And this sample which uses `cmp.Compare` for some reason, but the base
call lets you pass a compare function and would then let you pick fields
if slices are of more complex types.   Using a variable for the function
could be used in the first example above too:

```go
myCompare := func(a, b int) int {
    return cmp.Compare(a, b)
}
slices.SortFunc(a, myCompare)
slices.SortFunc(b, myCompare)
```

I don't need a special comparison, so I'll stick with the one with the
least code for now.

Thinking about the problem...   Maybe I could just total the lists
and return the difference.   I don't think that will work because
I think we need `abs(a-b)` with no negatives, I'll check the problem...
Yeah, it specifies 'how far apart' they are.

Now I'm thinking about what the second puzzle might be.   Each day there
are two puzzles.   The second one is more difficult, but usually
related closely to the first puzzle.   For example if this puzzle was the
hard one, the easy one might just ask you to total the two columns and
report the total difference so you wouldn't have to sort them.   I can't
think how to make this more difficult at the moment, so let's move on...

Well I found `math.Abs()`, but it only works with float64.  So I'll write
my own diff function.  I'll also change the function to return an int
and print it in the command:

```go
slices.Sort(a)
slices.Sort(b)
total := 0
distance := func(i, j int) int {
    if i < j {
        return j - i
    }
    return i - j
}
for i, v := range a {
    total += distance(v, b[i])
}
return total
```

This returns 11 for the sample (correct), let's try for my input download
too and submit it...

```go
fmt.Println("Solution:", solve(filepath.Join("cmd/day01/data/sample.txt")))
fmt.Println("Solution:", solve(filepath.Join("cmd/day01/data/in.txt")))
```

## Part 2

Ah, that's right.   The data is the same as in part 1, but you do something
more complicated.   And as I remember, you combine the data in an unintuitive
way to get a different answer.   Here we multiply each number in the first
list by how many times it appears in the second list, and total those results.

Sorting both lists still seems like it will let us process them efficiently.
There are a few things I'm thinking of doing:

1. a `while` loop (does go have while?) and track each index
    1. exit loop if `aIndex >= len(a)` or `bIndex >= len(b)`
    2. if `b` < `a`, move `bIndex` forward 1, continue loop (or keep moving forward until it is or we're past the end of `b`)
    2. if `a` == `b`, add `a` to total, move `b` forward 1, continue loop
    3. move 

> HAVE TO STOP

Ok, thinking and re-reading, what I was thinking won't work exactly,
but maybe something similar will.   Spit-balling:

1. Use something like a hashtable pre-populated with counts from `b`, and process `a`
2. Use binary search on `b` for each `a`
3. Better than last, pre-process `b` with sort, then make `c` with unique numbers from `b` and `d` with counts of those at same index as `c`, use `sort.Search` or `sort.Find` I saw earlier to do binary search to find index of number in `c` and use that index to get count in `d`

I see go has a `map` type for hashtables, I'll try that first.  For my project
structure I'll still use `day01` command since they use the same data and create
a `solve2` function.

Thinking about the data, it's possible the numbers will overflow a built-in int type,
but go int is 64 bits on a 64 bit system.   It also has a `big` package for arbitrary
precision so we'll worry about that if it's a problem.

First some quick code to checkout the map class:

```go
func solve2(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	a := takeElement(numbers, 0)
	b := takeElement(numbers, 1)
	slices.Sort(a)
	m := make(map[int]int)
	for _, v := range b {
		m[v]++
	}
	fmt.Printf("m: %v\n", m)
	return 0
}
```

    m: map[3:3 4:1 5:1 9:1]
    Solution (part 2 sample): 0

Seems good to me.   Bow that I think about it, no reason to sort `a` either...

```go
func solve2(filePath string) int {
	numbers := util.LoadNumbers(filePath)
	a := takeElement(numbers, 0)
	b := takeElement(numbers, 1)
	m := make(map[int]int)
	for _, v := range b {
		m[v]++
	}

	total := 0
	for _, v := range a {
		total += v * m[v]
	}

	return total
}
```

That works for sample, I'll submit:

    PS C:\git\go\go-advent-of-code-2024> go run . day01
    day01:
    Solution (part 1 sample): 11
    Solution (part 1 input ): 2904518
    Solution (part 2 sample): 31
    Solution (part 2 input ): 18650129

All good, day 1 complete!

## util

This gives me a chance to start my `util` package to handle some things.
To parse this puzzle I created three functions that I think can be used
in many more of the puzzles:

    func LoadString(filePath string) string
    func LoadLines(filePath string) []string
    func LoadNumbers(filePath string)

`LoadString` just takes the path and returns one string of the entire file's
contents.  It is used by `LoadLines`.

`LoadLines` calls `LoadString` and splits the contents into a slice (array)
of lines.   It also ignores a blank line at the end if there is one.
I had to use a regular expression to handle the file having `\r\n` as the
separator on Windows and to handle just `\n` if the file has that.

`LoadNumbers` calls `LoadLines` and parses each line into an array of
numbers separated by any number of spaces.

With these functions, I can use this call:

    numbers := util.LoadString("cmd/day01/data/sample.txt")

To take the input:

And return this:

    [[3 4] [4 3] [2 5] [1 3] [3 9] [3 3]]

Interestingly you can print things in go using `fmt.Printf` with
the `%v` specifier:

    fmt.Printf("%v\n", numbers)

And it shows a visual representation.   I wouldn't call it serialization,
but it's similar.   As you can probably tell above, the result is an
array of arrays.   Each array consists of two integers.   The reason it's
not really serialization is that it uses spaces as a separator and nothing
to note a string.   The first element `[3 4]` could be:

1. An array with two numbers, the integer 3 and the integer 4 (this is what it actually is)
2. An array with two strings, "3" and "4"
3. An array with one string, "3 4"
4. A string "[3 4]"

