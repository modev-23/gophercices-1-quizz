package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func initFile() map[string]string {
	file, err := os.Open("quizz.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	quizzMap := make(map[string]string)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	for _, record := range records {
		if len(record) >= 2 {
			quizzMap[record[0]] = record[1]
		} else {
			fmt.Println("Skipping malformed record:", record)
		}
	}
	return quizzMap
}

func repl(quizzMap map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	totalQuestions := len(quizzMap)

	fmt.Println("---------------------")
	fmt.Println("Welcome to the quizz cli")
	fmt.Println("---------------------")
	fmt.Println("Press Enter to start the quiz")
	fmt.Println("Type 'quit' to quit at any time")

	_, _ = reader.ReadString('\n')

	for question, correctAnswer := range quizzMap {
		fmt.Println(question)
		fmt.Print("Your answer: ")

		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		// Allow user to quit mid-quiz
		if strings.ToLower(answer) == "quit" {
			break
		}

		if strings.TrimSpace(answer) == strings.TrimSpace(correctAnswer) {
			fmt.Println("Correct Answer!")
			score++
		} else {
			fmt.Println("Wrong Answer!")
			fmt.Printf("The correct answer was: %s\n", correctAnswer)
		}
	}

	fmt.Printf("\nYou answered %d out of %d questions correctly!\n", score, totalQuestions)
}

func main() {
	quizzMap := initFile()
	if quizzMap != nil {
		repl(quizzMap)
	}
}
