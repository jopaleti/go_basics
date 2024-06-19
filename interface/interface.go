package main

import (
	"fmt"
	"math"
)

/*
	CREATING INTERFACE FOR THE SHAPE
	WHAT IS INTERFACE??
	Interface group types together based on their method
*/
type shape interface {
	area() float64
	circumf() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

// Square method and also a receiver function
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// Circle methods and also called a receiver function
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 50.00},
		circle{radius: 5.90},
		square{length: 16.11},
		circle{radius: 7.14},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("....")
	}
}