package main

import (
	"fmt"
)

func basics() {

	// Raw literal strings
	raw := `This is a "raw" literal string.\n`
	fmt.Println(raw)

	// Interpreted literal strings
	interpreted := "\nThis is on a new line!"
	fmt.Println(interpreted)

	// Concatenate strings
	hello := "Hello"
	world := "World"
	space := " "
	newline := "\n"
	// Print does not automatically add spaces between string operands.
	fmt.Print(hello + space + world + newline)
	// It only does so for non-string operands. Like this...
	fmt.Print(1, 2, 3, 4, " ...\n")
}
