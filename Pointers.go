package main

import "fmt"

func main() {
	namesMake()
}

// this is showing how to store a memory adress
// Note that the & must be used to reference to the
// memory location of the variable and not the actual value
func memaddressPointer(name string) {
	fmt.Println("Mem address", &name)
	//m variable stores the memory address of
	//the name variable but also has its own
	//memory address but stores names mem address
	//since we said m = &name
	m := &name
	fmt.Println(m)

	//this gets the value that is associated with the memory address
	//note that the astrix is needed for this
	//note that in this case it will return tifa
	//cause it is assosciated with the stored
	//memory address
	fmt.Println(*m)
	updatename(m)

}

func namesMake() {
	name := "tifa"

	memaddressPointer(name)

	fmt.Println(name)

}

// the * in front of the type is refering to
// accepting a pointer that has that data type
// stored within it
func updatename(x *string) {
	//the * infront of x is dereferncing the value
	//which is tifa and replacing it with wedge
	//note that is still keeps the same
	//memory address
	*x = "wedge"

	//this means that the original value from name is changed
	//this is beacause when we call name it looks for the
	//memory address and since we changed the value in the memory
	//address it changes the original value
}
