package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput() [][]string {
	// 	input :=
	// 		`............
	// ........0...
	// .....0......
	// .......0....
	// ....0.......
	// ......A.....
	// ............
	// ............
	// ........A...
	// .........A..
	// ............
	// ............`

	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	lineCharacters := [][]string{}
	for _, line := range lines {
		lineCharacters = append(lineCharacters, strings.Split(line, ""))
	}
	return lineCharacters
}

func createEmptyCopy(arr [][]string) [][]string {
	// Create new array with same dimensions
	copy := make([][]string, len(arr))
	for i := range arr {
		copy[i] = make([]string, len(arr[i]))
		// Fill with "."
		for j := range arr[i] {
			copy[i][j] = "."
		}
	}
	return copy
}

func Task1() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)
	antennaMap := map[string][][]int{}

	//traverse the map and find the antennas
	for i, line := range parsedInput {
		for j, char := range line {
			if char != "." {
				antennaMap[char] = append(antennaMap[char], []int{i, j})
			}
		}
	}
	fmt.Println(antennaMap)

	//make a copy of the input for antinodes same size as input
	emptyInput := createEmptyCopy(parsedInput)

	//traverse the node coordinates and find the antinodes
	for _, char := range antennaMap {
		fmt.Println(char)
		for _, coordinate := range char {
			//Check against each other coordinate
			for _, otherCoordinate := range char {
				if coordinate[0] == otherCoordinate[0] && coordinate[1] == otherCoordinate[1] {
					continue
				}
				//calculate distance
				x_distance := int(coordinate[0] - otherCoordinate[0])
				y_distance := int(coordinate[1] - otherCoordinate[1])
				fmt.Println(coordinate, otherCoordinate, x_distance, y_distance)

				//Attempt to place antinode in emptyInput
				combined_x := coordinate[0] + int(x_distance)
				combined_y := coordinate[1] + int(y_distance)
				if combined_x < len(emptyInput) && combined_y < len(emptyInput[0]) && combined_x >= 0 && combined_y >= 0 {
					emptyInput[combined_x][combined_y] = "#"
				}
			}
		}
	}

	count := 0
	for _, line := range emptyInput {
		fmt.Println(strings.Join(line, ""))
		for _, char := range line {
			if char == "#" {
				count++
			}
		}
	}

	fmt.Println(count)
}

func Task2() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)
	antennaMap := map[string][][]int{}
	//traverse the map and find the antennas
	for i, line := range parsedInput {
		for j, char := range line {
			if char != "." {
				antennaMap[char] = append(antennaMap[char], []int{i, j})
			}
		}
	}
	fmt.Println(antennaMap)

	//make a copy of the input for antinodes same size as input
	emptyInput := createEmptyCopy(parsedInput)

	//traverse the node coordinates and find the antinodes
	for c, char := range antennaMap {
		fmt.Println(char)
		if c == "#" {
			continue
		}
		for _, coordinate := range char {
			//Check against each other coordinate
			for _, otherCoordinate := range char {
				if coordinate[0] == otherCoordinate[0] && coordinate[1] == otherCoordinate[1] {
					continue
				}
				emptyInput[coordinate[0]][coordinate[1]] = c
				//calculate distance
				x_distance := int(coordinate[0] - otherCoordinate[0])
				y_distance := int(coordinate[1] - otherCoordinate[1])
				fmt.Println(coordinate, otherCoordinate, x_distance, y_distance)

				//Attempt to place antinode in emptyInput
				combined_x := coordinate[0] + int(x_distance)
				combined_y := coordinate[1] + int(y_distance)
				// if combined_x < len(emptyInput) && combined_y < len(emptyInput[0]) && combined_x >= 0 && combined_y >= 0 {
				// 	emptyInput[combined_x][combined_y] = "#"
				// }
				for combined_x < len(emptyInput) && combined_y < len(emptyInput[0]) && combined_x >= 0 && combined_y >= 0 {
					emptyInput[combined_x][combined_y] = "#"
					combined_x += x_distance
					combined_y += y_distance
				}
			}
		}
	}
	count := 0
	for _, line := range emptyInput {
		fmt.Println(strings.Join(line, ""))
		for _, char := range line {
			if char != "." {
				count++
			}
		}
	}

	fmt.Println(count)
}

func main() {
	// fmt.Println("Task 1:")
	// Task1()
	// fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
