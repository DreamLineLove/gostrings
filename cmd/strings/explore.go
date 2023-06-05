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

	s1 := "Ich habe meine Hausaufgaben noch nicht gemacht"
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
