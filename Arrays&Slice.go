package main

import "fmt"

func main() {
	arrs()

	slice()
}

func arrs() {

	//either way is fine
	//NB: arrays have a fixed length and can't be changed
	var ages [3]int = [3]int{1, 2, 3}
	var names = [3]string{"matt", "cool", "kile"}

	//this is how you change items in array
	names[1] = "Changed"

	fmt.Println(ages)
	fmt.Println(names)
}

func slice() {
	var ages = []int{1, 2, 3}
	var names = []string{"matt", "cool", "kile"}

	//slices are the same as arrays when editting an item
	ages[0] = 12

	//this is how to add item at end of slice
	names = append(names, "Wacko")

	//how to get the length of arrays/slices/string
	fmt.Println(ages, len(names))
	//////////////////////////////////
	//Slice ranges
	//////////////////////////////////
	//first num represents the start of range
	//second num represents end of range
	//NB: first is included, Second is excluded
	range1 := names[1:3]
	range2 := names[1:]
	range3 := names[:3]

	fmt.Println(range1, range2, range3)

	//slice range is still a range with less than OG
	//therfore we can treat as slice

	range1 = append(range1, "Yippe")

	fmt.Println(range1)
}
