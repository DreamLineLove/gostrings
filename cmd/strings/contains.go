package main

import (
	"fmt"
	"strings"
)

func contains() {

	// trying out Contains
	str1 := "Vorstellen"
	sub1 := "stellen"
	contains1 := strings.Contains(str1, sub1)
	fmt.Print(str1, " contains ", sub1, ": ", contains1)
	WriteSpace(1)

	// is it case insensitive?
	sub1AllCaps := strings.ToUpper(sub1)
	contains1Caps := strings.Contains(str1, sub1AllCaps)
	fmt.Print(str1, " contains ", sub1AllCaps, ": ", contains1Caps)
	WriteSpace(1)
	// Spoiler, it is case sensitive

	// Search case-insensitively
	str2 := "APPLE WORLD WIDE DEVELOPERS' CONFERENCE (WWDC)"
	sub2 := "World Wide Developers' Conference (wwdc)"
	str2searchable := strings.ToLower(str2)
	sub2useable := strings.ToLower(sub2)
	contains2 := strings.Contains(str2searchable, sub2useable)
	fmt.Print(str2, " contains ", sub2, ": ", contains2)
	WriteSpace(1)

	// does it search consecutively or not
	sub2v2 := "World Wide Developers' Conference wwdc"
	sub2v2useable := strings.ToLower(sub2v2)
	contains2v2 := strings.Contains(str2searchable, sub2v2useable)
	fmt.Print(str2, " contains ", sub2v2, ": ", contains2v2)
	WriteSpace(1)
	// Spoiler, it does NOT search everywhere

	// trying out ContainsAny with non-consecutive substring
	containsAny1 := strings.ContainsAny(str2searchable, sub2v2useable)
	fmt.Print(str2, " contains ", sub2v2, ": ", containsAny1)
	WriteSpace(1)
	// ContainsAny searches non-consecutively (aka performs a holistic search)

	// trying out ContainsAny case-insensitively
	case1 := "WWDC"
	casesub1 := "wwdc"
	containsAnyCase1 := strings.ContainsAny(case1, casesub1)
	// containsAny2 := strings.ContainsAny(str2, sub2)
	// fmt.Print(str2, " contains ", sub2, ": ", containsAny2)
	fmt.Print(case1, " contains ", casesub1, ": ", containsAnyCase1)
	WriteSpace(1)
	// ContainsAny IS case-sensitive!

	// trying out ContainsRune
	str3 := "Wie heisen Sie?"
	var rune1 rune = 63
	var rune2 rune = 67
	containsRune1 := strings.ContainsRune(str3, rune1)
	containsRune2 := strings.ContainsRune(str3, rune2)
	fmt.Print(str3, " contains rune", rune1, ": ", containsRune1)
	WriteSpace(1)
	fmt.Print(str3, " contains rune", rune2, ": ", containsRune2)
	WriteSpace(1)

	// Step by step check success condition
	str4 := "Ich mochte in Deutschland Informatik studieren."
	searchStringsSlice := []string{
		"Ich",
		"mochte",
		"in Deutschland",
		".",
	}
	foundAll := containsAll(str4, searchStringsSlice)
	fmt.Println(searchStringsSlice)
	fmt.Print("Contains the above slice: ", foundAll)
	WriteSpace(1)

	// Looking for phrases case-insensitively
	str5 := "WWDC 2023 Apple Reality Headset"
	str5lowercaps := strings.ToLower(str5)
	searchStringsSlice2 := []string{
		"apple",
		"wwdc",
		"reality",
	}
	foundAll2 := containsAll(str5lowercaps, searchStringsSlice2)
	fmt.Println(searchStringsSlice2)
	fmt.Print("Contains the above slice: ", foundAll2)
	WriteSpace(1)
}

func containsAll(operand string, slice []string) bool {
	count := 0
	for i := 0; i < len(slice); i++ {
		if strings.Contains(operand, slice[i]) == true {
			count++
		}
	}
	if count == len(slice) {
		return true
	} else {
		return false
	}
}
