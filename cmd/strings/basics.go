package main

import "fmt"

func basics() {

	// This is a raw literal.
	fmt.Println(`This should print "Hello World"`)

	// Printing everything including escape characters.
	fmt.Println(`The tab characters\tare also printed!\t`)

	// Multi-line strings without hackery
	multiLineLiteral :=
		`This should be
	on separate lines. Maybe it is. 
	Maybe it isn't!`
	fmt.Println(multiLineLiteral)

	// This is an interpreted literal.
	// Can be formatted and recognize escape characters.
	fmt.Println("\nthis\tis\tformatted!")

	// Concatenating strings in go.
	hello := "Hello"
	world := ` World`
	fmt.Println(hello + world)
}
