package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Println(n, "Good morning")
} 

// Passing a function into another function as an argument
func cycleNameGreeting(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

// Function that hs return
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// Multiple return values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	
	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}

func main() {
	sayGreeting("Oluwatobi")
	cycleNameGreeting([]string{"Oluwatobi", "Raphael"}, sayGreeting)
	fmt.Println(circleArea(20))

	// iNITIALIZING THE MULTIPLE RETURNS FUNCTION
	firstName, secondName := getInitials("Oluwatobi Opaleti")
	fmt.Println(firstName, secondName)
}