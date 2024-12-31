package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Move struct {
	x, y int
}

var (
	UpMove    Move = Move{0, -1}
	RightMove Move = Move{1, 0}
	DownMove  Move = Move{0, 1}
	LeftMove  Move = Move{-1, 0}
)

func parseInput() []string {
	// 	input :=
	// 		`###############
	// #.......#....E#
	// #.#.###.#.###.#
	// #.....#.#...#.#
	// #.###.#####.#.#
	// #.#.#.......#.#
	// #.#.#####.###.#
	// #...........#.#
	// ###.#.#####.#.#
	// #...#.....#.#.#
	// #.#.#.###.#.#.#
	// #.....#...#.#.#
	// #.###.#.#.#.#.#
	// #S..#.....#...#
	// ###############`
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	return lines
}

var pathCache = make(map[string]int)

type State struct {
	point     Point
	direction Move
	cost      int
	path      []Point
}

type CacheKey struct {
	point     Point
	direction Move
}

func dijkstra(maze []string, start Point, end Point) (int, [][]Point) {
	pq := make([]State, 0)
	pq = append(pq, State{start, RightMove, 0, []Point{start}})

	visited := make(map[CacheKey]int)
	minCost := math.MaxInt
	var allShortestPaths [][]Point

	for len(pq) > 0 {
		current := pq[0]
		pq = pq[1:]

		// If we've found a shorter path already, skip this one
		if current.cost > minCost {
			continue
		}

		cacheKey := CacheKey{current.point, current.direction}

		// Allow equal cost paths to continue
		if cost, exists := visited[cacheKey]; exists && cost < current.cost {
			continue
		}

		visited[cacheKey] = current.cost

		if string(maze[current.point.y][current.point.x]) == "E" {
			if current.cost < minCost {
				// Found a new shortest path, clear previous paths
				minCost = current.cost
				allShortestPaths = [][]Point{current.path}
			} else if current.cost == minCost {
				// Found another path with same minimum cost
				allShortestPaths = append(allShortestPaths, current.path)
			}
			continue
		}

		if string(maze[current.point.y][current.point.x]) == "#" {
			continue
		}

		right := rotate90(current.direction)
		left := rotate270(current.direction)

		nextPoint := Point{current.point.x + current.direction.x, current.point.y + current.direction.y}
		if isValidPoint(maze, nextPoint) {
			newPath := make([]Point, len(current.path))
			copy(newPath, current.path)
			newPath = append(newPath, nextPoint)
			pq = insertSorted(pq, State{nextPoint, current.direction, current.cost + 1, newPath})
		}

		rightPoint := Point{current.point.x + right.x, current.point.y + right.y}
		if isValidPoint(maze, rightPoint) {
			newPath := make([]Point, len(current.path))
			copy(newPath, current.path)
			newPath = append(newPath, rightPoint)
			pq = insertSorted(pq, State{rightPoint, right, current.cost + 1001, newPath})
		}

		leftPoint := Point{current.point.x + left.x, current.point.y + left.y}
		if isValidPoint(maze, leftPoint) {
			newPath := make([]Point, len(current.path))
			copy(newPath, current.path)
			newPath = append(newPath, leftPoint)
			pq = insertSorted(pq, State{leftPoint, left, current.cost + 1001, newPath})
		}
	}

	return minCost, allShortestPaths
}

func isValidPoint(maze []string, p Point) bool {
	return p.y >= 0 && p.y < len(maze) && p.x >= 0 && p.x < len(maze[0])
}

func insertSorted(pq []State, state State) []State {
	i := 0
	for i < len(pq) && pq[i].cost < state.cost {
		i++
	}
	pq = append(pq, State{})
	copy(pq[i+1:], pq[i:])
	pq[i] = state
	return pq
}

func solveMaze(maze []string, point Point, score int, direction Move, visited map[Point]bool) int {
	// Create a copy of the visited map for this path
	currentVisited := make(map[Point]bool)
	for k, v := range visited {
		currentVisited[k] = v
	}

	cacheKey := fmt.Sprintf("%d,%d,%d,%d", point.x, point.y, direction.x, direction.y)
	// if cachedScore, ok := pathCache[cacheKey]; ok {
	// 	// return cachedScore
	// }

	if string(maze[point.y][point.x]) == "E" {
		fmt.Println("End of maze", score)
		printMaze(maze, currentVisited)
		return score
	}

	if string(maze[point.y][point.x]) == "#" || currentVisited[point] {
		return math.MaxInt
	}

	currentVisited[point] = true

	//check positions left and right
	right := rotate90(direction)
	left := rotate270(direction)

	nextPoint := Point{point.x + direction.x, point.y + direction.y}
	leftPoint := Point{point.x + left.x, point.y + left.y}
	rightPoint := Point{point.x + right.x, point.y + right.y}

	// Try all possible moves and keep track of the maximum score
	nextScore := solveMaze(maze, nextPoint, score+1, direction, currentVisited)
	rightScore := solveMaze(maze, rightPoint, score+1001, right, currentVisited)
	leftScore := solveMaze(maze, leftPoint, score+1001, left, currentVisited)

	minScore := min(nextScore, rightScore)
	minScore = min(minScore, leftScore)

	pathCache[cacheKey] = minScore

	return minScore
}

func printMaze(maze []string, visited map[Point]bool) {
	for y, line := range maze {
		for x, char := range line {
			if visited[Point{x, y}] {
				fmt.Print("X")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func rotate270(direction Move) Move {
	switch direction {
	case UpMove:
		return LeftMove
	case RightMove:
		return UpMove
	case DownMove:
		return RightMove
	case LeftMove:
		return DownMove
	}
	return direction
}

func rotate90(direction Move) Move {
	switch direction {
	case UpMove:
		return RightMove
	case RightMove:
		return DownMove
	case DownMove:
		return LeftMove
	case LeftMove:
		return UpMove
	}
	return direction
}

func Task1() {
	parsedInput := parseInput()

	var start, end Point
	for i, line := range parsedInput {
		if strings.Contains(line, "S") {
			start.x = strings.Index(line, "S")
			start.y = i
		}
		if strings.Contains(line, "E") {
			end.x = strings.Index(line, "E")
			end.y = i
		}
	}

	score, paths := dijkstra(parsedInput, start, end)
	fmt.Printf("Score: %d\nNumber of shortest paths: %d\n", score, len(paths))

	// Print each path
	allVisited := make(map[Point]bool)
	for i, path := range paths {
		fmt.Printf("\nPath %d:\n", i+1)
		visited := make(map[Point]bool)
		for _, p := range path {
			visited[p] = true
			allVisited[p] = true
		}
		fmt.Println(len(visited))
	}
}

func Task2() {
	parsedInput := parseInput()

	var start, end Point
	for i, line := range parsedInput {
		if strings.Contains(line, "S") {
			start.x = strings.Index(line, "S")
			start.y = i
		}
		if strings.Contains(line, "E") {
			end.x = strings.Index(line, "E")
			end.y = i
		}
	}

	score, paths := dijkstra(parsedInput, start, end)
	fmt.Printf("Score: %d\nNumber of shortest paths: %d\n", score, len(paths))

	// Print each path
	allVisited := make(map[Point]bool)
	for i, path := range paths {
		fmt.Printf("\nPath %d:\n", i+1)
		visited := make(map[Point]bool)
		for _, p := range path {
			visited[p] = true
			allVisited[p] = true
		}
		fmt.Println(len(visited))
	}
	fmt.Println(len(allVisited))
}

func main() {
	fmt.Println("Task 1:")
	Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
