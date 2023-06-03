package main

import (
	"fmt"
	"strings"
)

func split() {
	longText := "Ich habe am Morgen einen Deutschkurs. Ich lerne Deutsch weil Deutsch sehr wichtig fur mein Studium ist. Ich mochte in Deutschland studieren."

	splitSlice := strings.Split(longText, ".")
	for i := 0; i < len(splitSlice); i++ {
		fmt.Print(splitSlice[i], "\"\n")
	}
	fmt.Println("Length:", len(splitSlice))
	WriteSpace(1)

	splitAfterSlice := strings.SplitAfter(longText, ". ")
	for i := 0; i < len(splitAfterSlice); i++ {
		trimSpace := strings.Trim(splitAfterSlice[i], " ")
		// No space whatsoever
		fmt.Print(trimSpace, "\"\n")
		fmt.Print(splitAfterSlice[i], "\"\n")
	}
	// Using ". " as separator does not separate the string after the last . because there is no space after that.
	fmt.Println("Length:", len(splitAfterSlice))
	WriteSpace(1)

	splitAfterNSlice := strings.SplitAfterN(longText, ".", 5)
	for i := 0; i < len(splitAfterNSlice); i++ {
		trimSpace := strings.Trim(splitAfterNSlice[i], " ")
		fmt.Print(trimSpace, "\"\n")
		fmt.Print(splitAfterNSlice[i], "\"\n")
	}
	fmt.Println("Length:", len(splitAfterNSlice))
	WriteSpace(1)

	var splitNSlice []string
	splitNSlice = strings.SplitN(longText, ".", 3)
	for i := 0; i < len(splitNSlice); i++ {
		fmt.Print(splitNSlice[i], "\"\n")
	}
	for i := 0; i < len(splitNSlice); i++ {
		trimmed := strings.Trim(splitNSlice[i], ". ")
		fmt.Print(trimmed, "\"\n")
	}
	fmt.Println("Length:", len(splitNSlice))
	WriteSpace(1)

	// Including the separator but removing it and other unwanted artifacts after the fact
	splitAfterNSliceTwo := strings.SplitAfterN(longText, ".", 3)
	var count int
	for i := 0; i < len(splitAfterNSliceTwo); i++ {
		fmt.Print(splitAfterNSliceTwo[i], "\"\n")
		count++
	}
	splitAfterNSliceTwoTrimmed := make([]string, 3)
	for i := 0; i < len(splitAfterNSliceTwo); i++ {
		trimmed := strings.Trim(splitAfterNSliceTwo[i], ". ")
		splitAfterNSliceTwoTrimmed[i] = trimmed
	}
	for i := 0; i < len(splitAfterNSliceTwoTrimmed); i++ {
		fmt.Print(splitAfterNSliceTwoTrimmed[i], "\"\n")
	}
	WriteSpace(1)

	textToTrim := "...The full-stops to the left as.well.as. spaces at the end should be trimmed. Not those inside   "
	trimmedText := strings.Trim(textToTrim, ". ")
	fmt.Print(textToTrim, "\"\n")
	fmt.Print(trimmedText, "\"\n")
	WriteSpace(1)

	// Testin edge cases
	str1 := "Does not include sep"
	sep1 := "Jeden Tag"
	test1slice := strings.Split(str1, sep1)
	// Should return a slice with original string str1
	fmt.Print(test1slice)
	WriteSpace(1)
	fmt.Println("Length: ", len(test1slice))
	WriteSpace(1)

	str2 := "Sep is empty"
	sep2 := ""
	fmt.Println(str2)
	test2slice := strings.SplitN(str2, sep2, 100)
	for i := 0; i < len(test2slice); i++ {
		fmt.Println(test2slice[i])
	}
	fmt.Println("Length: ", len(test2slice))
	WriteSpace(1)

	str3 := ""
	sep3 := ""
	test3slice := strings.SplitAfter(str3, sep3)
	fmt.Println(test3slice)
	WriteSpace(1)

	// Getting a slice of all the characters in a string
	str4 := "Aufwiedersehen"
	fmt.Println(str4)
	characters := strings.Split(str4, "")
	holen := []string{
		"h", "o", "l", "e", "n",
	}
	for j := 9; j < len(characters); j++ {
		characters[j] = holen[j-9]
	}
	fmt.Println(characters)
	str4Rejoined := strings.Join(characters, "")
	fmt.Println(str4Rejoined)
	WriteSpace(1)
}
