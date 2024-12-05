package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func parseInput() ([]string, [][]string) {
	// 	input :=
	// 		`47|53
	// 97|13
	// 97|61
	// 97|47
	// 75|29
	// 61|13
	// 75|53
	// 29|13
	// 97|29
	// 53|29
	// 61|53
	// 97|53
	// 61|29
	// 47|13
	// 75|47
	// 97|75
	// 47|61
	// 75|61
	// 47|29
	// 75|13
	// 53|13

	// 75,47,61,53,29
	// 97,61,53,29,13
	// 75,29,13
	// 75,97,47,61,53
	// 61,13,29
	// 97,13,75,29,47`
	input, _ := os.ReadFile("input.txt")

	parts := strings.Split(string(input), "\n\n")

	rules := strings.Split(string(parts[0]), "\n")

	pages := [][]string{}
	pageLines := strings.Split(string(parts[1]), "\n")
	for _, line := range pageLines {
		page := strings.Split(strings.TrimSpace(line), ",")
		pages = append(pages, page)
	}
	return rules, pages
}

func Task1() {
	rules, pages := parseInput()

	//make a map of rules
	rulesMap := make(map[string][]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		rulesMap[parts[0]] = append(rulesMap[parts[0]], parts[1])
	}
	fmt.Println(rulesMap)

	total := 0

	//loop back through pages and check if they match any rule
	for _, page := range pages {
		valid := true
		dangerNums := []string{}
		for i := len(page) - 1; i >= 0; i-- {
			if contains(dangerNums, page[i]) {
				valid = false
				break
			}
			if rulesMap[page[i]] != nil {
				dangerNums = append(dangerNums, rulesMap[page[i]]...)
			}
		}
		if valid {
			total += parseInt(page[(len(page)-1)/2])
		}
	}
	fmt.Println(total)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func moveAfter(arr []string, i, j int) []string {
	// Handle invalid indices
	if i < 0 || j < 0 || i >= len(arr) || j >= len(arr) {
		return arr
	}

	// Save the element we want to move
	elem := arr[i]

	// Remove the element at i
	copy(arr[i:], arr[i+1:])

	// If j is before i, we need to adjust j since removing i shifted everything
	if i < j {
		j--
	}

	// Make space after j and insert elem
	copy(arr[j+2:], arr[j+1:])
	arr[j+1] = elem

	return arr
}

func Task2() {
	rules, pages := parseInput()
	fmt.Println(pages)

	//make a map of rules
	rulesMap := make(map[string][]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		rulesMap[parts[0]] = append(rulesMap[parts[0]], parts[1])
	}

	total := 0
	for _, page := range pages {
		fmt.Println(page)
		invalid := false
		for i := len(page) - 1; i >= 0; i-- {
			fmt.Println(page[i])
			//iterate through previous pages and check if any of the values are in the previous indexes
			for j := len(page) - 1; j > i; j-- {
				if contains(rulesMap[page[j]], page[i]) {
					fmt.Println(page[i], "is in", page[j])
					//place page[j] after page[i]
					page = moveAfter(page, i, j)
					fmt.Println("SWAPPED", page)
					invalid = true
					i++
					break
				}
			}
		}
		if invalid {
			total += parseInt(page[(len(page)-1)/2])
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
