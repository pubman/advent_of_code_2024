package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	const i = 0
	// if statement
	if i == 0 {
		fmt.Println("i is 0")
	}

	// switch statement
	switch i {
	case 0:
		fmt.Println("i is 0")
	}

	// anonymous function
	func() {
		fmt.Println("Hello, World!")
	}()

	// anonymous function with return value
	func() int {
		return 0
	}()

	// map
	freq := make(map[int]int)

	freq[1] = 1
	freq[2] = 2

	fmt.Println(freq)

	// slice
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)

	// sort
	sort.Ints(slice)

	fmt.Println(slice)
}
