/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "Day 2: Red-Nosed Reports",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day02 called")
		fmt.Println("MyLogEnabled:", util.MyLogEnabled)
		// fmt.Println("sample:", day02.Solve("data/day02/sample.txt"))
		// fmt.Println("input :", day02.Solve("data/day02/input.txt"))
		// fmt.Println("Part 2 GOOD  :", day02.Solve2("data/day02/part2_goodtests.txt"))
		// fmt.Println("Part 2 sample:", day02.Solve2("data/day02/sample.txt"))
		// fmt.Println("Part 2 input :", day02.Solve2("data/day02/input.txt"))
	},
}

func init() {
	rootCmd.AddCommand(day02Cmd)
}
