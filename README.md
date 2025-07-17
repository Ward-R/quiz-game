# quiz-game

This is a simple command line game I made using the project idea from https://gophercises.com/exercises/quiz. This program got me more familiar with file reading and parsing,
using flags, channels, time package, getting user input, and using interfaces. All I can say is making a program like this was way different than my previous experience
in dynamic languages like Ruby, but it was fun!

This program will read a csv file in the format of question,answer on each row. I did not make it remove whitespace, and it is case sensitive. It does however allow
for both integer and string questions.

## Features

* **Timed Quiz:** The quiz runs for a set duration, stopping automatically when time expires.
* **Customizable Time Limit:** The time limit can be changed via a command-line flag.
* **Customizable Quiz File:** The CSV file containing questions can be specified via a command-line flag.
* **Interactive Start:** The quiz begins only after the user presses Enter.
* **Real-time Questioning:** Questions are presented one by one, allowing concurrent timing while waiting for user input.
* **Score Tracking:** Tracks correct answers and displays the final score.
* **Flexible Answers:** Supports both integer and string answers from the CSV.

## Flags

* Use `-limit` flag to change the time limit. Default is `30s`.
    * **Example:** `go run . -limit=60s` (for 60 seconds)
    * **Example:** `go run . -limit=2m` (for 2 minutes)
* Use `-file` flag to change the CSV file to read from. Default is `problems.csv`.
    * **Example:** `go run . -file=my_quiz.csv`

https://github.com/gophercises/quiz

## Example of how CSV file should be:

    5+5,10
    7+3,10
    1+1,2
    8+3,11
    1+2,3
    8+6,14
    3+1,4
    1+4,5
    5+1,6
    2+3,5
    3+3,6
    2+4,6
    5+2,7
    What is the Capital of France?,Paris

## Built With

* [Go](https://go.dev/) - Backend Programming Language

## Use

To run this quiz game:

1.  **Clone the repository:**
        git clone https://github.com/YOUR_GITHUB_USERNAME/quiz-game.git

2.  **Navigate to the project directory:**
        cd quiz-game

3.  **Run the game:**

    You can run the game directly using `go run` (which will compile and execute in one step):
        go run . -limit=30s -file=problems.csv

    Or, you can build an executable first (recommended for deployment or if `go run` has issues):
        go build -o quizgame
        ./quizgame -limit=30s -file=problems.csv
    *(Note: On Windows, use `.\quizgame.exe`)*

## Acknowledgements

This project is based on the [Quiz Game exercise](https://gophercises.com/exercises/quiz) from Gophercises. The purpose was to further my knowledge in Go, particularly with file I/O, command-line flags, concurrency (goroutines and channels), and basic input/output.
