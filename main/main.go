package main

import (
	"fmt"
	"strings"
	"sort"
)

func main()  {
	greeting := "Hello, good morning my Neighbours!"
	fmt.Println(strings.Contains(greeting, "morning"))
	fmt.Println(strings.ReplaceAll(greeting, "morning", "afternoon"))
	fmt.Println(greeting)
	fmt.Println(strings.Index(greeting, "N"))
	fmt.Println(strings.Split(greeting, " "))
	// Sort Package
	ages := []int{45, 9, 76, 0, 1}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 76)
	fmt.Println("Index of 76 in ages is ", index)
	x := strings.Split(greeting, " ")
	sort.Strings(x)
	fmt.Println(x)
	
	// var str []string = []string{"olu", "bade"}
	// str[0] = "oluwatobi"
	// str = append(str, "segun")
	// fmt.Println(str)
	// fmt.Println(len(str))
	// // Slicing in Go
	// rangeOne := str[1:]
	// fmt.Println(rangeOne)
}
