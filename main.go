package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvfile := flag.String("csv", "problems.csv", "questions")
	timelimit := flag.Int("time", 10, "time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvfile)

	if err != nil {
		fmt.Println("canot open file")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	correct := 0
	scanner := bufio.NewScanner(os.Stdin)
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
loop:
	for _, line := range lines {
		question := line[0]
		answer := strings.TrimSpace(line[1])
		fmt.Printf("question %s:\n", question)

		answerchan := make(chan string)
		go func() {
			scanner.Scan()
			answerchan <- scanner.Text()
		}()
		select {
		case <-timer.C:
			fmt.Println("\n timer is up!")
			break loop
		case youranswer := <-answerchan:
			if youranswer == answer {
				correct++
			}
		}

	}
	fmt.Printf("you have answered %d out of %d\n", correct, len(lines))

}
