package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	Question string
	Answer   string
}

func parseLinesToProblems(data [][]string) []problem {
	var problems []problem
	for _, line := range data {
		row := problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
		problems = append(problems, row)
	}
	return problems
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "Csv file in the format of 'question,answer'")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal("Error reading the file", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Error reading the data from the file")
	}

	problems := parseLinesToProblems(lines)
  correct := 0
	for _, row := range problems {
		fmt.Printf("What is the result of the following operation %s: ", row.Question)
		var usrAnswer string
		fmt.Scanf("%s", &usrAnswer)
		if usrAnswer == row.Answer {
      correct++
		}
	}
  fmt.Printf("You've scored %d out of %d\n", correct, len(problems))
}
