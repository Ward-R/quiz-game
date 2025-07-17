package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
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

func score(numbCorrect float64, totalQuestions float64) {
	percent := roundTo(((numbCorrect / totalQuestions) * 100), 1)
	fmt.Printf("You got %v/%v correct! Or %v%%\n", numbCorrect, totalQuestions, percent)
}

func main() {

	//flags name := flag.<type string/int/bool>(variable name, default val, description)
	timePtr := flag.Duration("limit", 30*time.Second, "quiz time limit")
	filePtr := flag.String("file", "problems.csv", "quiz filename")
	flag.Parse()

	fmt.Println("Press Enter to start the quiz...")
	var start string
	fmt.Scanln(&start) // This will block until the user presses Enter

	chTimeUp := time.After(*timePtr)
	chUserAnswer := make(chan interface{})

	fmt.Println("---Let the quiz begin!---")

	fileName := *filePtr

	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	totalQuestions := len(records)

	var numbCorrect float64 = 0

quizLoop:
	for i, record := range records {

		// parse the question and the answer from the CSV.
		// Further parse the CSV answer to either an int or a string as required.
		question := record[0]
		parsedQuizAnswer := parseQuizAnswer(record[1])
		fmt.Printf("Problem #%v: %v = ", i+1, question)

		// Get and parse user inputted answer, check if correct.
		go func() {
			var rawUserAnswer string
			fmt.Scanln(&rawUserAnswer)
			parsedUserAnswer := parseQuizAnswer(rawUserAnswer)
			chUserAnswer <- parsedUserAnswer
		}()

		select {
		case <-chTimeUp:
			fmt.Println("\nTime's up!")
			break quizLoop
		case userAnswerReceived := <-chUserAnswer:
			isCorrect := false
			// extract the type of value from the interface
			switch correctAnswer := parsedQuizAnswer.(type) {
			case int:
				if userAnswerInt, ok := userAnswerReceived.(int); ok {
					if userAnswerInt == correctAnswer {
						isCorrect = true
					}
				}
			case string:
				if userAnswerVal, ok := userAnswerReceived.(string); ok {
					if userAnswerVal == correctAnswer {
						isCorrect = true
					}
				}
			}
			if isCorrect {
				numbCorrect += 1
			}
		}
	}
	score(float64(numbCorrect), float64(totalQuestions))
}
