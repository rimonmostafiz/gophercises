package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Quiz struct {
	Question string
	Answer   string
}

func parseQuizes(lines [][]string) []Quiz {
	var quizes = make([]Quiz, len(lines))

	for i, line := range lines {
		quizes[i] = Quiz{
			Question: line[0],
			Answer:   line[1],
		}
	}

	return quizes
}

func main() {
	csvFileName := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer', file extension .csv is expected",
	)

	flag.Parse()

	csvFile, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open CSV file: %s\n", *csvFileName)
		os.Exit(1)
	}

	csvReader := csv.NewReader(csvFile)
	lines, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to parse provided csv file")
		os.Exit(1)
	}

	var quizes = parseQuizes(lines)

	totalQuestion := len(quizes)
	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	for i := 0; i < totalQuestion; i++ {
		fmt.Printf("Problem #%d: %s = ", i+1, quizes[i].Question)
		if scanner.Scan() {
			answer := scanner.Text()
			if quizes[i].Answer != answer {
				break
			}
			count = i + 1
		}
	}

	fmt.Printf("Total correct ans %d out of %d\n", count, totalQuestion)
}
