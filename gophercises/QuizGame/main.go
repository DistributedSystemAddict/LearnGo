package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func timeOut(t time.Duration) {
	<-time.After(t * time.Second)
	fmt.Println("Time Out")
	os.Exit(1)
}

func main() {
	pflag.IntP("time", "t", 30, "Time of games")
	pflag.BoolP("suffle", "s", false, "Time of games")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	go timeOut(time.Duration(viper.GetInt("time")))

	qR := 0
	qW := 0

	for _, r := range records {
		input := bufio.NewReader(os.Stdin)
		if err != nil {
			panic(err)
		}
		fmt.Printf("What is %s ?\n", r[0])
		answer, _ := input.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer == r[1] {
			qR++
			fmt.Print("Correct!, next question\n")
		} else {
			qW++
			fmt.Print("Incorrect!\n")
		}
	}

	fmt.Printf("The correct answer is: %d\n", qR)
	fmt.Printf("The wrong answer is %d\n", qW)
}
