package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	correct := 0

	// set the time limit
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s ", i+1, p.q)

		// use this to non-block fmt.Scanf
		answerCh := make(chan string)
		go func() {
			var answer string
			_, err = fmt.Scanf("%s\n", &answer)
			if err != nil {
				exit("Something is wrong!")
			}
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYour score: %d our of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("Your score: %d our of %d.\n", correct, len(problems))

}


func parseLines(lines [][]string) []problem {

	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}


type problem struct {
	q string
	a string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
