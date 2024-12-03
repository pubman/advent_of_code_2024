package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// const input = `3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3`

// Read input from file
func readInput() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(content)
}

var input = readInput()

func parseInput(list1 []int, list2 []int) ([]int, []int) {
	//parse input and add to lists
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}
	return list1, list2
}

func Task1() {
	// Solution: Consecutively add up the smallest numbers in each list, then sum those values.
	var list1 []int
	var list2 []int
	list1, list2 = parseInput(list1, list2)

	//sort lists
	sort.Ints(list1)
	sort.Ints(list2)

	var sum int

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	fmt.Println(sum)
}

func Task2() {
	//Solution: create a frequency list for list 2
	var list1 []int
	var list2 []int
	list1, list2 = parseInput(list1, list2)

	freq := make(map[int]int)

	for _, num := range list2 {
		freq[num]++
	}

	totalScore := 0
	for _, num := range list1 {
		totalScore += freq[num] * num
	}

	fmt.Println(totalScore)
}

func main() {
	fmt.Println("Task 1:")
	Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
