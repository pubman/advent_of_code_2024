package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Equation struct {
	ax, ay, bx, by, px, py int
}

func extractInts(s string) []int {
	re := regexp.MustCompile(`[+-]?\d+`)
	matches := re.FindAllString(s, -1)

	result := make([]int, 0, len(matches))
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func parseInput() []Equation {
	// input :=
	// 	`Button A: X+94, Y+34
	// Button B: X+22, Y+67
	// Prize: X=8400, Y=5400

	// Button A: X+26, Y+66
	// Button B: X+67, Y+21
	// Prize: X=12748, Y=12176

	// Button A: X+17, Y+86
	// Button B: X+84, Y+37
	// Prize: X=7870, Y=6450

	// Button A: X+69, Y+23
	// Button B: X+27, Y+71
	// Prize: X=18641, Y=10279`
	input, _ := os.ReadFile("input.txt")

	output := []Equation{}

	blocks := strings.Split(string(input), "\n\n")
	for _, block := range blocks {
		blockOutput := Equation{}
		lines := strings.Split(block, "\n")
		for i, line := range lines {
			numbers := extractInts(line)
			//Switch to add equation to output
			switch i {
			case 0:
				blockOutput.ax = numbers[0]
				blockOutput.ay = numbers[1]
			case 1:
				blockOutput.bx = numbers[0]
				blockOutput.by = numbers[1]
			case 2:
				blockOutput.px = numbers[0]
				blockOutput.py = numbers[1]
			}
		}
		output = append(output, blockOutput)
	}

	return output
}

func Task1() {
	parsedInput := parseInput()

	aCost := 3
	bCost := 1
	score := 0
	for _, equation := range parsedInput {
		fmt.Println(equation)
		a, b := solveEq(equation)
		fmt.Println(a, b)
		score += a*aCost + b*bCost
	}
	fmt.Println(score)
}

// Solve equation using Cramer's rule, return 0,0 if the result is not a whole
// number.
func solveEq(eq Equation) (int, int) {
	d := eq.ax*eq.by - eq.bx*eq.ay
	d1 := eq.px*eq.by - eq.py*eq.bx
	d2 := eq.py*eq.ax - eq.px*eq.ay

	if d1%d != 0 || d2%d != 0 {
		return 0, 0
	}

	return d1 / d, d2 / d
}

func Task2() {
	equations := parseInput()
	aCost := 3
	bCost := 1
	score := 0
	for _, equation := range equations {
		equation.px += 10000000000000
		equation.py += 10000000000000
		fmt.Println(equation)
		a, b := solveEq(equation)
		fmt.Println(a, b)
		score += a*aCost + b*bCost
		fmt.Println(score)
	}
	fmt.Println(score)
}

func main() {
	fmt.Println("Task 1:")
	Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
