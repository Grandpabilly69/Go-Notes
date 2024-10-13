package main

import "fmt"

var age int = 2

func main() {
	bool()
	ifStatement()
	cont_Break()
}

func bool() {

	//operators
	//less than or equal to : true
	fmt.Println(age <= 50)
	//greater than or equal to : false
	fmt.Println(age >= 50)
	//equal to : true
	fmt.Println(age == 45)
	//not equal to : true
	fmt.Println(age != 50)

}

func ifStatement() {

	//if statements are simple and similar/same as java
	if age < 30 {
		fmt.Println("age less than 30")
	} else if age < 40 {
		fmt.Println("age less than 40")
	} else {
		fmt.Println("age greater than 40")
	}
}

func cont_Break() {
	names := []string{"yoshi", "mario", "luigi", "peach", "bowser"}

	//for loop to showcase continue
	for index, value := range names {
		if index == 1 {
			fmt.Println("continue at ", index)
			//continue skips past this element and does the rest of elements
			continue
		}
		if index > 2 {
			//break breaks out the loop so there is less processing
			fmt.Println("break at ", index)
			break
		}
		fmt.Println(value, index)
	}
}
