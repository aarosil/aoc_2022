package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("1_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	elves := make([]int, 1)
	elf := 0
	fattest := 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" {
			elf++
			elves = append(elves, 0)
			continue
		}

		amount, _ := strconv.Atoi(text)
		elves[elf] += amount

		if elves[elf] > elves[fattest] {
			fmt.Println(elves[elf], fattest)
			fattest = elf
		}

	}

	fmt.Println(elves[fattest])
	sort.Ints(elves)
	top3 := 0
	for i := 1; i < 4; i++ {
		fmt.Println("asdf", i)
		top3 = top3 + elves[len(elves)-i]
	}

	fmt.Println(top3)
}
