package main

import (
	"fmt"
	"strings"
)

func parseInput() []string {
	input :=
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	// input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	return lines
}

func Task1() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)
}

func Task2() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)
}

func main() {
	fmt.Println("Task 1:")
	Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
