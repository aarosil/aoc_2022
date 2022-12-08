package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var values = map[string]int{}

func init() {
	fmt.Println("AoC day 3")
	priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, v := range strings.Split(priorities, "") {
		values[v] = i + 1
	}
}

func main() {
	readFile, _ := os.Open("3_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	badgeSum := 0
	sack1, sack2, sack3 := []string{}, []string{}, []string{}
	line := 1

	for fileScanner.Scan() {
		text := fileScanner.Text()
		chars := strings.Split(text, "")
		first := chars[:len(chars)/2]
		last := chars[len(chars)/2:]
		common := findCommon2(first, last)
		val := values[common]
		// fmt.Println(val)
		sum += val
		// fmt.Println(first, last, common)
		whichSack := line % 3
		switch whichSack {
		case 0:
			sack1 = strings.Split(text, "")
		case 1:
			sack2 = strings.Split(text, "")
		case 2:
			sack3 = strings.Split(text, "")
		}

		if whichSack == 0 {
			badge := findBadge(sack1, sack2, sack3)
			badgeSum += values[badge]
		}

		line++
	}

	fmt.Println(badgeSum)
}

func findCommon2(a, b []string) string {
	result := ""
	for _, v := range a {
		for _, w := range b {
			if v == w {
				result = v
			}
		}
	}

	return result
}

func findBadge(a, b, c []string) string {
	result := ""
	for _, v := range a {
		for _, w := range b {
			for _, x := range c {
				if v == w && w == x {
					result = v
				}
			}
		}
	}

	return result
}
