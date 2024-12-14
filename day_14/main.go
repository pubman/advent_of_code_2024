package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/term"
)

type Position struct {
	x int
	y int
}

type Velocity struct {
	x int
	y int
}

type Robot struct {
	position Position
	velocity Velocity
}

type Quadrant struct {
	q1 int
	q2 int
	q3 int
	q4 int
}

func parseInput() []Robot {
	// 	input :=
	// 		`p=0,4 v=3,-3
	// p=6,3 v=-1,-3
	// p=10,3 v=-1,2
	// p=2,0 v=2,-1
	// p=0,0 v=1,3
	// p=3,0 v=-2,-2
	// p=7,6 v=-1,-3
	// p=3,0 v=-1,-2
	// p=9,3 v=2,3
	// p=7,3 v=-1,2
	// p=2,4 v=2,-3
	// p=9,5 v=-3,-3`
	input, _ := os.ReadFile("input.txt")
	pattern := `p=(-?\d+),(-?\d+)\s*v=(-?\d+),(-?\d+)`
	re := regexp.MustCompile(pattern)
	lines := strings.Split(string(input), "\n")
	robots := []Robot{}
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 5 { // original match plus 4 capture groups
			px, _ := strconv.Atoi(matches[1])
			py, _ := strconv.Atoi(matches[2])
			vx, _ := strconv.Atoi(matches[3])
			vy, _ := strconv.Atoi(matches[4])
			robots = append(robots, Robot{position: Position{x: px, y: py}, velocity: Velocity{x: vx, y: vy}})
		}
	}
	return robots
}

func Task1() {
	robots := parseInput()
	fmt.Println(robots)

	gridWidth := 101
	gridHeight := 103
	midpointX := gridWidth / 2
	midpointY := gridHeight / 2

	time := 100
	quadrant := Quadrant{q1: 0, q2: 0, q3: 0, q4: 0}

	grid := make([][]int, gridHeight)
	for i := range grid {
		grid[i] = make([]int, gridWidth)
	}

	for _, robot := range robots {
		fmt.Println(robot.position.x, robot.position.y)
		finalX := (robot.position.x + robot.velocity.x*time) % gridWidth
		finalY := (robot.position.y + robot.velocity.y*time) % gridHeight
		if finalX < 0 {
			finalX += gridWidth
		}
		if finalY < 0 {
			finalY += gridHeight
		}
		fmt.Println(finalX, finalY)
		grid[finalY][finalX] = 1
		if finalX == midpointX || finalY == midpointY {
			fmt.Println("Midpoint detected")
			continue
		}
		if finalX < midpointX {
			if finalY < midpointY {
				quadrant.q1++
			} else {
				quadrant.q2++
			}
		} else {
			if finalY < midpointY {
				quadrant.q3++
			} else {
				quadrant.q4++
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println(quadrant)
	//multiply the number of robots in each quadrant
	totalRobots := quadrant.q1 * quadrant.q2 * quadrant.q3 * quadrant.q4
	fmt.Println(totalRobots)

}

func Task2() {
	robots := parseInput()

	gridWidth := 101
	gridHeight := 103

	currentTime := 0

	_, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Failed to set terminal to raw mode:", err)
		return
	}
	// defer term.Restore(int(os.Stdin.Fd()), oldState)
	// Print initial instructions

	// buffer := make([]byte, 1)
	// for {
	// Read a single character
	// os.Stdin.Read(buffer)
	// char := buffer[0]

	// switch char {
	// case 'q':
	// 	fmt.Println("\nExiting...")
	// 	return
	// case 'n':
	// 	currentTime++
	for currentTime < 10000 {
		grid := make([][]string, gridHeight)

		for i := range grid {
			grid[i] = make([]string, gridWidth)
		}

		for _, robot := range robots {
			finalY := (robot.position.y + robot.velocity.y*currentTime) % gridHeight
			finalX := (robot.position.x + robot.velocity.x*currentTime) % gridWidth
			if finalX < 0 {
				finalX += gridWidth
			}
			if finalY < 0 {
				finalY += gridHeight
			}
			grid[finalY][finalX] = "*"
		}

		// fmt.Println(grid)

		//displayGrid(grid, currentTime)
		saveGridToPng(grid, currentTime)
		currentTime++
	}
	// }
}

func saveGridToPng(grid [][]string, currentTime int) {
	// Create images directory if it doesn't exist
	if err := os.MkdirAll("images", 0755); err != nil {
		fmt.Println("Failed to create images directory:", err)
		return
	}

	file, err := os.Create(fmt.Sprintf("images/grid_%d.png", currentTime))
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()
	img := image.NewRGBA(image.Rect(0, 0, len(grid[0]), len(grid)))
	for y, row := range grid {
		for x, cell := range row {
			if cell == "*" {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	png.Encode(file, img)
}

func clearScreen() {
	// ANSI escape code to clear screen and move cursor to top-left
	fmt.Print("\033[H\033[2J")
}

func displayGrid(grid [][]string, currentTime int) {
	for _, row := range grid {
		fmt.Println(strings.Join(row, " "))
		fmt.Print("\r\n")
	}
	fmt.Println("Time:", currentTime)
}

func main() {
	// fmt.Println("Task 1:")
	// Task1()
	// fmt.Println("--------------------------------")
	// fmt.Println("Task 2:")
	Task2()
}
