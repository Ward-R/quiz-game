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

func toStr(s string) int {
	number, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Error converting to number", err)
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
		// questionString := strings.Split(record, " ")
		question := record[0]
		answer := toStr(record[1])
		userAnswer := 0 // would need to change this to an interface later if you want to hold strings or ints depending on the quiz.

		fmt.Printf("Problem #%v: %v = ", problemCounter, question)
		fmt.Scan(&userAnswer)
		if userAnswer == answer {
			numbCorrect += 1
		}
		fmt.Println(answer) // debug
		problemCounter += 1
	}
	percent := roundTo(((numbCorrect / (problemCounter - 1)) * 100), 1)
	fmt.Printf("You got %v/%v correct! Or %v%%\n", numbCorrect, problemCounter-1, percent)
}
