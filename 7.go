package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("AoC Day 7")
}

type Dir struct {
	name      string
	subdirs   map[string]*Dir
	parent    *Dir
	files     map[string]int
	totalSize int
}

var resultMap = map[string]int{}

func main() {
	readFile, _ := os.Open("7_input_s.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	root := &Dir{
		name:    "/",
		files:   map[string]int{},
		subdirs: map[string]*Dir{},
	}
	currentDir := root

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if match, _ := regexp.Match(`^\$ ls`, []byte(text)); match {
			continue
		}
		if match, _ := regexp.Match(`^\$ cd`, []byte(text)); match {
			name := parseNavigateLine(text)
			if currentDir.name == name {
				continue
			}
			if name == ".." {
				currentDir = currentDir.parent
				continue
			}
			currentDir = navigate(currentDir, name)
			continue
		}
		if match, _ := regexp.Match("^dir", []byte(text)); match {
			name := parseDirLine(text)
			currentDir = addSubdir(currentDir, name)
			continue
		}

		name, size := parseFileLine(text)
		currentDir = addFile(currentDir, name, size)

	}

	calculateTotalDirSize(root)
	result := 0

	for _, v := range resultMap {
		result += v
	}
	fmt.Println(result, resultMap)
}

func parseFileLine(text string) (string, int) {
	arr := strings.Split(text, " ")
	name := arr[1]
	size, _ := strconv.Atoi(arr[0])
	return name, size
}

func parseDirLine(text string) string {
	arr := strings.Split(text, " ")
	return arr[1]
}

func parseNavigateLine(text string) string {
	arr := strings.Split(text, " ")
	return arr[2]
}

func navigate(current *Dir, next string) *Dir {
	return current.subdirs[next]
}

func addFile(dir *Dir, name string, size int) *Dir {
	dir.files[name] = size
	return dir
}

func addSubdir(dir *Dir, name string) *Dir {
	if _, exists := dir.subdirs[name]; exists {
		return dir
	}
	newDir := &Dir{
		name:    name,
		files:   map[string]int{},
		subdirs: map[string]*Dir{},
		parent:  dir,
	}

	dir.subdirs[name] = newDir
	return dir
}

func calculateTotalDirSize(dir *Dir) int {
	result := 0
	for _, size := range dir.files {
		result += size
	}

	for _, subdir := range dir.subdirs {
		result += calculateTotalDirSize(subdir)
	}

	if result < 100000 {
		resultMap[dir.name] = result
	}
	return result
}

func accumulateDiskTotal(dir *Dir) int {
	result := 0
	return result
}
