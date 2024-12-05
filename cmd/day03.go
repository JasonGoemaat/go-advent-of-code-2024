/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day03"
	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day03 called")
		// fmt.Println("sample part1:", day03.SolveDay03("data/day03/sample.txt"))
		// fmt.Println("input part1 :", day03.SolveDay03("data/day03/input.txt"))
		fmt.Println("sample part2:", day03.SolveDay03Part2("data/day03/sample_part2.txt"))
		fmt.Println("input part2 :", day03.SolveDay03Part2("data/day03/input.txt"))
	},
}

func init() {
	rootCmd.AddCommand(day03Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day03Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day03Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
