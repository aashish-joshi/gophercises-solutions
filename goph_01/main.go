package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	// Commandline flags
	csvPtr := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	timePtr := flag.Int("limit", 30, "the time limit for quiz in seconds (default 30)")
	flag.Parse()

	// Read CSV
	csvData, csvErr := os.Open(*csvPtr)
	check(csvErr)
	r := csv.NewReader(csvData)

	// Get confirmation from user before starting timer
	fmt.Printf("Press enter to start the timer for %d seconds", *timePtr)
	fmt.Scanf("%s\n")

	// Setup the timer
	timer := time.NewTimer(time.Duration(*timePtr) * time.Second)

	var correct int
	var total int

problemloop:
	for {
		record, err := r.Read()
		if err == io.EOF {
			break problemloop
		}
		check(err)
		fmt.Print(record[0], " = ")
		total += 1

		ansChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-ansChan:
			// strip the answer of extra spaces
			if strings.ToLower(strings.TrimSpace(answer)) == strings.ToLower((record[1])) {
				correct++
			}
		}
	}

	fmt.Println("Correct: ", correct, "Total", total)
}
