package main

import (
	"fmt"
	"strconv"
	"strings"
)

func join() {
	hello := "Hello"
	world := "World"
	// Join operates on a "slice" of strings.
	// It does not accept individual strings as arguments.
	// It also requires argument for a separator that will be placed between elements.
	// Join returns a new string.
	helloWorldSlice := []string{
		hello,
		world,
	}
	// This does not work because the return value is not capture!
	strings.Join(helloWorldSlice, "---")
	helloWorld := strings.Join(helloWorldSlice, "---")
	fmt.Println(helloWorld)

	// Trying another example
	first10primes := [10]int{
		2,
		3,
		5,
		7,
		11,
		13,
		17,
		19,
		23,
		29,
	}
	var first10primesAsStrings [10]string
	for i := 0; i < 10; i++ {
		first10primesAsStrings[i] = strconv.Itoa(first10primes[i])
	}
	var first10primesAsStringsSlice = first10primesAsStrings[:10]
	var first10primesString = strings.Join(first10primesAsStringsSlice, ", ")
	fmt.Println(first10primesString)
}
