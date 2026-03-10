/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// charsCmd represents the chars command
var charsCmd = &cobra.Command{
	Use:   "chars",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			res, _ := wc(file)
			fmt.Printf("%7d %s\n", res.chars, file)
		}
	},
}

func init() {
	rootCmd.AddCommand(charsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// charsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// charsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
