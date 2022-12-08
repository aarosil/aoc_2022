package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println("AoC Day 8")
}

func main() {
	readFile, _ := os.Open("8_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		fmt.Println(text)
	}
}
