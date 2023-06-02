package main

import "fmt"

func WriteSpace(amount int) {
	for index := 0; index < amount; {
		fmt.Println("")
		index++
	}
}
