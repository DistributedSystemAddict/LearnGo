/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ioFs",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var embedFile embed.FS

var searchString string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(f embed.FS) {
	embedFile = f
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func walkFunction(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("Path=%q, isDir=%v\n", path, d.IsDir())
	return nil
}

func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.Name() == searchString {
		fileInfo, err := fs.Stat(embedFile, path)
		if err != nil {
			return err
		}
		fmt.Println("Found", path, "with size", fileInfo.Size())
		return nil
	}
	return nil
}

func writeToFile(s []byte, path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	n, err := fd.Write(s)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}

func list(f embed.FS) error {
	return fs.WalkDir(f, ".", walkFunction)
}

func search(f embed.FS) error {
	if searchString == "" {
		return nil
	}
	return fs.WalkDir(f, ".", walkSearch)
}

func extract(f embed.FS, filepath string) ([]byte, error) {
	s, err := fs.ReadFile(f, filepath)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func init() {

}
