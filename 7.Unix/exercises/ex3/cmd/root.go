/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

type Stats struct {
	lines int
	chars int
	words int
}

func wc(file string) (Stats, error) { //char, line, word, err
	f, err := os.Open(file)
	if err != nil {
		return Stats{}, err
	}

	stats := Stats{}
	defer f.Close()
	r := bufio.NewReader(f)
	re := regexp.MustCompile("[^\\s]+")
	for {
		l, err := r.ReadString('\n')
		if err == io.EOF {
			stats.lines += 1
			stats.chars += len(l)
			if len(l) != 0 {
				words := re.FindAllString(l, -1)
				stats.words += len(words)
			}
			break
		}
		if len(l) != 0 {
			stats.lines += 1
			stats.chars += len(l)
			words := re.FindAllString(l, -1)
			stats.words += len(words)

		}
	}
	return stats, nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mywc [file]",
	Short: "wc",
	Long:  `wx`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			res, _ := wc(file)
			fmt.Printf("%7d %7d %7d %s\n", res.lines, res.words, res.chars, file)
		}
	},
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
	rootCmd.PersistentFlags().BoolP("lines", "l", false, "Đếm số dòng")

	rootCmd.PersistentFlags().BoolP("words", "w", false, "Đếm số từ")

	rootCmd.PersistentFlags().BoolP("bytes", "c", false, "Đếm số byte")
}
