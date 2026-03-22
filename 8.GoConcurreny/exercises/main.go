package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/sync/semaphore"
)

type Stats struct {
	lines int
	chars int
	words int
}

type Result struct {
	filename string
	stats    Stats
	err      error
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
	pflag.StringP("output", "o", "", "ghi kết quả ra file")
	pflag.IntP("concurrency", "j", 4, "số lượng worker chạy song song (Semaphore)")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	files := pflag.Args()
	if !viper.GetBool("lines") && !viper.GetBool("words") && !viper.GetBool("chars") {
		viper.Set("lines", true)
		viper.Set("words", true)
		viper.Set("chars", true)
	}
	if len(files) == 0 {
		fmt.Println("Vui lòng cung cấp ít nhất một tệp tin.")
		return
	}

	sem := semaphore.NewWeighted(viper.GetInt64("concurrency"))
	resultChan := make(chan Result, len(files))

	var total Stats
	var mu sync.Mutex

	var wg sync.WaitGroup
	ctx := context.TODO()

	for _, filename := range files {
		wg.Add(1)
		err := sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("Cannot acquire semaphore:", err)
			break
		}

		go func(string) {
			defer wg.Done()
			defer sem.Release(1)
			stats, err := wc(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Lỗi đọc %s: %v\n", filename, err)
				return
			}
			resultChan <- Result{filename: filename, stats: stats, err: err}
		}(filename)
	}

	wg.Wait()
	close(resultChan)
	for res := range resultChan {
		if res.err != nil {
			fmt.Fprintf(os.Stderr, "Lỗi %s: %v\n", res.filename, res.err)
			continue
		}
		mu.Lock()
		total.lines += res.stats.lines
		total.words += res.stats.words
		total.chars += res.stats.chars
		mu.Unlock()
	}

	printStats(total, "total")
	if len(files) > 1 {
		printStats(total, "total")
		fmt.Println("Hello")
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
