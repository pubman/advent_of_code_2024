package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput() []string {
	// input :=
	// 	`125 17`
	input := "872027 227 18 9760 0 4 67716 9245696"
	// input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), " ")

	return lines
}

func Task1() {
	stones := parseInput()
	fmt.Println(stones)

	blinks := 25

	for i := 0; i < blinks; i++ {
		newStones := []string{}
		for _, stone := range stones {
			newStones = append(newStones, stoneLogic(stone)...)
		}
		fmt.Println(newStones)
		stones = newStones
	}
	fmt.Println(len(stones))
}

var lookupTable = map[string]int{}

func stoneLogic(stone string) []string {
	newStones := []string{}

	// if stone == "0" or stone is repeated 0s
	if stone == "0" || strings.Count(stone, "0") == len(stone) {
		return []string{"1"}
	} else if len(stone)%2 == 0 {
		left := stone[:len(stone)/2]
		right := stone[len(stone)/2:]

		//trim 0s from left and right
		if len(left) > 1 {
			left = strings.TrimLeft(left, "0")
		}
		if len(right) > 1 {
			right = strings.TrimLeft(right, "0")
		}

		newStones = append(newStones, left, right)
	} else {
		stoneInt, _ := strconv.Atoi(stone)
		newStones = append(newStones, strconv.Itoa(stoneInt*2024))
	}
	return newStones
}

func Task2() {
	stones := parseInput()
	fmt.Println(stones)

	blinks := 75

	count := 0
	for _, stone := range stones {
		count += stoneCountRecursive(stone, blinks)
	}
	fmt.Println(count)
}

func stoneCountRecursive(stone string, blinks int) int {
	if blinks == 0 {
		// lookupTable[stone+strconv.Itoa(blinks)] = 1
		return 1
	}

	lookup, ok := lookupTable[stone+","+strconv.Itoa(blinks)]
	if ok {
		return lookup
	}

	result := stoneLogic(stone)

	count := 0
	for _, newStone := range result {
		stoneCount := stoneCountRecursive(newStone, blinks-1)
		count += stoneCount
	}
	lookupTable[stone+","+strconv.Itoa(blinks)] = count

	return count
}

func main() {
	// fmt.Println("Task 1:")
	// Task1()
	fmt.Println("--------------------------------")
	fmt.Println("Task 2:")
	Task2()
}
