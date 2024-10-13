package main

import "fmt"

func main() {

	printFunc()

	printlnFunc()

	printfFunc()

	sprintfFunc()
}

func printFunc() {
	//This function does not use the \n automatically
	fmt.Print("Hello ")
	fmt.Print("World")
}

func printlnFunc() {
	//This function does use the \n automatically
	name := "Matt"
	fmt.Println("Hello ")
	fmt.Println("World")
	fmt.Println("Hello ", 44, "is my age", name, "is my name")
}

func printfFunc() {
	//This function does not use the \n automatically
	//The modulas is a format specifier

	age := 44
	name := "Matt"
	fmt.Printf("Hello %v , I am %v years old", name, age)

	//v = global format (works on any variable)
	//q = adding quotes (only works on strings)
	//T = outputs type of variable
	//f = float format
	//////add 0.2 inbetween modulas and f to get 2 decimel places
	//////eg. "%0.2f", 321,42221 = 321,42
	//% = returns percent sign/modulas
	//t = boolean
	//b = base 2(binary)
	//d = base 10(integer)
	//O = base 8(octal)
	//X = base 16(hex)
	//e = scientific notation
	//g = exponents
	//s = string
	//p = pointer

}

func sprintfFunc() {
	//This is the Printf function but returns the formated string
	name := "Matt"
	age := 33
	var str = fmt.Sprintf("Hello %v , I am %v years old", name, age)
	fmt.Println(str)
}
