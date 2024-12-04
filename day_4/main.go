package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput() []string {
	// 	input :=
	// 		`MMMSXXMASM
	// MSAMXMSMSA
	// AMXSXMAAMM
	// MSAMASMSMX
	// XMASAMXAMM
	// XXAMMXXAMA
	// SMSMSASXSS
	// SAXAMASAAA
	// MAMMMXMMMM
	// MXMXAXMASX`
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	return lines
}

func Task1() {
	wordsearchLines := parseInput()
	targetWord := "XMAS"
	count := 0

	for i := 0; i < len(wordsearchLines); i++ {
		for j := 0; j < len(wordsearchLines[i]); j++ {
			count += find(wordsearchLines, i, j, "UP", targetWord, 0)
			count += find(wordsearchLines, i, j, "DOWN", targetWord, 0)
			count += find(wordsearchLines, i, j, "LEFT", targetWord, 0)
			count += find(wordsearchLines, i, j, "RIGHT", targetWord, 0)
			count += find(wordsearchLines, i, j, "UP_LEFT", targetWord, 0)
			count += find(wordsearchLines, i, j, "UP_RIGHT", targetWord, 0)
			count += find(wordsearchLines, i, j, "DOWN_LEFT", targetWord, 0)
			count += find(wordsearchLines, i, j, "DOWN_RIGHT", targetWord, 0)
		}
	}
	fmt.Println(count)
}

func find(s []string, i int, j int, direction string, targetWord string, charPos int) int {
	count := 0
	if s[i][j] == targetWord[charPos] {
		if charPos == len(targetWord)-1 {
			return 1
		} else {
			switch direction {
			case "UP":
				if i > 0 {
					count += find(s, i-1, j, direction, targetWord, charPos+1)
				}
			case "DOWN":
				if i < len(s)-1 {
					count += find(s, i+1, j, direction, targetWord, charPos+1)
				}
			case "LEFT":
				if j > 0 {
					count += find(s, i, j-1, direction, targetWord, charPos+1)
				}
			case "RIGHT":
				if j < len(s[i])-1 {
					count += find(s, i, j+1, direction, targetWord, charPos+1)
				}
			case "UP_LEFT":
				if i > 0 && j > 0 {
					count += find(s, i-1, j-1, direction, targetWord, charPos+1)
				}
			case "UP_RIGHT":
				if i > 0 && j < len(s[i])-1 {
					count += find(s, i-1, j+1, direction, targetWord, charPos+1)
				}
			case "DOWN_LEFT":
				if i < len(s)-1 && j > 0 {
					count += find(s, i+1, j-1, direction, targetWord, charPos+1)
				}
			case "DOWN_RIGHT":
				if i < len(s)-1 && j < len(s[i])-1 {
					count += find(s, i+1, j+1, direction, targetWord, charPos+1)
				}
			}
		}

	}
	return count
}

func Task2() {
	parsedInput := parseInput()
	count := 0

	for i := 1; i < len(parsedInput)-1; i++ {
		for j := 1; j < len(parsedInput[i])-1; j++ {
			if string(parsedInput[i][j]) == "A" {
				//Check diagonals
				if (string(parsedInput[i-1][j-1]) == "M" && string(parsedInput[i+1][j+1]) == "S" ||
					string(parsedInput[i+1][j+1]) == "M" && string(parsedInput[i-1][j-1]) == "S") && (string(parsedInput[i+1][j-1]) == "M" && string(parsedInput[i-1][j+1]) == "S" ||
					string(parsedInput[i-1][j+1]) == "M" && string(parsedInput[i+1][j-1]) == "S") {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	fmt.Println("Task 1:")
	Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
