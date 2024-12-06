/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day01"
	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day02"
	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day03"
	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day04"
	"github.com/JasonGoemaat/go-advent-of-code-2024/cmd/day05"
	"github.com/JasonGoemaat/go-advent-of-code-2024/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var EnableLog bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-advent-of-code-2024",
	Short: "My Advent of Code 2024 Solutions",
	Long:  `Check README.md`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-advent-of-code-2024.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&util.MyLogEnabled, "log", "l", false, "enables mylog")
	rootCmd.PersistentFlags().BoolVarP(&util.StdinFlag, "stdin", "", false, "read test data from stdin")
	rootCmd.PersistentFlags().StringVarP(&util.InputFile, "file", "f", "", "read test data from stdin")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add days
	rootCmd.AddCommand(day01.Day01Cmd)
	rootCmd.AddCommand(day02.Day02Cmd)
	rootCmd.AddCommand(day03.Day03Cmd)
	rootCmd.AddCommand(day04.Day04Cmd)
	rootCmd.AddCommand(day05.Day05Cmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-advent-of-code-2024" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-advent-of-code-2024")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
