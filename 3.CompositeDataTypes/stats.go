package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

var globalMean float64

func readFile(filepath string) ([]float64, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0)
	for _, line := range lines {
		tmp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Println("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}

	return values, nil
}

func stdDev(x []float64) float64 {
	sum := 0.0
	for _, val := range x {
		sum = sum + val
	}

	meanValue := sum / float64(len(x))
	globalMean = meanValue

	var squared float64
	for i := 0; i < len(x); i++ {
		squared = squared + math.Pow((x[i]-meanValue), 2)
	}

	standardDevitation := math.Sqrt(squared / float64(len(x)))
	return standardDevitation
}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}

	return normalized
}

func main() {
	if len(os.Args) == 1 {
		log.Println("Need one argument!")
		return
	}

	file := os.Args[1]
	values, err := readFile(file)
	if err != nil {
		log.Println("Error reading:", file, err)
		os.Exit(0)
	}
	sort.Float64s(values)

	fmt.Println("Number of values:", len(values))
	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[len(values)-1])

	standardDeviation := stdDev(values)
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

	normalized := normalize(values, globalMean, standardDeviation)
	fmt.Println("Normalized:", normalized)
}
