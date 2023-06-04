package main

import (
	"fmt"
	"strings"
)

func explore() {

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

	// trying out ContainsAny case-insensitively
	containsAny2 := strings.ContainsAny(str2, sub2)
	fmt.Print(str2, " contains ", sub2, ": ", containsAny2)
	WriteSpace(1)
	// ContainsAny is both case-INsensitive and searches non-consecutively (aka performs a holistic search)

}
