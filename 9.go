package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println("AoC Day 9")
}

func main() {
	readFile, _ := os.Open("9_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		fmt.Println(text)
	}
}
