package main

import (
	"fmt"
	"sort"
	"strings"
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

	// dynamic programming fibonacci
	fib := make(map[int]int)
	fib[0] = 0
	fib[1] = 1

	n := 10
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println(fib[n])

	//contains
	fmt.Println(strings.Contains("Hello, World!", "World"))
}
