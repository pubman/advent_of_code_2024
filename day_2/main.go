package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 	input :=
	// 		`7 6 4 2 1
	// 1 2 7 8 9
	// 9 7 6 2 1
	// 1 3 2 4 5
	// 8 6 4 4 1
	// 1 3 6 7 9`

	input, _ := os.ReadFile("input.txt")

	count := 0

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		partsInt := make([]int, len(parts))
		//Convert the parts to ints
		for i, part := range parts {
			partsInt[i], _ = strconv.Atoi(part)
		}
		fmt.Println(partsInt)

		if isSafe(partsInt) {
			count++
		} else {
			fmt.Println("not safe", partsInt)
			valid := removeIndexes(partsInt)
			if valid {
				count++
			}
		}
	}
	fmt.Println(count)
}

func removeIndexes(partsInt []int) bool {
	for i := 0; i < len(partsInt); i++ {
		//remove the i index
		result := make([]int, 0)
		result = append(result, partsInt[:i]...)
		result = append(result, partsInt[i+1:]...)
		fmt.Println(result)
		valid := isSafe(result)
		if valid {
			return true
		}
	}
	return false
}

func isSafe(partsInt []int) bool {
	valid := true
	isIncreasing := true
	if partsInt[0] > partsInt[1] {
		isIncreasing = false
	}

	for i := 0; i < len(partsInt)-1; i++ {
		if isIncreasing && partsInt[i] > partsInt[i+1] || !isIncreasing && partsInt[i] < partsInt[i+1] {
			valid = false
			break
		}
		diff := math.Abs(float64(partsInt[i+1] - partsInt[i]))
		if diff > 3 || diff < 1 {
			valid = false
			break
		}

	}
	return valid
}
