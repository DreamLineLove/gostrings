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
	str1_upper := strings.ToUpper(str1)
	fmt.Println(str1_upper)

	// Return an all lower caps version of a string
	str1_lower := strings.ToLower(str1)
	fmt.Println(str1_lower)

	WriteSpace(2)

	var str2 string = "Toyotomi Hideyoshi was a shogun of Japan. During his time, he invaded Korea in what is known today as the largest sea-based invasion second only to the landings at Normandy."
	fmt.Println(str2)
	WriteSpace(1)

	// Returns true or false whether a string contains a substring
	contains_Normandy := strings.Contains(str2, "Normandy")
	fmt.Println("Contains \"Normandy\":", contains_Normandy)

	// Returns how many times a string contains a substring
	count_the := strings.Count(str2, "the")
	fmt.Println("Contains \"the\":", count_the, "times")

	WriteSpace(2)

	var str3 string = "https://github.com/DreamLineLove/gostrings"
	fmt.Println(str3)

	WriteSpace(1)

	// Returns true or false whether a strings starts with a string
	starts_with_https := strings.HasPrefix(str3, "https")
	fmt.Println("Starts with \"https\":", starts_with_https)

	// Returns true or false whether a string ends with a string
	ends_with_gostrings := strings.HasSuffix(str3, "gostrings")
	fmt.Println("Ends with \"gostrings\":", ends_with_gostrings)

	// Returns length of string
	str3_length := len(str3)
	fmt.Println("Length:", str3_length)
	WriteSpace(2)

	hello := "Hello"
	fmt.Println(hello)
	space := " "
	world := "World"
	fmt.Println(world)
	slice := []string{
		hello,
		space,
		world,
	}
	fmt.Println(slice)
	WriteSpace(1)

	hello_world_string := strings.Join(slice, "")
	fmt.Println("After using join:", hello_world_string)
	WriteSpace(2)

	var str4 string = "Yangon is the capital of Burma. It is teeming with people!"
	str5 := strings.Replace(str4, "is", "was", 1)
	str6 := strings.Replace(str5, "Burma", "Myanmar", 1)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)

	// I should make a command line quiz of some sort to test myself on content contained.
	// use make use of flags
	// should implement in the other repo as well
}
