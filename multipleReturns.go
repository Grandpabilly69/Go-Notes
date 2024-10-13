package main

import (
	"fmt"
	"strings"
)

func main() {
	//intitialize 2 vars at same time cause
	//it returns 2 values
	s1, s2 := getInitials("matt fairhurst")
	fmt.Println(s1, s2)
}

// Use parenthesis for multiple return types
func getInitials(n string) (string, string) {

	//uppercases the string
	s := strings.ToUpper(n)

	//splits so each slice is part of each name
	names := strings.Split(s, " ")

	var intitials []string
	//for each loop assigns the first char of each slice for initials
	for _, name := range names {
		intitials = append(intitials, name[:1])
	}
	//checks if there is 1 or 2 initials
	//and returns based on how many initials there are
	if len(intitials) > 1 {
		return intitials[0], intitials[1]
	} else {
		return intitials[0], "_"
	}
}
