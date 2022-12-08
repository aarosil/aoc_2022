package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stacks = [][]string{}
var moves = []string{}

func init() {
	fmt.Println("AoC Day 5")
	// testArray := [][]string{{"A", "B", "C"}, {"D", "E", "F"}, {"G", "H", "I"}}
	testArray := [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}
	fmt.Println(testArray)
	testArray = doMoves9001(testArray, 1, 1, 0)
	fmt.Println(testArray)
	testArray = doMoves9001(testArray, 3, 0, 2)
	fmt.Println(testArray)
	testArray = doMoves9001(testArray, 2, 1, 0)
	fmt.Println(testArray)
	testArray = doMoves9001(testArray, 1, 0, 1)
	fmt.Println(testArray)
	fmt.Println("---")
}

func main() {
	readFile, _ := os.Open("5_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	moving := false

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" {
			moving = true
			continue
		}

		if !moving {
			chars := strings.Split(text, "")
			stacks = append(stacks, chars)
			continue
		}

		moves = append(moves, text)
	}

	cranes := parseCranes(stacks)
	moves := parseMoves(moves)

	// moves = [][]int{}
	for _, move := range moves {
		amount, source, dest := move[0], move[1], move[2]
		cranes = doMoves9001(cranes, amount, source, dest)
	}

	fmt.Println(getResult(cranes))
}

func parseCranes(stacks [][]string) [][]string {
	indices := []int{}
	lastCol := len(stacks) - 1
	craneIndices := stacks[lastCol]
	boxesRows := stacks[:lastCol]

	for i, v := range craneIndices {
		if v != " " {
			indices = append(indices, i)
		}
	}

	cranes := make([][]string, len(indices))
	for _, boxes := range boxesRows {
		for craneNum, index := range indices {
			val := boxes[index]
			if val == " " {
				continue
			}

			cranes[craneNum] = append(cranes[craneNum], val)
		}
	}

	return cranes
}

func parseMove(mv string) []int {
	parts := strings.Split(mv, " ")
	amountS, sourceS, destS := parts[1], parts[3], parts[5]
	result := make([]int, 3)
	for i, v := range []string{amountS, sourceS, destS} {
		j, _ := strconv.Atoi(v)
		if i > 0 {
			result[i] = j - 1
		} else {
			result[i] = j
		}
	}
	return result
}

func parseMoves(moves []string) [][]int {
	result := make([][]int, 0)
	for _, v := range moves {
		mv := parseMove(v)
		result = append(result, mv)
	}
	return result
}

func getResult(stacks [][]string) string {
	result := ""
	for _, v := range stacks {
		result += v[0]
	}
	return result
}

func doMoves(cranes [][]string, amount, source, dest int) [][]string {
	for i := 0; i < amount; i++ {
		// remove 1 element from source array and add it to dest array
		box := cranes[source][0]
		cranes[source] = cranes[source][1:]
		cranes[dest] = append([]string{box}, cranes[dest]...)
	}

	return cranes
}

func doMoves9001(cranes [][]string, amount, source, dest int) [][]string {
	boxes := make([]string, amount) //cranes[source][:amount]
	copy(boxes, cranes[source][:amount])
	cranes[source] = cranes[source][amount:]
	fmt.Println("1", cranes[source])
	cranes[dest] = append(boxes, cranes[dest]...)
	fmt.Println("2", cranes[source])

	return cranes
}
