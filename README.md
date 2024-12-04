# Advent of Code 2024 in GO

Repo for my code figuring out the puzzles in [Advent of Code 2024](https://adventofcode.com/2024).

I thought this would be a good time to dive deeper into GO.
Using [cobra](https://github.com/spf13/cobra) and 
[cobra-cli](https://github.com/spf13/cobra-cli).

Each day should have a folder like `data/dayXX` with these files:

* `Instructions.md` - from the website
* `sample.txt` - the puzzle sample where the result is given
* `input.txt` - the input for the puzzle, unique per user
    * Direct download from like `https://adventofcode.com/2024/day/1/input`
* `Work.md` - my thoughts while figuring out the puzzle

Using `cobra-cli`, we create a command for each day running commands like this:

    cobra-cli add day01

This creates a go file such as `cmd/day01.go`.

There is a `data` directory with a subdirectory for each day (i.e. `day01`
that will have these files:

* `input.txt` - the downloadable solution input, unique for each user
* `Instructions.md` - the instructions from the site
* `sample.txt` - file for the sample data in the instructions that is explained with an answer
* `Work.md` - file for me to basically walk through my thought process and share what I learned about go

