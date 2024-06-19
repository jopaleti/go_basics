package main

import (
	"fmt"
)

func main() {
	x := 89
	m := &x

	fmt.Println(*m) 
	fmt.Println(m) 
}