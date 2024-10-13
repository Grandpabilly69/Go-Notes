package main

import (
	"fmt"
	"os"
)

// Note that this is connected to
// Structs&ConstumeTypes.go
// Note that this makes use of reciever functions
// Note that this makes use of reciver Functions with pointers
// Note that this makes use of saving files

// Structs are similar to objects
// Different from java OOP where
// we have to make a class
// Go we only need to creat a struct
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	//this intializes the object
	//and assigns values to the object
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	//and then returns the object
	return b
}

//Note below will be reciever functions
//These functions are associated to the Bill object

// the first set of parenthesis recieves a bill object
// this limits the functions association to only
// bill objects
// This function foramts the string
func (b *bill) format() string {
	fs := "Bill Breakdown:\n"
	var total float64 = 0

	//for each loop is a way to loop through
	//the map items from bill object and
	//increments total by the value in map item
	// the fs adds a new line with the key and value
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25s ...$%0.2f\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25s ...$%0.2f\n", "Tip:", b.tip)
	//the %-25 adds empty spaces for 25 characters long
	//it takes the length of the value and adds spaces
	//untill it has reached 25 characters long
	//Note that +/- for this tells weather the
	//spacing is added to the left or right of
	//the formatted variable
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total: ", total+b.tip)

	return fs
}

// this func updates Tip
// We use a pointer to reference to the pointer value
// Note that Go automatically dereferences pointers
// if they are being passed through a function
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Adds new item to map
// Note that GoLand automatically adds the astrix
// to create pointers in reciever functions
// as it doesn't make multiple copies of the
// object but only copies of the smaller things
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Saving bill using files
func (b *bill) save() {
	//we must save info into files in byte slices
	//this is a byte slice that takes bill format
	//as arguments
	data := []byte(b.format())

	//To deal with files we must make use of the os package

	//there are 3 arguments for the write file function
	//first is where: give it the directory and a name of the file
	//second is what you want to add which is data
	//third is the permissions
	//if the file already exists it will return an error
	err := os.WriteFile("Bill/"+b.name+".txt", data, 0644)

	//checks if there is an error
	if err != nil {
		//panic stops the flow of programme
		//and gives an error
		panic(err)
	}
	fmt.Println("Bill saved successfully!")
}
