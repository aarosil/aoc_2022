package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shapesEncrypted = map[string]string{
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

var shapes = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var shapeScores = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

var winScores = map[string]int{
	"Win":  6,
	"Draw": 3,
	"Lose": 0,
}

var defeats = map[string]string{
	"Rock":     "Paper",
	"Paper":    "Scissors",
	"Scissors": "Rock",
}

func main() {
	fmt.Println("AoC day 2")
	readFile, _ := os.Open("2_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	myScore := 0
	themScore := 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		moves := strings.Split(text, " ")
		them := shapes[moves[0]]
		me := shapesEncrypted[moves[1]]

		myScore += shapeScores[me]
		themScore += shapeScores[them]

		if me == them {
			myScore += winScores["Draw"]
			themScore += winScores["Draw"]
			continue
		}

		iWin := me == defeats[them]
		if iWin {
			myScore += winScores["Win"]
		} else {
			themScore += winScores["Win"]
		}
	}

	fmt.Println("my score:", myScore)
	fmt.Println("them score:", themScore)
}
