package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("AoC Day 9")
}

func main() {
	readFile, _ := os.Open("9_input_s.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	r := rope{
		head: position{x: 0, y: 0},
		tail: position{x: 0, y: 0},
	}

	for fileScanner.Scan() {
		text := fileScanner.Text()
		fmt.Println(text)
		arr := strings.Split(text, " ")
		dir := arr[0]
		amount, _ := strconv.Atoi(arr[1])
		for i := 0; i < amount; i++ {
			r.move(dir)
			fmt.Println(r)
		}
	}
}

type position struct {
	x int
	y int
}

type rope struct {
	head position
	tail position
}

var tailPositions []position

func (r *rope) move(dir string) {
	new := r.head
	switch dir {
	case "R":
		new.x += 1
	case "L":
		new.x -= 1
	case "U":
		new.y += 1
	case "D":
		new.y -= 1
	}
	r.head = new
}
