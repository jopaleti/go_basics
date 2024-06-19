package main

import (
	"fmt"
)

func updateVariable(n *string) {
	*n = "wedge"
}
func updateMenu(y map[string]float64) {
	y["soup"] = 7000.00
	y["coffee"] = 678
}

func main() {
	// Mapping
	menu := map[string]float64{
		"soup": 12.30,
		"pie": 90.00,
	}
	updateMenu(menu)
	fmt.Println(menu)

	for _, value := range menu {
		fmt.Println(value)
	}

	// Creating a phonebook map
	var phonebook map[int]string = map[int]string{
		8108603111: "Oluwatobi",
		7006603111: "Daddy",
	}

	// fmt.Println(phonebook[8108603111])
	phonebook[8108603111] = "Sawale"
	fmt.Println(phonebook)

	x := "Axe"
	updateVariable(&x)
	fmt.Println(x)
}