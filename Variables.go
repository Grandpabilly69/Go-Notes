package main

import "fmt"

func main() {
	strs()
	nums()
}

func strs() {
	var nameOne string = "Alice"
	var nameTwo string = "Bob"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"

	nameThree = "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Luigi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
}

func nums() {
	//integers
	/////////////////////////
	var age1 int = 44
	var age2 = 44
	age3 := 44

	var num1 int8 = 3 //from 8 -> 64 bits
	var num2 uint = 7 //only positive ints from 8 -> 64 bits if specifics needed

	fmt.Println(age1, age2, age3, num1, num2)
	////////////////////////////////////

	//Floats
	/////////////////////////////////
	var scrore float32 = 4534.3
	var score2 float64 = 3232.3

	fmt.Println(scrore, score2)
}
