package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	question, answer string
}

func main() {
	fileName, timeLimit := readFlags()
	file := openFile(fileName)
	defer file.Close()
	lines := readFile(file)
	problems := parseProblems(lines)
	runQuiz(problems, timeLimit)
}

func readFlags() (*string, *int) {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer', file extension .csv is expected")
	timeLimit := flag.Int("limit", 30, "timelimit in second for the quiz")
	flag.Parse()
	return fileName, timeLimit
}

func openFile(fileName *string) *os.File {
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Unable to open %s file\n", *fileName)
		os.Exit(1)
	}
	return file
}

func readFile(file *os.File) [][]string {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Error while rading csv file\n")
		os.Exit(1)
	}
	return records
}

func parseProblems(lines [][]string) []Problem {
	var problems = make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{line[0], line[1]}
	}
	return problems
}

func runQuiz(problems []Problem, timeLimit *int) {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correctAnsCount := 0

	for i, problem := range problems {

		fmt.Printf("Problem #%d %s = ", i+1, problem.question)
		scanner := bufio.NewScanner(os.Stdin)
		responseCh := make(chan string)

		go func() {
			scanner.Scan()
			response := scanner.Text()
			responseCh <- response
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time Out!!!\n")
			fmt.Printf("You scored %d out of %d\n", correctAnsCount, len(problems))
			os.Exit(1)
		case answer := <-responseCh:
			if problem.answer == answer {
				correctAnsCount++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correctAnsCount, len(problems))
	return
}
