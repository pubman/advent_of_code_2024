package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extractMulOperations(text string, includeConditionals bool) []string {
	// Create regex pattern for mul(number,number)

	pattern := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	conditionalPattern := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)|do\(\)|don't\(\)`)
	// Find all matches
	var matches []string
	if includeConditionals {
		matches = conditionalPattern.FindAllString(text, -1)
	} else {
		matches = pattern.FindAllString(text, -1)
	}
	return matches
}

func parseAndMultiply(text string) int {
	//remove the mul( and )
	text = strings.Replace(text, "mul(", "", 1)
	text = strings.Replace(text, ")", "", 1)
	parts := strings.Split(text, ",")
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])
	return num1 * num2
}

func Task1() {
	// input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	input, _ := os.ReadFile("input.txt")
	matches := extractMulOperations(string(input), false)
	total := 0
	for _, match := range matches {
		total += parseAndMultiply(match)
	}
	fmt.Println(total)
}

func Task2() {
	input, _ := os.ReadFile("input.txt")
	// input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	matches := extractMulOperations(string(input), true)
	fmt.Println(matches)

	total := 0
	multiplyingEnabled := true
	for _, match := range matches {
		if strings.Contains(match, "don't") {
			multiplyingEnabled = false
		} else if strings.Contains(match, "do") {
			multiplyingEnabled = true
		} else {
			if multiplyingEnabled {
				total += parseAndMultiply(match)
			}
		}
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("Task 1:")
	Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
