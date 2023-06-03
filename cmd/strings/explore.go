package main

import (
	"fmt"
	"strings"
)

func explore() {

	longText := "Ich habe am Morgen einen Deutschkurs. Ich lerne Deutsch weil Deutsch sehr wichtig fur mein Studium ist. Ich mochte in Deutschland studieren."

	// cloneLongText := strings.Clone(strings.Split(longText, ""))

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
}
