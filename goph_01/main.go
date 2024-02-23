package main

import (
	"bufio"
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
	csvData, csvErr := os.ReadFile(*csvPtr)

	check(csvErr)

	r := csv.NewReader(strings.NewReader(string(csvData)))

	var incorrect int
	var correct int
	var total int

	// Fix for part 2
	// Get confirmation from user before starting timer
	fmt.Println("Press enter to start the timer for ", *timePtr, "seconds")

	start := time.Now()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Print(record[0], ":")
		total += 1

		ansReader := bufio.NewReader(os.Stdin)

		answer, ansErr := ansReader.ReadString('\n')

		check(ansErr)

		// strip the answer
		if strings.ToLower(strings.TrimSpace(answer)) == strings.ToLower((record[1])) {
			correct += 1
		} else {
			incorrect += 1
		}
		t := time.Now()
		if t.Sub(start) >= time.Duration(*timePtr) {
			fmt.Println("Time over...")
			break
		}
	}

	fmt.Println("Correct: ", correct, "Total", total)

}
