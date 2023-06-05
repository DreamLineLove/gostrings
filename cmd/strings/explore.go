package main

import (
	"fmt"
	"strings"
	"time"
)

func explore() {

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

}
