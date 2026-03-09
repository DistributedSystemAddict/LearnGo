/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stats",
	Short: "Statistics application",
	Long:  `The statistics application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

type Entry struct {
	Filename string  `json:"filename"`
	Len      int     `json:"length"`
	Minimum  float64 `json:"minimum"`
	Maximum  float64 `json:"maximum"`
	Mean     float64 `json:"mean"`
	StdDev   float64 `json:"stddev"`
}

var logger *slog.Logger

// JSONFILE resides in the current directory
var JSONFILE = "./data.json"

type DFslice []Entry

var data = DFslice{}
var index map[string]int

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = Serialize(&data, f)
	return err
}

func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = DeSerialize(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Filename
		index[key] = i
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))

	if disableLogging == false {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}

	slog.SetDefault(logger)
}

var disableLogging bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&disableLogging, "log", "l", false, "Logging information")

	err := readJSONFile(JSONFILE)

	if err != nil && err != io.EOF {
		return
	}
	createIndex()
}
