package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid = [][]int{}
var visibleCount int
var highestSceneScore int

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

	for y, row := range grid {
		for x, height := range row {
			if isVisible(height, x, y) {
				visibleCount++
			}
			sceneScore := getSceneScore(x, y)
			if sceneScore > highestSceneScore {
				highestSceneScore = sceneScore
			}
		}
	}

	fmt.Println(visibleCount)
}

func getSceneScore(x, y int) int {
	return 0
}

func isVisible(height, x, y int) bool {
	if x == 0 ||
		y == 0 ||
		x == len(grid[0])-1 ||
		y == len(grid)-1 {
		return true
	}
	fmt.Println(x, y, height)

	if checkVizNorth(height, x, y, grid) ||
		checkVizSouth(height, x, y, grid) ||
		checkVizEast(height, x, y, grid) ||
		checkVizWest(height, x, y, grid) {
		return true
	}

	return false
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

func checkVizNorth(height, x, y int, grid [][]int) bool {
	for yPos := 0; yPos < y; yPos++ {
		if grid[yPos][x] >= height {
			return false
		}
	}
	fmt.Println("visible from north")
	return true
}

func checkVizSouth(height, x, y int, grid [][]int) bool {
	for yPos := y + 1; yPos < len(grid); yPos++ {
		if grid[yPos][x] >= height {
			return false
		}
	}
	fmt.Println("visible from south")
	return true
}

func checkVizEast(height, x, y int, grid [][]int) bool {
	for xPos := 0; xPos < x; xPos++ {
		if grid[y][xPos] >= height {
			return false
		}
	}
	fmt.Println("visible from east")
	return true
}

func checkVizWest(height, x, y int, grid [][]int) bool {
	for xPos := x + 1; xPos < len(grid[0]); xPos++ {
		if grid[y][xPos] >= height {
			return false
		}
	}
	fmt.Println("visible from west")
	return true
}
