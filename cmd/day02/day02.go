/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package day02

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var Day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "Day 2: Red-Nosed Reports",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day02 called")
		fmt.Println("sample:", Solve("cmd/day02/data/sample.txt"))
		fmt.Println("input :", Solve("cmd/day02/data/input.txt"))
		fmt.Println("Part 2 GOOD  :", Solve2("cmd/day02/data/part2_goodtests.txt"))
		fmt.Println("Part 2 sample:", Solve2("cmd/day02/data/sample.txt"))
		fmt.Println("Part 2 input :", Solve2("cmd/day02/data/input.txt"))
	},
}
