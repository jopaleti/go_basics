package main

import (
	"fmt"
)

func sendGreeting(nameList []string) {
	for _, value := range nameList {
		fmt.Printf("Hello %v, Good morning, How are you doing today? \n", value)
	}
}