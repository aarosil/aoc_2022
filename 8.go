package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid = [][]int{}

func init() {
	fmt.Println("AoC Day 8")
}

func main() {
	readFile, _ := os.Open("8_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		grid = append(grid, parseRow(text))
	}
}

func parseRow(text string) []int {
	result := []int{}
	arr := strings.Split(text, "")
	for _, v := range arr {
		d, _ := strconv.Atoi(v)
		result = append(result, d)
	}
	return result
}

func onEdge(x, y int) bool {
	if x == 0 || x == len(grid[0]) {
		return true
	}
	if y == 0 || y == len(grid) {
		return true
	}
	return false
}

func checkVizNorth(x, y int) bool {
	return false
}

func checkVizSouth(x, y int) bool {
	return false
}

func checkVizEast(x, y int) bool {
	return false
}

func checkVizWest(x, y int) bool {
	return false
}
