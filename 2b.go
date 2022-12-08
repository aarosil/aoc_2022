package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var choiceMap = map[string]map[string]string{}

var outcomeEncrypted = map[string]string{
	"X": "Lose",
	"Y": "Draw",
	"Z": "Win",
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

var loses = map[string]string{
	"Rock":     "Scissors",
	"Paper":    "Rock",
	"Scissors": "Paper",
}

func init() {
	outcomes := []string{"Win", "Lose", "Draw"}

	for _, outcome := range outcomes {
		choiceMap[outcome] = map[string]string{}

		for _, shape := range shapes {
			if outcome == "Win" {
				choice := defeats[shape]
				choiceMap[outcome][shape] = choice
			} else if outcome == "Lose" {
				choice := loses[shape]
				choiceMap[outcome][shape] = choice
			} else {
				choiceMap[outcome][shape] = shape
			}
		}
	}

	fmt.Println(choiceMap)
}

func getMe(outcome, them string) string {
	return choiceMap[outcome][them]
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
		outcome := outcomeEncrypted[moves[1]]
		me := getMe(outcome, them)

		myScore += shapeScores[me]
		themScore += shapeScores[them]

		if outcome == "Draw" {
			myScore += winScores["Draw"]
			themScore += winScores["Draw"]
			continue
		}

		if outcome == "Win" {
			myScore += winScores["Win"]
		} else {
			themScore += winScores["Win"]
		}
	}

	fmt.Println("my score:", myScore)
	fmt.Println("them score:", themScore)
}
