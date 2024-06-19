package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calling greeting.go modules inside main.go")
	names := []string{"Oluwatobi", "Johnson", "Raphael"}
	sendGreeting(names)
}