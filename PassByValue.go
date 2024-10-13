package main

//Note that func groupA will have coppied func
//that shows how to correctly deal with this issue
import "fmt"

// Data types are seperated into 2 groups
// Group A and Group B
func main() {
	groupA()
	groupAFix()

	groupB()
}

// Group A types
// strings, ints, floats, bools, arrays, structs
func groupA() {
	//this function shows that group A types
	//make a copy of themselves when called by another function
	//this does not change/save the value of the original variable
	//Group A is also known as Non-Pointer Values

	name := "tifa"
	updateName(name)
	fmt.Println(name)
}

func updateName(x string) {
	x = "wedge"
}

// ////////////////////////////////////////////////////
// note that this is the correct way of doing
// thr func groupA
func groupAFix() {
	name := "tifa"
	name = updateNameFix(name)
	fmt.Println(name)
}
func updateNameFix(x string) string {
	x = "wendie"
	return x
}

// /////////////////////////////////////
// Note that group B types are:
// slices, maps, functions
// Group B is also known as Pointer Wrapper Values
func groupB() {
	menu := map[string]float64{
		"pie":      4.33,
		"icecream": 8.88,
	}
	updateMenu(menu)
	fmt.Println(menu)

}

// In this case since this a group B data Type
// the value is added/updated to the original
// variable
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}
