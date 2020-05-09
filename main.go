package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.q)
		var answer string
		_, err = fmt.Scanf("%s\n", &answer)
		if err != nil {
			exit("Something is wrong!")
		}
		if answer == p.a {
			correct++
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
