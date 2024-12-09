package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() string {
	// input :=
	// 	`2333133121414131402`
	// input := `788998581870543740644523521555749467589`
	input, _ := os.ReadFile("input.txt")

	return string(input)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func assembleMemory(parsedInput string, isTask2 bool) []string {
	memory := []string{}

	//assemble memory
	for i, num := range parsedInput {
		if i%2 == 0 {
			id := (i + 1) / 2
			char := strconv.Itoa(id)
			numInt, _ := strconv.Atoi(string(num))
			if isTask2 {
				memory = append(memory, strings.Repeat(char, numInt))
			} else {
				for j := 0; j < numInt; j++ {
					memory = append(memory, char)
				}
			}
		} else {
			numInt, _ := strconv.Atoi(string(num))
			if isTask2 {
				memory = append(memory, strings.Repeat(".", numInt))
			} else {
				for j := 0; j < numInt; j++ {
					memory = append(memory, ".")
				}
			}
		}
	}

	return memory
}

func assembleMemory2(parsedInput string) [][]string {
	memory := [][]string{}

	//assemble memory
	for i, num := range parsedInput {
		numInt, _ := strconv.Atoi(string(num))
		if numInt == 0 {
			continue
		}
		if i%2 == 0 {
			id := (i + 1) / 2
			char := strconv.Itoa(id)

			group := []string{}
			for j := 0; j < numInt; j++ {
				group = append(group, string(char))
			}
			memory = append(memory, group)

		} else {
			group := []string{}
			for j := 0; j < numInt; j++ {
				group = append(group, ".")
			}
			memory = append(memory, group)

		}
	}

	return memory
}

func Task1() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)

	memory := assembleMemory(parsedInput, false)

	fmt.Println(memory)

	nextInt := 0

	//sort memory
	for i := len(memory) - 1; i >= 0; i-- {
		for j := nextInt; j < len(memory); j++ {
			if string(memory[j]) == "." && string(memory[i]) != "." {
				memory[i], memory[j] = memory[j], memory[i]
				nextInt = j
				break
			}
		}
		if !contains(memory[:i], ".") {
			break
		}
		// fmt.Println(memory)
	}

	fmt.Println(memory)

	//count memory
	count := 0
	for i, char := range memory {
		if string(char) != "." {
			num, _ := strconv.Atoi(string(char))
			count += num * i
		}
	}
	fmt.Println(count)
}

func Task2() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)

	memory := assembleMemory2(parsedInput)

	fmt.Println(memory)

	//sort by groups
	for i := len(memory) - 1; i >= 0; i-- {
		if contains(memory[i], ".") {
			continue
		}

		for j := 0; j < i; j++ {
			if contains(memory[j], ".") {
				if len(memory[j]) == len(memory[i]) {
					memory[i], memory[j] = memory[j], memory[i]
					break
				} else if len(memory[j]) > len(memory[i]) {

					tmp := memory[i]
					takenSpace := memory[j][:len(memory[i])]
					remainingSpace := memory[j][len(memory[i]):]
					memory[i] = takenSpace
					memory[j] = remainingSpace
					//add memory[i] to memory at index j
					memory = append(memory[:j], append([][]string{tmp}, memory[j:]...)...)
					i++
					break
				}
			}
		}
	}

	fmt.Println(memory)

	count := 0
	total := 0

	for i, group := range memory {
		fmt.Println(group, i)
		for _, char := range group {
			if string(char) != "." {
				num, _ := strconv.Atoi(string(char))
				total += num * count
			}
			count++
		}
	}

	fmt.Println(total)
}

func main() {
	// fmt.Println("Task 1:")
	// Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
