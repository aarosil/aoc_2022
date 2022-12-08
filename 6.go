package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	fmt.Println("AoC Day 6")
}

func main() {
	readFile, _ := os.Open("6_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	stream := []string{}

	for fileScanner.Scan() {
		text := fileScanner.Text()
		stream = strings.Split(text, "")
	}

	result := 0
	for i := 14; i < len(stream); i++ {
		sub := stream[i-14 : i]
		uniq := allElementsUnique(sub)
		if uniq {
			result = i
			break
		}
	}

	fmt.Println(result)
}

func allElementsUnique(arr []string) bool {
	keys := make(map[string]bool)

	for _, entry := range arr {
		if keys[entry] {
			return false
		}

		keys[entry] = true
	}

	return true
}
