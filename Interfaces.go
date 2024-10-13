package main

import (
	"fmt"
	"math"
)

func main() {
	//creates slice of shapes so we can do loop
	//and display multiple at once
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("----")
	}
}

// Shape interface
// Interface groups types together based on methods
// in this case inorder to be a shape
// the struct must have and area func
// and a circumf func
type shape interface {
	area() float64
	circumf() float64
}

// creating structs
type square struct {
	length float64
}

type circle struct {
	radius float64
}

// Square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// Circle Methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

//using shape allows us to use 1 method for multiple structs

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is %0.2f\n", s, s.area())
	fmt.Printf("circumfrance of %T is %0.2f\n", s, s.circumf())
}
