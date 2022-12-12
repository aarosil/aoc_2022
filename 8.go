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
var highestScenicScore int

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
			scenicScore := getScenicScore(height, x, y)
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	fmt.Println(visibleCount)
	fmt.Println(highestScenicScore)
}

func getScenicScore(height, x, y int) int {
	if onEdge(x, y) {
		return 0
	}

	north := getScenicScoreNorth(height, x, y)
	south := getScenicScoreSouth(height, x, y)
	east := getScenicScoreEast(height, x, y)
	west := getScenicScoreWest(height, x, y)

	return north * south * east * west
}

func getScenicScoreNorth(height, x, y int) int {
	score := 0
	for yPos := y - 1; yPos >= 0; yPos-- {
		score += 1
		if grid[yPos][x] >= height {
			return score
		}
	}

	return score
}

func getScenicScoreSouth(height, x, y int) int {
	score := 0
	for yPos := y + 1; yPos < len(grid); yPos++ {
		score += 1
		if grid[yPos][x] >= height {
			return score
		}
	}

	return score
}

func getScenicScoreEast(height, x, y int) int {
	score := 0
	for xPos := x - 1; xPos >= 0; xPos-- {
		score += 1
		if grid[y][xPos] >= height {
			return score
		}
	}

	return score
}

func getScenicScoreWest(height, x, y int) int {
	score := 0
	for xPos := x + 1; xPos < len(grid[0]); xPos++ {
		score += 1
		if grid[y][xPos] >= height {
			return score
		}
	}
	return score
}

func onEdge(x, y int) bool {
	return x == 0 ||
		y == 0 ||
		x == len(grid[0])-1 ||
		y == len(grid)-1
}

func isVisible(height, x, y int) bool {
	if onEdge(x, y) {
		return true
	}

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
	return true
}

func checkVizSouth(height, x, y int, grid [][]int) bool {
	for yPos := y + 1; yPos < len(grid); yPos++ {
		if grid[yPos][x] >= height {
			return false
		}
	}
	return true
}

func checkVizEast(height, x, y int, grid [][]int) bool {
	for xPos := 0; xPos < x; xPos++ {
		if grid[y][xPos] >= height {
			return false
		}
	}
	return true
}

func checkVizWest(height, x, y int, grid [][]int) bool {
	for xPos := x + 1; xPos < len(grid[0]); xPos++ {
		if grid[y][xPos] >= height {
			return false
		}
	}
	return true
}
