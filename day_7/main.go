package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInputOld() map[int][]int {
	input :=
		`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	// input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	targetMap := make(map[int][]int)
	for _, line := range lines {
		fmt.Println(line)
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		values := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, value := range values {
			valueInt, _ := strconv.Atoi(value)
			targetMap[target] = append(targetMap[target], valueInt)
		}
	}
	return targetMap
}

func Task1() {
	parsedInput := parseInputOld()
	fmt.Println(parsedInput)

	sumTotal := 0

	for target, values := range parsedInput {
		result := 0
		fmt.Println(target, values)
		result += calculate(target, values, 0, "+")
		result += calculate(target, values, 0, "*")

		//add the values
		if result > 0 {
			sumTotal += target
		}
	}
	fmt.Println(sumTotal)
}

func calculate(target int, values []int, count int, operator string) int {

	//recursively calculate the values
	//base case
	if len(values) == 0 {
		if count == target {
			return 1
		}
		return 0
	}

	result := 0

	//recursive case
	switch operator {
	case "+":
		count += values[0]
	case "*":
		count *= values[0]
	}

	result += calculate(target, values[1:], count, "+")
	result += calculate(target, values[1:], count, "*")

	fmt.Println("target", target, "result", result, "values", values)
	return result
}

func addSymbols(values []int, symbols string, combine bool, allSymbols []string) []string {
	//recursively add symbols to the values
	symbols += strconv.Itoa(values[0])
	//base case
	if len(values) == 1 {
		allSymbols = append(allSymbols, symbols)
		return allSymbols
	}

	//recursive case
	if combine {
		// symbols = append(symbols, "||")
	} else {
		symbols += "-"
	}

	allSymbols = addSymbols(values[1:], symbols, false, allSymbols)
	allSymbols = addSymbols(values[1:], symbols, true, allSymbols)

	return allSymbols
}

// Using a map approach
func removeDuplicates(arr []string) []string {
	//remove duplicate strings
	seen := make(map[string]bool)
	result := []string{}
	for _, str := range arr {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

var patGrammar = "[\\d]+"
var reGrammar *regexp.Regexp

func Task2() {

}

func testCalculate() {
	input := []int{6, 8, 6, 15}
	// fmt.Println(calculate2(7290, input, 0, "+"))
	// fmt.Println(calculate2(7290, input, 0, "*"))
	fmt.Println(calculate2(7290, input, 0, "||"))
}

func calculate2(target int, values []int, count int, operator string) int {

	//recursively calculate the values
	//base case
	if len(values) == 0 {
		if count == target {
			return 1
		}
		return 0
	}

	result := 0

	//recursive case
	switch operator {
	case "+":
		fmt.Println("+", values[0])
		count += values[0]
	case "*":
		fmt.Println("*", values[0])
		count *= values[0]

	case "||":
		if len(values) > 1 {
			combinedInt := strconv.Itoa(values[0]) + strconv.Itoa(values[1])
			fmt.Println(combinedInt, "||")
			values[1], _ = strconv.Atoi(combinedInt)
		}
	}

	result += calculate2(target, values[1:], count, "+")
	result += calculate2(target, values[1:], count, "*")
	result += calculate2(target, values[1:], count, "||")

	return result
}

type Equation struct {
	result   int
	operands []int
}

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}
func concat(a, b int) int {
	res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return res
}

func parseInput() []Equation {
	equations := []Equation{}
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create scanner from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		result, _ := strconv.Atoi(parts[0])
		values := strings.Split(parts[1], " ")
		operands := make([]int, len(values))
		for i, v := range values {
			o, _ := strconv.Atoi(v)
			operands[i] = o
		}
		equations = append(equations, Equation{result, operands})
	}
	return equations
}

// Depth-first search for equation rule, branching on each operation
func findOperations(equation *Equation, i int, partial int, ops []func(a, b int) int) bool {
	if partial > equation.result {
		return false
	} else if i == len(equation.operands) {
		return partial == equation.result
	} else {
		for _, op := range ops {
			if findOperations(equation, i+1, op(partial, equation.operands[i]), ops) {
				return true
			}
		}
		return false
	}
}

func main() {
	equations := parseInput()
	total1, total2 := 0, 0
	part1Ops := []func(a, b int) int{add, mul}
	part2Ops := []func(a, b int) int{add, mul, concat}

	for _, equation := range equations {
		if findOperations(&equation, 1, equation.operands[0], part1Ops) {
			total1 += equation.result
			total2 += equation.result
		} else if findOperations(&equation, 1, equation.operands[0], part2Ops) {
			total2 += equation.result
		}
	}

	fmt.Println(total1)
	fmt.Println(total2)
}
