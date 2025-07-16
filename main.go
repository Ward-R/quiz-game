package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// Program does not check for white space, case insensitivity, spaces, or nil user entry.

// this returns an interface holding either a string or an int type value.
func parseQuizAnswer(s string) interface{} {
	number, err := strconv.Atoi(s)

	if err != nil {
		return s
	}
	return number
}

func roundTo(f float64, n int) float64 {
	pow10_n := math.Pow(10, float64(n))
	return math.Round(f*pow10_n) / pow10_n
}

func main() {
	fmt.Println("---Let the quiz begin!---")

	file, err := os.Open("problems.csv") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var problemCounter float64 = 1
	var numbCorrect float64 = 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// parse the question and the answer from the CSV.
		// Further parse the CSV answer to either an int or a string as required.
		question := record[0]
		parsedQuizAnswer := parseQuizAnswer(record[1])

		// Get and parse user inputted answer, check if correct.
		var rawUserAnswer string
		fmt.Printf("Problem #%v: %v = ", problemCounter, question)
		fmt.Scanln(&rawUserAnswer)

		parsedUserAnswer := parseQuizAnswer(rawUserAnswer)

		isCorrect := false
		// extract the type of value from the interface
		switch quizAnswerVal := parsedQuizAnswer.(type) {
		case int:
			if userAnswerVal, ok := parsedUserAnswer.(int); ok {
				if userAnswerVal == quizAnswerVal {
					isCorrect = true
				}
			}
		case string:
			if userAnswerVal, ok := parsedUserAnswer.(string); ok {
				if userAnswerVal == quizAnswerVal {
					isCorrect = true
				}
			}
		}
		if isCorrect {
			numbCorrect += 1
		}
		problemCounter += 1
	}
	percent := roundTo(((numbCorrect / (problemCounter - 1)) * 100), 1)
	fmt.Printf("You got %v/%v correct! Or %v%%\n", numbCorrect, problemCounter-1, percent)
}
