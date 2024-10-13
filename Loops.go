package main

import "fmt"

func main() {

	forloop()
	trueForLoop()
	sliceLoop()
	forInRangeLoop()
}

// While loop structure
func forloop() {
	x := 0
	//for key word works like while loop
	for x < 5 {
		fmt.Println(x)
		//++ increments by 1
		x++
	}

}

// This is a true for loop in traditional sense
func trueForLoop() {
	//similar/same syntax as java
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// Uses traditional for loop and loops through slice
func sliceLoop() {
	names := []string{"mario", "luigi", "peach", "yoshi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

}

// Similar to for each loop
func forInRangeLoop() {
	names := []string{"mario", "luigi", "peach", "yoshi"}

	//first argument is the position
	//second argument is the value
	for index, name := range names {
		fmt.Println(index, name)
	}

	//use underscore if you only want the value and not position
	for _, name := range names {
		fmt.Println(name)

		//updating the value does not update the element in the slice
		//it only updates the argument/variable 2 which is name in this case
		name = "new str"
	}
}
