package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("filename", "problems.csv", "csv filename containing quiz")
	timelimit := flag.Int("time", 30, "time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	total := len(records)
	var correct int

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	for _, record := range records {
		fmt.Printf("Question: %v. Answer: ", record[0])
		answerCh := make(chan string)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			answer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTimes Up!\nResult: %v / %v\n", correct, total)
			return
		case answer := <-answerCh:
			if record[1] == strings.TrimSpace(answer) {
				correct++
			}
		}
	}
	fmt.Printf("Result: %v / %v\n", correct, total)
}
