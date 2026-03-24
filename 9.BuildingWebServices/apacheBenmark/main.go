package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	Duration time.Duration
	Status   int
	Error    error
}

func worker(client *http.Client, url string, jobs <-chan int, results chan<- Result) {
	for range jobs {
		start := time.Now()

		resp, err := client.Get(url)
		duration := time.Since(start)

		if err != nil {
			results <- Result{Duration: duration, Error: err}
			continue
		}

		resp.Body.Close()

		results <- Result{
			Duration: duration,
			Status:   resp.StatusCode,
		}
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <url> <total_requests> <concurrency>")
		return
	}

	url := os.Args[1]
	total := 0
	concurrency := 0

	fmt.Sscanf(os.Args[2], "%d", &total)
	fmt.Sscanf(os.Args[3], "%d", &concurrency)

	jobs := make(chan int, total)
	results := make(chan Result, total)

	client := &http.Client{}

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(client, url, jobs, results)
		}()
	}

	startBenchmark := time.Now()

	for i := 0; i < total; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var success, failed int
	var totalTime time.Duration
	minTime := time.Hour
	maxTime := time.Duration(0)

	for res := range results {
		if res.Error != nil {
			failed++
			continue
		}

		success++
		totalTime += res.Duration

		if res.Duration < minTime {
			minTime = res.Duration
		}
		if res.Duration > maxTime {
			maxTime = res.Duration
		}
	}

	elapsed := time.Since(startBenchmark)

	// Stats
	fmt.Println("---- Benchmark Results ----")
	fmt.Printf("Total Requests: %d\n", total)
	fmt.Printf("Concurrency: %d\n", concurrency)
	fmt.Printf("Success: %d\n", success)
	fmt.Printf("Failed: %d\n", failed)
	fmt.Printf("Total Time: %v\n", elapsed)

	fmt.Printf("Requests/sec: %.2f\n", float64(total)/elapsed.Seconds())

	if success > 0 {
		fmt.Printf("Avg Latency: %v\n", totalTime/time.Duration(success))
		fmt.Printf("Min Latency: %v\n", minTime)
		fmt.Printf("Max Latency: %v\n", maxTime)
	}
}
