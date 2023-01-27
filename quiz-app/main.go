package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

func scanUserAnswer(answerCh chan string) {
	var usrAnswer string
	fmt.Scanf("%s", &usrAnswer)
	answerCh <- usrAnswer
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "Csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for _, row := range problems {
		fmt.Printf("What is the result of the following operation %s: ", row.Question)

		answerCh := make(chan string)
		go scanUserAnswer(answerCh)

		select {
		case <-timer.C:
			fmt.Printf("\nTime expired! You've scored %d out of %d\n", correct, len(problems))
			return
		case usrAnswer := <-answerCh:
			if usrAnswer == row.Answer {
				correct++
			}
		}
	}
	fmt.Printf("You've scored %d out of %d\n", correct, len(problems))
}
