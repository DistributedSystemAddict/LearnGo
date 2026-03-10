package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

func main() {
	pflag.BoolP("lines", "l", false, "number of lines")
	pflag.BoolP("chars", "c", false, "number of chars")
	pflag.BoolP("words", "w", false, "number of words")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	files := pflag.Args()
	if !viper.GetBool("line") && !viper.GetBool("word") && !viper.GetBool("chars") {
		viper.Set("lines", true)
		viper.Set("words", true)
		viper.Set("chars", true)
	}
	if len(files) == 0 {
		fmt.Println("Vui lòng cung cấp ít nhất một tệp tin.")
		return
	}

	total := Stats{}

	for _, filename := range files {
		stats, err := wc(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Lỗi đọc %s: %v\n", filename, err)
			continue
		}
		printStats(stats, filename)

		total.lines += stats.lines
		total.words += stats.words
		total.chars += stats.chars
	}

	if len(files) > 1 {
		printStats(total, "total")
	}
}

func printStats(s Stats, label string) {
	result := ""
	if viper.GetBool("lines") {
		result += fmt.Sprintf("%7d ", s.lines)
	}
	if viper.GetBool("words") {
		result += fmt.Sprintf("%7d ", s.words)
	}
	if viper.GetBool("chars") {
		result += fmt.Sprintf("%7d ", s.chars)
	}
	fmt.Printf("%s%s\n", result, label)
}
