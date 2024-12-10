package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	// 	input := `89010123
	// 78121874
	// 87430965
	// 96549874
	// 45678903
	// 32019012
	// 01329801
	// 10456732`
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	grid := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			if char == '.' {
				row = append(row, -1)
			} else {
				num, _ := strconv.Atoi(string(char))
				row = append(row, num)
			}
		}
		grid = append(grid, row)
	}
	return grid
}

func Task1() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)

	start := []int{0, 0}
	count := 0

	for i, row := range parsedInput {
		//find the starting point, 0
		for j, num := range row {
			if num == 0 {
				start = []int{i, j}
				fmt.Println("start", start)
				visited = [][]int{}

				paths := [][][]int{}
				paths = append(paths, [][]int{start})
				paths = walk(parsedInput, start, [][]int{start}, paths, false)
				fmt.Println(len(visited))
				count += len(visited)
			}
		}
	}
	fmt.Println("count", count)

}

var visited = [][]int{}

func walk(grid [][]int, start []int, path [][]int, paths [][][]int, isAllPaths bool) [][][]int {
	//recursive dfs to find the easiest path

	//if we are at the peak, return the path
	if grid[start[0]][start[1]] == 9 {

		if isAllPaths {
			fmt.Println("peak detected", start)
			visited = append(visited, start)
		} else {
			if !contains(visited, start) {
				fmt.Println("peak detected", start)
				visited = append(visited, start)
			}
		}
		paths = append(paths, path)
		return paths
	}

	//select the next point from up, down, left, right
	next := [][]int{}
	options := [][]int{}
	options = append(options, []int{start[0] - 1, start[1]})
	options = append(options, []int{start[0] + 1, start[1]})
	options = append(options, []int{start[0], start[1] - 1})
	options = append(options, []int{start[0], start[1] + 1})

	//check options
	for _, option := range options {
		//check if the option is out of bounds
		if option[0] < 0 || option[0] >= len(grid) || option[1] < 0 || option[1] >= len(grid[0]) {
			continue
		}

		//check if the option is already in the path
		if contains(path, option) {
			continue
		}

		//check if the option is higher than the current point
		if grid[option[0]][option[1]] > grid[start[0]][start[1]] {
			heightDiff := grid[option[0]][option[1]] - grid[start[0]][start[1]]
			if heightDiff == 1 {
				next = append(next, option)
			}
		}
	}

	if len(next) == 0 {
		return paths
	}
	for _, option := range next {
		paths = walk(grid, option, append(path, option), paths, isAllPaths)
	}
	return paths

}

func contains(path [][]int, point []int) bool {
	for _, p := range path {
		if p[0] == point[0] && p[1] == point[1] {
			return true
		}
	}
	return false
}

func Task2() {
	parsedInput := parseInput()
	fmt.Println(parsedInput)

	start := []int{0, 0}
	count := 0

	for i, row := range parsedInput {
		//find the starting point, 0
		for j, num := range row {
			if num == 0 {
				start = []int{i, j}
				fmt.Println("start", start)
				visited = [][]int{}

				paths := [][][]int{}
				paths = append(paths, [][]int{start})
				paths = walk(parsedInput, start, [][]int{start}, paths, true)
				count += len(visited)
			}
		}
	}
	fmt.Println("count", count)

}

func main() {
	// fmt.Println("Task 1:")
	// Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
