package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	// Reading the cli arguments
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	// Parsing the csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv,")
	}

	problems := parseLines(lines)
	counter := setCounter(0)

	// Reading user's answers
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =\n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			counter()
			fmt.Println("Correct")
		} else {
			fmt.Println("Wrong")
		}
	}
	fmt.Println("TotalScore: ", counter()-1)
}

func parseLines(lines [][]string) []problem {
	prob := make([]problem, len(lines))
	for i, line := range lines {
		prob[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return prob
}

func setCounter(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
