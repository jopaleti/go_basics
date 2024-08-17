package main

import (
	"fmt"
)

func main() {
	// Example string containing Unicode characters
	str := "Hello, 世界!"

	// Count the number of runes (Unicode code points) in the string
	runeCount := 0
	for _, r := range str {
		runeCount++
		fmt.Printf("Rune: %c, Unicode: %U\n", r, r)
	}

	fmt.Printf("Total runes in string: %d\n", runeCount)
}
