package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Note that this is connected to
// Bill.go
// Note that this makes use of reciever functions
// Note that this makes use of reciver Functions with pointers
// Note this also goes through user input
// Note that this also makes use of switch statements
// Note that this also makes use of parsing floats
// Note that this also makes use of saving to files
func main() {
	mybill := userInput()
	promptOptions(mybill)
}

func userInput() bill {
	//bufio is an import to read from somewhere
	//os.STdin tells bufio to read from terminal
	reader := bufio.NewReader(os.Stdin)

	//We have done the above in a function so
	//we do not need to rewrite the code
	name, _ := getInput("Enter bill: ", reader)

	//this removes any white space the user has inputed
	//this is helpful for string manipulation
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("Bill created - ", b.name)

	return b
}
func getInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)

	//input is variable that is assigned the input
	//readString takes in the input
	//the arguments is to so it knows when
	//to start reading the input
	//Note the argument must be in single quotes
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("A - add,\nS - Save,\nT - add tip:\n", reader)

	//switch case is same as in java
	switch opt {
	case "A":
		//calls getInput to store the values and prompt
		//user for item and price
		name, _ := getInput("Enter Item: ", reader)
		price, _ := getInput("Enter Price: ", reader)
		fmt.Println(name, price)

		//string convert is the package that has the parsefloat
		//function, This takes 2 arguments
		//the string variable that we want to convert
		//and the bit size of the float
		//if it can't convert then it will through an error
		p, err := strconv.ParseFloat(price, 64)
		//nil is the same as null in java
		//checks if there is an error
		//if so it outputs the bellow and calls the function again
		if err != nil {
			fmt.Println("price must be a number")
			promptOptions(b)
		}
		//if there is no error
		//we add item to the map
		b.addItem(name, p)
		fmt.Println("Added - ", name, price)
		promptOptions(b)
	case "S":
		b.save()
		fmt.Println("You chose save the bill ", b)
	case "T":
		tip, _ := getInput("Enter Tip: ", reader)
		fmt.Println(tip)

		//this is the same as case a
		//except it does it for tip
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Added - ", tip)
		promptOptions(b)
		//default is if none of the other cases
		//have been applied
	default:
		fmt.Println("Not an option")
		//Using recursion so they can input a new value
		//Recycles the function
		promptOptions(b)
	}
}

func manuelInput() {
	//this instantiates the object and then calls the new bill
	//to create the object
	mybill := newBill("Marios Bill")

	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.85)
	mybill.addItem("toffee pudding", 4.95)
	mybill.addItem("coffee", 3.25)

	//this is how to call a reciever function
	//note that you cant call format without
	//mybill since it needs to recieve a bill object

	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
