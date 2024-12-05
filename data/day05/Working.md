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

## Part 1

This time it looke like we'll need a new loader, the file comes in two
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
and change the 'LoadXXX' functions to call those after 'LoadString'

```
func LoadLineGroups(string filePath)