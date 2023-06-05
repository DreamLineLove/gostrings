package main

import (
	"fmt"
	"strings"
)

func explore() {
	s1 := "WWDC 2023"
	s1_lower := strings.ToLower(s1)
	sub1 := "w"
	count_sub1 := strings.Count(s1_lower, sub1)
	fmt.Println(sub1, "appears", count_sub1, "times in", s1_lower)
	count_empty := strings.Count(s1_lower, "")
	fmt.Print("Count of ", "\"", s1_lower, "\": ", len(s1_lower))
	WriteSpace(1)
	fmt.Println("Empty substring should return 1 + length of operand string!")
	fmt.Println("Empty substring appears", count_empty, "times in", s1_lower)

	s2 := "I like ice cream. I like it because it is delicious!"
	sub2 := "like"
	count2 := strings.Count(s2, sub2)
	fmt.Println(sub2, "appears", count2, "times in", s2)
}
