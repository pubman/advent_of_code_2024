package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func parseInput() [][]string {
	input :=
		`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	// input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	output := [][]string{}
	for _, line := range lines {
		output = append(output, strings.Split(line, ""))
	}
	return output
}

func Task1() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)
	//find the starting point
	for i, row := range parsedInput {
		for j, col := range row {
			if col == "^" {
				x, y, mapInput := walk(parsedInput, "UP", j, i)
				fmt.Println(x, y)
				fmt.Println(mapInput)

				//count the number of Xs
				count := 0
				for _, row := range mapInput {
					for _, col := range row {
						if col == "X" {
							count++
						}
					}
				}
				fmt.Println(count)
			}
		}
	}
}

func walk(mapInput [][]string, direction string, x int, y int) (int, int, [][]string) {
	//recursively walk the path

	//base case
	if y == 0 || y == len(mapInput)-1 || x == 0 || x == len(mapInput[0])-1 {
		mapInput[y][x] = "X"
		return x, y, mapInput
	}

	nextDirection := ""

	if direction == "UP" {
		nextDirection = "RIGHT"
		for j := y - 1; j >= 0; j-- {
			mapInput[y][x] = "X"
			if mapInput[j][x] == "#" {
				break
			}

			y--
		}
	} else if direction == "RIGHT" {
		nextDirection = "DOWN"
		for j := x + 1; j < len(mapInput[0]); j++ {
			mapInput[y][x] = "X"

			if mapInput[y][j] == "#" {
				break
			}
			x++
		}
	} else if direction == "DOWN" {

		nextDirection = "LEFT"
		for j := y + 1; j < len(mapInput); j++ {

			mapInput[y][x] = "X"
			if mapInput[j][x] == "#" {
				break
			}

			y++
		}
	} else if direction == "LEFT" {
		nextDirection = "UP"
		for j := x - 1; j >= 0; j-- {

			mapInput[y][x] = "X"
			if mapInput[y][j] == "#" {
				break
			}

			x--
		}
	}

	fmt.Println(nextDirection, x, y)

	return walk(mapInput, nextDirection, x, y)
}

func walkObstacle(mapInput [][]string, direction string, x int, y int, obstaclePlaced bool, pathHistory []string) int {
	//recursively walk the path

	//base case
	if y == 0 || y == len(mapInput)-1 || x == 0 || x == len(mapInput[0])-1 {
		mapInput[y][x] = "X"
		return 0
	}

	if slices.Contains(pathHistory, fmt.Sprintf("%d,%d,%s", x, y, direction)) {
		// fmt.Println("loop", pathHistory, x, y, direction)
		if obstaclePlaced {
			return 1
		}
		return 0
	}

	pathHistory = append(pathHistory, fmt.Sprintf("%d,%d,%s", x, y, direction))

	nextDirection := ""
	result := 0

	if direction == "UP" {
		nextDirection = "RIGHT"
		for j := y - 1; j >= 0; j-- {
			mapInput[y][x] = "X"
			if mapInput[j][x] == "#" || mapInput[j][x] == "O" {
				break
			}
			if !obstaclePlaced {
				mapInput[j][x] = "#"
				count := walkObstacle(mapInput, nextDirection, x, y, true, pathHistory)
				result += count
				if count > 0 {
					fmt.Println("result", j, y, direction)
				}
				// mapInput[j][x] = "."
			}

			y--
		}
	} else if direction == "RIGHT" {
		nextDirection = "DOWN"
		for j := x + 1; j < len(mapInput[0]); j++ {
			mapInput[y][x] = "X"

			if mapInput[y][j] == "#" || mapInput[y][j] == "O" {
				break
			}
			if !obstaclePlaced {
				mapInput[y][j] = "#"
				count := walkObstacle(mapInput, nextDirection, x, y, true, pathHistory)
				result += count
				if count > 0 {
					fmt.Println("result", x, j, direction)
				}
				// mapInput[y][j] = "."
			}
			x++
		}
	} else if direction == "DOWN" {

		nextDirection = "LEFT"
		for j := y + 1; j < len(mapInput); j++ {

			mapInput[y][x] = "X"
			if mapInput[j][x] == "#" || mapInput[j][x] == "O" {
				break
			}
			if !obstaclePlaced {
				mapInput[j][x] = "#"
				count := walkObstacle(mapInput, nextDirection, x, y, true, pathHistory)
				result += count
				if count > 0 {
					fmt.Println("result", x, j, direction)
				}
				// mapInput[j][x] = "."
			}

			y++
		}
	} else if direction == "LEFT" {
		nextDirection = "UP"
		for j := x - 1; j >= 0; j-- {

			mapInput[y][x] = "X"
			if mapInput[y][j] == "#" || mapInput[y][j] == "O" {
				break
			}
			if !obstaclePlaced {
				mapInput[y][j] = "#"
				count := walkObstacle(mapInput, nextDirection, x, y, true, pathHistory)
				result += count
				if count > 0 {
					fmt.Println("result", j, y, direction)
				}
				// mapInput[y][j] = "."
			}

			x--
		}
	}

	result += walkObstacle(mapInput, nextDirection, x, y, obstaclePlaced, pathHistory)
	return result
}

func Task2(input string) {
	m := Parse(input)

	var partTwo int
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m.At(x, y) == '.' {
				m2 := Parse(input)
				m2[y][x] = '#'

				_, loop := Simulate(m2)
				if loop {
					partTwo++
				}
			}
		}
	}

	fmt.Println(partTwo)
}

type Position struct {
	X int
	Y int
}

type PositionDirection struct {
	X    int
	Y    int
	Cell uint8
}

type Map [][]uint8

func (m Map) FindGuard() (uint8, int, int) {
	for y, row := range m {
		for x, cell := range row {
			if cell == '^' || cell == 'v' || cell == '<' || cell == '>' {
				return cell, x, y
			}
		}
	}

	panic("guard not found")
}

func (m Map) At(x, y int) uint8 {
	if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
		return 0
	}

	return m[y][x]
}

func Parse(input string) Map {
	var m Map

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		m = append(m, []uint8(line))
	}

	return m
}

func Simulate(m Map) (int, bool) {
	var positions = make(map[Position]bool)
	var positionsAndDirections = make(map[PositionDirection]bool)

	cell, x, y := m.FindGuard()

	for x >= 0 && x < len(m[0]) && y >= 0 && y < len(m) {
		// infinite loop
		if positionsAndDirections[PositionDirection{x, y, cell}] {
			return -1, true
		}

		positions[Position{x, y}] = true
		positionsAndDirections[PositionDirection{x, y, cell}] = true

		// Turn if needed
		turn := true
		for turn {
			turn = false

			switch cell {
			case '^':
				if m.At(x, y-1) == '#' {
					cell = '>'
					turn = true
				}
			case '>':
				if m.At(x+1, y) == '#' {
					cell = 'v'
					turn = true
				}
			case 'v':
				if m.At(x, y+1) == '#' {
					cell = '<'
					turn = true
				}
			case '<':
				if m.At(x-1, y) == '#' {
					cell = '^'
					turn = true
				}
			default:
				panic("invalid cell")
			}
		}

		// Move
		switch cell {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		default:
			panic("invalid cell")
		}
	}

	return len(positions), false
}

func main() {
	fmt.Println("Task 1:")
	Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	Task2(string(input))
}
