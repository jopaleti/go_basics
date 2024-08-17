package main

import (
	"fmt"
)

func main() {
	defer safeExist()

	const one int = 2
	if one != 1 {
		panic("One is not 1. This is not good!!!!")
	}
	
}

func safeExist() {
	if r := recover(); r != nil {
		fmt.Println("Panic is recovered!!!!")
	}
}