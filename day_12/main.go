package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseInput() [][]string {
	// 	input :=
	// 		`AAAAAA
	// AAABBA
	// AAABBA
	// ABBAAA
	// ABBAAA
	// AAAAAA`
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	chars := [][]string{}
	for _, line := range lines {
		chars = append(chars, strings.Split(line, ""))
	}
	return chars
}

func Task1() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)

	total := 0

	for x := 0; x < len(parsedInput); x++ {
		for y := 0; y < len(parsedInput[0]); y++ {
			perimeter := explore(parsedInput, x, y, parsedInput[x][y])
			if perimeter > 0 {
				//calculate price
				total += area[parsedInput[x][y]] * perimeter
				area[parsedInput[x][y]] = 0
			}
		}
	}
	fmt.Println(area)
	fmt.Println(total)
}

var visited = map[string]bool{}
var area = map[string]int{}

func explore(chars [][]string, x, y int, char string) int {
	if x < 0 || y < 0 || x >= len(chars) || y >= len(chars[0]) {
		return 1
	}
	currentChar := chars[x][y]
	if currentChar != char {
		return 1
	}

	if visited[fmt.Sprintf("%d,%d", x, y)] {
		return 0
	}

	visited[fmt.Sprintf("%d,%d", x, y)] = true

	area[currentChar]++

	localPerimeter := 0
	localPerimeter += explore(chars, x+1, y, char)
	localPerimeter += explore(chars, x-1, y, char)
	localPerimeter += explore(chars, x, y+1, char)
	localPerimeter += explore(chars, x, y-1, char)

	return localPerimeter
}

func Task2() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)

	total := 0

	for x := 0; x < len(parsedInput); x++ {
		for y := 0; y < len(parsedInput[0]); y++ {
			sideCount := 0
			upSides = map[int][]int{}
			downSides = map[int][]int{}
			leftSides = map[int][]int{}
			rightSides = map[int][]int{}
			exploreSides(parsedInput, x, y, parsedInput[x][y], "")
			if len(upSides) == 0 && len(downSides) == 0 && len(leftSides) == 0 && len(rightSides) == 0 {
				continue
			}
			//sort xSides
			sideCount += sortAndCount(upSides)
			sideCount += sortAndCount(downSides)
			sideCount += sortAndCount(leftSides)
			sideCount += sortAndCount(rightSides)

			fmt.Println(sideCount)
			total += sideCount * area[parsedInput[x][y]]
			area[parsedInput[x][y]] = 0
		}
	}

	fmt.Println(total)
	fmt.Println(area)
}

func sortAndCount(sides map[int][]int) int {
	for x := range sides {
		sort.Ints(sides[x])
	}
	fmt.Println(sides)
	sideCount := 0
	for x := range sides {
		sideCount++
		for i := 0; i < len(sides[x])-1; i++ {
			if sides[x][i+1]-sides[x][i] != 1 {
				sideCount++
			}
		}
	}
	return sideCount
}

var upSides = map[int][]int{}
var downSides = map[int][]int{}
var leftSides = map[int][]int{}
var rightSides = map[int][]int{}

func addToSide(x, y int, dir string) {
	if dir == "up" {
		upSides[x] = append(upSides[x], y)
	} else if dir == "down" {
		downSides[x] = append(downSides[x], y)
	} else if dir == "left" {
		leftSides[y] = append(leftSides[y], x)
	} else if dir == "right" {
		rightSides[y] = append(rightSides[y], x)
	}
}

func exploreSides(chars [][]string, x, y int, char string, dir string) int {
	if x < 0 || y < 0 || x >= len(chars) || y >= len(chars[0]) {
		addToSide(x, y, dir)
		return 1
	}
	currentChar := chars[x][y]
	if currentChar != char {
		addToSide(x, y, dir)
		return 1
	}

	if visited[fmt.Sprintf("%d,%d", x, y)] {
		return 0
	}

	visited[fmt.Sprintf("%d,%d", x, y)] = true

	area[currentChar]++

	localPerimeter := 0
	localPerimeter += exploreSides(chars, x+1, y, char, "up")
	localPerimeter += exploreSides(chars, x-1, y, char, "down")
	localPerimeter += exploreSides(chars, x, y+1, char, "right")
	localPerimeter += exploreSides(chars, x, y-1, char, "left")

	return localPerimeter
}

func main() {
	// fmt.Println("Task 1:")
	// Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
