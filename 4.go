package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var values = map[string]int{}

func init() {}

func main() {
	readFile, _ := os.Open("4_input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		ranges := strings.Split(text, ",")
		rangeA := [][]int{{}, {}}
		for i, val := range ranges {
			limits := strings.Split(val, "-")
			for _, v := range limits {
				d, _ := strconv.Atoi(v)
				rangeA[i] = append(rangeA[i], d)
			}
		}

		for i, arr := range rangeA {
			new := []int{}
			for i := arr[0]; i <= arr[1]; i++ {
				new = append(new, i)
			}
			rangeA[i] = new
		}

		fmt.Println(rangeA)
		result += arraysOverlap(rangeA[0], rangeA[1])
	}

	fmt.Println(result)
}

func arraysOverlap(a, b []int) int {
	// ...456
	// .....678
	// -- OR --
	// .234...
	// 12....
	// if a[len(a)-1] <= b[0] ||
	// 	a[0] >= b[len(b)-1] {
	// 	return 1
	// }
	// return 0

	for _, v := range a {
		for _, w := range b {
			if w == v {
				return 1
			}
		}
	}

	return 0
}
