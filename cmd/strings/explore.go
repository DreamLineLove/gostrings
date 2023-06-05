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

	s2before, s2after, s2result := strings.Cut(s2, ".")
	s2beforeE, s2afterE, s2resultE := strings.Cut(s2, "...")
	fmt.Println(s2before + ".")
	fmt.Println(strings.Trim(s2after, " "))
	fmt.Println(s2result)
	fmt.Println(s2beforeE)
	fmt.Println(s2afterE)
	fmt.Println(s2resultE)
	WriteSpace(1)

	google := "https://google.com"
	protocol, path, googleResult := strings.Cut(google, "://")
	fmt.Println(protocol)
	fmt.Println(path)
	fmt.Println(googleResult)
	WriteSpace(1)

	justDomain, err := cutPrefixAndSuffixAndPrint(google, "https://", ".com")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("Just the domain:", justDomain)
	}

}

func cutPrefixAndSuffixAndPrint(operand string, prefix string, suffix string) (string, error) {
	noPrefix, found := strings.CutPrefix(operand, prefix)
	if found {
		noSuffix, found := strings.CutSuffix(noPrefix, suffix)
		if found {
			return noSuffix, nil
		} else {
			return "", fmt.Errorf("Suffix not found!")
		}
	} else {
		return "", fmt.Errorf("Prefix not found!")
	}
}
