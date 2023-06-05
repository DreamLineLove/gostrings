package main

import (
	"fmt"
	"strings"
	"time"
)

func explore() {
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
	for i := 0; i < len(s4slice); i++ {
		fmt.Println("\"" + s4slice[i] + "\"")
	}
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
