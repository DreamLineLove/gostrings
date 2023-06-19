package main

import (
	"fmt"
	"strings"
	"time"
)

func explore() {
	// legacy()
	str1 := "this is a good old interpreted raw string"
	str2 := "should not contain at all"
	if strings.Contains(str1, str2) {
		fmt.Println("str1 contains str2!")
	} else {
		fmt.Println("str1 does not contain str2!")
	}

	str3 := "apple"
	var a rune = 'a'
	if strings.ContainsRune(str3, a) {
		fmt.Println(str3, "contains a")
	}
	str3_space := str3 + " "
	var space rune = ' '
	if strings.ContainsRune(str3_space, space) {
		fmt.Println("\""+str3_space+"\"", "contains a whitespace")
	}

	charSet1 := " oriter"
	if strings.ContainsAny(str1, charSet1) {
		fmt.Println("str1 contains any of", strings.TrimPrefix(charSet1, " "))
	}

	str4 := "the Go Standard Library!"
	str4_noprefix := strings.TrimPrefix(str4, "the ")
	str4_nosuffix := strings.TrimSuffix(str4, "!")
	str4_noprefix_nosuffix := strings.TrimSuffix(str4_noprefix, "!")
	fmt.Println()
	fmt.Println(str4_noprefix_nosuffix)
	fmt.Println(str4_nosuffix, "is the best standard library of any language!")

	fmt.Println()
	str5 := "hello it's me, i've been blah blah blah..."
	var slice1 []string
	slice1 = strings.Fields(str5)
	for _, value := range slice1 {
		fmt.Println("\t\"" + value + "\"")
	}
}

func legacy() {
	// weekdayGuessingGame()
	WriteSpace(1)

	protocols := []string{
		"http",
		"https",
	}
	url := "http://google.com"
	usesKnownProtocol := false
	var usedprotocol string
	for i := 0; i < len(protocols); i++ {
		if strings.HasPrefix(url, protocols[i]) {
			usesKnownProtocol = true
			protocol, _, _ := strings.Cut(url, "://")
			usedprotocol = protocol
		}
	}
	fmt.Println("url:", url)
	fmt.Println("url uses a known protocol:", usesKnownProtocol)
	fmt.Println("protocol:", usedprotocol)

	s1 := " Ich habe meine Hausaufgaben noch nicht gemacht"
	fmt.Println(s1)
	s1_doubleWhiteSpace := strings.ReplaceAll(s1, " ", "  ")
	s1words := strings.Fields(s1)
	s1_double_words := strings.Fields(s1_doubleWhiteSpace)
	fmt.Println(s1words)
	fmt.Println(s1_doubleWhiteSpace)
	fmt.Println(s1_double_words)
	WriteSpace(1)

	s2 := "users/old/60plus/alive"
	s2fields := strings.Fields(s2)
	s2directories := strings.Split(s2, "/")
	fmt.Println(s2fields)
	for i := 0; i < len(s2directories); i++ {
		fmt.Println(s2directories[i])
	}
	WriteSpace(1)

	s3 := "You are the King of Scotland. You are the law in this country. You are blessed by God."
	s3_after_conquest := strings.ReplaceAll(s3, "You", "I")
	s3_semifinal := strings.Replace(s3_after_conquest, "are", "am", 3) + " I am better than any Kings."
	s3_final := strings.Replace(s3_semifinal, "King", "Queen", 1)
	fmt.Println(s3)
	fmt.Println(s3_semifinal)
	fmt.Println(s3_final)
	WriteSpace(1)

	s4 := "/users/type/admin/id/"
	s4directories := strings.FieldsFunc(s4, func(c rune) bool {
		return c == '/'
	})
	s4slice := strings.Split(s4, "/")
	for i := 0; i < len(s4directories); i++ {
		fmt.Println("\"" + s4directories[i] + "\"")
	}
	WriteSpace(1)
	for i := 0; i < len(s4slice); i++ {
		fmt.Println("\"" + s4slice[i] + "\"")
	}
	WriteSpace(1)

	s5 := "/users/type/admin/with?id=123/"
	s5fields := strings.FieldsFunc(s5, func(c rune) bool {
		return c == '/' || c == '?' || c == '='
	})
	for i := 0; i < len(s5fields); i++ {
		fmt.Println("\"" + s5fields[i] + "\"")
	}
	WriteSpace(1)

}

func weekdayGuessingGame() {
	expectedInput := time.Now().Weekday().String()
	var userInput string
	fmt.Print("What day is it today...")
	fmt.Scan(&userInput)

	for !strings.EqualFold(expectedInput, userInput) {
		WriteSpace(1)
		fmt.Println("Pleae try again!")
		fmt.Print("What day is it today...")
		fmt.Scan(&userInput)
	}
	return
}
