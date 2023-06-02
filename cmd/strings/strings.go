package main

import (
	"fmt"
	"strings"
)

func stringsPkg() {

	WriteSpace(2)
	var str1 string = "First String"
	fmt.Println(str1)
	WriteSpace(1)

	// Return an all caps version of a string
	var str1_upper = strings.ToUpper(str1)
	fmt.Println(str1_upper)

	// Return an all lower caps version of a string
	var str1_lower = strings.ToLower(str1)
	fmt.Println(str1_lower)

	WriteSpace(2)

	var str2 string = "Toyotomi Hideyoshi was a shogun of Japan. During his time, he invaded Korea in what is known today as the largest sea-based invasion second only to the landings at Normandy."
	fmt.Println(str2)
	WriteSpace(1)

	// Returns true or false whether a string contains a substring
	var contains_Normandy = strings.Contains(str2, "Normandy")
	fmt.Println("Contains \"Normandy\":", contains_Normandy)

	// Returns how many times a string contains a substring
	var count_the = strings.Count(str2, "the")
	fmt.Println("Contains \"the\":", count_the, "times")
}
