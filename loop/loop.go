package main

import (
	"fmt"
)

func main() {
	// for x := 0; x < 5; x++ {
	// 	fmt.Printf("Value of x is %v \n", x)
	// }
	names := []string{"oluwatobi", "Johnson", "opaleti", "segun"}
	// for index, value := range names {
	// 	fmt.Printf("Name of index %v is %v \n", index, value)
	// }
	for _, value := range names {
		if value == "oluwatobi" {
			fmt.Println("Yoo! my name is printed")
			continue
		}
		fmt.Println(value)
	}
}