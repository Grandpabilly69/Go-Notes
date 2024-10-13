package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strsPackage()
	sortPackage()
}

func strsPackage() {
	greeting := "Hello Friends"
	//.contains returns boolean value if the string has the substring in it
	fmt.Println(strings.Contains(greeting, "Hellow"))

	//replace all searchs first argument string
	//for the second argument string
	//and replaces the second argument string
	//with the third argument string
	//Note it does not change the original string just returns the new string
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"))
	fmt.Println(greeting)

	//returns string in uppercase
	//Note this does not change original string
	fmt.Println(strings.ToUpper(greeting))

	//returns the position of the first character is the sub string only
	//Note starts from 0
	fmt.Println(strings.Index(greeting, "ll"))

	//splits the first string argument into different strings
	//based on the second argument
	//Note it formats it in array format
	//Note will split multiple times if possible
	fmt.Println(strings.Split(greeting, " "))
}

func sortPackage() {
	//Ints
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	//this takes the slice argument and sorts in ascending order
	//Note this changes original slice
	sort.Ints(ages)
	fmt.Println(ages)

	//it searchs of the int in argument 2
	//from the slice in argument 1
	//And returns position of the element found
	//Note in this example it will search through the sorted slice
	//Note if number is larger than the rest and not found in slice
	//it returns an index 1 greater than the last element
	//else if it less than the largest element
	//it returns 0
	index := sort.SearchInts(ages, 700)
	fmt.Println(index)

	//Strings
	names := []string{"yoshi", "mario", "luigi", "peach", "bowser"}

	//Sorts string slice in typical alphabetical order
	sort.Strings(names)
	fmt.Println(names)

	//this searches the slice and returns the position of the element searched
	fmt.Println(sort.SearchStrings(names, "mario"))
}
