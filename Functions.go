package main

import (
	"fmt"
	"math"
)

func main() {
	//calling functions
	//fill parameters if applicable
	sayGreeting("Mario")
	sayBye("goo")

	//calling function
	//remember when calling func within func
	//don't put arguments in interior func
	cycleNames([]string{"matt", "ams", "abd"}, sayGreeting)

	//remember that since it returns value you must use the value
	//if you don't it is redundant
	//formating to get 2 decimal places
	fmt.Printf("circle is %0.2f", circleArea(10.23))
}

// func is keyword
// without func it is not a function
func sayGreeting(s string) {
	fmt.Println("hello,", s)
}

// needs parenthesis and a body
// parameters/arguments are within parenthesis
func sayBye(s string) {
	fmt.Println("bye, ", s)
}

// arguments are slice of string
// and function with a argument of string
func cycleNames(n []string, f func(string)) {

	//for each loop goes through each element of slice
	//then calls the function that is in the arguments
	//and gives interior func argument from
	//slice element
	for _, value := range n {
		f(value)
	}
}

// function returns a float
// you must specify return type after parenthesis
// you must use return in func if you specify a return type
func circleArea(n float64) float64 {
	//imported math package for pi
	return math.Pi * n * n
}
