package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type quiz struct {
	question string
	answer   string
}

func main() {
	var problemFile string
	flag.StringVar(&problemFile, "from", "problems.csv", "specifiy problems file")
	flag.Parse()

	problems := []quiz{}

	file, err := os.Open(problemFile)
	if err != nil {
		fmt.Printf("Error: Coudln't open problem file '%s'\n", problemFile)
		flag.Usage()
		os.Exit(1)
	}

	r := csv.NewReader(file)
	for i := 0; ; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		problems = append(problems, quiz{
			question: record[0],
			answer:   record[1],
		})
	}

	correct := 0
	incorrect := 0
	reader := bufio.NewReader(os.Stdin)

	for i, problem := range problems {
		fmt.Printf("Question #%d: %s\nAnswer: ", i+1, problem.question)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}

		if strings.TrimSpace(answer) == problem.answer {
			correct += 1
			fmt.Printf("-> Correct\n\n")
		} else {
			incorrect += 1
			fmt.Printf("-> Incorrect. The answer is %s\n\n", problem.answer)
		}
	}

	fmt.Printf("-------------------\nCorrect answers: %d\n", correct)
	fmt.Printf("Incorrect answers: %d\n", incorrect)
}
